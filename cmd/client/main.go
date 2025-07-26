package main

import (
	"context"
	"fmt"
	"io"

	"golang-gin/api"
)

func main() {
	apiClient, _ := api.NewClient("http://0.0.0.0:8080/api/v1")

	createAlbumResponse, _ := apiClient.CreateAlbum(context.Background(), api.CreateAlbumJSONRequestBody{
		Title: "test",
	})
	fmt.Println(createAlbumResponse.Status)

	getResponse, _ := apiClient.GetAlbum(context.Background(), 1)
	body, _ := io.ReadAll(getResponse.Body)
	fmt.Println(string(body))
}