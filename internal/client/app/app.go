package app

import (
	"fmt"
	"log"
	"os"

	"github.com/dimk00z/grpc-filetransfer/internal/client/service"
	"github.com/spf13/cobra"
)

var (
	serverAddr string
	filePath   string
	rootCmd    = &cobra.Command{
		Use:   "transfer_client",
		Short: "Sending files via gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			clientService := service.New(serverAddr, filePath)
			if err := clientService.SendFile(); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&serverAddr, "addr", "a", "", "server address")
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "file path")
	if err := rootCmd.MarkFlagRequired("file"); err != nil {
		log.Fatal(err)
	}
	if err := rootCmd.MarkFlagRequired("addr"); err != nil {
		log.Fatal(err)
	}
}
