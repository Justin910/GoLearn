package conf

import "time"

type Config struct {
	Server *Server `json:"server"`
	Data   *Data   `json:"data"`
}

type Data struct {
	Database *Database `json:"database"`
}

type Server struct {
	Grpc *GRPC `json:"grpc"`
}

type GRPC struct {
	Addr    string         `json:"addr"`
	Timeout *time.Duration `json:"timeout"`
}

type Database struct {
	Driver          string         `json:"driver"`
	Source          string         `json:"source"`
	MaxOpenConns    *int           `json:"maxopenconns"`
	ConnMaxLifeTime *time.Duration `json:"connmaxlifetime"`
}

type Registry struct {
	Etcd *Etcd `json:"etcd"`
}

type Etcd struct {
	Address    []string `json:"address"`
	User       string   `json:"user"`
	Pwd        string   `json:"pwd"`
	PemPath    *string  `json:"pem_path"`
	KeyPemPath *string  `json:"key_pem_path"`
	CaPemPath  *string  `json:"ca_pem_path"`
}
