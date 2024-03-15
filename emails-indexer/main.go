package main

import (
	"email-challenge/utils"
	"emails-indexer/handlers"

	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	utils.LoadEnvVars()

	// Start Profiling
	cpu, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	log.Println("Starting indexer...")
	startTime := time.Now()

	nameIndex := os.Getenv("ZS_INDEX")

	// Check index exists
	res := handlers.CkeckIndexExists()
	if res.StatusCode == http.StatusOK {
		log.Printf("Index %s does exist", nameIndex)
		err := handlers.DeleteIndex(nameIndex)
		if err != nil {
			log.Fatal("Error DELETE index: ", err)
		}
	} else {
		log.Printf("Index %s does not exist", nameIndex)
	}

	handlers.ProcessDataIndexer()

	duration := time.Since(startTime)
	log.Printf("Finished indexing.\nDuration: %.2f minutes", duration.Minutes())

	// Finish Profiling
	runtime.GC()
	mem, err := os.Create("memory.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer mem.Close()

	if err := pprof.WriteHeapProfile(mem); err != nil {
		log.Fatal(err)
	}
}
