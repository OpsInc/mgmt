package forms

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func formToJSON() string {
  return string("test")
  // Need to add a function to fetch forms data and transforms to JSON
  // https://github.com/gin-gonic/gin/issues/364 ---> Example with from data
}

func testfunc(c *gin.Context) {

  // contentType := c.Header.Get("Content-Type")
  contentType := c.Request.Header
   
  c.String(http.StatusOK, contentType.Get("Content-Type"))
  // Will show Content-Type: application/x-www-form-urlencoded
}

func testfunc2(c *gin.Context) {
  name := c.PostForm("fname")
  c.JSON(http.StatusOK, gin.H{
    "jsonName": name,
  })
}

func Routes(route *gin.Engine){
  forms := route.Group("/forms")
  {
    forms.POST("/", testfunc) //Default same has /forms
    forms.POST("/test", testfunc2) // same as /forms/test
  }
}
