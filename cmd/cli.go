package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var sender, recipient string
var amount float64

var rootCmd = &cobra.Command{
	Use:   "zamrud",
	Short: "Zamrud is a simple blockchain CLI",
}

var transactionCmd = &cobra.Command{
	Use:   "addtx",
	Short: "Add a transaction",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure required flags are provided
		if sender == "" || recipient == "" || amount <= 0 {
			fmt.Println("Please provide valid sender, recipient, and amount.")
			return
		}

		// Logic to add transaction
		// Example: blockchain.AddTransaction(sender, recipient, amount)
		fmt.Println("Transaction added successfully.")
	},
}

var mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Mine a block",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to mine a block
		// Example: blockchain.Mine()
		fmt.Println("Block mined successfully.")
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to display the blockchain
		// Example: print each block in the blockchain
	},
}

//var startCmd = &cobra.Command{
//	Use:   "start",
//	Short: "Start HTTP server",
//	Run: func(cmd *cobra.Command, args []string) {
//		port, _ := cmd.Flags().GetString("port")
//		log.Printf("Starting HTTP server on port %s", port)
//		handler.StartHTTPServer(port)
//	},
//}

//var nodeCmd = &cobra.Command{
//	Use:   "node",
//	Short: "Start as a nodeManager",
//	Run: func(cmd *cobra.Command, args []string) {
//		log.Println("Starting as nodeManager")
//		// Initialize and start the nodeManager here
//		nodeManager := services.NewNodeManager()
//		//... further logic
//	},
//}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	transactionCmd.Flags().StringVarP(&sender, "sender", "s", "", "Sender's address")
	transactionCmd.Flags().StringVarP(&recipient, "recipient", "r", "", "Recipient's address")
	transactionCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Amount to send")
}
