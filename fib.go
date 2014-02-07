package main

import (
    "fmt"
    "flag"
    "os"
    "strconv"
)

func usage() {
     fmt.Fprintf(os.Stderr, "usage: hello [name]\n")
     flag.PrintDefaults()
     os.Exit(2)
}

func fibonacci(num int64) int64 {
     if num == 0 {
     	return 0
     } else if num == 1 {
        return 1
     }

     return fibonacciR(num-2, 0, 1)
}

func fibonacciR(num int64, twoAgo int64, oneAgo int64) int64 {
     if num == 0 {
     	return twoAgo + oneAgo
     }

     return fibonacciR(num - 1, oneAgo, twoAgo + oneAgo)
}

func main() {
     flag.Usage = usage
     flag.Parse()

     args := flag.Args()

     if len(args) < 1 {
          fmt.Printf("must enter a number")
	  os.Exit(1)
     } else {
          num, _ := strconv.ParseInt(args[0], 0, 64)
	  fmt.Printf("got: %d\n",  fibonacci(num))
     }

}