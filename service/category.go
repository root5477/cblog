package service

import (
	"goStudy/blog/dao/db"
	"goStudy/blog/model"
)

//获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategorys()
	if err != nil {
		return
	}
	return
}



