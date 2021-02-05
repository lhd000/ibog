package controller

import (
	"encoding/json"
	"iblog/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Front(ctx *gin.Context) {

	list := model.GetCateList()

	js, _ := json.Marshal(list)

	ctx.HTML(http.StatusOK, "front_index.html", gin.H{
		"data": "this is front page",
		"cat":  string(js),
	})
}

func AboutUs(ctx *gin.Context) {

	list := model.GetCateList()

	js, _ := json.Marshal(list)

	ctx.HTML(http.StatusOK, "front_about.html", gin.H{
		"data": "this is front about us page",
		"cat":  string(js),
	})
}

func Contacts(ctx *gin.Context) {

	list := model.GetCateList()

	js, _ := json.Marshal(list)

	ctx.HTML(http.StatusOK, "front_contact.html", gin.H{
		"data": "this is front contact page",
		"cat":  string(js),
	})
}

func FrontCategory(ctx *gin.Context) {

	list := model.GetCateList()

	js, _ := json.Marshal(list)

	c := &model.Category{}
	ctx.Bind(&c)
	alist := c.GetArticleListByCid()

	catinfo := c.GetCategoryByCid()

	catjs, _ := json.Marshal(catinfo)

	ajson, _ := json.Marshal(alist)

	ctx.HTML(http.StatusOK, "front_category.html", gin.H{
		"data":  "this is front category page",
		"cat":   string(js),
		"ajs":   string(ajson),
		"catjs": string(catjs),
	})
}

func Author(ctx *gin.Context) {

	list := model.GetCateList()

	js, _ := json.Marshal(list)

	ctx.HTML(http.StatusOK, "front_author.html", gin.H{
		"data": "this is front category page",
		"cat":  string(js),
	})
}

func BlogPost(ctx *gin.Context) {

	list := model.GetCateList()

	js, _ := json.Marshal(list)

	a := &model.Article{}
	ctx.Bind(&a)
	article := a.FindArticleByID()

	ajson, _ := json.Marshal(article)

	cat := &model.Category{ID: article.Cid}
	catinfo := cat.GetCategoryByCid()

	cjson, _ := json.Marshal(catinfo)

	ctx.HTML(http.StatusOK, "front_blog-post.html", gin.H{
		"data": "this is front blog_post page",
		"cat":  string(js),
		"ajs":  string(ajson),
		"cjs":  string(cjson),
	})
}

func Blank(ctx *gin.Context) {

	list := model.GetCateList()

	js, _ := json.Marshal(list)

	ctx.HTML(http.StatusOK, "front_blank.html", gin.H{
		"data": "this is front blog_post page",
		"cat":  string(js),
	})
}
