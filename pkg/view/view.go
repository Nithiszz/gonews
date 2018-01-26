package view

import (
	"net/http"

	"github.com/Nithiszz/gonews/pkg/model"
)

type IndexData struct {
	List []*model.News
}

// Index renders index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// AdminLogin renders admin login
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminList renders admin List
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// AdminCreate renders admin Create
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit renders admin Edit
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminEdit, w, data)
}
