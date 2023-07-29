package cat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	api_url = "https://api.thecatapi.com/v1/images"
	client  = *http.DefaultClient
)

func GetCatImages(count int) ([]CatImage, error) {
	url := fmt.Sprintf("%s/search?limit=%d", api_url, count)

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// req.Header.Set("x-api-key", os.Getenv("API_KEY"))

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

func GetCatImage(id string) (CatImage, error) {
	return CatImage{}, fmt.Errorf("method not implemented")
}

func CreateCatImage(c CatImage) error {
	return fmt.Errorf("method not implemented")
}
