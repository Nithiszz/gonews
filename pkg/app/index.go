package app

import (
	"log"
	"net/http"

	"github.com/Nithiszz/gonews/pkg/model"
	"github.com/Nithiszz/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}
	list := model.ListNews()
	log.Println(list)
	view.Index(w, &view.IndexData{
		List: list,
	})
}
