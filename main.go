package main

import (
	"fmt"
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
		if strings.HasPrefix(entry.Ctt, args.Date) && entry.Ser != "" { // filter for events that are on the selected date & have associated details page

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
