package main

import (
	"io/ioutil"
	"log"
	"github.com/gookit/color"
	"fmt"
	"github.com/shirou/gopsutil/host"
)


func formatEntry(title string, item string, c color.Color256) string {
	return c.Sprintf(title + ": ") + item
}

func printEntry(title string, item string, c color.Color256) {
	fmt.Println(formatEntry(title, item, c))
}

func main() {
	pumpkin_orange := color.C256(208)

	file := "./pumpkin_one"

	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n")
	text := string(fileContent)
	pumpkin_orange.Println(text);

	hostStat, _ := host.Info()

	printEntry("OS", hostStat.Platform + " " + hostStat.OS + " "+ hostStat.KernelArch, pumpkin_orange)
	printEntry("Kernel", hostStat.KernelVersion, pumpkin_orange)
	printEntry("Hostname", hostStat.Hostname, pumpkin_orange)
	fmt.Printf("\n")
}
