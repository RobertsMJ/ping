package main

import (
	"log"
	"os"

	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
)

func main() {
	f, err := os.Open("out/replay.dem")
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	finfo, err := f.Stat()
	if err != nil {
		log.Fatalf("File stat error: %s", err)
	}
	log.Printf("filename: %v", finfo)

	p, err := manta.NewStreamParser(f)
	if err != nil {
		log.Fatalf("unable to create parser: %s", err)
	}

	p.Callbacks.OnCUserMessageSayText(func(m *dota.CUserMessageSayText) error {
		log.Printf("%d said: %s\n", m.GetPlayerindex(), m.GetText())
		return nil
	})

	p.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
		log.Printf("%s said: %s\n", m.GetParam1(), m.GetParam2())
		return nil
	})

	p.Callbacks.OnCDOTAUserMsg_AbilityPing(func(ping *dota.CDOTAUserMsg_AbilityPing) error {
		log.Printf("%d pinged %d\n", ping.GetCasterId(), ping.GetAbilityId())
		return nil
	})

	p.Start()

	log.Printf("finished parsing\n")
}
