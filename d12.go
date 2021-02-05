package main

import (
	"bufio"
	"december12/gotst"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const waitexit_s = 1000

func main() {

	// getArgs2()
	// testHTTP()
	filename, taskgap_ms, multicnt := gotst.GetArgs3()
	urlindex, taskid := 0, 0

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("file pen Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		urlindex++
		fmt.Printf("$$url index %d, input line: %s\n", urlindex, string(a))

		sliceline := strings.Split(string(a), ",")

		// var downtime string = sliceline[0]
		// var rate, _ = strconv.Atoi(sliceline[1])
		var url string = sliceline[2]

		for i := 0; i < multicnt; i++ {
			taskid++
			fmt.Printf("$$ task id: %d, launch\n", taskid)

			go gotst.TestDownloadURL(taskid, url)

			time.Sleep(time.Duration(taskgap_ms * 1000000))
		}

	}

	fmt.Println("\r\n est \r\n")

	time.Sleep(waitexit_s * 1000000000)

	fmt.Println("\n exit")

}
