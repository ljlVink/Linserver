package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

type BaseReq struct {
	Client_version string `json:"client_version"`
	Method         string `json:"method"`
	Params         string `json:"params"`
}

var ip string
var port string

func HandleUserLogin(ctx *gin.Context) {
	log.Println("Start handle user login")
	ctx.String(200, `{"code":0,"type":"a","data":{"id":123123,"email":"123","name":"123","groupinfo":[],"schoolinfo":{"name":"123","school_id":"123"}}}`)
}
func HandleCommand(ctx *gin.Context) {
	log.Println("Start Handle Command")
	ctx.String(200, `{"code":0,"data":[{"command":"command_release_control","active":"1","type":1}]}`)
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	log.Println("Linserver ReleaseControl Demo 5.04")
	flag.StringVar(&ip, "ip", "", "cust ip address")
	flag.StringVar(&port, "port", "1243", "cust port number")
	flag.Parse()
	if ip != "" {
		log.Println("ip addr:" + ip)
	}
	log.Println("port:" + port)
	r := gin.Default()
	r.POST("/public-interface.php", func(ctx *gin.Context) {
		var req BaseReq
		err := ctx.BindJSON(&req)
		if err != nil {
			log.Println(err)
		}
		log.Println("method:", req.Method)
		if req.Method == "com.linspirer.user.login" {
			HandleUserLogin(ctx)
		}
		if req.Method == "com.linspirer.device.getcommand" {
			HandleCommand(ctx)
		}
	})
	r.Run(ip + ":" + port)
}
