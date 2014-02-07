package main

import (
    "fmt"
    "flag"
    "os"
)

func usage() {
     fmt.Fprintf(os.Stderr, "usage: hello [name]\n")
     flag.PrintDefaults()
     os.Exit(2)
}

func main() {
     flag.Usage = usage
     flag.Parse()

     args := flag.Args()

     if len(args) < 1 {
          fmt.Printf("hello, world\n")
     } else {
          fmt.Printf("hello, %s\n", args[0])
     }

}