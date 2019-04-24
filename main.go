package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/cvhariharan/Utils/utils"
	"github.com/joho/godotenv"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"encoding/json"
)

var Session *r.Session

func getfeed(w http.ResponseWriter, r *http.Request) {
	var jsonString string
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	username := r.Form.Get("username")
	// fmt.Println(username)
	result := utils.GetFeed(username, Session)
	jsonRes, _ := json.Marshal(result)
	jwt := utils.GenerateJWT(username, Session)
	jsonString = `{ "result": ` + string(jsonRes) + `, "token": "` + jwt + "\"}"
	w.Write([]byte(jsonString))
}

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	endpoints := os.Getenv("DBURL")
	url := strings.Split(endpoints, ",")
	dbpass := os.Getenv("DBPASS")
	s, err := r.Connect(r.ConnectOpts{
		Addresses: url,
		Password: dbpass,
	})
	if err != nil {
		log.Fatalln(err)
	}
	Session = s
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/feed/main", utils.AuthMiddleware(getfeed, Session))
	log.Fatal(http.ListenAndServe(port, nil))
}
