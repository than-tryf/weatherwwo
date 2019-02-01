package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sources := []string{
		//Insert "http://links here"
	}

	wg.Add(len(sources))

	for _, source := range sources {
		go func(source string) {
			resp, _ := http.Get(source)
			//weatherObject := entities.WeatherPayload{}
/*
			json.NewDecoder(resp.Body).Decode(&weatherObject)

			for _, item := range weatherObject.Data.Weather{

				fmt.Println(item)

			}
*/
			fmt.Println(resp.Body)

			if resp.StatusCode == http.StatusOK {
				bodyBytes, _ := ioutil.ReadAll(resp.Body)
				bodyString := string(bodyBytes)
				fmt.Println(bodyString)
			}



			wg.Done()
		}(source)
	}

	wg.Wait()
}
