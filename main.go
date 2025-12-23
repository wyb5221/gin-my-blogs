package main

import (
	"fmt"
	"gin-my-blogs/blog/interfaces/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("---开始---")

	// 加密密码
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	// fmt.Println("--hashedPassword:", string(hashedPassword))
	// service.Test1()

	// interfaces.Test1()
	// base.CreateTable()
	r := gin.Default()
	router.RegisterRoutes(r)
	router.Test1RegisterRoutes(r)
	//创建分组
	// gp := r.Group("/blogs")
	// interfaces.Test1RegisterRoutes(gp)

	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
