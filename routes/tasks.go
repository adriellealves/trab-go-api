package routes

import (
	"io/ioutil"
	"trab-Go-API/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

var tasks []Task

func ShowTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		conn := database.SetConnection()
		ImpDB, err := conn.Query("select * from task")

		if err != nil {
			fmt.Println("Não encontrada", err)
		}
		for ImpDB.Next() {
			var task Task
			err = ImpDB.Scan(&task.ID, &task.Description)
			tasks = append(tasks, task)
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(tasks)
	}
}

func NewTask(w http.ResponseWriter, r *http.Request){
	if r.Method=="Post"{
		w.WriteHeader(http.StatusCreated)
		conn:=database.SetConnection()
		defer conn.Close()
		var cadnew Task
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, "Bad Request")
		}
		json.Unmarshal(body, &cadnew)
		action, err := conn.Prepare("insert into task (description) VALUES (?)")
		action.Exec(cadnew.Description)
		encoder := json.NewEncoder(w)
		encoder.Encode(cadnew)

	}

}
func GetTaskID(w http.ResponseWriter, r *http.Request){

		if r.Method == "GET" {
			conn := database.SetConnection()
			defer conn.Close()
			var task Task
			vars := mux.Vars(r)
			id := vars["id"]
			ImpDB := conn.QueryRow("SELECT id, description FROM task WHERE id=" + id)

			ImpDB.Scan(&task.ID, &task.Description)
			encoder := json.NewEncoder(w)
			encoder.Encode(task)
		}
}


func UpTask(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			w.WriteHeader(http.StatusCreated)
			conn := database.SetConnection()
			defer conn.Close()

			var uptask Task
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprint(w, "Bad Request")
			}
			vars := mux.Vars(r)
			id := vars["id"]

			json.Unmarshal(body, &uptask)
			action, err := conn.Prepare("UPDATE task SET description = ? WHERE id = " + id)
			action.Exec(uptask.Description)

			encoder := json.NewEncoder(w)
			encoder.Encode(uptask)
		}
}

func DelTask(w http.ResponseWriter, r *http.Request) {

		if r.Method == "DELETE" {
			w.WriteHeader(http.StatusOK)
			conn := database.SetConnection()
			defer conn.Close()

			var idtask Task
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprint(w, "Bad Request")
			}
			vars := mux.Vars(r)
			id := vars["id"]

			json.Unmarshal(body, &idtask)
			action, err := conn.Prepare("delete from task where id = ?")
			action.Exec(id)
			fmt.Fprint(w, "A tarefa foi deletada com sucesso")

		}
}