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
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return fmt.Errorf("failed to create docker client: %v", err)
	}

	_, err = client.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("docker engine is not running or API is not reachable: %v", err)
	}

	return nil
}

/*
Check if Docker container image is present
*/
func ImageExists(ctx context.Context, cli *client.Client, image string) (bool, error) {
	// Query API for image list
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
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
