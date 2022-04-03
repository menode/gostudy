package controller

import (
	"ginstudy/logger"
	"ginstudy/model"
	"ginstudy/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	products := service.SelectProduct()
	c.HTML(http.StatusOK, "products.html", products)
}

func GetProduct(context *gin.Context) {
	idstring := context.Query("id")
	if idstring == "" {
		idstring = "0"
	}
	var product model.Product
	if id, err := strconv.ParseUint(idstring, 10, 64); err != nil {
		logger.Debugf("转换参数错误： %v", err)
	} else {
		product = service.GetProduct(id)
	}
	context.HTML(http.StatusOK, "product.html", product)
}

func SaveProduct(context *gin.Context) {
	var product model.Product
	if err := context.ShouldBind(&product); err != nil {
		logger.Error("模型参数绑定错误： %v", err)
	} else {
		service.SaveProduct(product)
	}
	context.Redirect(http.StatusFound, "/products")
}

func DeleteProduct(context *gin.Context) {
	idstring := context.Query("id")
	if idstring == "" {
		idstring = "0"
	}
	if id, err := strconv.ParseUint(idstring, 10, 64); err != nil {
		logger.Debugf("转换参数错误： %v", err)
	} else if id > 0 {
		service.DeleteProduct(id)
	}
	context.Redirect(http.StatusFound, "/products")
}
