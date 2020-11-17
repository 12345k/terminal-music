package main

import  (
        "os"
        "path"
        "fmt"
        "net/http"
        "log"
        "io/ioutil"
	)

func main(){
	query_url,query_name:=get_query()
	MakeRequest(query_url,query_name)
	fmt.Println(query_url)
}


func get_query() (string,string){
	var  search_name string
	fmt.Print("Enter search name: ")
	fmt.Scanf("%s", &search_name)
	search_url:="https://www.youtube.com/results?search_query="+search_name
	return search_url,search_name
}

func MakeRequest(url ,name string) {
        resp, err := http.Get(url)
        dir := "url_result"
        os.Mkdir(dir, 0777)
        if err != nil {
                log.Fatalln(err)
        }
	body, err := ioutil.ReadAll(resp.Body)
        data := []byte(string(body))
        if err != nil {
                log.Fatalln(err)
        }
        fileName := path.Join(dir, name+".txt")
        errURL := ioutil.WriteFile(fileName, data, 0)
        if errURL != nil {
		log.Fatalln(errURL)
        }
}
