package main

import (
	"MyGoExamples/concurrency/html"
	"fmt"
	"time"
)

func faster(url1, url2, url3 string) string{
	c1 := html.Title(url1)
	c2 := html.Title(url2)
	c3 := html.Title(url3)

	//Control structure specific for deal with concurrency.
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "All lose!"
	//default:
	//	return "No answer!"
	}
}

func main()  {
	champion := faster(
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.amazon.com",
	)
	fmt.Println(champion)
}