package task

import (
	"fmt"
	"healthcheck/helper/csv"
	"healthcheck/helper/line"
	"healthcheck/helper/util"
	"time"

	"github.com/thoas/go-funk"
)

func CheckWebsite(fileName string) {
	start := time.Now()
	fmt.Println("**********************************************************")
	fmt.Println("")
	raw, err := csv.ReadFileCSV(fmt.Sprintf("public/testfile/%s", fileName))
	if err != nil {
		fmt.Println("Fail to read file")
		return
	}

	r := funk.Map(raw, func(x []string) string {
		return x[0]
	})

	websiteList := r.([]string)
	res := util.CheckWebsite(websiteList)

	duration := time.Since(start)
	fmt.Println(fmt.Sprintf("Total times to finished checking website: %s", util.FmtDuration(duration)))

	res.TotalTime = duration.Nanoseconds()
	status, err := line.SendReport(res)

	fmt.Println("")
	if err != nil || status != 200 {
		fmt.Println("Fail to send the report via Healtcheck Report API")
		return
	}
	fmt.Println("Send the report via Healcheck Report API success")
	fmt.Println("")
	fmt.Println("**********************************************************")
}
