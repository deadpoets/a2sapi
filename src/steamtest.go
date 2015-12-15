package main

import (
	"flag"
	"fmt"
	"os"
	"steamtest/src/config"
	"steamtest/src/constants"
	"steamtest/src/steam"
	"steamtest/src/steam/filters"
	"steamtest/src/util"
	"steamtest/src/web"
)

var (
	doConfig  bool
	runSilent bool
)

const (
	configFlag = "config"
	silentFlag = "silent"
)

func init() {
	flag.BoolVar(&doConfig, configFlag, false, "Generate the configuration file")
	flag.BoolVar(&runSilent, silentFlag, false,
		"Launch without displaying startup information.")
}

func main() {
	flag.Parse()
	if doConfig {
		if !util.FileExists(constants.GameFileFullPath) {
			filters.DumpDefaultGames()
		}
		config.CreateConfig()
		os.Exit(0)
	}
	checkRequiredConfigs()
	cfg := config.ReadConfig()

	if !runSilent {
		printStartInfo(cfg)
	}

	if cfg.SteamConfig.AutoQueryMaster {
		autoQueryGame := filters.GetGameByName(
			cfg.SteamConfig.AutoQueryGame)
		if autoQueryGame == filters.GameUnspecified {
			fmt.Println("Invalid game specified for automatic timed query!")
			fmt.Printf(
				"You may need to delete: '%s' and/or recreate the config with: %s --%s",
				constants.GameFileFullPath, constants.GameFileFullPath)
			os.Exit(1)
		}
		// HTTP server + API + Steam auto-querier
		go web.Start(runSilent)
		// + Steam querier
		filter := filters.NewFilter(autoQueryGame, filters.SrAll, nil)
		stop := make(chan bool, 1)
		go steam.StartMasterRetrieval(stop, filter, 7,
			cfg.SteamConfig.TimeBetweenMasterQueries)
		<-stop
	} else {
		// HTTP server + API standalone
		web.Start(runSilent)
	}
}

func printStartInfo(cfg *config.Config) {
	fmt.Printf("Launching %s\n", constants.AppInfo)

	if cfg.SteamConfig.AutoQueryMaster {
		fmt.Println("Automatic timed master server queries: enabled")
		fmt.Printf("Automatic timed master server queries every %d seconds\n",
			cfg.SteamConfig.TimeBetweenMasterQueries)
		fmt.Printf("Automatic timed master server query game: %s\n",
			cfg.SteamConfig.AutoQueryGame)
		fmt.Printf("Automatic timed master server query max hosts to receive: %d\n",
			cfg.SteamConfig.MaximumHostsToReceive)
	} else {
		fmt.Println("Automatic timed master server queries: disabled")
	}
}

func checkRequiredConfigs() {
	if !util.FileExists(constants.GameFileFullPath) {
		filters.DumpDefaultGames()
	}
	if !util.FileExists(constants.ConfigFilePath) {
		fmt.Printf("Could not read configuration file '%s' in the '%s' directory.\n",
			constants.ConfigFilename, constants.ConfigDirectory)
		fmt.Printf("You must generate the configuration file with: %s --%s\n",
			constants.GameFileFullPath)
		os.Exit(1)
	}
}