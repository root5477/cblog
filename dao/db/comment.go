package db

import "goStudy/blog/model"

func AddComment(comment *model.Comment) (commentId int64, err error) {
	if comment == nil {
		return
	}
	sqlStr := ` insert into comment
				(id,content,username,status,article_id)
				values
				(?,?,?,?,?);`
	res, err := DB.Exec(sqlStr,comment.Id, comment.Content, comment.UserName, comment.Status, comment.ArticleId)
	if err != nil {
		return
	}
	commentId, err = res.LastInsertId()
	return
}

func GetCommentListByArticleId(articleId int64) (commentList []*model.Comment, err error) {
	if articleId <= 0 {
		return
	}
	sqlstr := ` select
				id,content,username,status,article_id,create_time
				from comment
				where
				article_id = ?
				and
				status = 1;`
	err = DB.Select(&commentList, sqlstr, articleId)
	return
}
