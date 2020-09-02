package service

import (
	"fmt"
	"goStudy/blog/dao/db"
	"goStudy/blog/model"
)

//获取主页文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleList []*model.ArticleRecord, err error)  {
	// 1.获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2.获取文章对应的分类(多个)
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	// 聚合信息，返回页面
	for _, article := range articleInfoList {
		tmpArticleRecord := &model.ArticleRecord{}
		tmpArticleRecord.ArticleInfo = *article
		for _, category := range categoryList {
			if category.CategoryId == article.CategoryId {
				tmpArticleRecord.Category = *category
				break
			}
		}
		articleList = append(articleList, tmpArticleRecord)
	}
	return
}

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
	//遍历文章，得到每个文章
	for _, v := range articleInfoList {
		//从当前文章取出文章分类id
		existTag := false
		for _, id := range ids {
			if id == v.CategoryId {
				existTag = true
			}
		}
		if existTag {
			continue
		} else {
			ids = append(ids, v.CategoryId)
		}
	}
	return ids
}

//根据分类id, 获取该类文章和他们对应的分类信息
func GetArticleRecordListByCategoryId(categoryId int64, pageNum, pageSize int) (articleRecordList [] *model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticlesByCategoryId(categoryId, pageNum, pageSize)
	fmt.Println("articleInfoList:", articleInfoList, err)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2.获取文章对应的分类(多个)
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	// 聚合信息，返回页面
	for _, article := range articleInfoList {
		tmpArticleRecord := &model.ArticleRecord{}
		tmpArticleRecord.ArticleInfo = *article
		for _, category := range categoryList {
			if category.CategoryId == article.CategoryId {
				tmpArticleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, tmpArticleRecord)
	}
	return
}

//根据文章id，获取文章详情
func GetArticleDetailById(atricleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(atricleId)
	if err != nil {
		return
	}
	return
}

//获取上下一篇文章
func GetNearArticle(atricleId int64) (nearArticles []*model.ArticleInfo, err error) {
	nearArticles, err = db.GetNearArticleFromDb(atricleId)
	return
}

//创建文章
func CreateArticle(article *model.ArticleDetail) (articleId int64, err error) {
	articleId, err = db.InsertArticle(article)
	return
}