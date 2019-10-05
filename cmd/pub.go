package cmd

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "publish <topic> <message>",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		message := []byte(args[1])

		url := fmt.Sprintf("nats://%s:%d", host, port)
		nc, err := nats.Connect(url)
		if err != nil {
			log.Printf("Failed to connect to: %s\n", url)
			return
		}

		nc.Publish(topic, []byte(message))
		nc.Flush()
		if nc.LastError() != nil {
			log.Printf("Failed to publish to: %s\n", topic)
			return
		}

		log.Printf("Published a message to topic: %s\n", topic)
		log.Printf("Message: %s\n", message)
	},
}

func init() {
	rootCmd.AddCommand(pubCmd)
}
