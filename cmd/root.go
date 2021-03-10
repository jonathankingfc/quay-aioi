package cmd

import (
	_ "embed" // embed package is used to embed service files
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

//go:embed "assets/quay.service"
var quayServiceBytes []byte

//go:embed "assets/postgres.service"
var postgresServiceBytes []byte

//go:embed "assets/redis.service"
var redisServiceBytes []byte

type service struct {
	name     string
	image    string
	location string
	bytes    []byte
}

var services = []service{
	{
		"quay-app", "registry.redhat.io/quay/quay-rhel8:v3.4.1", "/etc/systemd/system/quay-app.service", quayServiceBytes,
	},
	{
		"quay-postgres", "registry.redhat.io/rhel8/postgresql-10:1", "/etc/systemd/system/quay-postgres.service", postgresServiceBytes,
	},
	{
		"quay-redis", "registry.redhat.io/rhel8/redis-5:1", "/etc/systemd/system/quay-redis.service", redisServiceBytes,
	},
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

func check(err error) {
	if err != nil {
		log.Fatalf("An error occurred: %s", err.Error())
	}
}

// verbose is the optional command that will display INFO logs
var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Display verbose logs")
}

var (
	rootCmd = &cobra.Command{
		Use: "quay-installer",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				log.SetLevel(log.InfoLevel)
			} else {
				log.SetLevel(log.WarnLevel)
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	fmt.Println(`
   __   __
  /  \ /  \     ______   _    _     __   __   __
 / /\ / /\ \   /  __  \ | |  | |   /  \  \ \ / /
/ /  / /  \ \  | |  | | | |  | |  / /\ \  \   /
\ \  \ \  / /  | |__| | | |__| | / ____ \  | |
 \ \/ \ \/ /   \_  ___/  \____/ /_/    \_\ |_|
  \__/ \__/      \ \__
                  \___\ by Red Hat
 Build, Store, and Distribute your Containers
	`)
	return rootCmd.Execute()
}
