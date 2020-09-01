package db

import (
	"goStudy/blog/model"
	"testing"
	"time"
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

func TestAddComment(t *testing.T) {
	comment := &model.Comment {
		Id:4,
		Content:"大神666777",
		UserName:"cq2",
		Status:1,
		ArticleId:8,
		CreateTime:time.Now(),
	}
	commentId, err := AddComment(comment)
	if err != nil {
		t.Errorf("AddComment failed, err is:%v", err)
	} else {
		t.Logf("AddComment success, commentId:%v", commentId)
	}
}