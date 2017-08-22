// MediaCloud-API project main.go

/*
MediaCloud-API document
*/
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

var users = map[string]string{"zhangli": "lehoo2016", "yinyingxia": "lehoo2016"}

func main() {
	http.HandleFunc("/mdwiki", mdwiki) //设置访问的路由
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe: " + err.Error())
	}
}

func mdwiki(w http.ResponseWriter, r *http.Request) {
	if !authenticate(r.Header.Get("Authorization")) {
		w.Header().Set("WWW-Authenticate", "Basic realm="+r.URL.String())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	} else {
		path := "./mdwiki.html"
		http.ServeFile(w, r, path)
	}
}

func authenticate(s string) bool {
	result := false
	if s != "" {
		ss := strings.Split(s, " ")
		if len(ss) == 2 {
			var coder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")
			info, err := coder.DecodeString(ss[1])
			if err == nil {
				user := strings.Split(string(info), ":")
				if len(user) == 2 {
					if users[user[0]] != "" && users[user[0]] == user[1] {
						result = true
					}
				}
			}
		}
	}

	return result
}
