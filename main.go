package main

import (
	"github.com/gin-gonic/gin"
	"goStudy/blog/controller"
	"goStudy/blog/dao/db"
)

func main()  {
	router := gin.Default()
	//数据库初始化
	dsn := "root:hanitt5477@tcp(127.0.0.1:3306)/blogger?charset=utf8mb4&parseTime=True"
	err := db.Init(dsn)
	if err != nil {
		panic(err)
	}
	//加载静态文件
	router.Static("/static/", "./static")
	//加载模板
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.Categorylist)
	router.GET("/article/detail/", controller.DetailHandle)
	router.GET("/article/new/", controller.NewArticleHandler)

	router.POST("/comment/submit/", controller.CommentAddHandler)
	router.POST("/article/submit/", controller.ArticleSubmitHandler)

	router.Run(":8000")
}
