package mgmt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func formToJSON() string {
  return string("test")
  // Need to add a function to fetch mgmt data and transmgmt to JSON
  // https://github.com/gin-gonic/gin/issues/364 ---> Example with from data
}

func testfunc(c *gin.Context) {

  // header := c.Header.Get("Content-Type")
  header := c.Request.Header
   
  c.String(http.StatusOK, header.Get("Content-Type"))
  // Will show Content-Type: application/x-www-form-urlencoded
  // if Content-Type = form type ---> formToJSON func
}

func testfunc2(c *gin.Context) {
  name := c.PostForm("fname")

  // write to DB
  //

  // Response to client
  c.JSON(http.StatusOK, gin.H{
    "jsonName": name,
  })
}

func Routes(route *gin.Engine){
  mgmt := route.Group("/mgmt")
  {
    mgmt.POST("/orders", testfunc) // Default same as /mgmt
    mgmt.POST("/buy", testfunc2) // same as /mgmt/test
  }
}
