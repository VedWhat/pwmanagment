/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// createWalletCmd represents the createWallet command
var createWalletCmd = &cobra.Command{
	Use:   "createWallet",
	Short: "Command to create a wallet which can interact with the Solana blockchain",
	Long:  `Command to create a wallet which can interact with the Solana blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		wallet := createNewWallet(rpc.DevnetRPCEndpoint)
		fmt.Println("Public Key:" + wallet.account.PublicKey.ToBase58())
		fmt.Println("Private Key: ac35tg...dont worry it's saved in a secure location :P")
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createWalletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createWalletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
