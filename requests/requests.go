package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"main.go/creds"
)

func Req_Dogs() ([]DogImage, error) {
	Creds := creds.Init_Creds()
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	fmt.Println(Creds.Api_URL)
	request, err := http.NewRequest("GET", Creds.Api_URL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	request.Header.Set("x-api-key", Creds.Api_KEY)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var dog_images []DogImage

	err = json.Unmarshal(body, &dog_images)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	for _, image := range dog_images {
		fmt.Println("URL:", image.URL)
	}

	// Print the response body
	return dog_images, nil
}
