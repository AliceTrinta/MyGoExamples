package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Google I/O 2012 - Go Concurrency Patterns
//Concurrency pattern "Generators"

// <- chan - channel only-reading
func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls{
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			aReturn := r.FindStringSubmatch(string(html))

			if cap(aReturn) == 0 {
				c <- "Erro ao ler página " + url
				return
			}
			c <- aReturn[1]
		}(url)
	}
	return c
}

func main() {
	//When the function generate something to you, and in the main there's no
	// worry with channels ans goroutines.
	t1 := title("https://www.google.com", "https://www.youtube.com")
	t2 := title("https://www.google.com.br/drive", "https://www.amazon.com")
	fmt.Println("First:", <-t1, "|", <-t2)
	fmt.Println("Second:", <-t1, "|", <-t2)
}
