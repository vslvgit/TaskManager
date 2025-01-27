package handlers

import (
	"TaskManager/stracts"
	"encoding/json"
	"io"
	"net/http"
)

func PostTask(res http.ResponseWriter, req *http.Request) {

	jsonTask, err := io.ReadAll(req.Body)

	if err != nil {

		http.Error(res, "Bad Request", http.StatusBadRequest)
	}
	var task stracts.Task
	json.Unmarshal(jsonTask, task)

}
