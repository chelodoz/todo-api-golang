package integration

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"todo-api-golang/internal/config"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func startMongoDB(pool *dockertest.Pool, mongoVersion string, network *dockertest.Network, config *config.Config) (*dockertest.Resource, error) {
	r, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       config.MongoHost,
		Repository: "mongo",
		Tag:        mongoVersion,
		Mounts:     []string{getProjectRootPath() + "/internal/mongo/initdb.d:/docker-entrypoint-initdb.d:ro"},
		Networks:   []*dockertest.Network{network},
		Env: []string{
			fmt.Sprintf("MONGO_INITDB_ROOT_USERNAME=%s", config.MongoUsername),
			fmt.Sprintf("MONGO_INITDB_ROOT_PASSWORD=%s", config.MongoPassword),
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		fmt.Printf("Could not start Mongodb: %v \n", err)
		return r, err
	}

	err = r.Expire(60)
	if err != nil {
		fmt.Printf("Could set expiration time: %v \n", err)
	}

	mongoPort := r.GetPort("27017/tcp")

	fmt.Printf("mongo-%s - connecting to : %s \n", mongoVersion, fmt.Sprintf("mongodb://localhost:%s", mongoPort))
	if err := pool.Retry(func() error {
		var err error

		url := fmt.Sprintf("mongodb://%s:%s@localhost:%s", config.MongoUsername, config.MongoPassword, mongoPort)
		clientOptions := options.Client().ApplyURI(url)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return err
		}

		err = client.Ping(context.TODO(), nil)
		if err == nil {
			fmt.Println("successfully connected to Mongodb.")
		}
		return err

	}); err != nil {
		fmt.Printf("Could not connect to mongodb container: %v \n", err)
		return r, err
	}

	return r, nil
}
func startAPI(pool *dockertest.Pool, network *dockertest.Network, config *config.Config) (*dockertest.Resource, error) {
	apiContainerName := "todo-integration-test"

	r, err := pool.BuildAndRunWithBuildOptions(
		&dockertest.BuildOptions{
			ContextDir: "../",
			Dockerfile: "Dockerfile",
			BuildArgs:  []docker.BuildArg{{Name: "test", Value: "-t mysuperimage -f MyDockerfile ."}},
		},
		&dockertest.RunOptions{
			Name:       apiContainerName,
			Repository: apiContainerName,
			Networks:   []*dockertest.Network{network},
		}, func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		})
	if err != nil {
		fmt.Printf("Could not start %s: %v \n", apiContainerName, err)
		return r, err
	}

	err = r.Expire(60)
	if err != nil {
		fmt.Printf("Could set expiration time: %v \n", err)
	}

	waiter, err := pool.Client.AttachToContainerNonBlocking(docker.AttachToContainerOptions{
		Container:    apiContainerName,
		OutputStream: log.Writer(),
		ErrorStream:  log.Writer(),
		RawTerminal:  true,
		Logs:         true,
		Stream:       true,
		Stdout:       true,
		Stderr:       true,
	})
	if err != nil {
		fmt.Println("unable to get LOGS: ", err)
	}
	defer waiter.Close()

	appPort = r.GetPort("8080/tcp")
	basePath = fmt.Sprintf("http://localhost:%s/api/v1", appPort)
	if err := pool.Retry(func() error {

		resp, err := http.Get(fmt.Sprintf("%s/health", basePath))
		if err != nil {
			fmt.Printf("trying to connect to %s on localhost:%s, got : %v \n", apiContainerName, appPort, err)
			return err
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("trying to connect to %s on localhost:%s, got : %v , status: %v \n", apiContainerName, appPort, err, resp.StatusCode)
			return err
		}

		fmt.Println("status: ", resp.StatusCode)
		rs, _ := io.ReadAll(resp.Body)
		fmt.Printf("RESPONSE: %s \n", rs)

		return nil
	}); err != nil {
		fmt.Printf("Could not connect to %s container: %v \n", apiContainerName, err)
		return r, err
	}

	return r, nil
}

func getProjectRootPath() string {
	p, err := os.Getwd()

	if err != nil {
		panic("Unable to get project root path")
	}
	parent := filepath.Dir(p)
	return strings.ReplaceAll(parent, "\\", "/")
}

func cleanUp(code int, network *dockertest.Network, mongoRes *dockertest.Resource, apiRes *dockertest.Resource) {
	fmt.Println("removing resources.")
	if mongoRes != nil {
		if err := pool.Purge(mongoRes); err != nil {
			log.Fatalf("Could not purge resource: %s\n", err)
		}
	}

	if apiRes != nil {
		if err := pool.Purge(apiRes); err != nil {
			log.Fatalf("Could not purge resource: %s\n", err)
		}
	}

	if network != nil {
		if err := pool.RemoveNetwork(network); err != nil {
			log.Fatalf("Could not remove network: %s\n", err)
		}
	}
	os.Exit(code)
}
