package main

import (
	"fmt"
	_ "gin-my-blogs/blogs/base"
	"gin-my-blogs/blogs/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("---开始---")
	interfaces.Test1()
	// base.CreateTable()
	r := gin.Default()
	//创建分组
	gp := r.Group("/blogs")
	interfaces.Test1RegisterRoutes(gp)

	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
