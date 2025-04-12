package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DateAdd string `json:"dateAdd"`
	DateTo  string `json:"dateTo"`
	Content string `json:"content"`
	IsDone  bool   `json:"isDone"`
	RefUser int    `json:"refUser"`
}

var taskSetup = map[string]string{
	"payload": "id,date_add,date_to,title,content,is_done,ref_user",
	"table":   "task",
}

func CreateTask(wrapper *Wrapper) {
	fmt.Printf("Db value is %v", db)
	if err := wrapper.wrapData("title"); err != nil {
		wrapper.Error(err.Error())
		return
	}
	task := Task{
		Title:   wrapper.data["title"].(string),
		DateAdd: time.Now().UTC().Truncate(time.Second).String(),
		IsDone:  false,
		RefUser: wrapper.ReturnUser(),
	}
	smtp, err := db.Prepare("INSERT INTO " + taskSetup["table"] + "(title,date_add,is_done,ref_user) VALUES(?,?,?,?)")
	if err != nil {
		wrapper.Error(err.Error(), 400)
		return
	}
	defer smtp.Close()
	result, err := smtp.Exec(task.Title, task.DateAdd, task.IsDone, task.RefUser)
	if err != nil {
		wrapper.Error(err.Error(), 400)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		wrapper.Error(err.Error(), 400)
		return
	}
	task.ID = int(lastInsertID)

	wrapper.Render(map[string]any{
		"data": task,
	}, 200)
}

func GetTasks(wrapper *Wrapper) {
	rows, err := db.Query("SELECT "+taskSetup["payload"]+" FROM "+taskSetup["table"]+" WHERE ref_user=? ORDER BY date_add DESC", wrapper.ReturnUser())
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	data := []Task{}
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.DateAdd, &task.DateTo, &task.Title, &task.Content, &task.IsDone, &task.RefUser); err != nil {
			wrapper.Error(err.Error(), http.StatusBadGateway)
			return
		}
		if task.DateTo == "0000-00-00 00:00:00" {
			task.DateTo = ""
		}
		data = append(data, task)
	}
	wrapper.Render(map[string]any{
		"data": data,
	}, 200)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	wrapper := NewWrapper(w, r)
	rows, err := db.Query("SELECT "+taskSetup["payload"]+" FROM "+taskSetup["table"]+" WHERE id=? ORDER BY date_add DESC", chi.URLParam(wrapper.request, "id"))
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	task := Task{}
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.DateAdd, &task.DateTo, &task.Title, &task.Content, &task.IsDone); err != nil {
			wrapper.Error(err.Error(), http.StatusBadGateway)
			return
		}
		if task.DateTo == "0000-00-00 00:00:00" {
			task.DateTo = ""
		}
	}
	wrapper.Render(map[string]any{
		"task": task,
	}, 200)
}

func PatchTask(wrapper *Wrapper) {
	rows, err := db.Query(
		"UPDATE "+taskSetup["table"]+" SET is_done = NOT is_done WHERE id=? AND ref_user=?",
		chi.URLParam(wrapper.request, "id"), wrapper.ReturnUser(),
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()
	wrapper.Render(map[string]any{
		"message": "Update successfull",
	})
}
