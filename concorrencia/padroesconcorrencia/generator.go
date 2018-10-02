//Padrões de concorrência - Generator

//Para mais informações sobre padrões de concorrência, visitar a documentação 
//Google I/O 2012 - Go Concurrency Patterns

package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"net/http"
)

func tituloURL(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls{
		go func(url string){
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return ch
}


func main(){
	t1 := tituloURL("https://www.gitlab.com", "https://www.linkedin.com")
	t2 := tituloURL("https://www.instagram.com", "https://www.youtube.com")
	fmt.Println("Prmeiros títulos:", <-t1, "|", <-t2)	
	fmt.Println("Segundos títulos:", <-t1, "|", <-t2)
	
}