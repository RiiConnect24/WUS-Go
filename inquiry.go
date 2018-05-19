package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Inquiry(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Parse form.
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	uid := r.Form.Get("uid")
	pwd := r.Form.Get("pwd")
	if keys[uid] != pwd {
		w.Header().Add("X-RESULT", "101")
		w.Header().Add("X-WUS-Host", "wus.wii.rc24.xyz")
		log.Println("Unknown game - uid = ", uid, ", pwd = ", pwd)
		return
	}

	chkno := strings.Split(r.Form.Get("chkno"), ",")
	hasGame := ""

	for _, v := range chkno {
		var wiino int
		err := db.QueryRow("SELECT * FROM `wus` WHERE `wiino` =? AND `uid` =?", v, uid).Scan(&wiino, &uid)

		if err == sql.ErrNoRows {
			hasGame += "0"
		} else if err != nil {
			w.Header().Add("X-RESULT", "110")
			w.Header().Add("X-WUS-Host", "wus.wii.rc24.xyz")
			log.Fatal(err)
			return
		} else {
			hasGame += "1"
		}
	}

	w.Header().Add("X-RESULT", "001")
	w.Header().Add("X-WUS-Host", "wus.wii.rc24.xyz")

	fmt.Fprint(w, hasGame)
}
