package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	//fmt.Println(photos[0:3])

	dir := "myDonwloadImage"
	if _, err = os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModeDir)
	}

	//chImg := make(chan ?? ,len(photos))

	for _, v := range photos[1:] {

		img, err := donwloadImage(v.ThumbnailURL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(img)

		format, err := decodeImage(img)
		fmt.Println("Format img : ", format)
		if err != nil {
			log.Fatal(err)
		}
		//fileName := "abc" + "." + format
		fileName := fmt.Sprintf("%d.%s", v.ID, format)
		err = saveImage(filepath.Join("myDonwloadImage", fileName), img)
		if err != nil {
			log.Println(err)
		}

	}

	//for  range photos{
	//	v := <-chImg
	//	saveImage(v.filepath, v.img)
	//}
}
func saveImage(fileName string, img []byte) error {

	f, err := os.Create(fileName)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewReader(img))
	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}
func decodeImage(img []byte) (string, error) {
	_, format, err := image.Decode(bytes.NewBuffer(img))
	return format, err
}
func donwloadImage(url string) ([]byte, error) {
	errMsg := func(err error) error {
		return fmt.Errorf("downloadImage : %v", err)

	}
	res, err := http.Get(url)
	if err != nil {
		return nil, errMsg(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errMsg(err)
	}
	return body, nil
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
