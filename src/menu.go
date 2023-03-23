/**************************************************************************************************/
/* Copyright (C) mc2lab.com, SSE@USTC, 2014-2015                                                  */
/*                                                                                                */
/*  FILE NAME             :  menu.c                                                               */
/*  PRINCIPAL AUTHOR      :  Mengning                                                             */
/*  SUBSYSTEM NAME        :  menu                                                                 */
/*  MODULE NAME           :  menu                                                                 */
/*  LANGUAGE              :  GO                                                                   */
/*  TARGET ENVIRONMENT    :  ANY                                                                  */
/*  DATE OF FIRST RELEASE :  2023/03/13                                                           */
/*  DESCRIPTION           :  This is a menu program                                               */
/**************************************************************************************************/

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

// 命令数组
var cmdArray = []*tDataNode{
	&tDataNode{"help", "Display available commands.", nil, nil},
	&tDataNode{"version", "Display version information.", version, nil},
	&tDataNode{"exit", "Exit the program.", exit, nil},
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

func main() {

	cmdArray[0].handler = help

	var input string
	for {
		fmt.Print(">> ")
		fmt.Scanln(&input)
		node := findCmd(input)
		if node == nil {
			fmt.Printf("Command not found: %s\n", input)
			continue
		}
		node.handler()
		if node.cmd == "exit" {
			break
		}
	}
}
