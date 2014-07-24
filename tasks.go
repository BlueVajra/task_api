package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"os"
)

const fileName = "tasks.json"

type Tasks []Task
type Task struct {
	Name        string
	Description string
}

func (tasks Tasks) toBytes() ([]byte, error) {
	bytes, err := json.Marshal(tasks)
	return bytes, err
}

func (tasks Tasks) save() error {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error saving: %s", err))
		return err
	} else {
		return ioutil.WriteFile(fileName, bytes, 0600)
	}

}

// try refactorings so we are passing the data and a writer to the method.

func (tasks *Tasks) load() error {
//	var (
//		err      error
//		bytes    []byte
//	)
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error loading: %s", err))
		return err
	}

	err = json.Unmarshal(bytes, tasks)
	if err != nil {
		fmt.Println(fmt.Sprintf("Bytes: %s", bytes))
		return err
	}

	return nil
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	log.Println("Testing: ", r.Body)
	if r.Method == "POST" {
//		email := req.FormValue("email")
//		// ...
//		fmt.Println("Sending confirmation email to:", email)
		taskJSON, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//	}

		var task Task
		{
			err := json.Unmarshal(taskJSON, &task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		log.Println("New Task: ", task)


		var tasks Tasks
		{
			err := tasks.load()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}


		tasks = append(tasks, task)

		log.Println("All Tasks: ", tasks)

		if err := tasks.save(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

//		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(taskJSON)
	}else{

		var tasks Tasks
		err := tasks.load()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

//		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		bytes, err := tasks.toBytes()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)

	}



}

//func newTasksHandler(w http.ResponseWriter, r *http.Request) {
//
//	// read json post parameters
////	var taskJSON []byte
//
////	{
//		//		var err error
//		taskJSON, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
////	}
//
//	var task Task
//	{
//		err := json.Unmarshal(taskJSON, &task)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
//	log.Println("New Task: ", task)
//
//
//	var tasks Tasks
//	{
//		err := tasks.load()
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//	}
//
//
//	tasks = append(tasks, task)
//
//	log.Println("All Tasks: ", tasks)
//
//	if err := tasks.save(); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(200)
//	w.Write(taskJSON)
//}

func main() {
	// make sure tasks.json exists.
	_, err := ioutil.ReadFile(fileName)
	if os.IsNotExist(err) {
		ioutil.WriteFile(fileName, []byte("[]"), 0600)
	}
	log.Println("STARTING")
	http.HandleFunc("/api/tasks", tasksHandler)
//	http.HandleFunc("/api/task/new", newTasksHandler)
	http.ListenAndServe(":8080", nil)
}
