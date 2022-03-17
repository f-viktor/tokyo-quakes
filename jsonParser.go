package main

import (
	"log"
	"strconv"
	"strings"
)

func pickHigherNumericalValue(a string, b string) (string, bool) {

	ai := convertIntensityToFloat(a)
	bi := convertIntensityToFloat(b)

	if ai == -1 || bi == -1 {
		return "0", false
	}

	if ai > bi {
		return string(a), false
	} else {
		return string(b), true
	}

}

func convertIntensityToFloat(a string) float64 {
	if strings.Contains(a, "-") {
		tmp := strings.Trim(a, "-")
		f, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			log.Println(err)
		}
		f = f - 0.5
		return f
	}

	if strings.Contains(a, "+") {
		tmp := strings.Trim(a, "+")
		f, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			log.Println(err)
		}
		f = f + 0.5

		return f
	}

	f, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Println(err)
	}
	return f

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
