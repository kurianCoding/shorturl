package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kurianCoding/shorturl/shortFunc"
)

type shortUrlReq struct {
	LongUrl string `json:"longUrl"`
}
type shortUrlResp struct {
	ShortUrl string `json:"shortUrl"`
}

const (
	BadUrl = http.StatusBadRequest
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shortUrl", shortUrl).Methods("POST")
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	srv.ListenAndServe()
	return
}
func shortUrl(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var t shortUrlReq
	defer r.Body.Close()
	err := jsonDecoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	jsonEncoder := json.NewEncoder(w)
	if good := validate(t.LongUrl); good {
		jsonEncoder.Encode(shortUrlResp{ShortUrl: "http://" + shortFunc.ShortUrl(t.LongUrl, 10) + ".me"})
	} else {
		w.WriteHeader(BadUrl)
		jsonEncoder.Encode(shortUrlResp{ShortUrl: "invalid url"})
	}
}
func validate(url string) bool {
	httpClient := &http.Client{}
	httpReq, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return false
	}
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return false
	}
	if resp.StatusCode > 399 || resp.StatusCode < 200 {
		return false
	}
	return true
}
