package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type userinfo struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/", indexHandler)
	r.POST("/login", formHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func formHandler(c *gin.Context) {
	var userinfo userinfo
	// name := c.PostForm("name")
	// password := c.PostForm("password")
	// fmt.Printf("name:%s,password:%s\n",name,password)
	if c.ShouldBind(&userinfo) == nil {
		fmt.Println(userinfo.Name, userinfo.Password)
		if userinfo.Name == "liyan" && userinfo.Password == "123" {
			// c.JSON(200, gin.H{"status": "you are logged in"})
			c.HTML(200, "welcome.html", gin.H{
				"status": "you are logged in",
				"name":   userinfo.Name,
			})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}

}
