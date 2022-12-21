package Web

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"unsafe"
)

//字节转string黑魔法 性能比[]byte()好很多
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//日志中间件
func LogMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[ logs ]  %v %v %v %v \n", time.Now().Format("2006-01-02 15:04:05"), r.Host, r.RequestURI, r.RemoteAddr)
}

//跨域中间件
func CrossOriginMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //http://localhost:6062
}

//后端业务
//base64编码与解码
func Base64(w http.ResponseWriter, r *http.Request) {

	if strings.ToUpper(r.Method) != "POST" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	//根据QueryString确定编码或解码
	choice := r.URL.Query().Get("choice")
	if choice == "" {
		w.WriteHeader(400)
		io.WriteString(w, "QueryString read error")
		return
	}
	//绑定请求体中的数据
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "body read error")
		return
	}
	defer r.Body.Close()

	//请求体的数据为空
	if len(buf) == 0 {
		w.WriteHeader(400)
		io.WriteString(w, "body empty")
		return
	}

	switch choice {
	case "encode":
		w.WriteHeader(200)
		io.WriteString(w, base64.StdEncoding.EncodeToString(buf))
		break
	case "decode":
		respBuf, err := base64.StdEncoding.DecodeString(Bytes2String(buf))
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, "DecodeString error")
			break
		}
		//测试
		// fmt.Println(string(respBuf))
		w.Write(respBuf)
		break
	}

}

//文件编码成base64字符串
func FilesToBase64(w http.ResponseWriter, r *http.Request) {

	if strings.ToUpper(r.Method) != "POST" {
		w.WriteHeader(400)
		io.WriteString(w, "请求方法错误")
		return
	}

	//读取前端传来的文件
	file, _, err := r.FormFile("b64_upload")
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()

	//进行base64编码
	fileBuf, err4 := ioutil.ReadAll(file)
	if err4 != nil {
		w.WriteHeader(500)
		io.WriteString(w, err4.Error())
		return
	}

	//返回响应
	io.WriteString(w, base64.StdEncoding.EncodeToString(fileBuf))

}

//前端业务
const HtmlFilePath = "./Page/index.html"

func Page(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(400)
		io.WriteString(w, "bad request!")
		return
	}
	http.ServeFile(w, r, HtmlFilePath)
}

//集成的HandlerFunc切片，按顺序执行
func CallFunc(order ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, o := range order {
			o(w, r) //依次执行自定义的HandlerFunc
		}
	}
}
