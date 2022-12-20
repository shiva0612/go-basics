package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// request body → JSON, XML, Form with files, Form wo files
// query(using form tag),
// params(using uri tag)
// header(using header tag)
func main() {
	var c *gin.Context
	var a any
	c.ShouldBindUri(a)
	c.ShouldBindQuery(a)
	c.ShouldBindHeader(a)
	c.ShouldBind(a) //default is form
	c.ShouldBindJSON(a)
	c.ShouldBindXML(a)
	c.ShouldBindBodyWith(a, binding.JSON)
	// c.ShouldBindBodyWith(a, binding.Form)
}
