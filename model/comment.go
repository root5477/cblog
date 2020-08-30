package model

import "time"

type Comment struct {
	Id         int64     `db:"id"`
	Content    string    `db:"content"`
	UserName   string    `db:"username"`
	Status     int       `db:"status"`
	ArticleId  int64     `db:"article_id"`
	CreateTime time.Time `db:"create_time"`
}
