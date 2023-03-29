package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Photos []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func main() {
	photos := Photos{}
	err := getjson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	//fmt.Println(photos)
	fmt.Println(len(photos))
	fmt.Println(photos[0:3])

}
func getjson(url string, structType interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	switch v := structType.(type) {
	case *Photos:
		fmt.Println("in photos")
		decoder := json.NewDecoder(res.Body)
		//photos := Photos{}
		photos := structType.(*Photos)
		decoder.Decode(&photos)
		return nil
	default:
		return fmt.Errorf("getJson : not support type %v ", v)
	}
}
