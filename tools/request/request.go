package request

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Get(remoteUrl string) (result []byte, err error) {
	client := &http.Client{}
	uri, err := url.Parse(remoteUrl)
	if err != nil {
		log.Println(err)
		return
	}
	reqest, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		log.Println(err)
		return
	}
	response, err := client.Do(reqest)
	if err != nil {
		log.Println(err)
		return
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		result = body
	}
	defer response.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
	return
}
