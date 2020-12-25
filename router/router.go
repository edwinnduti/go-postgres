package router

import(
	"github.com/edwinnduti/go-postgres/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/newuser",middlewares.CreateUser).Methods("POST","OPTIONS")
	router.HandleFunc("/api/user/{id}",middlewares.GetUser).Methods("GET","OPTIONS")
	router.HandleFunc("/api/users",middlewares.GetAllUser).Methods("GET","OPTIONS")
	router.HandleFunc("/api/user/{id}",middlewares.UpdateUser).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/user/{id}",middlewares.DeleteUser).Methods("DELETE","OPTIONS")

	return router
}
