package main


import (
	"log"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func myhandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hehhehehehheheheheheh!")
}

func hahaHander(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "haha twicefdsafdsafdsafdsafds")

}
func RootHandler(res http.ResponseWriter, req *http.Request) {
	file, _ := ioutil.ReadFile("templates/index.tmpl.html")
	res.Write(file)
}

func main() {

	http.HandleFunc("/", RootHandler)

	http.Handle("/public/", http.StripPrefix("/public/",
		http.FileServer(http.Dir("public"))))

	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	http.HandleFunc("/hello", myhandler)

	http.HandleFunc("/haha", hahaHander)


	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Listening...")
	http.ListenAndServe(":" + port,nil)

	}