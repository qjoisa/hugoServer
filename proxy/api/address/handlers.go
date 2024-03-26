package address

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

var ErrDadataApiKeysNotAvailable error

func init() {
	err := godotenv.Load("dadata.env")
	if err != nil {
		fmt.Println("env failed to load, api dadata.ru won't work")
		ErrDadataApiKeysNotAvailable = errors.New("dadata api unavailable")
	}
}

func AddressRoutes(r chi.Router) {
	if ErrDadataApiKeysNotAvailable != nil {
		fmt.Println(ErrDadataApiKeysNotAvailable)
		return
	}
	r.Post("/geocode", GeocodeHandle)
	r.Post("/search", SearchHandle)
}

func GeocodeHandle(w http.ResponseWriter, r *http.Request) {
	gc := GeocodeRequest{}

	err := json.NewDecoder(r.Body).Decode(&gc)
	if err != nil {
		log.Printf("error: %v\nBody must be: '{\"lat\": \"anylatitude\", \"lng\":\"anuLongitude\"}'", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Body must be: '{\"lat\": \"anylatitude\", \"lng\":\"anuLongitude\"}'"))
		return
	}
	defer r.Body.Close()

	data, err := gc.GetGeocodeData()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	address := ResponseAddress{}

	w.WriteHeader(http.StatusOK)
	address.TransformData(data)
	err = json.NewEncoder(w).Encode(address)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SearchHandle(w http.ResponseWriter, r *http.Request) {
	var address ResponseAddress
	sq := SearchRequest{}
	err := json.NewDecoder(r.Body).Decode(&sq)
	if err != nil {
		fmt.Printf("error: %v\nBody must be: '{\"query\": \"any address\"}'", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Body must be: '{\"query\": \"any address\"}'"))
		return
	}
	defer r.Body.Close()
	data, err := sq.getSearchData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	address.TransformData(data)
	err = json.NewEncoder(w).Encode(address)
	if err != nil {
		log.Println(err)
	}
}
