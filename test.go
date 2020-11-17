package main

import (
	//"encoding/json"
	"os"
	"path"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"regexp"
)

func main() {
	//MakeRequest()
	ReadFile()
}

func ReadFile(){
	dir:="url_result"
	fileName := path.Join(dir, "test.txt")
	dat,_ := ioutil.ReadFile(fileName)
	stringData:=string(dat)
    	r, _ := regexp.Compile("/watch([a-z?=0-9A-Z]+)")
	url_list:=r.FindAllString(stringData, -1)
	fmt.Println(url_list[0])
	fmt.Println(len(url_list))
}

func MakeRequest() {
	resp, err := http.Get("https://www.youtube.com/results?search_query=test")
	if err != nil {
	        log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	data := []byte(string(body))
	if err != nil {
	        log.Fatalln(err)
	}

	log.Println("---------------")
	
	dir := "test_dir"

    	os.Mkdir(dir, 0777)

	fileName := path.Join(dir, "file.txt")
	
	err_test := ioutil.WriteFile(fileName, data, 0)
	if err_test != nil {
        log.Fatal(err_test)
    	}



}
