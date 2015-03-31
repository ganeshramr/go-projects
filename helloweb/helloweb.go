package main

import (
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	resp, _ := http.Get("https://maps.googleapis.com/maps/api/geocode/json?latlng=" + values.Get("lat") + "," + values.Get("log"))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.Write(body)

}

func main() {

	http.HandleFunc("/sayHello", handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "geoHello.html")
	})
	http.ListenAndServe(":4444", nil)

}
