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
	jsonEncoder.Encode(shortUrlResp{ShortUrl: shortFunc.ShortUrl(t.LongUrl, 10)})
}
