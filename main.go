package main

import (
	"log"
	"time"
	"github.com/StackExchange/wmi"
)

type Win32_Process struct {
	CommandLine string
}

func main() {
	var prevCommandLines = make([]string, 0)

	for {
		var ytdlProcs []Win32_Process
		q := wmi.CreateQuery(&ytdlProcs, "where name = 'youtube-dl.exe'")

		if err := wmi.Query(q, &ytdlProcs); err != nil {
			log.Fatal(err)
		}

		commandLines := make([]string, 0)

		for _, p := range ytdlProcs {
			for _, cl := range prevCommandLines {
				if p.CommandLine == cl {
					goto AVOID_DUPLICATE_PRINTING_IN_A_SHORT_TIME 
				}
			}

			for _, cl := range commandLines {
				if p.CommandLine == cl {
					goto AVOID_DUPLICATE_PRINTING_IN_A_SHORT_TIME
				}
			}

			log.Println(p.CommandLine)

AVOID_DUPLICATE_PRINTING_IN_A_SHORT_TIME:
			commandLines = append(commandLines, p.CommandLine)
		}

		prevCommandLines = commandLines

		time.Sleep(time.Second / 2)
	}
}

