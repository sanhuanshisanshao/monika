package database

import (
	"monika/model"
	"testing"
	"time"
)

func TestNewDatabase(t *testing.T) {
	url := "root:123456@tcp(localhost:3306)/monika?charset=utf8"
	db, err := NewDatabase(url)
	if err != nil {
		t.Errorf("NewDatabase error:%v", err)
	}

	err = db.db.DB().Ping()
	if err != nil {
		t.Errorf("Ping error:%v", err)
	}

	//err = db.CreateTable(&model.Weibo{}, &model.WeiboForward{})
	//
	//if err != nil {
	//	t.Errorf("CreateTable error:%v", err)
	//}

	err = db.InsertWeibo(model.Weibo{Name: "gaoq", Content: "你好！", Time: time.Now(), Client: "iPhone"})
	if err != nil {
		t.Errorf("InsertWeibo error:%v", err)
	}
}
