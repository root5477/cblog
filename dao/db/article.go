package db

import (
	_ "github.com/go-sql-driver/mysql"
	"goStudy/blog/model"
)

//添加文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	//加个验证
	if article == nil {
		return
	}

	sqlStr := `insert into article 
			  (category_id, summary, content, title, view_count, comment_count, username)
			   values
			   (?,?,?,?,?,?,?);`
	result, err := DB.Exec(sqlStr, article.Category.CategoryId, article.Summary, article.Content, article.Title,
		article.ViewCount, article.CommentCount, article.UserName)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

//查询文章
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum <0 || pageSize <0 {
		return
	}

	//按时间降序排序
	sqlStr :=  `select
					id, summary, title, view_count, create_time, comment_count, username
				from article
				where
					status = 1
				order by create_time desc
				limit ?,?`
	start := (pageNum - 1) * pageSize
	size := pageSize
	err = DB.Select(&articleList, sqlStr, start, size)
	return
}

//根据文章id查询单个文章
func GetArticleDetail(articleId int64) (articleDetails *model.ArticleDetail, err error) {
	articleDetails = &model.ArticleDetail{}
	if articleId < 0 {
		return
	}
	sqlStr :=  `select
					id, summary, content, title, view_count, create_time, comment_count, username, category_id
				from article
				where
					id=?
				and
					status=1`

	err = DB.Get(articleDetails, sqlStr, articleId)
	return
}

//根据分类id,查询这一类的文章
func GetArticlesByCategoryId(categoryId int64, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	sqlStr :=  `select
					id, summary, title, view_count, create_time, comment_count, username
				from article
				where
					status = 1
				and
					category_id =?
				order by create_time desc
				limit ?,?`
	start := (pageNum - 1) * pageSize
	size := pageSize
	err = DB.Select(&articleList, sqlStr, categoryId, start, size)
	if err != nil {
		return
	}
	return
}