package core

import (
	"EmoSearch/utils"
	"flag"
	"fmt"
	"sync"

	"github.com/panjf2000/ants/v2"
)

var U string
var Wg sync.WaitGroup
var URL = flag.String("url", "", "input url")
var Urllist = flag.String("file", "", "input path to urllist.txt")
var FILE = flag.String("f", "", "save to file")
var Dic = flag.String("d", "", "choose the dictionary")

type taskFunc func()

var statuscode int

func Run1(url string) {
	lines := Getline()
	for _, line := range lines {
		statuscode = Get_statuscode(*URL + line)
		if statuscode == 200 {
			if *FILE != "" {
				utils.Resulte_write(url + line)
			} else {
				fmt.Println(url + line)
			}

		}
	}
}

func Runs() {
	lines := Get_urllist()
	p, _ := ants.NewPool(10)
	defer p.Release()
	for _, line := range lines {
		Wg.Add(1)
		p.Submit(taskFuncWrapper(line, &Wg))
	}
}

func taskFuncWrapper(line string, wg *sync.WaitGroup) taskFunc {
	return func() {
		Run2(line)
		wg.Done()
	}
}

func Run2(url string) {
	lines := Getline()
	for _, line := range lines {
		statuscode = Get_statuscode(url + line)
		if statuscode == 200 {
			if *FILE != "" {
				utils.Resulte_write(url + line)
			} else {
				fmt.Println(url + line)
			}
		}
	}
}
