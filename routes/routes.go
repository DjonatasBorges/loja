package routes

import (
	"net/http"

	"loja/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/produtos", controllers.Index)
	http.HandleFunc("/produtos/new", controllers.New)
	http.HandleFunc("/produtos/insert", controllers.Insert)
	http.HandleFunc("/produtos/delete", controllers.Delete)
	http.HandleFunc("/produtos/edit", controllers.Edit)
	http.HandleFunc("/produtos/update", controllers.Update)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/produtos", http.StatusFound)
	})
}
