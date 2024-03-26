package address

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func (g *GeocodeRequest) GetGeocodeData() (GeocodeResponse, error) {
	req, err := g.prepareRequest()
	if err != nil {
		return GeocodeResponse{}, err
	}

	data, err := g.getRequestJsonData(req)

	suggestions := GeocodeResponse{}

	err = json.Unmarshal(data, &suggestions)
	if err != nil {
		fmt.Println(err)
		return GeocodeResponse{}, err
	}

	return suggestions, nil
}

func (g *GeocodeRequest) prepareData() (io.Reader, error) {
	requestBody := struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}{
		Lat: g.Lat,
		Lon: g.Lng,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonData), nil
}

func (g *GeocodeRequest) prepareRequest() (*http.Request, error) {
	body, err := g.prepareData()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "http://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", os.Getenv("DADATA_API_KEY")))

	return req, nil
}

func (g *GeocodeRequest) getRequestJsonData(req *http.Request) ([]byte, error) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return data, nil
}
