package main

import (
	"database/sql"
	"log"
	"net/http"
)

func Notify(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	wiino := r.Form.Get("wiino")
	_, err = db.Exec("INSERT IGNORE INTO `wus` (`wiino`, `uid` ) VALUES (?, ?)", wiino, uid)
	if err != nil {
		w.Header().Add("X-RESULT", "110")
		w.Header().Add("X-WUS-Host", "wus.wii.rc24.xyz")
		log.Fatal(err)
		return
	}

	w.Header().Add("X-RESULT", "001")
	w.Header().Add("X-WUS-Host", "wus.wii.rc24.xyz")
}
