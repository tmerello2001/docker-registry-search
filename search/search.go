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

func SearchImage(searchTerm string, registryUrl string, useHttps bool) []DockerImage {
	baseUrl := fmt.Sprintf("http://%s", registryUrl)
	if useHttps {
		baseUrl = fmt.Sprintf("https://%s", registryUrl)
	}

	var foundImages []DockerImage

	var repositories DockerRepoList
	queryRegistry(fmt.Sprintf("%s/v2/_catalog?n=1000", baseUrl), &repositories)

	for _, repo := range repositories.Repositories {
		if strings.Contains(repo, searchTerm) {
			var image DockerImage
			queryRegistry(fmt.Sprintf("%s/v2/%s/tags/list", baseUrl, repo),
				&image)

			for _, tag := range image.Tags {
				image.Refs = append(image.Refs, fmt.Sprintf("%s/%s:%s", registryUrl, image.Name, tag))
			}
			foundImages = append(foundImages, image)
		}
	}
	return foundImages
}
