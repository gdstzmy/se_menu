package main

import (
	"fmt"
	"strings"
)

type tDataNode struct {
	cmd     string
	desc    string
	handler func() int
	next    *tDataNode
}

func help() int {
	fmt.Println("Available commands:")
	for _, cmd := range cmdArray {
		fmt.Printf("    %-10s%s\n", cmd.cmd, cmd.desc)
	}
	return 0
}

func version() int {
	fmt.Println("Version 1.0")
	return 0
}

func exit() int {
	fmt.Println("Exiting...")
	return 1
}

func findCmd(cmd string) *tDataNode {
	for _, node := range cmdArray {
		if strings.ToLower(node.cmd) == strings.ToLower(cmd) {
			return node
		}
	}
	return nil
}
