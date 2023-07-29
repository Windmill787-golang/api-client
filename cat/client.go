package cat

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	api_url = "https://api.thecatapi.com/v1/images/search"
	client  = *http.DefaultClient
)

func GetCatImages() ([]CatImage, error) {
	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", os.Getenv("API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var catImages []CatImage
	if err = json.Unmarshal(body, &catImages); err != nil {
		return nil, err
	}

	return catImages, nil
}
