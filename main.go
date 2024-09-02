/*

 Auhor Gaurav Sablok
 Universitat Potsdam
 Date 2024-9-2
 handler for all the functions

*/

package main

import (
	 "net/http"
	 "fmt"
	 "html/template"
	 "os"
	 "log"
)

func main () {

      http.HandleFunc("/", nil)
			http.HandleFunc("/prepare", nil)
			http.HandleFunc("/expediction", nil)
}
