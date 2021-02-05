package gotst

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func TestDownloadURL(taskid int, url string) {
	/*
		userFile := "ttt.ts"
		fout, err := os.Create(userFile)
		defer fout.Close()
	*/

	bgntime := time.Now()

	resp, err := http.Get(url)

	// resptime := time.Now()

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	// fmt.Println(resp.ContentLength)

	buf := make([]byte, 1024)
	times := 0
	for {
		size, _ := resp.Body.Read(buf)
		//读到文件结尾
		if size == 0 {
			break
		} else {
			// fout.Write(buf[:size])
		}
		times++
	}

	endtime := time.Now()

	fmt.Printf("$$ task id: %v, finish, time used %v\n", taskid, endtime.Sub(bgntime).Milliseconds())

}

func getArgs() {
	fmt.Println("命令行参数数量:", len(os.Args))
	for k, v := range os.Args {
		fmt.Printf("args[%v]=[%v]\n", k, v)
	}
}

func GetArgs3() (string, int, int) {

	var filename string
	var taskgap int
	var multicnt int

	// StringVar用指定的名称、控制台参数项目、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量
	flag.StringVar(&filename, "f", "", "输入文件名,必填")
	flag.IntVar(&taskgap, "g", 100, "启动任务的间隔时间，单位ms,默认100")
	flag.IntVar(&multicnt, "m", 1, "URL重复次数,默认为1")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	// 打印
	fmt.Printf("filename=%v taskgap=%v multicnt=%v\n\n", filename, taskgap, multicnt)

	return filename, taskgap, multicnt
}
