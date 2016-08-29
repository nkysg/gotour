package main

import (
//"golang.org/x/tour/wc"
"strings"
)

func WordCount(s string) map[string]int {
    smap := make(map[string]int)

     as := strings.Split(s, " ")
     for _, value := range as {
	if _, ok := smap[value]; ok {
		smap[value] += 1
	 } else {
	    smap[value] = 1
	}
      }

	return smap
}

func main() {
 //wc.Test(WordCount)
}
