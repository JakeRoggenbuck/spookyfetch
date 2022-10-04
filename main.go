package main

import (
	"github.com/gookit/color"
	"fmt"
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
	return c.Sprintf(title + ": ") + item
}

func printEntry(title string, item string, c color.Color256) {
	fmt.Println(formatEntry(title, item, c))
}

func main() {
	pumpkin_orange := color.C256(208)

	pumpkin_orange.Println(PUMPKIN_ONE);

	hostStat, _ := host.Info()
	_, month, day := time.Now().Date()

	var toSpooky string
	if month == 10 {
		toSpooky = fmt.Sprintf("%d days until spooky day", 31 - day)
	} else if month > 10 {
		toSpooky = fmt.Sprintf("%d months and %d days until spooky day", month + (month - 10), 31 - day)
	} else if month < 10 {
		toSpooky = fmt.Sprintf("%d months and %d days until spooky day", 10 - month, 31 - day)
	}

	printEntry("OS", hostStat.Platform + " " + hostStat.OS + " "+ hostStat.KernelArch, pumpkin_orange)
	printEntry("Kernel", hostStat.KernelVersion, pumpkin_orange)
	printEntry("Hostname", hostStat.Hostname, pumpkin_orange)
	printEntry("To Spooky", toSpooky, pumpkin_orange)
	fmt.Printf("\n")
}
