/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"runtime"

	"github.com/UtkarshM-hub/Routine-Planner/cmd"
)

func main() {
	cmd.Execute()
	// handle err
	fmt.Println(runtime.GOOS)
}
