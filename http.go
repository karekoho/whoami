package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	s "strings"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

// ContainerInfo Some info from types.Container
type ContainerInfo struct {
	name string
	ID   string
}

// Find a container by image name
func (ci *ContainerInfo) findByImage(image string, cli *docker.Client) (*ContainerInfo, error) {

	// Get the containers running on host
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		return ci, err
	}

	for _, container := range containers {
		if container.Image == image {
			ci.name = container.Names[0][1:] // Get the first name in list and skip leading "/"
			ci.ID = container.ID[:12]        // Get the first 12 characters
			return ci, err
		}
	}

	ci.name = "unknown"
	ci.ID = "unknown"
	return ci, err
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the host name of node
		hostnameNode, err := ioutil.ReadFile("/etc/hostname")
		if err != nil {
			panic(err)
		}

		cli, err := docker.NewEnvClient()
		if err != nil {
			panic(err)
		}

		// Find the container by image name
		ci := new(ContainerInfo)
		ci, err = ci.findByImage("karek/whoami", cli)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Node name: %s\nContainer ID: %s\nContainer name: %s",
			s.TrimRight(string(hostnameNode), "\n"), ci.ID, ci.name)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
