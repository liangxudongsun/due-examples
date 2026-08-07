package main

import (
	"context"
	"time"

	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/registry"
	"github.com/dobyte/due/v2/utils/xconv"
	"github.com/dobyte/due/v2/utils/xuuid"
)

func main() {
	var (
		reg   = etcd.NewRegistry()
		id    = xuuid.UUID()
		name  = "game-server"
		alias = "mahjong"
		ins   = &registry.ServiceInstance{
			ID:       id,
			Name:     name,
			Kind:     cluster.Node.String(),
			Alias:    alias,
			State:    cluster.Work.String(),
			Endpoint: "grpc://127.0.0.1:2000",
		}
	)

	// 注册服务
	if err := reg.Register(context.Background(), ins); err != nil {
		log.Fatal(err)
	}

	log.Infof("register service success: %s", xconv.Json(ins))

	time.Sleep(2 * time.Second)

	// 更新服务
	ins.State = cluster.Busy.String()

	if err := reg.Register(context.Background(), ins); err != nil {
		log.Fatal(err)
	}

	log.Infof("update service success: %s", xconv.Json(ins))

	time.Sleep(30 * time.Second)
}
