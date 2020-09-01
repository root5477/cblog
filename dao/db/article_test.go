package db

import (
	"fmt"
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

func TestInsertArticle(t *testing.T) {
	articleInfo := model.ArticleInfo{
		Id:3,
		Summary:"golang入门教程",
		Title:"go语言之旅",
		ViewCount:1277,
		CreateTime:time.Now(),
		CommentCount:3456,
		UserName:"cq",
	}



	article := &model.ArticleDetail{
		ArticleInfo:articleInfo,
		Content:"balabalabalabalabalabalabalabalabalabalabalabalabalabalabalabala",
		Category:model.Category{
			CategoryId:3,
		},
	}
	articleId, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("articleId:%v", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 3)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Logf("articleList:%v", articleList)
}

func TestGetArticleDetail(t *testing.T) {
	articles, err := GetArticleDetail(8)
	if err != nil {
		t.Errorf("err:%v", err)
	} else {
		t.Logf("articles:%v", articles)
	}
}

func TestGetNearArticleFromDb(t *testing.T) {
	res, err := GetNearArticleFromDb(2)
	if err != nil {
		t.Errorf("GetNearArticleFromDb err:%v", err)
	} else {
		t.Logf("GetNearArticleFromDb resp:%v", res)
		fmt.Println(*res[0])
		fmt.Println(*res[1])
	}
}

