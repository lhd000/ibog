package controller

import (
	"encoding/json"
	"iblog/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	//ctx.Render("/admin/login")

	session := sessions.Default(ctx)
	u := session.Get("admin")
	if u != nil {
		ctx.Redirect(302, "/admin/index")
	}

	ctx.HTML(http.StatusOK, "login.html", gin.H{

		"data": "this is index",
	})
}

func Admin(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin_index.html", gin.H{
		"data": "this is blog backend",
	})
}

func Article(ctx *gin.Context) {

	a := &model.Article{}
	alist := a.List()

	js, _ := json.Marshal(alist)

	ctx.HTML(http.StatusOK, "admin_article.html", gin.H{
		"data": "this is the atcile page",
		"ali":  string(js),
	})
}

func AdminCategory(ctx *gin.Context) {

	list := model.GetCateList()

	js, _ := json.Marshal(list)

	log.Println("the json data is :", string(js))
	//log.Println("this list data is :", list)

	ctx.HTML(http.StatusOK, "admin_category.html", gin.H{
		"data": string(js),
	})
}

func CategoryAdd(ctx *gin.Context) {

	var cat model.Category

	err := ctx.Bind(&cat)

	if err != nil {
		log.Println("the data valiate err :", err)
	}

	cat.AddTime = time.Now()

	log.Println("the form data is :", cat)

	e := model.AddCate(&cat)

	if e != nil {
		log.Println("the add cat error :", e)
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(http.StatusOK, "<script>alert('ADD FAIL');location.href='/admin/category';</script>")
		return

	} else {

		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(http.StatusOK, "<script>alert('添加成功');location.href='/admin/category';</script>")
		//ctx.Redirect(302, "/admin/category")
		return
	}
	/*
		if model.AddCate(&cat) == true {
			log.Println("the add catgory duccess")
		} else {
			log.Println("add the category fail ")
		}
	*/

	// //ctx.JSON(http.StatusOK, gin.H{
	// 	"data": "this is add  category page",
	// })
}

func Categoryupdate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin_update-category.html", gin.H{
		"data": "update the category page",
	})
}

func CategoryDel(ctx *gin.Context) {

	var p model.Category
	ctx.Bind(&p)

	log.Println("the bind data is :", p)

	model.DeleteCategory(p.ID)

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "<script>alert('删除成功');location.href='/admin/category';</script>")
	//ctx.Redirect(302, "/admin/category")
	return
}

func ArticleAdd(ctx *gin.Context) {
	list := model.GetCateList()
	js, _ := json.Marshal(list)

	ctx.HTML(http.StatusOK, "admin_add-article.html", gin.H{
		"data": "this is the add article page ",
		"cat":  string(js),
	})
}

func ArticleAddDo(ctx *gin.Context) {

	formData := &model.Article{}

	ctx.Bind(&formData)

	formData.AddTime = time.Now()
	formData.Add()

	log.Println("the bind data is :", formData)

	log.Println("add the article do function ")
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "<script>alert('添加成功');location.href='/admin/article';</script>")
	//ctx.Redirect(302, "/admin/category")
	return
}

func ArticleDelete(ctx *gin.Context) {

	a := &model.Article{}

	ctx.Bind(&a)

	log.Println("the delete article is : ", a)

	err := a.DeleteArticleByID()

	if err != nil {

		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(http.StatusOK, "<script>alert('删除失败');location.href='/admin/article';</script>")
		//ctx.Redirect(302, "/admin/category")
		return
	}

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "<script>alert('删除成功');location.href='/admin/article';</script>")
	//ctx.Redirect(302, "/admin/category")
	return
}

func ArticleUpdatePage(ctx *gin.Context) {

	a := &model.Article{}
	ctx.Bind(&a)
	log.Println("the bind update article data is :", a)

	data := a.FindArticleByID()

	log.Println("the article data is : ", data)

	js, _ := json.Marshal(data)

	log.Println("the article update page aricle json is :", string(js))

	m := StructToMap(*data)

	clist := model.GetCateList()

	cjson, _ := json.Marshal(clist)

	ctx.HTML(http.StatusOK, "admin_update-article.html", gin.H{
		"data": "this is article update page ",
		"js":   string(js),
		"map":  m,
		"cjs":  string(cjson),
	})
}

func ArticleUpdate(ctx *gin.Context) {

	formData := &model.Article{}
	ctx.Bind(&formData)

	log.Println("the bind form data is ", formData)

	formData.AddTime = time.Now()

	err := formData.UpdateArticleByID()

	if err != nil {
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(http.StatusOK, "<script>alert('更新失败');location.href='/admin/article';</script>")
		//ctx.Redirect(302, "/admin/category")
		return
	}

	log.Println("update form data is :", formData)
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "<script>alert('修改成功');location.href='/admin/article';</script>")
	//ctx.Redirect(302, "/admin/category")
	return
}

func Manageuser(ctx *gin.Context) {

	list := model.AdminList()

	js, _ := json.Marshal(list)

	log.Println("THE JSON DATA IS :", js)

	ctx.HTML(http.StatusOK, "admin_manage-user.html", gin.H{
		"data": "this is the  manager user page",
		"js":   string(js),
	})
}

func UserAdd(ctx *gin.Context) {

	u := &model.Admin{}
	ctx.Bind(&u)

	log.Println("THE USER IS :", u)

	err := u.Add()

	if err != nil {
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(http.StatusOK, "<script>alert('添加失败');location.href='/admin/manageuser';</script>")
		//ctx.Redirect(302, "/admin/category")
		return
	}

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "<script>alert('添加成功');location.href='/admin/manageuser';</script>")
	//ctx.Redirect(302, "/admin/category")
	return
}

func UserDel(ctx *gin.Context) {

	u := &model.Admin{}

	ctx.Bind(&u)
	log.Println("the bind data is :", u)

	u.DeleteByID()

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "<script>alert('删除成功');location.href='/admin/manageuser';</script>")
	return
}

func LoginAction(ctx *gin.Context) {

	u := &model.Admin{}

	ctx.Bind(&u)
	log.Println("the login bind data is :", u)

	re := u.Login()

	log.Println("login result is :", re)

	if re <= 0 {

		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(http.StatusOK, "<script>alert('登录失败');location.href='/login';</script>")
		return
	}

	session := sessions.Default(ctx)
	js, _ := json.Marshal(model.Admin{ID: re, UserName: u.UserName, PassWord: u.PassWord})
	jsstr := string(js)

	session.Set("admin", jsstr)
	e := session.Save()

	if e != nil {
		log.Println("session save error is", e)
	}

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "<script>alert('登录成功');location.href='/admin/index';</script>")
	return

}

func LogOut(ctx *gin.Context) {

	session := sessions.Default(ctx)
	//session.Set("admin", nil)
	session.Delete("admin")
	session.Save()

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "<script>alert('退出成功');location.href='/login';</script>")
	return
}
