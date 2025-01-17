package main

import (
	"context"
	"fmt"
	"os"

	"github.com/getlago/lago-go-client"
)

func main() {
	apiKey := os.Getenv("LAGO_API_KEY")

	lagoClient := lago.New().
		SetApiKey(apiKey)

	walletListInput := &lago.WalletListInput{
		PerPage:            10,
		Page:               1,
		ExternalCustomerID: "orion-deploy-s:customer:org_01jhk7jb6vfj68hk90stmesmm0",
	}

	walletResult, err := lagoClient.Wallet().GetList(context.Background(), walletListInput)
	if err != nil {
		// Error is *lago.Error
		panic(err)

		// Error:
		// panic: {"status":0,"error":"","code":"","err":"json: cannot unmarshal object into Go struct field RecurringTransactionRuleResponse.wallets.recurring_transaction_rules.transaction_metadata of type []lago.WalletTransactionMetadata"}
	}

	// wallet is *lago.WalletResult
	fmt.Println(walletResult)
}
