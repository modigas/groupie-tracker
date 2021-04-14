package main

import (
	"net/http"
	"strconv"
)

func errHandler(w http.ResponseWriter, r *http.Request, er errType) {
	w.WriteHeader(er.Status)
	tpl.ExecuteTemplate(w, "error", er)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errHandler(w, r, errType{Status: http.StatusNotFound, Message: "Page Not Found"})
		return
	}
	tpl.ExecuteTemplate(w, "layout", artists)
}

func artistInfo(w http.ResponseWriter, r *http.Request) {
	artistID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(artistID)
	if err != nil {
		errHandler(w, r, errType{Status: http.StatusBadRequest, Message: "Bad Request"})
		return
	}
	if id > len(artists) {
		errHandler(w, r, errType{Status: http.StatusBadRequest, Message: "Bad Request"})
		return
	}
	var dl datesLocations
	if err = dl.dataFromURL(artists[id-1].Relations); err != nil {
		errHandler(w, r, errType{Status: http.StatusInternalServerError, Message: "Internal Error"})
	}

	d := struct {
		ArtistInfo     artist
		DatesLocations datesLocations
	}{
		ArtistInfo:     artists[id-1],
		DatesLocations: dl,
	}
	tpl.ExecuteTemplate(w, "artist.html", d)
}
