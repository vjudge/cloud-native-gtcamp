package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main ()  {
	HTTP_PORT := "8080"
	http.HandleFunc("/", resHeader)
	http.HandleFunc("/healthz", CheckHealth)

	fmt.Printf("Server listening on %s: http://localhost:%s\n", HTTP_PORT, HTTP_PORT)
	log.Fatal(http.ListenAndServe(":" + HTTP_PORT, nil))
}

// 问题 1
// http://localhost:8080
func resHeader (res http.ResponseWriter, req *http.Request)  {
	//fmt.Println("req: ", req.Header)
	for key, val := range req.Header {
		res.Header().Set(key, val[0])
	}
	// 问题 2 : 添加环境变量
	res.Header().Set("VERSION", GetSystemEnv("VERSION"))
	res.WriteHeader(http.StatusOK)
	PrintLog(req)
	//fmt.Println("=== status: ", res.Header().Get("statusCode"))
	//fmt.Println("res: ", res.Header())
	return
}

// 问题 2: 获取环境变量
func GetSystemEnv (name string) string {
	version := os.Getenv(name)
	if  version == "" {
		version = "default version"
	}
	return version
}

// 问题 3
func PrintLog(req *http.Request)  {
	fmt.Println("HTTP PATH: ", req.URL.Path) // 请求路径
	fmt.Println("HTTP IP:", req.RemoteAddr) // 请求地址
	//fmt.Println("HTTP StatusCode: ", http.StatusOK) // 请求返回状态码
}

// 问题 4
// http://localhost:8080/healthz
func CheckHealth (res http.ResponseWriter, req *http.Request)  {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("200"))
	PrintLog(req)
	return
}
