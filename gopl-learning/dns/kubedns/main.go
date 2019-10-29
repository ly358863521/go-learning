package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var port = flag.String("p", "8080", "the listening port")

type db struct {
	db   map[string]string
	path string
}

var (
	trueIPmu  sync.RWMutex
	falseIPmu sync.RWMutex
	rulemu    sync.RWMutex
)
var (
	trueIP  = db{db: make(map[string]string), path: "trueIP.json"}
	falseIP = db{db: make(map[string]string), path: "falseIP.json"}
	rule    = db{db: make(map[string]string), path: "rule.json"}
)

//初始化数据
func initial() bool {
	f, err := ioutil.ReadFile(trueIP.path)
	if err != nil {
		log.Fatalf("load trueip.json failed: %s", err)
		return false
	}
	json.Unmarshal(f, &trueIP.db)

	f, err = ioutil.ReadFile(falseIP.path)
	if err != nil {
		log.Fatalf("load falseip.json failed: %s", err)
		return false
	}
	json.Unmarshal(f, &falseIP.db)

	f, err = ioutil.ReadFile(rule.path)
	if err != nil {
		log.Fatalf("load rule.json failed: %s", err)
		return false
	}
	json.Unmarshal(f, &rule.db)
	return true
}

//写回
func (p db) writeback() bool {
	str, err := json.Marshal(p.db)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	if ioutil.WriteFile(p.path, str, 0777) == nil {
		fmt.Println("suceess!")
	}
	return true
}
func iptoseg(ip string) string {
	return ip
}
func addTrueIP(c *gin.Context) {
	trueIPmu.Lock()
	defer trueIPmu.Unlock()
	name := c.Param("name")
	ip := c.Param("ip")
	trueIP.db[name] = ip
	if ok := trueIP.writeback(); !ok {
		c.String(http.StatusOK, "add failed!")
	}
	c.String(http.StatusOK, "add success!")
}
func delTrueIP(c *gin.Context) {
	trueIPmu.Lock()
	defer trueIPmu.Unlock()
	name := c.Param("name")
	if _, ok := trueIP.db[name]; ok {
		delete(trueIP.db, name)
		c.String(http.StatusOK, "del %s success!", name)
		trueIP.writeback()
	} else {
		c.String(http.StatusOK, "%s not exist!", name)
	}
}
func addFalseIP(c *gin.Context) {
	falseIPmu.Lock()
	defer falseIPmu.Unlock()
	name := c.Param("name")
	ip := c.Param("ip")
	falseIP.db[name] = ip
	if ok := falseIP.writeback(); !ok {
		c.String(http.StatusOK, "add failed!")
	}
	c.String(http.StatusOK, "add success!")
}
func delFalseIP(c *gin.Context) {
	falseIPmu.Lock()
	defer falseIPmu.Unlock()
	name := c.Param("name")
	if _, ok := falseIP.db[name]; ok {
		delete(falseIP.db, name)
		c.String(http.StatusOK, "del %s success!", name)
		falseIP.writeback()
	} else {
		c.String(http.StatusOK, "%s not exist!", name)
	}
}
func getFalseIP(name string) (string, bool) {
	falseIPmu.RLock()
	defer falseIPmu.RUnlock()
	if ip, ok := falseIP.db[name]; ok {
		return ip, true
	}
	return "", false
}
func getTrueIP(c *gin.Context) {
	src := c.Param("src")
	dst := c.Param("dst")
	srcIP, ok := getFalseIP(src)
	if !ok {
		c.String(http.StatusOK, "%s not exist!", src)
		return
	}
	dstIP, ok := getFalseIP(dst)
	if !ok {
		c.String(http.StatusOK, "%s not exist!", dst)
		return
	}
	trueIPmu.RLock()
	rulemu.RLock()
	defer rulemu.RUnlock()
	defer trueIPmu.RUnlock()
	seg1 := iptoseg(srcIP)
	seg2 := iptoseg(dstIP)
	if seg1 == seg2 {
		c.String(http.StatusOK, trueIP.db[dst])
		return
	}
	rule1 := seg1 + "@" + seg2
	if _, ok := rule.db[rule1]; ok {
		c.String(http.StatusOK, trueIP.db[dst])
	} else {
		c.String(http.StatusOK, "255.255.255.255")
	}
}

func setRule(c *gin.Context) {
	rulemu.Lock()
	defer rulemu.Unlock()
	rule1 := c.Param("seg1") + "@" + c.Param("seg2")
	rule2 := c.Param("seg2") + "@" + c.Param("seg1")
	rule.db[rule1] = "ok"
	rule.db[rule2] = "ok"
	if ok := rule.writeback(); !ok {
		c.String(http.StatusOK, "set rule failed!")
		return
	}
	c.String(http.StatusOK, "add success!")
}
func delRule(c *gin.Context) {
	rulemu.Lock()
	defer rulemu.Unlock()
	seg1 := c.Param("seg1")
	seg2 := c.Param("seg2")
	rule1 := seg1 + "@" + seg2
	rule2 := seg2 + "@" + seg1
	if _, ok := rule.db[rule1]; ok {
		delete(rule.db, rule1)
		delete(rule.db, rule2)
		c.String(http.StatusOK, "del %s success!", rule1)
		rule.writeback()
	} else {
		//fmt.Println(rule.db, rule1)
		c.String(http.StatusOK, "%s not exist!", rule1)
	}
}
func main() {

	if ok := initial(); !ok {
		return
	}

	router := gin.Default()

	router.GET("/addTrueIp/:name/:ip", addTrueIP)

	router.GET("/delTrueIp/:name", delTrueIP)

	router.GET("/addFalseIp/:name/:ip", addFalseIP)

	router.GET("delFasleIp/:name", delFalseIP)

	router.GET("/getTrueIp/:src/:dst", getTrueIP)

	router.GET("/setRule/:seg1/:seg2", setRule)

	router.GET("/delRule/:seg1/:seg2", delRule)
	fmt.Println(":" + *port)
	router.Run(":" + *port)

}
