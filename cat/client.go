package cat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	_ "github.com/joho/godotenv/autoload"
)

type uploadResponse struct {
	CatImage
	Pending          int    `json:"pending"`
	Approved         int    `json:"approved"`
	OriginalFilename string `json:"original_filename"`
}

var (
	api_url = "https://api.thecatapi.com/v1/images"
	client  = &http.Client{
		Timeout: 10 * time.Second,
	}
)

func executeRequest(method, url string) (*http.Response, error) {
	fmt.Println(url)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", os.Getenv("CAT_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetCatImages(count int) ([]CatImage, error) {
	url := fmt.Sprintf("%s/search?limit=%d", api_url, count)

	resp, err := executeRequest("GET", url)
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

func GetMyCatImages() ([]CatImage, error) {
	url := fmt.Sprintf("%s/", api_url)

	resp, err := executeRequest("GET", url)
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

func GetCatImage(id string) (*CatImage, error) {
	url := fmt.Sprintf("%s/%s", api_url, id)

	resp, err := executeRequest("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var catImage CatImage
	if err = json.Unmarshal(body, &catImage); err != nil {
		return nil, err
	}

	return &catImage, nil
}

func UploadCatImage(filename string) (string, error) {
	u := fmt.Sprintf("%s/upload", api_url)

	client := resty.New()

	resp, err := client.R().
		SetHeader("x-api-key", os.Getenv("CAT_API_KEY")).
		SetFile("file", filename).
		Post(u)

	if err != nil {
		return "", err
	}

	var r uploadResponse
	if err = json.Unmarshal(resp.Body(), &r); err != nil {
		return "", err
	}

	return r.ID, nil
}
