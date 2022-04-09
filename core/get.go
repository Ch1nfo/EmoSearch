package core

import (
	"EmoSearch/assets"
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
)

func Get_random_ua() string {
	ass := assets.Info
	file, _ := ass.Open("user-agents.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var UA []string
	for scanner.Scan() {
		UA = append(UA, scanner.Text())
	}
	USER_AGENTS := UA
	length := len(USER_AGENTS)
	index := rand.Intn(length)
	return USER_AGENTS[index]

}

func Getline() (lines []string) {
	ass := assets.Info
	file, err := ass.Open(*Dic)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Get_urllist() (urls []string) {
	file, err := os.Open(*Urllist)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls
}

func Get_statuscode(U string) (resCode int) {

	u, _ := url.Parse(U)
	q := u.Query()
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	res.Header.Add("User_Agent", Get_random_ua())
	if err != nil {

		fmt.Println("0")
		return
	}
	resCode = res.StatusCode
	res.Body.Close()
	if err != nil {

		fmt.Println("0")
		return
	}
	return resCode
}
