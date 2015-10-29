package web

import (
	"fmt"
	"net/http"
)

func managecenterCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		fmt.Fprintf(w, "Welcome to visit managecenter callback method")
	} else {
		// 接受数据
		// partnerId := r.Form["PartnerId"]
		// ...

		fmt.Fprintf(w, "Success")
	}
}
