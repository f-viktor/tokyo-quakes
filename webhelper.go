package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type ListJson []struct {
	Ctt string `json:"ctt"`
	Ser string `json:"ser"`
	Int []struct {
		City []struct {
			Code string `json:"code"`
			Maxi string `json:"maxi"`
		} `json:"city"`
		Code string `json:"code"`
		Maxi string `json:"maxi"`
	} `json:"int"`
	Anm   string    `json:"anm"`
	At    time.Time `json:"at"`
	EnTTL string    `json:"en_ttl"`
	TTL   string    `json:"ttl"`
	Ift   string    `json:"ift"`
	Rdt   time.Time `json:"rdt"`
	Acd   string    `json:"acd"`
	Mag   string    `json:"mag"`
	JSON  string    `json:"json"`
	Maxi  string    `json:"maxi"`
	Eid   string    `json:"eid"`
	EnAnm string    `json:"en_anm"`
	Cod   string    `json:"cod"`
}

type ListJsonEntry struct {
	Ctt string `json:"ctt"`
	Ser string `json:"ser"`
	Int []struct {
		City []struct {
			Code string `json:"code"`
			Maxi string `json:"maxi"`
		} `json:"city"`
		Code string `json:"code"`
		Maxi string `json:"maxi"`
	} `json:"int"`
	Anm   string    `json:"anm"`
	At    time.Time `json:"at"`
	EnTTL string    `json:"en_ttl"`
	TTL   string    `json:"ttl"`
	Ift   string    `json:"ift"`
	Rdt   time.Time `json:"rdt"`
	Acd   string    `json:"acd"`
	Mag   string    `json:"mag"`
	JSON  string    `json:"json"`
	Maxi  string    `json:"maxi"`
	Eid   string    `json:"eid"`
	EnAnm string    `json:"en_anm"`
	Cod   string    `json:"cod"`
}

type langTable struct {
	Japanese   string `json:"japanese"`
	English    string `json:"english"`
	ChineseZs  string `json:"chinese_zs"`
	ChineseZt  string `json:"chinese_zt"`
	Korean     string `json:"korean"`
	Portuguese string `json:"portuguese"`
	Spanish    string `json:"spanish"`
	Vietnamese string `json:"vietnamese"`
	Thai       string `json:"thai"`
	Indonesian string `json:"indonesian"`
	Tagalog    string `json:"tagalog"`
	Nepali     string `json:"nepali"`
	Khmer      string `json:"khmer"`
	Burmese    string `json:"burmese"`
	Mongolian  string `json:"mongolian"`
}

//get raw earthquake data
func getEarthquakeData(url string, target interface{}) error {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)

}

//get raw earthquake data
func getFilterKeys(url string, filterKeys []string, areaType string, quiet bool) []string {

	areaMap := map[string]langTable{}
	areaCodes := []string{}

	if len(filterKeys[0]) > 0 {

		err := getEarthquakeData(url, &areaMap)
		if err != nil {
			log.Println(err)
		}

		// get a all the area codes that contain filter sting (in English)
		if !quiet {
			fmt.Println("Searching for " + areaType + ": ")
		}
		for code, name := range areaMap {
			for _, areaFilter := range filterKeys {
				if strings.Contains(strings.ToLower(name.English), strings.ToLower(areaFilter)) {
					areaCodes = append(areaCodes, code)
					if !quiet {
						fmt.Println("\t" + name.English + " [" + string(code) + "], ")
					}
				}
			}
		}
	}

	return areaCodes

}
