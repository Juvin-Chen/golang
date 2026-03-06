// 在底层的 Socket 中，我们只知道接收字符串。但在 Web 框架中，服务端需要根据用户请求的不同 **URL 路径**（如 `/login`, `/register`）和 **方法**（如 GET, POST），执行不同的 Go 函数。这个分发请求的机制，就叫 **路由**。

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	// 判断请求方法
	if r.Method != 
}

func main(){

}