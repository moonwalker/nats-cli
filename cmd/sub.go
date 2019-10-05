package cmd

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "subscribe <topic>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]

		url := fmt.Sprintf("nats://%s:%d", host, port)
		nc, err := nats.Connect(url)
		if err != nil {
			log.Printf("Failed to connect to: %s\n", url)
			return
		}

		_, err = nc.Subscribe(topic, func(msg *nats.Msg) {
			log.Printf("Received a message to topic: %s\n", msg.Subject)
			log.Printf("Message: %s\n", string(msg.Data))

		})
		if err != nil {
			log.Printf("Failed to subscribe to: %s\n", topic)
			return
		}

		log.Printf("Listening for messages on: %s\n", topic)
		select {}
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
