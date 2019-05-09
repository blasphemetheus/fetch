package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
  "strings"
)

func main() {
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    if ! strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
      url = "http://" + url
    }
    go fetch(url, ch) // starts goroutine
  }
  for range os.Args[1:] {
    fmt.Println(<-ch) // receive from channel ch
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err) // send to channel ch
    return
  }

  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  //fmt.Fprintf(os.Stderr, "Response-Status:" + resp.Status + " ")
  resp.Body.Close() // to not leak resources I guess
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
