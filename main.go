package main
 
import (
    "log"
    "net/http"
	"html/template"
)
var index = template.Must(template.ParseFiles("index.html"))
var first = template.Must(template.ParseFiles("first.html"))
func home(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

func firstimage(w http.ResponseWriter, r *http.Request) {
	first.Execute(w, nil)
}
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
	mux.HandleFunc("/first/", firstimage)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}