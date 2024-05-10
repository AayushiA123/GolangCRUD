package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			// Logic
			taskname := r.FormValue("taskname")
			deadline := r.FormValue("deadline")

			log.Println(taskname, deadline)
			models.TaskAddView{}.Add(taskname, deadline)
		}
	}
	Helper.RenderTemplate(w, r, "addTask", nil)
}
