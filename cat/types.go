package cat

import "fmt"

type CatImage struct {
	ID     string `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   string `json:"size"`
}

func (c *CatImage) Info() string {
	return fmt.Sprintf("ID: %s, Url: %s [%dx%d]", c.ID, c.Url, c.Width, c.Height)
}
