package models

import (
	"github.com/PamelaMarinho/go-schedule/pkg/config"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name    string `json: "name"`
	Urgency string `json: "urgency"`
	Status  string `json: "status"`
}

var db *gorm.DB

func init() {
	config.Connection()
	db := config.GetDB()
	db.AutoMigrate(db)
}

func (t *Task) Create() *Task {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetTask() []Task {
	var task []Task
	db.Find(&task)
	return task
}

func GetTaskById(ID int64) Task {
	var task Task
	db.Where("ID=?", ID).Find(task)
	return task
}

func Delete(ID int64) Task {
	var task Task
	db.Where("ID=?", ID).Delete(task)
	return task
}

// func Update(ID int64) Task{
// 	var task Task
//  	db.Where("ID=?", ID).Find(task)

// }
