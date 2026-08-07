package main

import (
	"context"
	"time"

	"github.com/dobyte/due/registry/nacos/v2"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xconv"
)

func main() {
	var (
		reg  = nacos.NewRegistry()
		name = "game-server"
	)

	// 监听
	watch(reg, name, 1)

	// 监听
	watch(reg, name, 2)

	select {}
}

func watch(reg *nacos.Registry, serviceName string, goroutineID int) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	watcher, err := reg.Watch(ctx, serviceName)
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Infof("goroutine %d: startup", goroutineID)

		for {
			services, err := watcher.Next()
			if err != nil {
				log.Fatalf("goroutine %d: %v", goroutineID, err)
				return
			}

			log.Infof("goroutine %d: %s", goroutineID, xconv.Json(services))
		}
	}()
}
