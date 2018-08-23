package cmd

import (
	"log"

	"github.com/nats-io/go-nats"
	"github.com/spf13/cobra"
)

var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "publish <topic> <message>",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		message := []byte(args[1])

		nc, err := nats.Connect(nats.DefaultURL)
		if err != nil {
			log.Printf("Failed to connect to: %s\n", nats.DefaultURL)
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
