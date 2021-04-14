package main

import (
	"html/template"
	"log"
)

var tpl *template.Template

var cats categories
var artists artistsArr

func init() {
	go initData()
	tpl = template.Must(template.ParseGlob("templates/*html"))

}

func initData() {
	err := cats.dataFromURL(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	err = artists.dataFromURL(cats.Artists)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	baseURL = "https://groupietrackers.herokuapp.com/api"
)

type categories struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type artistID struct {
	ID int `json:"id"`
}

type artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type datesLocations struct {
	ID             int `json:"id"`
	DatesLocations map[string][]string
}

type artistsArr []artist

type errType struct {
	Status  int
	Message string
}
