package main

import (
	"Base64/Config"
	"Base64/Web"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const backend_addr = ":6060"
const frontend_addr = ":6062"

func main() {

	backend := http.NewServeMux()
	frontend := http.NewServeMux()

	//后端
	backend.Handle("/base64", Web.CallFunc(Web.LogMiddleware, Web.CrossOriginMiddleware, Web.Base64))
	backend.Handle("/file2base64", Web.CallFunc(Web.LogMiddleware, Web.CrossOriginMiddleware, Web.FilesToBase64))
	//前端
	frontend.Handle("/", Web.CallFunc(Web.LogMiddleware, Web.Page))

	fmt.Println("【 Sweet Base64 】")

	var wg sync.WaitGroup

	wg.Add(2)

	//监听后端
	go func() {
		log.Printf(" [Backend] Listening %s \n\n", Config.Cfg.BackendAddr)
		err := http.ListenAndServe(Config.Cfg.BackendAddr, backend)
		if err != nil {
			log.Fatalf("error listening %v", err)
			return
		}
		wg.Done()
	}()

	//监听前端
	go func() {
		log.Printf(" [Frontend] Listening %s \n\n", Config.Cfg.FrontendAddr)
		err := http.ListenAndServe(Config.Cfg.FrontendAddr, frontend)
		if err != nil {
			log.Fatalf("error listening %v", err)
			return
		}
		wg.Done()
	}()

	wg.Wait()
}
