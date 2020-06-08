package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type controller struct {
	app     []string `json:"String"`
	version []string `json:"Numberofcharacters"`
}

var c = controller{}

func (c *controller) addapp(n string, v string) {
	c.app = append(c.app, n)
	c.version = append(c.version, v)
}

func (c controller) checkver(s string) string {
	for index := range c.app {
		if s == c.app[index] {
			return c.version[index]
		}
	}
	return ""
}

func getPort() string {
	p := os.Getenv("PORT")
	fmt.Println(p)
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func ver(w http.ResponseWriter, r *http.Request) {
	response := c.checkver(mux.Vars(r)["appname"])
	fmt.Fprintf(w, response)
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version Control API")
}

func main() {
	port := getPort()
	c.addapp("test", "1")
	fmt.Println("API has started.")
	fmt.Println("Running on port " + port)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", start)
	router.HandleFunc("/ver/{appname}", ver).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
