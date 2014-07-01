package main

import (
    "strings"
    "fmt"
    "flag"
    "os"
    "strconv"             
    "math"
)

func usage() {
     fmt.Fprintf(os.Stderr, "usage: roman[number]\n")
     flag.PrintDefaults()
     os.Exit(2)
}

func trans(num int64)string {
     numerals := map[int64]string {
      1:"I",
      5:"V",
      10:"X",
      50:"L",
      100:"C",
      500:"D",
      1000:"M",
     }

     keys := []int64{1000, 500, 100, 50, 10, 5, 1}

     var r string = ""

     for i,k := range keys {
     	if k <= num {
	    s := numerals[k]
	    d := int(math.Floor(float64(num)/float64(k)))
	    if d == 4 {
	       r = r + s + numerals[keys[i-1]]
	    } else {
	       r += strings.Repeat(s,d)
	    }
	    num -= int64(d*int(k))
	    if num <= 0 {
	       return r
	    }
        }
     } 
     return r    
}


func main() {
     flag.Usage = usage
     flag.Parse()

     args:= flag.Args()
     if len(args) < 1 {
        fmt.Printf("must enter a positive number\n")
	os.Exit(1)
     } else {
       num, _ := strconv.ParseInt(args[0], 0, 64)
       if (num > 0) {
              fmt.Printf("got: %s\n", trans(num))
       } else {
        fmt.Printf("must enter a positive number\n")
	os.Exit(1)
       }
     }
}