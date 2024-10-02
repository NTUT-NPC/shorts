package shorts

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type Stat struct {
	Visitors    int       `json:"visitors"`
	LastVisited time.Time `json:"last_visited"`
}

type Stats map[string]Stat

var (
	stats      Stats = make(Stats)
	statsMutex sync.Mutex
)

const statsFile = "config/stats.json"

func UpdateStat(slug string) {
	statsMutex.Lock()
	defer statsMutex.Unlock()

	visitors := 1
	if stat, ok := stats[slug]; ok {
		visitors = stat.Visitors + 1
	}

	stats[slug] = Stat{
		Visitors:    visitors,
		LastVisited: time.Now(),
	}

	// Write stats to file on every update for simplicity
	writeStats()
}

func ReadStats() {
	file, err := os.ReadFile(statsFile)
	if err != nil {
		log.Println("No stats file found, starting with an empty stats map")
		return
	}

	err = json.Unmarshal(file, &stats)
	if err != nil {
		log.Fatal(err)
	}
}

func writeStats() {
	statsJson, err := json.MarshalIndent(stats, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(statsFile)
	if err != nil {
		log.Fatal(err)
	}

	file.Write(statsJson)
}
