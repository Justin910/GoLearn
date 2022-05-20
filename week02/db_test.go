package main

import (
	"database/sql"
	"errors"
	"testing"
)

// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 答: 需要; 因为可能在接收到请求以后，在中间的逻辑处理时有些问题，导致的向数据库请求时参数不对，返回了查不到数据，这种情况下需要Wrap查看到整个调用链方便排查问题, 并且可以携带错误参数返回
func TestQueryUserName_ErrNoRows(t *testing.T) {

	dd := &StruDBDelegate{
		Err: sql.ErrNoRows,
	}
	uid := "123"

	name, err := QueryUserNameFromDB(dd, uid)
	if err != nil {
		// Mysql使用Rows查询报错时，需要区分开是数据库报错，还是查不到导致的报错。
		// 根据具体业务实际情况选择是否处理
		if errors.Is(err, sql.ErrNoRows) {
			// 这个人不存在
			//fmt.Printf("%+v\n", err)
			t.Logf("%+v", err)
		} else {
			// 数据库查询异常，需要触发告警紧急处理
			t.Fatalf("%+v", err)
		}
		return
	}

	t.Logf("UserName: %s", name)
}

func TestQueryUserName_ErrConnDone(t *testing.T) {

	dd := &StruDBDelegate{
		Err: sql.ErrConnDone,
	}
	uid := "123"

	name, err := QueryUserNameFromDB(dd, uid)
	if err != nil {

		// Mysql使用Rows查询报错时，需要区分开是数据库报错，还是查不到导致的报错。
		// 根据具体业务实际情况选择是否处理
		if errors.Is(err, sql.ErrNoRows) {
			// 这个人不存在
			t.Logf("%+v", err)
			return
		} else {
			// 数据库查询异常，需要触发告警紧急处理
			t.Fatalf("%+v", err)
		}
	}

	t.Logf("UserName: %s", name)
}

func TestQueryUserName_Success(t *testing.T) {

	dd := &StruDBDelegate{}
	uid := "123"

	name, err := QueryUserNameFromDB(dd, uid)
	if err != nil {

		// Mysql使用Rows查询报错时，需要区分开是数据库报错，还是查不到导致的报错。
		// 根据具体业务实际情况选择是否处理
		if errors.Is(err, sql.ErrNoRows) {
			// 这个人不存在
			t.Logf("%+v", err)
			return
		} else {
			// 数据库查询异常，需要触发告警紧急处理
			t.Fatalf("%+v", err)
		}
	}

	t.Logf("UserName: %s", name)
}
