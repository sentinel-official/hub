package types

import (
	"reflect"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestNode_Bytes(t *testing.T) {
	type fields struct {
		GigabytePrices sdk.Coins
	}
	type args struct {
		coin sdk.Coin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sdk.Int
		wantErr bool
	}{
		{
			"nil gigabyte_prices and empty coin",
			fields{
				GigabytePrices: nil,
			},
			args{
				coin: sdk.Coin{},
			},
			sdk.NewInt(0),
			true,
		},
		{
			"empty gigabyte_prices and empty coin",
			fields{
				GigabytePrices: sdk.Coins{},
			},
			args{
				coin: sdk.Coin{},
			},
			sdk.NewInt(0),
			true,
		},
		{
			"1one gigabyte_prices and empty coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.Coin{},
			},
			sdk.NewInt(0),
			true,
		},
		{
			"1one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"1one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"1one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(2000000000),
			false,
		},
		{
			"1one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(3000000000),
			false,
		},
		{
			"1one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(4000000000),
			false,
		},
		{
			"1one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(5000000000),
			false,
		},
		{
			"1one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(6000000000),
			false,
		},
		{
			"1one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(7000000000),
			false,
		},
		{
			"1one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(8000000000),
			false,
		},
		{
			"1one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(9000000000),
			false,
		},
		{
			"2one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"2one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(500000000),
			false,
		},
		{
			"2one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"2one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(1500000000),
			false,
		},
		{
			"2one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(2000000000),
			false,
		},
		{
			"2one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(2500000000),
			false,
		},
		{
			"2one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(3000000000),
			false,
		},
		{
			"2one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(3500000000),
			false,
		},
		{
			"2one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(4000000000),
			false,
		},
		{
			"2one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(4500000000),
			false,
		},
		{
			"3one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"3one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(333333333),
			false,
		},
		{
			"3one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(666666666),
			false,
		},
		{
			"3one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(999999999),
			false,
		},
		{
			"3one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(1333333332),
			false,
		},
		{
			"3one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(1666666665),
			false,
		},
		{
			"3one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(1999999998),
			false,
		},
		{
			"3one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(2333333331),
			false,
		},
		{
			"3one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(2666666664),
			false,
		},
		{
			"3one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(2999999997),
			false,
		},
		{
			"4one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"4one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(250000000),
			false,
		},
		{
			"4one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(500000000),
			false,
		},
		{
			"4one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(750000000),
			false,
		},
		{
			"4one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"4one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(1250000000),
			false,
		},
		{
			"4one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(1500000000),
			false,
		},
		{
			"4one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(1750000000),
			false,
		},
		{
			"4one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(2000000000),
			false,
		},
		{
			"4one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(2250000000),
			false,
		},
		{
			"5one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"5one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(200000000),
			false,
		},
		{
			"5one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(400000000),
			false,
		},
		{
			"5one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(600000000),
			false,
		},
		{
			"5one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(800000000),
			false,
		},
		{
			"5one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"5one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(1200000000),
			false,
		},
		{
			"5one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(1400000000),
			false,
		},
		{
			"5one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(1600000000),
			false,
		},
		{
			"5one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(1800000000),
			false,
		},
		{
			"6one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"6one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(166666666),
			false,
		},
		{
			"6one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(333333332),
			false,
		},
		{
			"6one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(499999998),
			false,
		},
		{
			"6one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(666666664),
			false,
		},
		{
			"6one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(833333330),
			false,
		},
		{
			"6one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(999999996),
			false,
		},
		{
			"6one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(1166666662),
			false,
		},
		{
			"6one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(1333333328),
			false,
		},
		{
			"6one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(1499999994),
			false,
		},
		{
			"7one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"7one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(142857142),
			false,
		},
		{
			"7one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(285714284),
			false,
		},
		{
			"7one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(428571426),
			false,
		},
		{
			"7one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(571428568),
			false,
		},
		{
			"7one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(714285710),
			false,
		},
		{
			"7one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(857142852),
			false,
		},
		{
			"7one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(999999994),
			false,
		},
		{
			"7one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(1142857136),
			false,
		},
		{
			"7one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(1285714278),
			false,
		},
		{
			"8one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"8one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(125000000),
			false,
		},
		{
			"8one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(250000000),
			false,
		},
		{
			"8one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(375000000),
			false,
		},
		{
			"8one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(500000000),
			false,
		},
		{
			"8one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(625000000),
			false,
		},
		{
			"8one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(750000000),
			false,
		},
		{
			"8one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(875000000),
			false,
		},
		{
			"8one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"8one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(1125000000),
			false,
		},
		{
			"9one gigabyte_prices and 0one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"9one gigabyte_prices and 1one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(111111111),
			false,
		},
		{
			"9one gigabyte_prices and 2one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(222222222),
			false,
		},
		{
			"9one gigabyte_prices and 3one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(333333333),
			false,
		},
		{
			"9one gigabyte_prices and 4one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(444444444),
			false,
		},
		{
			"9one gigabyte_prices and 5one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(555555555),
			false,
		},
		{
			"9one gigabyte_prices and 6one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(666666666),
			false,
		},
		{
			"9one gigabyte_prices and 7one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(777777777),
			false,
		},
		{
			"9one gigabyte_prices and 8one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(888888888),
			false,
		},
		{
			"9one gigabyte_prices and 9one coin",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(999999999),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				GigabytePrices: tt.fields.GigabytePrices,
			}
			got, err := m.Bytes(tt.args.coin)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.NodeAddress
	}{
		{
			"empty",
			fields{
				Address: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			hubtypes.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				Address: tt.fields.Address,
			}
			if got := m.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GigabytePrice(t *testing.T) {
	type fields struct {
		GigabytePrices sdk.Coins
	}
	type args struct {
		denom string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sdk.Coin
		want1  bool
	}{
		{
			"nil gigabyte_prices and empty denom",
			fields{
				GigabytePrices: nil,
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"empty gigabyte_prices and empty denom",
			fields{
				GigabytePrices: sdk.Coins{},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one gigabyte_prices and empty denom",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one gigabyte_prices and one denom",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "one",
			},
			sdk.NewInt64Coin("one", 1),
			true,
		},
		{
			"1one gigabyte_prices and two denom",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "two",
			},
			sdk.Coin{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				GigabytePrices: tt.fields.GigabytePrices,
			}
			got, got1 := m.GigabytePrice(tt.args.denom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GigabytePrice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GigabytePrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_HourlyPrice(t *testing.T) {
	type fields struct {
		HourlyPrices sdk.Coins
	}
	type args struct {
		denom string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sdk.Coin
		want1  bool
	}{
		{
			"nil hourly_prices and empty denom",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"empty hourly_prices and empty denom",
			fields{
				HourlyPrices: sdk.Coins{},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one hourly_prices and empty denom",
			fields{
				HourlyPrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one hourly_prices and one denom",
			fields{
				HourlyPrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "one",
			},
			sdk.NewInt64Coin("one", 1),
			true,
		},
		{
			"1one hourly_prices and two denom",
			fields{
				HourlyPrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "two",
			},
			sdk.Coin{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				HourlyPrices: tt.fields.HourlyPrices,
			}
			got, got1 := m.HourlyPrice(tt.args.denom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HourlyPrice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HourlyPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_Validate(t *testing.T) {
	type fields struct {
		Address        string
		GigabytePrices sdk.Coins
		HourlyPrices   sdk.Coins
		RemoteURL      string
		Status         hubtypes.Status
		StatusAt       int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty address",
			fields{
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				Address: "sentnode",
			},
			true,
		},
		{
			"invalid address prefix",
			fields{
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgse4wwrm",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"nil gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"empty gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"nil hourly_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"empty hourly_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{},
			},
			true,
		},
		{
			"empty denom hourly_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom hourly_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative hourly_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero hourly_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive hourly_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"empty remote_url",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			true,
		},
		{
			"length 72 remote_url",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"unspecified status",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusUnspecified,
				StatusAt:       1,
			},
			true,
		},
		{
			"active status",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
		{
			"inactive_pending status",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusInactivePending,
				StatusAt:       1,
			},
			true,
		},
		{
			"inactive status",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusInactive,
				StatusAt:       1,
			},
			false,
		},
		{
			"negative status_at",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       -1,
			},
			true,
		},
		{
			"zero status_at",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       0,
			},
			true,
		},
		{
			"positive status_at",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				Address:        tt.fields.Address,
				GigabytePrices: tt.fields.GigabytePrices,
				HourlyPrices:   tt.fields.HourlyPrices,
				RemoteURL:      tt.fields.RemoteURL,
				Status:         tt.fields.Status,
				StatusAt:       tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
