package controller

import (
	"iblog/model"

	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {

	db := model.DB

	ulist := make([]model.Admin, 0)

	db.Select("id,username,password,addtime").Find(&ulist)

	ctx.JSON(200, gin.H{
		"data": ulist,
	})

}
