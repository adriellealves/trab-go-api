package trab_Go_API

import (
	"trab-Go-API/middlewares"
	"trab-Go-API/routes"
	"fmt"

	"net/http"
	"github.com/gorilla/mux"
)
func iniroute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
func setRoutes(router *mux.Router) {

	router.HandleFunc("/", iniroute)
	router.HandleFunc("/tasks/", routes.ShowTasks)
	router.HandleFunc("/newTask",routes.NewTask)
	router.HandleFunc("/tasks/{id}", routes.GetTaskID)
	router.HandleFunc("/uptask/{id}", routes.UpTask)
	router.HandleFunc("/deltask/{id}", routes.DelTask)

}
func main()  {
	var router *mux.Router
	router = mux.NewRouter()
	router.Use(middlewares.JsonMiddleware)
	http.Handle("/",router)

	setRoutes(router)
	err := http.ListenAndServe(":3306",router)
	if err != nil {
		fmt.Println("Error", err)
	}
}