package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goStudy/blog/service"
	"net/http"
	"strconv"
)

//访问主页的控制器
func IndexHandle(c *gin.Context) {
	//从service取数据
	//1.加载文章数据
	articleList, err := service.GetArticleRecordList(1, 10)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//2.加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H {
		//gin.H本质上一个map
		"article_list":articleList,
		"category_list":categoryList,
	})
}

//点击分类云进行分类
func Categorylist(c *gin.Context)  {
	categoryIdStr := c.Query("category_id")
	//转换类型
	cateforyId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusNonAuthoritativeInfo, "views/500.html", nil)
		return
	}
	articleRecords, err := service.GetArticleRecordListByCategoryId(cateforyId, 1, 10)
	if err != nil {
		c.HTML(http.StatusNonAuthoritativeInfo, "views/500.html", nil)
		return
	}

	//再次加载所有分类数据，用于分类云显示
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusNonAuthoritativeInfo, "views/500.html", nil)
		return
	}
	fmt.Println("articleRecords:", articleRecords)
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":articleRecords,
		"category_list":categoryList,
	})
}

//点击首页文章标题，进入文章详情页
func DetailHandle(c *gin.Context)  {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		fmt.Printf("ParseInt failed, err is:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleDetail, err := service.GetArticleDetailById(articleId)
	if err != nil {
		fmt.Printf("GetArticleDetailById failed, err is:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	commentList, err := service.GetCommentListByArticleId(articleId)
	if err != nil {
		fmt.Printf("GetCommentListByArticleId failed, err is:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//获取上下篇
	nearArticles, err := service.GetNearArticle(articleId)
	if err != nil || len(nearArticles) != 2 {
		fmt.Printf("GetNearArticle failed, err is:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	prevArticle := nearArticles[0]
	nextArticle := nearArticles[1]
	c.HTML(http.StatusOK, "views/detail.html", gin.H{
		"detail":articleDetail,
		"comment_list":commentList,
		"prev":prevArticle,
		"next":nextArticle,
	})
}

