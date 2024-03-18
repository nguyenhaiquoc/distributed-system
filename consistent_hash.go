package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

// HashRing represents the consistent hash ring
type HashRing struct {
	nodes       []string
	replication int
	keys        map[uint32]string
}

// NewHashRing creates a new HashRing with the given nodes and replication factor
func NewHashRing(nodes []string, replication int) *HashRing {
	hashRing := &HashRing{
		nodes:       nodes,
		replication: replication,
		keys:        make(map[uint32]string),
	}

	hashRing.generateKeys()

	return hashRing
}

// generateKeys generates the keys for each node in the hash ring
func (hr *HashRing) generateKeys() {
	for _, node := range hr.nodes {
		for i := 0; i < hr.replication; i++ {
			key := hr.hashKey(node, i)
			hr.keys[key] = node
		}
	}

	// Sort the keys for efficient lookup
	sortKeys := make([]uint32, 0, len(hr.keys))
	for key := range hr.keys {
		sortKeys = append(sortKeys, key)
	}
	sort.Slice(sortKeys, func(i, j int) bool {
		return sortKeys[i] < sortKeys[j]
	})
}

// hashKey hashes the given node and index to generate a key
func (hr *HashRing) hashKey(node string, index int) uint32 {
	str := node + strconv.Itoa(index)
	return crc32.ChecksumIEEE([]byte(str))
}

// GetNode returns the node responsible for the given key
func (hr *HashRing) GetNode(key string) string {
	hash := crc32.ChecksumIEEE([]byte(key))

	// Binary search to find the node responsible for the key
	nodes := hr.nodes
	i := sort.Search(len(nodes), func(i int) bool {
		return nodes[i] >= hr.keys[hash]
	})

	if i == len(nodes) {
		i = 0
	}

	return nodes[i]
}

func main_dev() {
	// Create a new HashRing with 3 nodes and replication factor of 2
	nodes := []string{"node1", "node2", "node3"}
	replication := 2
	hashRing := NewHashRing(nodes, replication)

	// Get the node responsible for the given key
	key := "some_key"
	node := hashRing.GetNode(key)

	fmt.Printf("Key '%s' is mapped to node '%s'\n", key, node)
}
