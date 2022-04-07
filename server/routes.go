package server

import (
	"fmt"
	"net/http"
)

func InitRoute() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Listcountries(w, r)

		case http.MethodPost:
			AddCountries(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "metodo no permitido")
			return
		}
	})

}
