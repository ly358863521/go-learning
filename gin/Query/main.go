package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)
func main() {
	router := gin.Default()

// url 为 http://localhost:8080/welcome?name=ningskyer时
// 输出 Hello ningskyer
// url 为 http://localhost:8080/welcome时
// 输出 Hello Guest
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest") //可设置默认值
		// 是 c.Request.URL.Query().Get("lastname") 的简写
		lastname := c.Query("name") 
		fmt.Printf("Hello %s\n", name)
		fmt.Println(lastname)
		c.String(http.StatusOK, "Hello %s", name)

		
	})

	router.Run(":8080")
}