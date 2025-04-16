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

const format = "2006-01-02 15:04:05"

func CreateTask(wrapper *Wrapper) {
	fmt.Printf("Db value is %v", db)
	if err := wrapper.wrapData("title"); err != nil {
		wrapper.Error(err.Error())
		return
	}
	task := Task{
		Title:   wrapper.data["title"].(string),
		DateAdd: time.Now().UTC().Truncate(time.Second).Format(format),
		IsDone:  false,
		RefUser: wrapper.ReturnUser(),
	}
	smtp, err := db.Prepare("INSERT INTO " + taskSetup["table"] + "(title,date_add,is_done,ref_user) VALUES(?,?,?,?)")
	if err != nil {
		wrapper.Error(err.Error(), 400)
		return
	}
	defer smtp.Close()
	_, err = smtp.Exec(task.Title, task.DateAdd, task.IsDone, task.RefUser)
	if err != nil {
		wrapper.Error(err.Error(), 400)
		return
	}
	GetTasks(wrapper)
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
		} else {
			task.DateTo = wrapFormat(task.DateTo)
		}
		task.DateAdd = wrapFormat(task.DateAdd)

		data = append(data, task)
	}
	wrapper.Render(map[string]any{
		"data": data,
	}, 200)
}

func GetTask(wrapper *Wrapper) {
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
	rows, err := db.Query("SELECT "+taskSetup["payload"]+" FROM "+taskSetup["table"]+" WHERE id=? ORDER BY date_add DESC", chi.URLParam(wrapper.request, "id"))
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	task := Task{}
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.DateAdd, &task.DateTo, &task.Title, &task.Content, &task.IsDone, &task.RefUser); err != nil {
			wrapper.Error(err.Error(), http.StatusBadGateway)
			return
		}
		if task.DateTo == "0000-00-00 00:00:00" {
			task.DateTo = ""
		}
	}

	task.DateTo = time.Now().UTC().Truncate(time.Second).Format(format)
	if task.IsDone {
		task.DateTo = ""
	}
	task.IsDone = !task.IsDone
	rows, err = db.Query(
		"UPDATE "+taskSetup["table"]+" SET is_done = ?,date_to=? WHERE id=? AND ref_user=?",
		task.IsDone, task.DateTo, chi.URLParam(wrapper.request, "id"), wrapper.ReturnUser(),
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	task.DateTo = wrapFormat(task.DateTo)

	wrapper.Render(map[string]any{
		"message": "Update successfull",
		"result":  task,
	})
}

func wrapFormat(dateStr string) string {
	loc, _ := time.LoadLocation("Europe/Paris")
	parsed, _ := time.ParseInLocation(format, dateStr, loc)
	return parsed.Format(format)
}

func DeleteTask(wrapper *Wrapper) {
	rows, err := db.Exec(
		"DELETE FROM "+taskSetup["table"]+" WHERE id=? AND ref_user=?",
		chi.URLParam(wrapper.request, "id"), wrapper.ReturnUser(),
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := rows.RowsAffected(); err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}
	wrapper.Render(map[string]any{
		"message": "Delete successfull",
	})
}
