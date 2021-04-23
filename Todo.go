package main

import (
	"html/template"
	"net/http"
	"strconv"
)

//структура задачи
type Task struct {
	Id     int
	Text   string
	Status bool
}

// карта всех задач
var (
	tasks = make(map[int]Task)
)

//функция добавления новой задачи
func AddTask(str string) {
	if str != "" {
		tasks[len(tasks)] = Task{Id: len(tasks), Text: str, Status: false}
	}
}

//функция смены статуса на противоположный
func NewStatusTask(t Task) Task {
	if t.Status == true {
		t.Status = false
	} else {
		t.Status = true
	}
	return t
}

//функция редактирования задачи
func EditTask(t Task, str string) Task {
	t.Text = str
	return t
}

//функция удаления задачи
func DeleteTask(k int) {
	t := tasks[k]
	for i := k; i < len(tasks)-1; i++ {
		t.Status = tasks[i+1].Status
		t.Text = tasks[i+1].Text
		tasks[k] = t
	}
	delete(tasks, len(tasks)-1)
}

//функция обработки запросов главной страницы
func Hendler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("Home_page.html")
	//обработка кнопки добавления
	if r.FormValue("tsk") != "" {
		AddTask(r.FormValue("tsk"))
	}
	if len(tasks) > 1 {
		//обработка кнопки нового статуса задачи
		btn1 := r.FormValue("btn1")
		id, _ := strconv.Atoi(btn1)
		if btn1 != "" {
			tasks[id] = NewStatusTask(tasks[id])
		}
		//обработка кнопки удаления задачи
		btn2 := r.FormValue("btn2")
		id2, _ := strconv.Atoi(btn2)
		if btn2 != "" {
			DeleteTask(id2)
		}
		//обработка кнопки редактирования задачи
		btn3 := r.FormValue("btn3")
		id3, _ := strconv.Atoi(btn3)
		if btn3 != "" {
			tasks[id3] = EditTask(tasks[id3], r.FormValue(strconv.Itoa(id3)))
		}
	}

	tmpl.Execute(w, tasks)
}

func main() {
	AddTask("0")
	http.HandleFunc("/", Hendler)
	http.ListenAndServe(":8080", nil)
}
