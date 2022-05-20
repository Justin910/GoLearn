package main

import (
	"fmt"
	"github.com/pkg/errors"
	_ "github.com/pkg/errors"
)

/**
* @Author: 徐家圳
* @Date: 2022/5/20 22:46
 */

type DBInterface interface {
	QueryRow(query string, args ...interface{}) DBRow
}

type DBRow interface {
	Scan(dest ...interface{}) error
}

func QueryUserNameFromDB(dbconn DBInterface, uid string) (string, error) {
	var name string
	err := dbconn.QueryRow("select name from hm_user_info where uid=?", uid).Scan(&name)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Can't Get UserName, UID: %s", uid))
		//return "", err
	}
	return name, err
}
