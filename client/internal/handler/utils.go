package handler

import (
	"fmt"
	"strconv"

	"github.com/F24-CSE535/2pcbyz/client/internal/config"
)

// localAddress returns the address of client's server.
func localAddress(port int) string {
	return fmt.Sprintf("127.0.0.1:%d", port)
}

// findClientShard returns the client shard by checking the shards configs.
func findClientShard(client string, shards []config.ShardConfig) string {
	target, _ := strconv.Atoi(client)

	for _, shard := range shards {
		if target >= shard.Range[0] && target <= shard.Range[1] {
			return shard.Cluster
		}
	}

	return ""
}
