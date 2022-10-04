package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/shirou/gopsutil/host"
	"time"
)

const PUMPKIN_ONE string = `
                       .,'
                     .''.'
                    .' .'
        _.ood0Pp._ ,'  '.~ .q?00doo._
    .od00Pd0000Pdb._. . _:db?000b?000bo.
  .?000Pd0000Pd0000PdbMb?0000b?000b?0000b.
.d0000Pd0000Pd0000Pd0000b?0000b?000b?0000b.
d0000Pd0000Pd00000Pd0000b?00000b?0000b?000b.
00000Pd0000Pd0000Pd00000b?00000b?0000b?0000b
?0000b?0000b?0000b?00000Pd00000Pd0000Pd0000P
?0000b?0000b?0000b?00000Pd00000Pd0000Pd000P
'?0000b?0000b?0000b?0000Pd0000Pd0000Pd000P'
  '?000b?0000b?000b?0000Pd000Pd0000Pd000P
    '~?00b?000b?000b?000Pd00Pd000Pd00P'
        '~?0b?0b?000b?0Pd0Pd000PdP~'
`

func formatEntry(title string, item string, c color.Color256) string {
	return c.Sprintf(title+": ") + item
}

func printEntry(title string, item string, c color.Color256) {
	fmt.Println(formatEntry(title, item, c))
}

func getToSpooky() string {
	_, month, day := time.Now().Date()
	var toSpooky string
	if month == 10 {
		toSpooky = fmt.Sprintf("%d days until spooky day", 31-day)
	} else if month > 10 {
		toSpooky = fmt.Sprintf("%d months and %d days until spooky day", month+(month-10), 31-day)
	} else if month < 10 {
		toSpooky = fmt.Sprintf("%d months and %d days until spooky day", 10-month, 31-day)
	}
	return toSpooky
}

type Theme int

const (
	Pumpkin Theme = iota + 1
)

func (i Theme) Image() string {
	images := [...]string{PUMPKIN_ONE}
	return images[i-1]
}

func (i Theme) Color() color.Color256 {
	pumpkin_orange := color.C256(208)

	colors := []color.Color256{pumpkin_orange}
	return colors[i-1]
}

func decideTheme() Theme {
	return Pumpkin
}

func main() {
	theme := decideTheme()

	color := theme.Color()
	image := theme.Image()

	color.Println(image)

	hostStat, _ := host.Info()
	toSpooky := getToSpooky()

	printEntry("OS", hostStat.Platform+" "+hostStat.OS+" "+hostStat.KernelArch, color)
	printEntry("Kernel", hostStat.KernelVersion, color)
	printEntry("Hostname", hostStat.Hostname, color)
	printEntry("To Spooky", toSpooky, color)
	fmt.Printf("\n")
}
