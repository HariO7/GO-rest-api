package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicError(err error, msg string) {
	if err != nil {
		fmt.Println("Panic Error: --------------")
		fmt.Println(err)
		panic(msg)
	}
}

func ContextErrors(err error, context *gin.Context, msg string) {
	if err != nil {
		fmt.Println("---------------------")
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": msg})
		return
	}
}
