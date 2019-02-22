package app

import "github.com/gorilla/mux"

func (hawk *App) LoadRoutes () {

	hawk.router = mux.NewRouter()
	hawk.router.HandleFunc ("/api/addUser", hawk.addUser).Methods ("POST")
	hawk.router.HandleFunc ("/api/login", hawk.login).Methods ("POST")
	hawk.router.HandleFunc ("/api/logout", hawk.logout).Methods ("POST")
	hawk.router.HandleFunc ("/api/forgotPassword", hawk.forgotPassword).Methods ("POST")
	hawk.router.HandleFunc ("/api/resetPassword", hawk.resetPassword).Methods ("POST")
}
