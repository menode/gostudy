package router

import (
	"ginstudy/controller"

	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	engine.GET("/login", controller.LoginGetAction)
	engine.POST("/login", controller.LoginPostAction)
	engine.PUT("/login", controller.LoginPutAction)

	engine.GET("/products", controller.ProductList)
	engine.GET("/product", controller.GetProduct)
	engine.GET("/delete-product", controller.DeleteProduct)
	engine.POST("/products", controller.SaveProduct)
}
