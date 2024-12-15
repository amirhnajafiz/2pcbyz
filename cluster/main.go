package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/config"
	"github.com/F24-CSE535/2pcbyz/cluster/internal/handler"
	"github.com/F24-CSE535/2pcbyz/cluster/internal/server"
	"github.com/F24-CSE535/2pcbyz/cluster/internal/storage"
	"github.com/F24-CSE535/2pcbyz/cluster/pkg/logger"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("at least two arguments are needed (./main <config-path> <iptable>)")
	}

	// load config file
	cfg := config.New(args[1])

	// load iptable file
	ipt := config.NewIPTable(args[2])

	// create replicas based on the information provided in config
	wg := sync.WaitGroup{}
	for _, replica := range cfg.Replicas {
		// add one value to wait-group
		wg.Add(1)

		log.Printf("starting replica: %s on %d\n", replica.Name, replica.Port)

		// start a go-routine
		go func(name string, port int) {
			defer func() {
				wg.Done() // call done on wait-group
			}()

			// create a new logger instance
			logr := logger.NewLogger(cfg.LogLevel, fmt.Sprintf("%s_logs.csv", name))
			if logr == nil {
				log.Fatal("failed to initialize zap logger")
				return
			}

			// open database connection
			stg, err := storage.NewStorage(cfg.Storage.URI, cfg.Storage.Database, name)
			if err != nil {
				log.Fatal(err)
				return
			}

			// update shards of the node
			if err := stg.DeleteShards(); err != nil {
				log.Fatal(err)
				return
			}
			if err := stg.InsertShards(cfg.Shard.From, cfg.Shard.To); err != nil {
				log.Fatal(err)
				return
			}

			// create a handler queue
			queue := make(chan context.Context, cfg.Handler.QueueSize)

			// create a new handler instance
			hdl := handler.Handler{
				Sequence: int(time.Now().Unix()),
				Port:     port,
				Cfg:      &cfg,
				Ipt:      &ipt,
				Logger:   logr.Named("handler"),
				Storage:  stg,
				Queue:    queue,
			}

			// start the handler instances in a go-routine
			for i := 0; i < cfg.Handler.Instances; i++ {
				go func() {
					logr.Info("handler started")
					hdl.Start()
				}()
			}

			// create a bootstrap instance
			bts := server.Bootstrap{
				ServicePort: port,
				Logger:      logr.Named("grpc"),
				Storage:     stg,
				Queue:       queue,
			}

			// start the gRPC server
			if err := bts.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}(replica.Name, replica.Port)
	}

	// wait for all replicas
	fmt.Printf("total %d replicas running. see logs.\n", len(cfg.Replicas))
	wg.Wait()
}
