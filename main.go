package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type BaseReq struct {
	Client_version string `json:"client_version"`
	Method         string `json:"method"`
	Params         string `json:"params"`
}
type UserLoginReq struct {
	Appstore_version string `json:"appstore_version"`
	Email            string `json:"email"`
	Launcher_version string `json:"launcher_version"`
	Model            string `json:"model"`
	Rom_version      string `json:"rom_version"`
	Sn               string `json:"sn"`
	Swdid            string `json:"swdid"`
}

var ip string
var port string

func Decrypt(dec string) (string, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(dec)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte("1191ADF18489D8DA"))
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, []byte("5E9B755A8B674394"))
	plaintext := make([]byte, len(decodedCiphertext))
	mode.CryptBlocks(plaintext, decodedCiphertext)
	padding := plaintext[len(plaintext)-1]
	plaintext = plaintext[:len(plaintext)-int(padding)]
	return string(plaintext), nil
}
func HandleUserLogin(DecodedPara string, ctx *gin.Context) {
	log.Println("Start handle user login")
	var dec UserLoginReq
	json.Unmarshal([]byte(DecodedPara), &dec)
	log.Println("swdid", dec.Swdid)
	log.Println("email", dec.Email)
	ctx.String(200, `{"code":0,"type":"a","data":{"id":123123,"email":"123","name":"123","groupinfo":[],"schoolinfo":{"name":"123","school_id":"123"}}}`)
}
func HandleCommand(ctx *gin.Context) {
	log.Println("Start Handle Command")
	ctx.String(200, `{"code":0,"data":[{"command":"command_release_control","active":"1","type":1}]}`)
}
func main() {
	gin.SetMode(gin.DebugMode)
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
			fmt.Println(err)
		}
		fmt.Println("client version:", req.Client_version)
		fmt.Println("method:", req.Method)
		fmt.Println("Params:", req.Params)
		decode, _ := Decrypt(req.Params)
		fmt.Println("Decode:", decode)
		if req.Method == "com.linspirer.user.login" {
			HandleUserLogin(decode, ctx)
		}
		if req.Method == "com.linspirer.device.getcommand" {
			HandleCommand(ctx)
		}
	})
	r.Run(ip + ":" + port)
}