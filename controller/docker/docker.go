package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

/*
Check if Docker engine and API are present and reachable on this system
*/
func CheckEngine() error {
	// Create new client
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// Ping Docker engine server
	_, err = client.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("docker engine is not running or API is not reachable: %v", err)
	}

	return nil
}

/*
Check if Docker container image is present
*/
func ImageExists(image string) (bool, error) {
	// Create new client
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return false, err
	}
	images, err := client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return false, err
	}
	for _, imageSum := range images {
		for _, tag := range imageSum.RepoTags {
			// Image is present
			if tag == image {
				return true, nil
			}
		}
	}

	return false, nil
}

/*
Pull Docker image
*/
func PullImage(image string) error {
	// Create new client
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	// Pull the image
	_, err = client.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	return nil
}
