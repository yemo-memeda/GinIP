package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetRealIP(c *gin.Context) (string, error) {
	var ip string
	var err error
	// X-Real-IP 获取真实ip
	ip = strings.TrimSpace(c.Request.Header.Get("X-Real-IP"))
	if ip != "" {
		return ip, nil
	}
	// X-Forwarded-For 获取真实ip
	ip = c.Request.Header.Get("X-Forwarded-For")
	if index := strings.IndexByte(ip, ','); index >= 0 {
		ip = ip[0:index]
	}
	if ip != "" {
		return ip, nil
	}
	// RemoteAddr 获取真实ip
	//ip = c.Request.RemoteAddr
	ip = c.ClientIP()
	if ip != "" {
		return ip, nil
	}
	return ip, err
}

func main() {
	r := gin.Default()                // 使用Default创建路由
	r.GET("/", func(c *gin.Context) { // 这个c就是上下文,这个路由接收GET请求
		ip, _ := GetRealIP(c)
		//ip := strings.TrimSpace(c.Request.Header.Get("X-Real-IP"))
		//if ip == "" {
		//	ip = c.Request.Header.Get("X-Forward-For")
		//}
		//if ip != "" {
		//	ip = c.Request.RemoteAddr
		//}
		fmt.Println(ip)
		c.String(http.StatusOK, ip) // 返回状态码和数据
	})
	_ = r.Run(":8000") //监听端口默认为8080
}
