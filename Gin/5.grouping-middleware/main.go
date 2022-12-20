package main

import "github.com/gin-gonic/gin"

func main() {
	// Per route middleware, you can add as many as you desire.
	r := gin.Default()
	r.GET("/form", middleware1, middleware2, acutalHandler)

	authorized := r.Group("/auth")
	authorized.Use(middleware1, middleware2)
	{
		authorized.POST("/signup", acutalHandler) // -> /auth/signup
		authorized.POST("/login", acutalHandler)  // -> /auth/login

		// nested group
		testing := authorized.Group("testing")
		// visit 0.0.0.0:8080/testing/analytics
		testing.GET("/analytics", acutalHandler)
	}
}

func middleware1(c *gin.Context) {

}

func middleware2(c *gin.Context) {
}

func acutalHandler(c *gin.Context) {
}
