// Path: internal/apiserver/router/v1
// FileName: message.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/18$ 15:55$

package v1

import (
	"alert-manager/internal/apiserver/controller"
	"github.com/gin-gonic/gin"
)

func initMessageRouter(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("message")
	messageController := controller.MessageController{}
	api.POST("create", messageController.Create) //创建
	api.POST("list", messageController.List)     //列表

}
