package router

import(
	"github.com/edwinnduti/go-postgres/lib"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api",lib.PostDataHandler).Methods("POST","OPTIONS")
	router.HandleFunc("/api/user/{user_id}", lib.GetUserHandler).Methods("GET","OPTIONS")
	router.HandleFunc("/api/users", lib.GetAllUsersHandler).Methods("GET","OPTIONS")
	router.HandleFunc("/api/user/{user_id}", lib.UpdateUserHandler).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/user/{id}", lib.DeleteUserHandler).Methods("DELETE","OPTIONS")

	return router
}
