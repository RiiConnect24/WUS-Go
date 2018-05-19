package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-systemd/daemon"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

// Config structure for `config.json`.
type Config struct {
	Port     int
	Host     string
	Username string
	Password string
	DBName   string
	BindTo   string
	Debug    bool
}

var global Config
var db *sql.DB

var keys = map[string]string{
	"RVL-HCIJ": "Jb3Mp3Sg", // Wii no Ma
	"RVL-R64J": "r3tWuGcq", // Wii Music
	"RVL-RUDJ": "2C5NHqCv", // Animal Crossing Wii
	"RVL-WA4E": "XHrACw4r", // WarioWare D.I.Y.: Showcase
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if global.Debug {
			// Dump path, etc
			log.Printf("%s %s", r.Method, r.URL)
		}
		handler.ServeHTTP(w, r)
	})
}

func inquiryHandler(w http.ResponseWriter, r *http.Request) {
	Inquiry(w, r, db)
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	Notify(w, r, db)
}

func main() {
	file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&global)
	if err != nil {
		panic(err)
	}
	testDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		global.Username, global.Password, global.Host, global.Port, global.DBName))
	if err != nil {
		panic(err)
	}
	err = testDb.Ping()
	if err != nil {
		panic(err)
	}

	// If we've reached here, we're working fine.
	db = testDb

	log.Println("Running...")
	http.HandleFunc("/inquiry", inquiryHandler)
	http.HandleFunc("/notify", notifyHandler)

	// Allow systemd to run as notify
	// Thanks to https://vincent.bernat.im/en/blog/2017-systemd-golang
	// for the following things.
	daemon.SdNotify(false, "READY=1")

	// We do this to log all access to the page.
	log.Fatal(http.ListenAndServe(global.BindTo, logRequest(http.DefaultServeMux)))
}
