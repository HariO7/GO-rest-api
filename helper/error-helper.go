package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PanicError(err error, msg string) {
	if err != nil {
		fmt.Println("Panic Error: --------------")
		fmt.Println(err)
		panic(msg)
	}
}

func ContextErrors(err error, context *gin.Context, statusCode int, msg string) bool {
	if err != nil {
		fmt.Println("---------------------")
		fmt.Println(err)
		context.JSON(statusCode, gin.H{"message": msg})
		return true
	}
	return false
}
