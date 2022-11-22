package v1

import (
	"alert-manager/internal/apiserver/config"
	"alert-manager/internal/pkg/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// InitApi
// @description: 初始化路由
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/28 15:27
// @success:
func InitApi() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//是否跨域
	if config.Config.Web.Cors {
		router.Use(middleware.CORS)
	}

	apiV1 := router.Group("alert/v1")
	//ping服务检测接口
	apiV1.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	{
		//加载其他接口
		initPolicyRouter(apiV1)  //策略模块
		initMessageRouter(apiV1) //告警消息模块
	}

	//启动gin路由服务
	log.Println("端口号：", config.Config.Web.Port)
	err := router.Run(fmt.Sprintf(":%s", config.Config.Web.Port))
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic("http服务启动失败")
	}
}
