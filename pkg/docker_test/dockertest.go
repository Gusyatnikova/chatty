package docker_test

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type DockerResource struct {
	Pool     *dockertest.Pool
	Resource *dockertest.Resource
}

type DockerConfig struct {
	HostIP        string
	HostPort      string
	DockerPort    string
	ContainerName string
	ImageName     string
	ImageTag      string
	NetworkID     string
	Env           []string
	Links         []string
}

func NewDockerResource(cfg *DockerConfig) (*DockerResource, error) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, err
	}

	err = pool.RemoveContainerByName(cfg.ContainerName)
	if err != nil {
		return nil, err
	}

	portBinding := make(map[docker.Port][]docker.PortBinding)

	portBinding[docker.Port(cfg.DockerPort)] = []docker.PortBinding{{
		HostIP:   cfg.HostIP,
		HostPort: cfg.HostPort,
	}}

	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Name:       cfg.ContainerName,
			Repository: cfg.ImageName,
			Tag:        cfg.ImageTag,
			NetworkID:  cfg.NetworkID,
			Env:        cfg.Env,
			Links:      cfg.Links,
		}, func(hostConfig *docker.HostConfig) {
			hostConfig.PortBindings = portBinding
		})
	if err != nil {
		return nil, err
	}

	return &DockerResource{Resource: resource, Pool: pool}, nil
}

func NewRedis() *DockerResource {
	dockerCfg := &DockerConfig{
		HostIP:        "localhost",
		HostPort:      "6379",
		DockerPort:    "6379/tcp",
		ContainerName: "redis-task-scheduler",
		ImageName:     "redis",
		ImageTag:      "latest",
	}

	dockerResource, err := NewDockerResource(dockerCfg)
	if err != nil {
		panic(err)
	}

	return dockerResource
}
func NewPostgres() *DockerResource {
	dockerCfg := &DockerConfig{
		HostIP:        "",
		HostPort:      "5432",
		DockerPort:    "5432/tcp",
		ContainerName: "postgres-chatty",
		ImageName:     "postgres",
		ImageTag:      "latest",
		Env: []string{
			"POSTGRES_USER=pgu",
			"POSTGRES_DB=chatty",
			"POSTGRES_PASSWORD=pg000PG",
		},
	}

	dockerResource, err := NewDockerResource(dockerCfg)
	if err != nil {
		panic(err)
	}

	return dockerResource
}
