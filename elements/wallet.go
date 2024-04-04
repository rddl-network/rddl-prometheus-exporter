package elements

import (
	"errors"
	"fmt"

	elements "github.com/rddl-network/elements-rpc"
	"github.com/rddl-network/elements-rpc/types"
)

func GetWalletBalance(url string, wallet string) (balance float64, err error) {
	res, err := elements.GetBalance(url, []string{})
	if err != nil {
		return 0, fmt.Errorf("error getting balance for wallet %s: %w", wallet, err)
	}
	balance, err = GetAsset(res)
	if err != nil {
		return 0, fmt.Errorf("bitcoin balance not found for wallet %s", wallet)
	}
	return balance, nil
}

func GetAsset(balance types.GetBalanceResult) (float64, error) {
	btc, ok := balance["bitcoin"]
	if !ok {
		return 0, errors.New("bitcoin balance not found")
	}
	return btc, nil
}
