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

// ContainerInfo Some info from types.Container
type ContainerInfo struct {
	name string
	ID   string
}

// Find a container by image name
func (ci *ContainerInfo) findByImage(image string, cli *docker.Client) *ContainerInfo {

	// Get the containers running on host
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == image {
			ci.name = container.Names[0][1:] // Get the first name in list and skip leading "/"
			ci.ID = container.ID[:12]        // Get the first 12 characters
			return ci
		}
	}

	ci.name = "unknown"
	ci.ID = "unknown"
	return ci
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get the host name of node
	// TODO: read the first line and strip off newline
	hostnameNode, err := ioutil.ReadFile("/etc/hostname_node")

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cli, err := docker.NewEnvClient()
		if err != nil {
			panic(err)
		}

		// Find the container by image name
		ci := new(ContainerInfo)
		ci.findByImage("karek/whoami", cli)

		fmt.Fprintf(w, "Node name: %sContainer ID: %s\nContainer name: %s",
			string(hostnameNode), ci.ID, ci.name)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
