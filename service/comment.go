package service

import (
	"goStudy/blog/dao/db"
	"goStudy/blog/model"
)

func GetCommentListByArticleId(articleId int64) (commentList []*model.Comment, err error) {
	commentList, err = db.GetCommentListByArticleId(articleId)
	return
}

func CreateComment(comment *model.Comment) (commentId int64, err error) {
	commentId, err = db.AddComment(comment)
	if err != nil {
		return
	}
	return
}