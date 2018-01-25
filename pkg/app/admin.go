package app

import (
	"log"
	"net/http"

	"github.com/Nithiszz/gonews/pkg/view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminLogin(w, nil)
}
func adminList(w http.ResponseWriter, r *http.Request) {
	view.AdminList(w, nil)
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		detail := r.FormValue("detail")
		log.Println(title)
		log.Println(detail)
		http.Redirect(w, r, "/admin/create", http.StatusSeeOther)
		return
	}
	view.AdminCreate(w, nil)
}

func adminEdit(w http.ResponseWriter, r *http.Request) {
	view.AdminEdit(w, nil)
}
