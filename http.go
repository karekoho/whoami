package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    hostname_node, err := ioutil.ReadFile("/etc/hostname_node")

    if err != nil {
        panic(err)
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 	      // fmt.Fprintf(w, "I'm %s\n", os.Hostname()) // Container id
        fmt.Fprintf(w, "Node name: %s\nContainer ID: %s\nContainer name: %s",
          string(hostname_node), "???", "???")
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}
