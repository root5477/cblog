package model

import "time"

type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	UserName     string    `db:"username"`
}

//用于文章详情页的实体
//为了提升效率
type ArticleDetail struct {
	ArticleInfo
	//文章内容
	Content string `db:"content"`
	Category
}

//用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
