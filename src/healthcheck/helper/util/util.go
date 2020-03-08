package util

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

// WebsiteCheckStatusOutput -
type WebsiteCheckStatusOutput struct {
	TotalWebsite int   `json:"totalWebsite"`
	Success      int   `json:"success"`
	Failure      int   `json:"failure"`
	TotalTime    int64 `json:"totalTime"`
}

const (
	SIZE = 300
)

func CheckWebsite(list []string) WebsiteCheckStatusOutput {
	fmt.Println("Perform website checking...")
	total := 0
	success := 0
	failure := 0

	var chans = []chan int{}

	for i := 0; i < SIZE; i++ {
		ch := make(chan int)
		chans = append(chans, ch)
	}

	j := 0
	for i := 0; i < len(list); i++ {
		if j == SIZE {
			j = 0
		}
		go ping(list[i], chans[j])
		j++
	}

	cases := make([]reflect.SelectCase, len(chans))
	for i, ch := range chans {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	remaining := len(list)
	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)

		if !ok {
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining--
			continue
		}
		if value.Interface().(int) == 200 {
			success++
		} else {
			failure++
		}
		total++
		remaining--
	}

	fmt.Println("Done")
	fmt.Println("")
	fmt.Println(fmt.Sprintf("Checked webistes: %d", total))
	fmt.Println(fmt.Sprintf("Successful websites: %d", success))
	fmt.Println(fmt.Sprintf("Failure websites: %d", failure))

	return WebsiteCheckStatusOutput{TotalWebsite: total, Success: success, Failure: failure}
}

func ping(url string, ch chan int) {
	tr := &http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(r)
	if err != nil {
		ch <- 500
		return
	}
	ch <- resp.StatusCode
}

func FmtDuration(d time.Duration) string {
	timedif := d
	min := timedif / time.Minute
	timedif -= min
	sec := timedif / time.Second
	timedif -= sec
	milli := timedif.Milliseconds()

	return fmt.Sprintf("%d/%d/%d", milli, sec, min)
}
