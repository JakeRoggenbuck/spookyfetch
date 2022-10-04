package main

import (
	"io/ioutil"
	"log"
	"github.com/gookit/color"
	"fmt"
	"github.com/shirou/gopsutil/host"
)

func main() {
	pumpkin_orange := color.C256(208)

	file := "./pumpkin_one"

	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	pumpkin_orange.Println(text);

	hostStat, _ := host.Info()
	fmt.Printf("OS: %s %s %s\n", hostStat.Platform, hostStat.OS, hostStat.KernelArch)
	fmt.Println("Kernel: " + hostStat.KernelVersion)
	fmt.Println("Hostname: " + hostStat.Hostname)
}
