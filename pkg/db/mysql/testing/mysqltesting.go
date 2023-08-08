package mysqltesting

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var HostPort string

var (
	imageName     = "mysql/mysql-server:8.0.28"
	containerPort = nat.Port("3306/tcp")
	containerName = "mysql-test"
	Username      = "root"
	Password      = "123456"
)

// RunWithMongoInDocker 在 docker 容器中运行 MongoDB
func RunWithMongoInDocker(m *testing.M) int {
	c, err := client.NewClientWithOpts(client.WithVersion("1.41"))

	if err != nil {
		panic(fmt.Sprintf("new docker client err:%v", err))
	}

	ctx := context.Background()

	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"3306/tcp": {},
		},
		Env: []string{fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", Password), "MYSQL_ROOT_HOST=%"},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0", // 使用随机端口，防止端口冲突
				},
			},
		},
	}, nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	containerID := resp.ID
	defer func() {
		err := c.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			panic(fmt.Sprintf("error removing container: %v", err))
		}
	}()

	err = c.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("container started")

	inspRes, err := c.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}

	// 获取监听的端口
	ctIpPort := inspRes.NetworkSettings.Ports[containerPort][0]
	HostPort = fmt.Sprintf("%s:%s", ctIpPort.HostIP, ctIpPort.HostPort)

	return m.Run()
}
