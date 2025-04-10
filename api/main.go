package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mobile-api/initializers"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Wrapper struct {
	writer  http.ResponseWriter
	request *http.Request
	data    map[string]interface{}
}

func init() {
	godotenv.Load()
	var err error
	db, err = initializers.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	GET("/", Index)
	POST("/task", CreateTask)
	http.ListenAndServe(":"+port, nil)
}

func (wrapper Wrapper) Render(data map[string]interface{}, status ...int) {
	wrapper.writer.Header().Set("Content-type", "application/json")
	code := http.StatusOK
	if len(status) > 0 {
		code = status[0]
	}
	wrapper.writer.WriteHeader(code)
	dataJSON, err := json.Marshal(data)
	if err != nil {
		wrapper.Error(err.Error())
		return
	}
	wrapper.writer.Write(dataJSON)
}

func (wrapper Wrapper) Error(error string, code ...int) {
	wrapper.writer.Header().Set("Content-type", "application/json")
	statusCode := 400
	if len(code) > 0 {
		statusCode = code[0]
	}
	dataJSON, _ := json.Marshal(map[string]string{
		"error": error,
	})
	wrapper.writer.WriteHeader(statusCode)
	wrapper.writer.Write(dataJSON)
}

func GET(path string, handler func(w *Wrapper)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		wrapper := NewWrapper(w, r)
		if r.Method != "GET" {
			wrapper.Error("Not authorized", http.StatusMethodNotAllowed)
			return
		}
		handler(wrapper)
	})
}

func POST(path string, handler func(w *Wrapper)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		wrapper := NewWrapper(w, r)
		if r.Method != "POST" {
			wrapper.Error("Not authorized", http.StatusMethodNotAllowed)
			return
		}
		if err := wrapper.request.ParseMultipartForm(10 >> 20); err != nil {
			wrapper.Error(err.Error(), http.StatusBadGateway)
			return
		}
		wrapper.data = make(map[string]interface{})
		for key, values := range wrapper.request.MultipartForm.Value {
			if len(values) <= 0 {
				continue
			}
			wrapper.data[key] = values[0]
		}
		if len(wrapper.data) <= 0 {
			wrapper.Error("No data received", http.StatusBadGateway)
			return
		}
		handler(wrapper)
	})
}

func NewWrapper(w http.ResponseWriter, r *http.Request) *Wrapper {
	return &Wrapper{
		writer:  w,
		request: r,
	}
}

func Index(wrapper *Wrapper) {
	wrapper.Render(map[string]interface{}{
		"message": "Hello world",
	})
}

func (wrapper *Wrapper) wrapData(data string) error {
	if wrapper.data[data] == nil || wrapper.data[data] == "" {
		return fmt.Errorf("you have to set a value for %v", data)
	}
	return nil
}

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
