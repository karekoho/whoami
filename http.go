package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

func test(w http.ResponseWriter) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Fprintf(w, "\ncontainer.ID[:10]: %s\ncontainer.Image: %s", container.ID[:10], container.Image)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	hostnameNode, err := ioutil.ReadFile("/etc/hostname_node")

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "I'm %s\n", os.Hostname()) // Container id
		fmt.Fprintf(w, "Node name: %s\nContainer ID: %s\nContainer name: %s",
			string(hostnameNode), "???", "???")
		test(w)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
