package server

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io/ioutil"
	"store/service/internal/conf"
	"time"
)

var ProviderSet = wire.NewSet(NewRegistrar, NewGRPCServer)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	cli, err := createEtcdClient(conf)
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)
	return r
}

func createEtcdClient(conf *conf.Registry) (*clientv3.Client, error) {
	var _tlsConfig *tls.Config
	if conf.Etcd.PemPath != nil && conf.Etcd.KeyPemPath != nil && conf.Etcd.CaPemPath != nil {
		// load cert
		cert, err := tls.LoadX509KeyPair(*conf.Etcd.PemPath, *conf.Etcd.KeyPemPath)
		if err != nil {
			return nil, err
		}

		// load root ca
		caData, err := ioutil.ReadFile(*conf.Etcd.CaPemPath)
		if err != nil {
			return nil, err
		}

		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caData)

		_tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      pool,
		}
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Etcd.Address,
		DialTimeout: 5 * time.Second,
		Username:    conf.Etcd.User,
		Password:    conf.Etcd.Pwd,
		TLS:         _tlsConfig,
	})

	return cli, err
}
