package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PamelaMarinho/go-schedule/pkg/models"
	"github.com/PamelaMarinho/go-schedule/pkg/utils"
	"github.com/gorilla/mux"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetTask()
	res, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("error while parsing")
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	valuesMap := mux.Vars(r)
	idstr := valuesMap["ID"]
	id, err := strconv.ParseInt(idstr, 0, 0)
	task, _ := models.GetTaskById(id)
	res, err := json.Marshal(task)
	if err != nil {
		fmt.Println("error while parsing")
	}
	w.Header().Set("Content - Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	valuesMap := mux.Vars(r)
	idstr := valuesMap["ID"]
	id, err := strconv.ParseInt(idstr, 0, 0)
	task, _ := models.GetTaskById(id)
	if err != nil {
		panic(err)
	}
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var taskUpdated *models.Task
	utils.ParseBody(r, taskUpdated)
	valuesMap := mux.Vars(r)
	idstr := valuesMap["ID"]
	id, err := strconv.ParseInt(idstr, 0, 0)
	if err != nil {
		panic(err)
	}
	task, db := models.GetTaskById(id)
	if task.Name != "" {
		task.Name = taskUpdated.Name
	}
	if task.Status != "" {
		task.Status = taskUpdated.Status
	}
	if task.Urgency != "" {
		task.Urgency = taskUpdated.Urgency
	}
	db.Save(taskUpdated)
	res, _ := json.Marshal(taskUpdated)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
