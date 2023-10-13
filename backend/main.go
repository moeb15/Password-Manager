package main

import (
	"fmt"
	"log"

	"pwdmanager_api/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	var pwd string
	fmt.Scanln(&pwd)
	fmt.Println(helpers.HashPwd(pwd))
}

func Listen() {
	router := gin.Default()
	log.Fatal(router.Run(":8000"))
}
