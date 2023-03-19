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
	"os"
)

func main() {
	fmt.Println(("hello,word!"))

	var cmd string = ""

	for true {
		fmt.Scan(&cmd)

		if cmd == "help" {
			fmt.Printf("this is help cmd!\n")
		} else if cmd == "quit" {
			os.Exit(0)
		} else {
			fmt.Printf("wrong cmd!\n")
		}
	}

}
