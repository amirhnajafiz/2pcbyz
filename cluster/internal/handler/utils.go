package handler

import (
	"fmt"
	"strconv"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/config"
)

// localAddress returns the address of client's server.
func localAddress(port int) string {
	return fmt.Sprintf("127.0.0.1:%d", port)
}

// findClientShard returns the client shard by checking the shards configs.
func findClientShard(client string, shards []config.ShardConfig) string {
	target, _ := strconv.Atoi(client)

	for _, shard := range shards {
		if target >= shard.From && target <= shard.To {
			return shard.Cluster
		}
	}

	return ""
}
