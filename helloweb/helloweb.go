package main

import (
	"io/ioutil"
	"net/http"
)

/*
   A handler function , observe the input params
*/
func handler(w http.ResponseWriter, r *http.Request) {
	/*Get the query params from the URL. Retruns the query params as go map
	  wait .. i did not define 'values' as map yet.
	  Go infers the type since you wanted go to do so. :=  does the trick.
	*/
	values := r.URL.Query()

	//maps has Get(key) func to read the values
	resp, _ := http.Get("https://maps.googleapis.com/maps/api/geocode/json?latlng=" + values.Get("lat") + "," + values.Get("log"))

	//defer says - Do this when you exit this func. like finally in Java
	defer resp.Body.Close()

	contentAsByteArray, _ := ioutil.ReadAll(resp.Body)
	w.Write(contentAsByteArray)

}

func main() {

	/*
	  HandleFunc defines the handler function for a specific route (URL).
	  The handler function of type Handler an Interface
	  Go functions are First-Class. YAY.
	*/
	http.HandleFunc("/sayHello", handler)

	/*
		Another approach of defining handler function
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//See how it returns the HTML. My favourite
		http.ServeFile(w, r, "geoHello.html")
	})

	/*
		Start the server and listen on the port 4444
		Setting nil as the handler here tells go to use default ServerMux
		In simple terms ServerMux is the router of requests to handlers based on url

	*/
	http.ListenAndServe(":4444", nil)

}
