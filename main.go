package main
import(
	"net/http"
	
	
)
func main(){
	fileserver := http.FileServer(http.Dir("./template"))
    http.Handle("/",fileserver)
	http.ListenAndServe(":8080",nil)
}