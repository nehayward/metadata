package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gobwas/glob"
	"github.com/julienschmidt/httprouter"
	"github.com/nehayward/metadata/models"
)

func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	keys, ok := r.URL.Query()["title"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'title' is missing")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("search query '?title='search'"))
		return
	}

	title := keys[0]
	apps := lookup(title)

	if len(apps) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No apps with title: '" + title + "' exists"))
		return
	}

	encoded, err := json.Marshal(apps)
	if err != nil {
		log.Println("encode failed: " + err.Error())
	}

	w.WriteHeader(http.StatusFound)
	w.Write(encoded)
}

func lookup(title string) []models.App {
	g := glob.MustCompile(title)

	var foundApps []models.App
	for _, a := range apps {
		if g.Match(a.Title) {
			foundApps = append(foundApps, a)
		}
	}

	return foundApps
}
