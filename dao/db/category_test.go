package db

import "testing"
import _ "github.com/go-sql-driver/mysql"


func init() {
	err := Init("root:hanitt5477@tcp(127.0.0.1:3306)/blogger?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
}

func TestGetArticleById(t *testing.T) {
	res, err := GetCategoryById(2)
	if err != nil {
		t.Errorf("GetArticleById failed, err is:%v\n", err)
	} else {
		t.Logf("GetArticleById success, res is:%v\n", res)
	}
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		t.Errorf("err:%v", err)
	} else {
		for _, v := range categoryList {
			t.Logf("%v", v)
		}
	}
}

func TestGetAllCategorys(t *testing.T) {
	categoryList, err := GetAllCategorys()
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("%v", v)
	}}