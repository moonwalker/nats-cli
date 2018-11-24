package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	host string
	port int
)

var rootCmd = &cobra.Command{
	Use:   "nats-cli",
	Short: "cli for nats",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "s", "0.0.0.0", "bind to host address")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 4222, "use port for clients")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
