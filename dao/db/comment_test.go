package db

import (
	"testing"
)

func init()  {
	err := Init("root:hanitt5477@tcp(127.0.0.1:3306)/blogger?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
}

func TestGetCommentListByArticleId(t *testing.T) {
	commentList, err := GetCommentListByArticleId(8)
	if err != nil {
		t.Errorf("GetCommentListByArticleId failed, err:%v", err)
	} else {
		t.Logf("commentList:%v", *commentList[0])
	}
}