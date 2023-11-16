package docker

import "github.com/docker/docker/client"

/*
Builder struct

Contains info on payload builder container.

Client -> Docker client connection
*/
type Builder struct {
	Client *client.Client
}
