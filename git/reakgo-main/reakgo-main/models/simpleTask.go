package models

import (
	"reakgo/utility"
)

type TaskAddView struct {
	Id       int32
	TaskName string
	DeadLine string
}

func (form TaskAddView) Add(taskname string, deadline string) {
	utility.Db.MustExec("INSERT INTO addTask (taskname, deadline) VALUES (?, ?)", taskname, deadline)
}
