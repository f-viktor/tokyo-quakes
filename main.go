package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	// parse command-line arguments
	args := ParseArgs()

	//get earthquake list
	list := ListJson{}
	err := getEarthquakeData(listUrl, &list)
	if err != nil {
		//		log.Println(err)
	}

	// get the code mapping of prefectures
	prefCodes := getFilterKeys(prefectureUrl, args.Prefectures, "Prefectures", args.Quiet)

	cityCodes := getFilterKeys(cityUrl, args.Cities, "Cities", args.Quiet)

	highestIntensity := "0"
	highestEvent := ""

	for _, entry := range list {
		if strings.HasPrefix(entry.Ctt, args.Date) && entry.Ser != "" { // filter for events that are on the selected date

			highestIntensity, highestEvent = checkPrefectureList(entry, prefCodes, highestIntensity, highestEvent)

			highestIntensity, highestEvent = checkCityList(entry, cityCodes, highestIntensity, highestEvent)

		}
	}

	if args.Quiet {
		fmt.Println(highestIntensity + " " + queryUrl + highestEvent)
	} else {
		fmt.Println("Highest intensity event on " + args.Date + ": " + queryUrl + highestEvent + " \nMaximum Earthquake intensity: " + highestIntensity)
	}
}

func pickHigherNumericalValue(a string, b string) (string, bool) {

	ai, err := strconv.Atoi(a)
	if err != nil {
		log.Println(err)
	}

	bi, err := strconv.Atoi(b)
	if err != nil {
		log.Println(err)
	}

	if ai > bi {
		return string(a), false
	} else {
		return string(b), true
	}

}

// filter events from the selected prefecture (if any)
func checkPrefectureList(entry ListJsonEntry, prefCodes []string, highestIntensity string, highestEvent string) (string, string) {
	changed := false

	for _, listprefecture := range entry.Int {
		for _, filterPrefecture := range prefCodes {
			if listprefecture.Code == filterPrefecture {
				highestIntensity, changed = pickHigherNumericalValue(highestIntensity, listprefecture.Maxi)
				if changed {
					highestEvent = entry.Ctt
				}
			}
		}
	}
	return highestIntensity, highestEvent

}

// filter events from the selected cities (if any)
func checkCityList(entry ListJsonEntry, citycodes []string, highestIntensity string, highestEvent string) (string, string) {
	changed := false
	for _, listprefecture := range entry.Int {
		for _, listCity := range listprefecture.City {
			for _, filterCity := range citycodes {
				if listCity.Code == filterCity { // filter events from the selected prefecture (if any)
					highestIntensity, changed = pickHigherNumericalValue(highestIntensity, listCity.Maxi)
					if changed {
						highestEvent = entry.Ctt
					}
				}
			}
		}
	}
	return highestIntensity, highestEvent
}
