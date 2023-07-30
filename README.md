# Http client package to interract with [cat api](https://docs.thecatapi.com/)

### Implemented functions:

```
1. Get list of cat images
2. Get single cat image
3. Upload new cat image
4. Delete cat image
```

### Setting up

Create `.env` file by copying `.env.example` and replace `CAT_API_KEY` with valid key.\
If you don't have your personal key, you can signup [here](https://thecatapi.com/)\
You can also ignore this step, only requesting of cat images will be available

### Example usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/Windmill787-golang/api-client/cat"
)

func main() {
	// Requesting 5 cat images
	catImages, err := cat.GetCatImages(5)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range catImages {
		fmt.Println(c.Info())
	}

	// Requesting single cat image by id
	catImage, err := cat.GetCatImage("MalgUNr6D")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(catImage.Info())

	// Uploading new cat image
	// You need to set valid CAT_API_KEY in .env file
	uploadedCatImage, err := cat.UploadCatImage("/path/to/file.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uploadedCatImage)

	// Deleting cat image
	// You can delete only your personal cat images
	// You need to set valid CAT_API_KEY in .env file
	err = cat.DeleteCatImage("MalgUNr6D")
	if err != nil {
		log.Fatal(err)
	}
}
```
