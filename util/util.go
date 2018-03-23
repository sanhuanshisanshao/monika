package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Split(sep string, str string) []string {
	return strings.SplitN(str, sep, 2)
}

func ConvertTime(t string) (time.Time, error) {
	b := strings.Contains(t, "月")

	switch b {
	case true:
		//03月18日 22:08
		t = strings.Replace(t, "月", "-", -1)
		t = strings.Replace(t, "日", "", -1)
		//03-18 22:08
		t = strconv.Itoa(time.Now().Year()) + "-" + t + ":00"
		//2018-3-18 22:08:00
		break
	default:
		//2015-08-07 10:04:52
	}

	date := strings.Split(t, " ")
	if len(date) != 2 {
		return time.Now(), fmt.Errorf("split %v format error: sezi %v is not equal 2", date, len(date))
	}

	head := strings.Split(date[0], "-")
	if len(head) != 3 {
		return time.Now(), fmt.Errorf("split %v format error: size %v is not equal 3", date[0], len(date))
	}

	tail := strings.Split(date[1], ":")
	if len(tail) != 3 {
		return time.Now(), fmt.Errorf("split %v format error: size %v is not equal 3", date[1], len(date))
	}

	year, _ := strconv.Atoi(head[0])
	mon, _ := strconv.Atoi(head[1])
	day, _ := strconv.Atoi(head[2])

	hour, _ := strconv.Atoi(tail[0])
	min, _ := strconv.Atoi(tail[1])
	sec, _ := strconv.Atoi(tail[2])

	result := time.Date(year, time.Month(mon), day, hour, min, sec, 0, time.Local)

	return result, nil
}
