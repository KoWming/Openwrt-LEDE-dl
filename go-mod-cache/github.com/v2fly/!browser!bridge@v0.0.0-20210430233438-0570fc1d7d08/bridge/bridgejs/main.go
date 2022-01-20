package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/v2fly/BrowserBridge/bridge"
	"log"
	"net/url"
)

func main() {
	fmt.Println("V3")
	go func() {
		link := js.Global.Get("window").Get("location").Get("href").String()
		urlw, err := url.Parse(link)
		if err != nil {
			log.Println(err)
		}
		host := urlw.Host
		log.Println(fmt.Sprintf("ws://%v/link", host))
		bridge.Bridge(&bridge.Settings{DialAddr: fmt.Sprintf("ws://%v/link", host)})
		js.Global.Get("location").Get("reload").Invoke()
	}()
}

//export GOPHERJS_GOROOT="$(go1.12.16 env GOROOT)"
