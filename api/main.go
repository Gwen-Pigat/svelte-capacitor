package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mobile-api/initializers"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Wrapper struct {
	writer  http.ResponseWriter
	request *http.Request
}

func init() {
	godotenv.Load()
	db, err := initializers.ConnectDB()
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
	statusCode := 404
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
