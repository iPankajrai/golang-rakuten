package main

import (
	"fmt"

	"github.com/JitenPalaparthi/rakutenshapes/rect"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r1, err := rect.New(10.11, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Area of rect:", r1.Area())
	}

	r2, err := rect.New(10.11, 12.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Area of rect:", r2.Area())
	gin.Default() // not doing any thing just to understand the third party pacakge related stuff

	var m gorm.Model // not doing any thing just to understand the third party pacakge related stuff
	fmt.Println(m)
}
