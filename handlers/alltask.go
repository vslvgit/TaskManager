package handlers

import (
	"encoding/json"
	"net/http"

	"TaskManager/db"
)

func GetAllTask(res http.ResponseWriter, req http.Request) {

	tasks, err := db.GetAllTasks()

	if err != nil {

		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}

	jsonTask, err := json.Marshal(tasks)

	if err != nil {

		http.Error(res, "JSON Error", http.StatusConflict)

	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonTask)

}
