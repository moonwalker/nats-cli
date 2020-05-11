package cmd

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var repCmd = &cobra.Command{
	Use:   "rep",
	Short: "reply <topic> <message>",
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

		_, err = nc.Subscribe(topic, func(msg *nats.Msg) {
			if len(msg.Reply) > 0 {
				log.Printf("Received a request on topic: %s\n", msg.Subject)
				log.Printf("Request: %s\n", string(msg.Data))
				log.Printf("Reply: %s\n", string(message))
				nc.Publish(msg.Reply, message)
			}
		})

		if err != nil {
			log.Printf("Failed to subscribe to: %s\n", topic)
			return
		}

		log.Printf("Listening for requests on: %s\n", topic)
		select {}
	},
}

func init() {
	rootCmd.AddCommand(repCmd)
}
