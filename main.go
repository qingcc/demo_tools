package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/demo_tools/util"
)

func main(){
    fmt.Printf("hellow world")
    fmt.Println("string:", util.GetSjCode(16))
    gin.SetMode(gin.DebugMode)
    

}
