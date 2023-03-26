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
)

// 命令数组
var cmdArray = []*tDataNode{
	&tDataNode{"help", "Display available commands.", nil, nil},
	&tDataNode{"version", "Display version information.", version, nil},
	&tDataNode{"exit", "Exit the program.", exit, nil},
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
