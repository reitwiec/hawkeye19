package app

import "github.com/gorilla/mux"

func (hawk *App) LoadRoutes () {

	hawk.router = mux.NewRouter()
	hawk.router.HandleFunc ("/api/addUser", hawk.addUser).Methods ("POST")
	hawk.router.HandleFunc ("/api/login", hawk.login).Methods ("POST")
}
