package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var reqCmd = &cobra.Command{
	Use:   "req",
	Short: "request <topic> [message]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]

		var message []byte
		if len(args) == 2 {
			message = []byte(args[1])
		}

		url := fmt.Sprintf("nats://%s:%d", host, port)
		nc, err := nats.Connect(url)
		if err != nil {
			log.Printf("Failed to connect to: %s\n", url)
			return
		}

		msg, err := nc.Request(topic, message, time.Millisecond*250)
		if err != nil {
			log.Printf("Failed to send a request to: %s\n", topic)
			return
		}
		
		log.Printf("Sent a request to topic: %s\n", topic)
		log.Printf("Request: %s\n", string(message))
		log.Printf("Reply: %s\n", string(msg.Data))

		err = nc.Flush()
		if err != nil {
			log.Printf("Failed to flush nats connection")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(reqCmd)
}
