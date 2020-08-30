package db

import (
	"github.com/jmoiron/sqlx"
	"goStudy/blog/model"
)

//添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := `insert into category (category_name, category_no) values (?,?);`
	res, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = res.LastInsertId()
	return
}


//获取单个的文章分类
func GetArticleById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := "select id, category_name, category_no from category where id=?;"
	err = DB.Get(category, sqlStr, id)
	if err != nil {
		return
	}
	return
}

//获取多个分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error)  {
	//构建sql
	sqlStr, args, err := sqlx.In("select id, category_name, category_no from category where id in (?);", categoryIds)
	if err != nil {
		return
	}
	err = DB.Select(&categoryList, sqlStr, args...)
	if err != nil {
		return
	}
	return
}

//查询所有分类
func GetAllCategorys() (allCategoryList []*model.Category, err error)  {
	sqlStr := "select id, category_name, category_no from category order by category_no asc;"
	err = DB.Select(&allCategoryList, sqlStr)
	return
}