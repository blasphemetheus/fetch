package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "github.com/blasphemetheus/giffy"
)

func main() {
  gifler := func(w http.ResponseWriter, r *http.Request) {
    // first print this out
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

    u, err := url.Parse(r.URL.Path)
    if err != nil {
      fmt.Println("hello")
      panic(err)
    }
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    //fmt.Println(m["k"][0])

    giffy.Lissajous(w)
  }
  //http.HandleFunc("/", handler)
  http.HandleFunc("/", gifler)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s %s %s/n", r.Method, r.URL, r.Proto)
  for k, v := range r.Header {
    fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
  }
  fmt.Fprintf(w, "Host = %q\n", r.Host)
  fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
  if err := r.ParseForm(); err != nil {
    log.Print(err)
  }
  for k, v := range r.Form {
    fmt.Fprintf(w, "Form[%q] %q\n", k, v)
  }

  // done with the info stuff
}
