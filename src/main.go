package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
)

func main() {
	var files []string
	err := filepath.Walk("replays", func(path string, info os.FileInfo, err error) error {
		// skip folders
		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal("error walking replay dir")
	}

	for _, file := range files {
		log.Printf("replay: %s", file)
		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("unable to open file: %s", err)
		}
		defer f.Close()

		p, err := manta.NewStreamParser(f)
		if err != nil {
			log.Fatalf("unable to create parser: %s", err)
		}

		p.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
			log.Printf("%s > %s said: %s\n", file, m.GetParam1(), m.GetParam2())
			return nil
		})

		p.Start()

		log.Printf("parse complete\n")
	}

}
