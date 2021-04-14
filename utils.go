package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *categories) dataFromURL(url string) (err error) {
	pageContent, err := http.Get(url)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(pageContent.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, c)
	return
}

func (a *artistsArr) dataFromURL(url string) (err error) {
	pageContent, err := http.Get(url)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(pageContent.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, a)
	return
}

func (d *datesLocations) dataFromURL(url string) (err error) {
	pageContent, err := http.Get(url)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(pageContent.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, d)
	return
}
