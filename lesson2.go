package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	HttpServerStart(8080)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func HttpServerStart(port int) {
	log.SetPrefix("Info:")                  //为每条日志文本前增加一个info:前缀
	log.SetFlags(log.Ldate | log.Llongfile) //可以获取当前设置的选项,Ldate：输出当地时区的日期;Llongfile：输出长文件名+行号

	http.HandleFunc("/", httpAccessFunc) //http.HandleFunc接收两个参数,一个是路由匹配的字符串，另外一个是 func(ResponseWriter, *Request) 类型的函数
	http.HandleFunc("healthz", healthzFunc)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err) //输出日志后，调用os.Exit(1)退出程序
	}
}

func healthzFunc(w http.ResponseWriter, r *http.Request) {
	HealthzCode := "200"
	w.Write([]byte(HealthzCode))
}

func httpAccessFunc(w http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			log.Printf("%s=%s", k, v[0])

			//1. request header写入response header
			w.Header().Set(k, v[0])
		}
	}

	log.Printf("\n\n\n")

	r.ParseForm() //解析所有请求数据，否则无法获取数据
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			log.Printf("%s=%s", k, v[0])
		}
	}
	log.Printf("\n\n\n")

	os.Setenv("VERSION", " JDK version 1.11.0") //设置环境值的值

	//2. 获取环境变量"VERSION"
	name := os.Getenv("VERSION")
	log.Printf("VERSION Env: ", name)

	log.Printf("\n\n\n")

	//3.获取Client IP，并且打印出来
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}

	if net.ParseIP(ip) != nil {
		fmt.Println("ip ===>>%s\n", ip)
		log.Println(ip)
	}

	fmt.Println("http Status Code ===>>%s\n", http.StatusOK)
	log.Println(http.StatusOK)

	//response响应
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Server Access,Success!"))
}
