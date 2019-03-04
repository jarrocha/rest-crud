package controllers

import (
	"net/http"

	"gitlab.com/codelittinc/golang-interview-project-jaime/rest-crud/view"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bks, err := HC.AllHires()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	if view.TPL != nil {
		view.TPL.ExecuteTemplate(w, "hires.gohtml", bks)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := HC.OneHire(r)
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	if view.TPL != nil {
		view.TPL.ExecuteTemplate(w, "show.gohtml", bk)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if view.TPL != nil {
		view.TPL.ExecuteTemplate(w, "create.gohtml", nil)
	}
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := HC.PutHire(r)
	if err != nil {
		http.Error(w, http.StatusText(406)+err.Error(), http.StatusNotAcceptable)
		return
	}

	if view.TPL != nil {
		view.TPL.ExecuteTemplate(w, "created.gohtml", bk)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := HC.OneHire(r)
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	if view.TPL != nil {
		view.TPL.ExecuteTemplate(w, "update.gohtml", bk)
	}
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := HC.UpdateHire(r)
	if err != nil {
		http.Error(w, http.StatusText(406)+err.Error(), http.StatusBadRequest)
		return
	}

	if view.TPL != nil {
		view.TPL.ExecuteTemplate(w, "updated.gohtml", bk)
	}
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := HC.DeleteHire(r)
	if err != nil {
		http.Error(w, http.StatusText(400)+err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
