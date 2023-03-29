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
	"strconv"
	"sync"
	"time"
)

type Photos []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
type Image struct {
	filePath string
	img      []byte
}

func main() {
	defer func() {
		fmt.Println("Main program exit successfully")
	}()
	log.SetFlags(log.Ltime)

	dir := "myDonwloadImage" + time.Now().Format("15_04_05")
	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModeDir)
	}

	photos := Photos{}
	err := getJson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	//fmt.Println(photos)
	fmt.Println(len(photos))
	//fmt.Println(photos[0:3])

	//	dir := "myDonwloadImage" + time.Now().Format("15_04_05")

	chImg := make(chan Image, len(photos))
	counter := sync.WaitGroup{}

	for _, v := range photos[1:100] {
		v := v
		counter.Add(1)
		go func() {
			defer counter.Done()
			if v.ID > 50 {
				v.ThumbnailURL = "http://abc.jpg"
			}
			img, err := donwloadImage(v.ThumbnailURL)
			if err != nil {
				//log.Fatal(err)
				log.Println(err)
				return
			}
			//fmt.Println(img)
			format, err := decodeImage(img)
			//fmt.Println("Format img : ", format)
			if err != nil {
				log.Fatal(err)
			}
			//log.Panicf("Downloaded : %v\n",v.ID)
			//fileName := "abc" + "." + format
			//absoluteFileName := filepath.Join(dir, fmt.Sprintf("%d.%s", v.ID, format))
			absoluteFileName := strconv.Itoa(v.ID) + "." + format

			/*err = saveImage(filepath.Join("myDonwloadImage", fileName), img)
			if err != nil {
				log.Println(err)
			}*/
			chImg <- Image{filePath: absoluteFileName, img: img}

		}()

	}
	go func() {
		counter.Wait()
		close(chImg)
	}()

	for v := range chImg {
		err := saveImage(v.filePath, v.img)
		if err != nil {
			log.Println(err)
		}
	}
	/*for range photos {
		v := <-chImg
		saveImage(v.filePath, v.img)
		if err != nil {
			log.Println(err)
		}
	}*/

}
func saveImage(fileName string, img []byte) error {

	f, err := os.Create(fileName)
	if err != nil {
		//log.Fatal(err)
		return fmt.Errorf("saveImage - cannot create file : %v", err)
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewReader(img))
	if err != nil {
		//log.Fatal(err)
		return fmt.Errorf("saveImage - cannot save file : %v", err)
	}
	log.Printf("\tSaved : %v\n", fileName)
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
func getJson(url string, structType interface{}) error {
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
