package main

import (
	"flag"
	"strings"
	"time"
)

type inputArgs struct {
	Prefectures      []string
	Cities           []string
	MinimumIntensity int
	Date             string
	Quiet            bool
}

// parse command line arguments
func ParseArgs() inputArgs {
	currentTime := time.Now()
	prefectures := flag.String("pref", "", "Which prefectures to query, (comma separated: Tokyo, Hokkaido, Aomori, etc)")
	cities := flag.String("city", "", "Which cities to query (leave blank for every city in prefecture, otherwise comma separated: Noboribetsu, Chiyoda, Bunkyo, etc))")
	minimumIntensity := flag.Int("mini", 2, "Minimum intensity of earthquakes to query")
	date := flag.String("d", currentTime.Format("20060102"), "Query the earthquake data for a given day (default is today YYYYMMDD)")
	quiet := flag.Bool("q", false, "Quiet mode for integrating into other scripts/polybar")

	flag.Parse()

	prefs := splitOnComma(*prefectures)
	cits := splitOnComma(*cities)

	args := inputArgs{prefs, cits, *minimumIntensity, *date, *quiet}

	return args
}

func splitOnComma(csv string) []string {

	splitValues := strings.Split(csv, ",")

	for idx, _ := range splitValues {
		splitValues[idx] = strings.TrimSpace(splitValues[idx])
	}
	return splitValues

}
