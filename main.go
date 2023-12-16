package main

import (
	"fmt"
	"test/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载模板目录下模板文件
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/static/file1", "static/数.txt")

	r.GET("/", rtHTML)
	r.POST("/dingxiang", rRedirect)
	r.GET("/json", rtJSON)
	//导入文件写的路由
	router.InitApi(r)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/*以下是一些示例*/
// 直接响应json
func rtJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "lkz",
		"age":  23,
	})
}

// 响应html
func rtHTML(c *gin.Context) {
	//获取get请求的参数
	// fmt.Println(c.Query("user")) //使用变量名获取参数值
	c.HTML(200, "index.html", gin.H{"time": "菜花"})
}

// 重定向
func rRedirect(c *gin.Context) {
	//获取post请求的参数值
	message := c.PostForm("user") //方法1  建议使用该方法
	// forms, err := c.MultipartForm() //方法2， 接收所有的form参数 注意：使用这个方法要在form表单属性中设置enctype

	fmt.Println("post信息为：", message)
	c.HTML(200, "test.html", gin.H{"message": message})
}
