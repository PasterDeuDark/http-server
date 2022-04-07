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

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &newcountry)
	if err != nil {
		panic(err)
	}

	Countries = append(Countries, newcountry)

}
func Listcountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Countries)

	return
}
