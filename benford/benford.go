package benford 

import (
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func Calculate_from_string(text string) map[string]int {
	records := strings.Split(text, ";")
	benford := make(map[string]int)

	for _, value := range records {
		value = strings.Replace(value, ",", "", -1)
		_, e := strconv.ParseFloat(value, 64)
		if e == nil {
			numberStr := string([]rune(value)[0])
			if _, ok := benford[numberStr]; ok {
			    benford[numberStr]++
			} else {
				benford[numberStr] = 1
			}
		}
	}
	return benford

}

func Calculate_from_file() map[string]int {
	dat, err := ioutil.ReadFile("../report.csv")
	check(err)
	n := len(dat)
	text := string(dat[:n])
	return Calculate_from_string(text)
}
