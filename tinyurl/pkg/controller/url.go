package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rakshitha31/urlshortnerchallenge/pkg/helper"
	"github.com/rakshitha31/urlshortnerchallenge/pkg/service"
)

type InputUrl struct {
	LongUrl string `json:"longUrl"`
}

func ShortenUrl(w http.ResponseWriter, req *http.Request) {
	var input InputUrl
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	key := helper.GenerateShortUrl(input.LongUrl)
	fmt.Println(key)
	shortUrl := "http://localhost:8080/" + key
	doc := service.AddUrl(input.LongUrl, shortUrl, key)
	jsonBytes, err := json.Marshal(doc)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

}

func RedirectUrl(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path[1:]
	longUrl, err := service.GetLongUrl(key)
	if err != nil {
		http.Error(w, "Error finding long url", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, longUrl, http.StatusMovedPermanently)
}
