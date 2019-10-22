package route

import (
	"../app/controller"
	"net/http"
)

// Blogs handler route
func Blogs(w http.ResponseWriter, r *http.Request) {
	controller.BlogAll(w, r)
}
