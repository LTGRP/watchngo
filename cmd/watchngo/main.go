package main

import (
	"log"
	"os"

	"github.com/LTGRP/watchngo/pkg/conf"
	"github.com/LTGRP/watchngo/pkg/watcher"

	"flag"
)

func run(watchers []*watcher.Watcher) {
	forever := make(chan bool, 1)
	working := 0

	for _, watcher := range watchers {
		if err := watcher.Find(); err != nil {
			log.Printf("error: watcher.Find: %s: %v", watcher.Name, err)
			continue
		}

		go watcher.Work()

		working++
	}

	if working > 0 {
		<-forever
	} else {
		log.Fatalf("error: no watcher working")
	}
}

func main() {
	flagCfg := flag.String("conf", "watchngo.ini", "configuration file path")

	flagMatch := flag.String("match", "", "file or directory to watch")
	flagFilter := flag.String("filter", "", "filter as a regex supported by golang")
	flagCommand := flag.String("command", "", "command to run. see configuration example for supported variables")
	flagExecutor := flag.String("executor", "unixshell", "executors: unixshell, raw, stdout")
	flagDebug := flag.Bool("debug", false, "debug")

	flag.Parse()

	logger := log.New(os.Stderr, "", log.LstdFlags)

	log.SetOutput(os.Stderr)

	if *flagCommand != "" && *flagMatch != "" {
		executor, err := conf.ExecutorFrom(*flagExecutor)
		if err != nil {
			log.Fatal(err)
		}
		w, err := watcher.NewWatcher(
			"on the fly",
			*flagMatch,
			*flagFilter,
			*flagCommand,
			executor,
			*flagDebug,
			logger,
		)

		if err != nil {
			log.Fatalf("error: on the fly: %v", err)
		}

		run([]*watcher.Watcher{w})
	} else {

		watchers, err := conf.WatchersFromPath(*flagCfg, logger)
		if err != nil {
			log.Fatalf("error: WatchersFromPath: %v", err)
		}

		run(watchers)
	}
}
