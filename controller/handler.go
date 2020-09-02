package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goStudy/blog/model"
	"goStudy/blog/service"
	"net/http"
	"strconv"
	"time"
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
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		//gin.H本质上一个map
		"article_list":  articleList,
		"category_list": categoryList,
	})
}

//点击分类云进行分类
func Categorylist(c *gin.Context) {
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
		"article_list":  articleRecords,
		"category_list": categoryList,
	})
}

//点击首页文章标题，进入文章详情页
func DetailHandle(c *gin.Context) {
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
		"detail":       articleDetail,
		"comment_list": commentList,
		"prev":         prevArticle,
		"next":         nextArticle,
	})
}

func NewArticleHandler(c *gin.Context) {
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/post_article.html", gin.H{
		"category_list": categoryList,
	})
	return
}

type AddArticleReq struct {
	Author     string `json:"author"`
	Title      string `json:"title"`
	CategoryId string `json:"category_id"`
	Content    string `json:"content"`
}

func ArticleSubmitHandler(c *gin.Context) {
	addreq := &AddArticleReq{}
	//author := c.Query("author")
	//title := c.Query("title")
	//categoryIdStr := c.Query("category_id")
	//content := c.Query("content")
	err := c.ShouldBindJSON(addreq)
	if err != nil {
		fmt.Println("0000000")
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return	}
	fmt.Println("categoryIdStr:", addreq.CategoryId)
	//转换类型
	categoryId, err := strconv.ParseInt(addreq.CategoryId, 10, 64)
	if err != nil {
		fmt.Println("1111111")
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleInfo := model.ArticleInfo{
		Title:      addreq.Title,
		CreateTime: time.Now(),
		UserName:   addreq.Author,
	}
	articleDetail := &model.ArticleDetail{
		ArticleInfo: articleInfo,
		Content:     addreq.Content,
		Category: model.Category{
			CategoryId: categoryId,
		},
	}
	articleId, err := service.CreateArticle(articleDetail)
	if err != nil {
		fmt.Println("2222222")
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	fmt.Println("create article success, articleId:", articleId)
	c.HTML(http.StatusOK, "/", nil)
	return
}

func CommentAddHandler(c *gin.Context) {
	commentStr := c.Query("comment")
	author := c.Query("author")
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	comment := &model.Comment{
		Id:         4,
		Content:    commentStr,
		UserName:   author,
		Status:     1,
		ArticleId:  articleId,
		CreateTime: time.Now(),
	}
	_, err = service.CreateComment(comment)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
}
