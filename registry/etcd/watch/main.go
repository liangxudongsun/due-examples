package main

import (
	"context"

	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xconv"
)

func main() {
	var (
		reg  = etcd.NewRegistry()
		name = "game-server"
	)

	// 监听
	watch(reg, name, 1)

	// 监听
	watch(reg, name, 2)

	select {}
}

func watch(reg *etcd.Registry, serviceName string, goroutineID int) {
	watcher, err := reg.Watch(context.Background(), serviceName)
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
