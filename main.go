/*

 Auhor Gaurav Sablok
 Universitat Potsdam
 Date 2024-9-2
 handler for all the functions

*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main () {



	 // a struct for storing the information from the seagrass expedition or experiments
   type prepare struct {
		 temp int
		 benthic string
		 colletor string
		 depth string
		 latitutde string
		 longitude string
		 Organization string
		 site string
		 date string
		 time time.Time
	 }

	 // func to store and process all the experimental layounts
	 prepareexp := func (w http.http.ResponseWriter, r *http.Request) {
	      message := r.PostForm("prepareexp")
				templ := template.Must(template.ParseGlob("../index.html"))
				templ.Execute(w, ,&prepare)

		 }

	// function to store and process the sequencing runs for the genome

	sequencing : = func (w http.ResponseWriter, r *http.Request) {
      message := r.PostForm("sequencing")
			templ := template.Must(template.ParseGlob("../sequencing.html"))
			templ.Execute(w, &sequencing)
	}

	    http.HandleFunc("/", nil)
			http.HandleFunc("/preparesheet",prepare)
			http.HandleFunc("/sequencing", sequencing)




	}
