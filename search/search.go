package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type DockerRepoList struct {
	Repositories []string `json:"repositories"`
}

type DockerImage struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
	Refs []string
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func queryRegistry(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func SearchImage(searchTerm string, registry string, allTags bool) []DockerImage {
	var foundImages []DockerImage

	registryUrl := "docker.jampp.com"
	// registryUrl := "127.0.0.1:5000"

	var repositories DockerRepoList
	queryRegistry(fmt.Sprintf("https://%s/v2/_catalog", registryUrl), &repositories)

	for _, repo := range repositories.Repositories {
		if strings.Contains(repo, searchTerm) {
			var image DockerImage
			queryRegistry(fmt.Sprintf("https://%s/v2/%s/tags/list", registryUrl, repo),
				&image)

			for _, tag := range image.Tags {
				image.Refs = append(image.Refs, fmt.Sprintf("%s/%s:%s", registryUrl, image.Name, tag))
			}
			foundImages = append(foundImages, image)
		}
	}
	return foundImages
}
