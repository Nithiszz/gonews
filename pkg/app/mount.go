package app

import (
	"net/http"
)

// Mount mounts handler to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	mux.HandleFunc("/new/", newsView)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/edit", adminEdit)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
