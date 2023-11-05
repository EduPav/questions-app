package routes

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//Shows index.html
	http.ServeFile(w, r, "./static/index.html")
}
