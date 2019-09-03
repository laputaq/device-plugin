package main

import (
	"fmt"
	"strings"
)

func encodeID(uuid string, index uint) string {
	return fmt.Sprintf("%s_%d", uuid, index)
}

func decodeID(id string) string {
	index := strings.LastIndex(id, "_")
	return id[:index]
}