package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Name string
	Lang string
}

var Countries []Country = []Country{
	{
		Name: "Colombia",
		Lang: "Spanish",
	},
	{
		Name: "USA",
		Lang: "Ingles",
	},
}

func Index(w http.ResponseWriter, r *http.Request) {

}

func AddCountries(w http.ResponseWriter, r *http.Request) {
	var newcountry Country
	
	err := json.NewDecoder(r.Body).Decode(&newcountry)
	if err != nil {
		panic(err)
	}

	Countries = append(Countries, newcountry)
	
	fmt.Fprint(w, "Fue agregado con exirto")


}
func Listcountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Countries)
	if err != nil {
		panic(err)
	}

	return
}
