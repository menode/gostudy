package controller

import (
	"ginstudy/logger"
	"ginstudy/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginGetAction(c *gin.Context) {
	logger.Debug("您进入了登录页面")
	if userName, err := c.Request.Cookie("user_name"); err != nil || userName.Value == "" {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		var user model.User
		user.UserName = userName.Value
		c.HTML(http.StatusOK, "results.html", user)
	}

}

func LoginPostAction(c *gin.Context) {
	logger.Info("点击了登录按钮")
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		logger.Errorf("登录错误，错误信息：%v", err)
		//把错误返回i前端
		user.UserName = "用户名不能为空"
		user.Password = "密码不能为空"
		c.HTML(http.StatusOK, "results.html", user)
		return
	}
	c.SetCookie("user_name", user.UserName, 30, "/", "localhost", false, false)
	c.HTML(http.StatusOK, "results.html", user)
}

func LoginPutAction(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}
