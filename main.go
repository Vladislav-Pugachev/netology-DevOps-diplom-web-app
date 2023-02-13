package main
 
import (
   "log"
   "net/http"
	"html/template"
   aws "web-app/utils"
)

func home(w http.ResponseWriter, r *http.Request) {
http.ServeFile(w, r, "index.html")
}
type ViewData struct{
 
    Path string
    Link string
}
func renderImage(w http.ResponseWriter, r *http.Request) {
    data:=new(ViewData)
    link:=new(aws.Link)
    Path:= r.URL.String()
    switch Path {
    case "/first":
        data.Path="Первая картинка"
    case "/second":
        data.Path="Вторая картинка"
    }
    Link:=link.ReturnLink(Path)
    data.Link=Link
    template.Must(template.ParseFiles("template/template.html")).Execute(w, data)
}
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
	mux.HandleFunc("/first", renderImage)
    mux.HandleFunc("/second", renderImage)
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
