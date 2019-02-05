package main

import (
	"encoding/json"
	"fmt"
	"github.com/mohae/struct2csv"
	"github.com/weatherwwo/entities"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 1, 0)
		return date
	}
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

/*	sources := []string{
		//Insert "http://links here"
		//"http://api.worldweatheronline.com/premium/v1/past-weather.ashx?key=3dacac42551c42d2aa2192139192101&q=Nicosia&format=json&date=2009-09-01&enddate=2009-09-30",
	}
*/

	var sources []string

	baseUrl := "http://api.worldweatheronline.com/premium/v1/past-weather.ashx"
	apiKey :=  "?key=3dacac42551c42d2aa2192139192101"
	location := "&q=Nicosia"
	format := "&format=json"

	end := time.Now()
	start := end.AddDate(-10,0,0)
	//end := start.AddDate(0, 0, 6)
	fmt.Println(start.Format("2006-01-02"), "-", end.Format("2006-01-02"))

	for rd := rangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}

		fromDate := date.AddDate(0,1,0)
		endDate := date

		url := baseUrl+apiKey+location+format+"&date="+endDate.Format("2006-01-02")+"&enddate="+fromDate.Format("2006-01-02")

		//fmt.Println(date.Format("2006-01-02"))
		//fmt.Println(url)
		sources = append(sources, url)

	}


	//wg.Add(len(sources))
	wg.Add(1)

	go func(source string){

		log.Printf("Collecting Weather data using call: %v\n", source)

		resp, _ := http.Get(source)
		weatherObject := entities.WeatherPayload{}

		json.NewDecoder(resp.Body).Decode(&weatherObject)

		//fmt.Println(weatherObject.Data.Weather)

/*		for _, item := range weatherObject.Data.Weather{

			fmt.Println(item)

		}
*/
		csvFile, _ := os.OpenFile("weather.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer csvFile.Close()
		//
		//csvstruct.ReadStruct(weatherObject.Data.Weather)

		writer := struct2csv.NewWriter(csvFile)

		_ = writer.WriteStructs(weatherObject.Data.Weather)


		wg.Done()
	}(sources[0])


	//for _, source := range sources {
	//	go func(source string) {
			//resp, _ := http.Get(source)
			//weatherObject := entities.WeatherPayload{}

			//json.NewDecoder(resp.Body).Decode(&weatherObject)

			//fmt.Println(len(weatherObject.Data.Weather))
			//fmt.Println(len(sources))
			//fmt.Println("---------------------------------")

			//for _, item := range weatherObject.Data.Weather{
			//
			//	fmt.Println(item)
			//
			//}

			//fmt.Println(resp.Body)

/*			if resp.StatusCode == http.StatusOK {
				bodyBytes, _ := ioutil.ReadAll(resp.Body)
				bodyString := string(bodyBytes)
				fmt.Println(bodyString)
			}
*/


			//wg.Done()
		//}(source)
	//}

	wg.Wait()
}
