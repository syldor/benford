package main 

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func main() {
	benford := make(map[string]int)
	dat, err := ioutil.ReadFile("report.csv")
	check(err)
	n := len(dat)
	text := string(dat[:n])
	records := strings.Split(text, ";")

	for _, value := range records {
		value = strings.Replace(value, ",", "", -1)
		_, e := strconv.ParseFloat(value, 64)
		if e == nil {
			numberStr := string([]rune(value)[0])
			if _, ok := benford[numberStr]; ok {
			    benford[numberStr]++
			} else {
				benford[numberStr] = 0
			}
		}
	}
	fmt.Println(benford)
}