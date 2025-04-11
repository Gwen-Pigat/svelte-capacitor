package main

import (
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	DateAdd time.Time `json:"dateAdd"`
	DateTo  time.Time `json:"dateTo"`
	Content string    `json:"content"`
	IsDone  bool      `json:"isDone"`
}

func CreateTask(wrapper *Wrapper) {
	fmt.Printf("Db value is %v", db)
	if err := wrapper.wrapData("title"); err != nil {
		wrapper.Error(err.Error())
		return
	}
	task := Task{
		Title:   wrapper.data["title"].(string),
		DateAdd: time.Now().UTC().Truncate(time.Second),
		IsDone:  false,
	}
	smtp, err := db.Prepare("INSERT INTO tasks(title,date_add,is_done) VALUES(?,?,?)")
	if err != nil {
		wrapper.Error(err.Error(), 400)
		return
	}
	defer smtp.Close()
	result, err := smtp.Exec(task.Title, task.DateAdd, task.IsDone)
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

	wrapper.Render(map[string]interface{}{
		"task": task,
	}, 200)
}

func GetTasks(wrapper *Wrapper) {
	rows, err := db.Query("SELECT id,date_add,date_to,title,content,is_done FROM tasks ORDER BY date_add DESC")
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	data := make(map[int]interface{})
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.DateAdd, &task.DateTo, &task.Title, &task.Content, &task.IsDone); err != nil {
			wrapper.Error(err.Error(), http.StatusBadGateway)
			return
		}
		data[task.ID] = task
	}
	wrapper.Render(map[string]interface{}{
		"tasks": data,
	}, 200)
}
