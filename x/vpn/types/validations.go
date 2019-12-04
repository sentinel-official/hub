package types

import (
	"fmt"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	reDnmString = `[a-z][a-z0-9]{2,15}`
	reAmt       = `[[:digit:]]+`
	reSpc       = `[[:space:]]*`
	reDnm       = regexp.MustCompile(fmt.Sprintf(`^%s$`, reDnmString))
	reCoin      = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnmString))
)

func validateDenom(denom string) error {
	if !reDnm.MatchString(denom) {
		return fmt.Errorf("invalid denom: %s", denom)
	}
	return nil
}

func isValidCoins(coins sdk.Coins) bool {
	switch len(coins) {
	case 0:
		return true
	case 1:
		if err := validateDenom(coins[0].Denom); err != nil {
			return false
		}
		return !coins[0].IsNegative()
	default:
		if !isValidCoins((sdk.Coins{coins[0]})) {
			return false
		}

		lowDenom := coins[0].Denom
		for _, coin := range coins[1:] {
			if strings.ToLower(coin.Denom) != coin.Denom {
				return false
			}
			if coin.Denom <= lowDenom {
				return false
			}
			if coin.IsNegative() {
				return false
			}

			lowDenom = coin.Denom
		}

		return true
	}
}

func ParseCoin(coinStr string) (coin sdk.Coin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		return sdk.Coin{}, fmt.Errorf("invalid coin expression: %s", coinStr)
	}

	denomStr, amountStr := matches[2], matches[1]

	amount, ok := sdk.NewIntFromString(amountStr)
	if !ok {
		return sdk.Coin{}, fmt.Errorf("failed to parse coin amount: %s", amountStr)
	}

	if err := validateDenom(denomStr); err != nil {
		return sdk.Coin{}, fmt.Errorf("invalid denom cannot contain upper case characters or spaces: %s", err)
	}

	return sdk.NewCoin(denomStr, amount), nil
}

func ParseCoins(coinsStr string) (sdk.Coins, error) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return nil, nil
	}

	coinStrs := strings.Split(coinsStr, ",")
	coins := make(sdk.Coins, len(coinStrs))
	for i, coinStr := range coinStrs {
		coin, err := ParseCoin(coinStr)
		if err != nil {
			return nil, err
		}

		coins[i] = coin
	}

	coins.Sort()

	if !isValidCoins(coins) {
		return nil, fmt.Errorf("parseCoins invalid: %#v", coins)
	}

	return coins, nil
}
