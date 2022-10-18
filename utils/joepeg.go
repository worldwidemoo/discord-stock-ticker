package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	JoePegURL = "https://api.joepegs.dev/v2/collections/slug/%s"
)

type JoePegCollection struct {
	Stats struct {
		TotalVolume           float64 `json:"volume"`
		FloorPrice            float64 `json:"floor"`
	} `json:"stats"`
}

func GetJoePegData(collection string) (OpenSeaCollection, error) {
	var result JoePegCollection

	reqUrl := fmt.Sprintf(JoePegURL, collection)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Set("X-Joepegs-Api-Key", "ubBvIkshbMe2J0l9Xtz7YiMsRBtiim5yL7Wd")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(results, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
