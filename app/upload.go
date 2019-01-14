package app

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, _ := ioutil.ReadAll(r.Body)

	metadata, err := IsValid(body)
	if err != nil {
		log.Println("request invalid: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := Save(metadata); err != nil {
		log.Println("request invalid: " + err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Added: " + metadata.Title))
}
