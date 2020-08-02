package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
)

func main() {
	// Create a new parser instance from a file. Alternatively see NewParser([]byte)
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

	log.Print("starting")
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

		// Register a callback, this time for the OnCUserMessageSayText2 event.
		p.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
			log.Printf("%s said: %s\n", m.GetParam1(), m.GetParam2())
			return nil
		})

		// Register a callback for map pings
		p.Callbacks.OnCDOTAUserMsg_LocationPing(func(ping *dota.CDOTAUserMsg_LocationPing) error {
			location := ping.GetLocationPing()
			log.Printf("%d pinged at (%d, %d)", ping.GetPlayerId(), location.GetX(), location.GetY())
			return nil
		})

		// Never called?
		p.Callbacks.OnCDOTAUserMsg_AbilityPing(func(ping *dota.CDOTAUserMsg_AbilityPing) error {
			log.Printf("%d ability-pinged %d", ping.GetPlayerId(), ping.GetAbilityId())
			return nil
		})

		// Never called?
		p.Callbacks.OnCDOTAUserMsg_ItemAlert(func(itemalert *dota.CDOTAUserMsg_ItemAlert) error {
			alert := itemalert.GetItemAlert()
			log.Printf("%d item-alerted %d", itemalert.GetPlayerId(), alert.GetItemAbilityId())
			return nil
		})

		// Never called?
		p.Callbacks.OnCDOTAUserMsg_EnemyItemAlert(func(itemalert *dota.CDOTAUserMsg_EnemyItemAlert) error {
			log.Printf("%d enemy-item-alerted %d's %d", itemalert.GetPlayerId(), itemalert.GetTargetPlayerId(), itemalert.GetItemAbilityId())
			return nil
		})

		p.Start()

		log.Printf("parse complete\n")

	}

}
