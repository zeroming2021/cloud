package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, string(r.Header))
	fmt.Fprintf(w, "<h1>This is a test</h1>")
	fmt.Println(r.Header)
	for k, v := range r.Header {
		//fmt.Println(k)
		//fmt.Println(v)
		for _, info := range v {
			//fmt.Println(info)
			w.Header().Set(k, info)
		}
	}

	os.Setenv("VERSION", "v0.0.0.0")
	version := os.Getenv("VERSION")
	fmt.Println(version)

	ipAddr := ClientIP(r)
	fmt.Println("IP: ", ipAddr, ", Http response: 200")
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
}

func main() {
	http.HandleFunc("/test", indexHandler)
	http.HandleFunc("/healthz", healthCheck)
	http.ListenAndServe(":8000", nil)
}
