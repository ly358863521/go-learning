package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

var port = flag.String("p", "8080", "the listening port")

type status struct {
	NodeID     string   `json:"node_id"`
	Roles      string   `json:"roles,omitempty"`
	NodeStatus string   `json:"node_status"`
	CPU        string   `json:"cpu"`
	Memory     string   `json:"memory"`
	Speed      string   `json:"speed"`
	Age        string   `json:"age"`
	Pod        []string `json:"pod,omitempty"`
}

var edge1 = []status{
	{NodeID: "node95a", Roles: "master", NodeStatus: "NotReady", CPU: "0m", Memory: "20Mi", Speed: "100Mbps", Age: "5d17h"},
	{NodeID: "node95b", NodeStatus: "NotReady", CPU: "0m", Memory: "20Mi", Speed: "100Mbps", Age: "5d17h", Pod: []string{"nginx_1", "nginx_2"}},
}

func main() {
	router := gin.Default()

	router.GET("/getstatus", func(c *gin.Context) {
		c.JSON(http.StatusOK, edge1)

	})
	router.Run(":" + *port)

}
