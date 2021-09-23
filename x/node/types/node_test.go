package types

import (
	"reflect"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestNode_BytesForCoin(t *testing.T) {
	type fields struct {
		Address   string
		Provider  string
		Price     sdk.Coins
		RemoteURL string
		Status    hubtypes.Status
		StatusAt  time.Time
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
			"nil price and empty coin",
			fields{
				Price: nil,
			},
			args{
				coin: sdk.Coin{},
			},
			sdk.NewInt(0),
			true,
		},
		{
			"empty price and empty coin",
			fields{
				Price: sdk.Coins{},
			},
			args{
				coin: sdk.Coin{},
			},
			sdk.NewInt(0),
			true,
		},
		{
			"1one price and empty coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.Coin{},
			},
			sdk.NewInt(0),
			true,
		},
		{
			"1one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"1one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"1one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(2000000000),
			false,
		},
		{
			"1one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(3000000000),
			false,
		},
		{
			"1one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(4000000000),
			false,
		},
		{
			"1one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(5000000000),
			false,
		},
		{
			"1one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(6000000000),
			false,
		},
		{
			"1one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(7000000000),
			false,
		},
		{
			"1one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(8000000000),
			false,
		},
		{
			"1one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(9000000000),
			false,
		},
		{
			"2one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"2one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(500000000),
			false,
		},
		{
			"2one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"2one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(1500000000),
			false,
		},
		{
			"2one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(2000000000),
			false,
		},
		{
			"2one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(2500000000),
			false,
		},
		{
			"2one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(3000000000),
			false,
		},
		{
			"2one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(3500000000),
			false,
		},
		{
			"2one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(4000000000),
			false,
		},
		{
			"2one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 2)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(4500000000),
			false,
		},
		{
			"3one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"3one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(333333333),
			false,
		},
		{
			"3one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(666666666),
			false,
		},
		{
			"3one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(999999999),
			false,
		},
		{
			"3one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(1333333332),
			false,
		},
		{
			"3one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(1666666665),
			false,
		},
		{
			"3one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(1999999998),
			false,
		},
		{
			"3one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(2333333331),
			false,
		},
		{
			"3one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(2666666664),
			false,
		},
		{
			"3one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 3)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(2999999997),
			false,
		},
		{
			"4one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"4one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(250000000),
			false,
		},
		{
			"4one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(500000000),
			false,
		},
		{
			"4one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(750000000),
			false,
		},
		{
			"4one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"4one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(1250000000),
			false,
		},
		{
			"4one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(1500000000),
			false,
		},
		{
			"4one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(1750000000),
			false,
		},
		{
			"4one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(2000000000),
			false,
		},
		{
			"4one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 4)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(2250000000),
			false,
		},
		{
			"5one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"5one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(200000000),
			false,
		},
		{
			"5one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(400000000),
			false,
		},
		{
			"5one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(600000000),
			false,
		},
		{
			"5one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(800000000),
			false,
		},
		{
			"5one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"5one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(1200000000),
			false,
		},
		{
			"5one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(1400000000),
			false,
		},
		{
			"5one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(1600000000),
			false,
		},
		{
			"5one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 5)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(1800000000),
			false,
		},
		{
			"6one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"6one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(166666666),
			false,
		},
		{
			"6one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(333333332),
			false,
		},
		{
			"6one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(499999998),
			false,
		},
		{
			"6one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(666666664),
			false,
		},
		{
			"6one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(833333330),
			false,
		},
		{
			"6one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(999999996),
			false,
		},
		{
			"6one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(1166666662),
			false,
		},
		{
			"6one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(1333333328),
			false,
		},
		{
			"6one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 6)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(1499999994),
			false,
		},
		{
			"7one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"7one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(142857142),
			false,
		},
		{
			"7one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(285714284),
			false,
		},
		{
			"7one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(428571426),
			false,
		},
		{
			"7one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(571428568),
			false,
		},
		{
			"7one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(714285710),
			false,
		},
		{
			"7one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(857142852),
			false,
		},
		{
			"7one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(999999994),
			false,
		},
		{
			"7one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(1142857136),
			false,
		},
		{
			"7one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 7)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(1285714278),
			false,
		},
		{
			"8one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"8one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(125000000),
			false,
		},
		{
			"8one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(250000000),
			false,
		},
		{
			"8one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(375000000),
			false,
		},
		{
			"8one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(500000000),
			false,
		},
		{
			"8one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(625000000),
			false,
		},
		{
			"8one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(750000000),
			false,
		},
		{
			"8one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(875000000),
			false,
		},
		{
			"8one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(1000000000),
			false,
		},
		{
			"8one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 8)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 9),
			},
			sdk.NewInt(1125000000),
			false,
		},
		{
			"9one price and 0one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 0),
			},
			sdk.NewInt(0),
			false,
		},
		{
			"9one price and 1one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 1),
			},
			sdk.NewInt(111111111),
			false,
		},
		{
			"9one price and 2one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 2),
			},
			sdk.NewInt(222222222),
			false,
		},
		{
			"9one price and 3one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 3),
			},
			sdk.NewInt(333333333),
			false,
		},
		{
			"9one price and 4one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 4),
			},
			sdk.NewInt(444444444),
			false,
		},
		{
			"9one price and 5one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 5),
			},
			sdk.NewInt(555555555),
			false,
		},
		{
			"9one price and 6one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 6),
			},
			sdk.NewInt(666666666),
			false,
		},
		{
			"9one price and 7one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 7),
			},
			sdk.NewInt(777777777),
			false,
		},
		{
			"9one price and 8one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
			},
			args{
				coin: sdk.NewInt64Coin("one", 8),
			},
			sdk.NewInt(888888888),
			false,
		},
		{
			"9one price and 9one coin",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 9)},
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
			n := &Node{
				Address:   tt.fields.Address,
				Provider:  tt.fields.Provider,
				Price:     tt.fields.Price,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				StatusAt:  tt.fields.StatusAt,
			}
			got, err := n.BytesForCoin(tt.args.coin)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesForCoin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesForCoin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GetAddress(t *testing.T) {
	type fields struct {
		Address   string
		Provider  string
		Price     sdk.Coins
		RemoteURL string
		Status    hubtypes.Status
		StatusAt  time.Time
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
			n := &Node{
				Address:   tt.fields.Address,
				Provider:  tt.fields.Provider,
				Price:     tt.fields.Price,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				StatusAt:  tt.fields.StatusAt,
			}
			if got := n.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GetProvider(t *testing.T) {
	type fields struct {
		Address   string
		Provider  string
		Price     sdk.Coins
		RemoteURL string
		Status    hubtypes.Status
		StatusAt  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.ProvAddress
	}{
		{
			"empty",
			fields{
				Provider: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			hubtypes.ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Address:   tt.fields.Address,
				Provider:  tt.fields.Provider,
				Price:     tt.fields.Price,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				StatusAt:  tt.fields.StatusAt,
			}
			if got := n.GetProvider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_PriceForDenom(t *testing.T) {
	type fields struct {
		Address   string
		Provider  string
		Price     sdk.Coins
		RemoteURL string
		Status    hubtypes.Status
		StatusAt  time.Time
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sdk.Coin
		want1  bool
	}{
		{
			"nil price and empty denom",
			fields{
				Price: nil,
			},
			args{
				s: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"empty price and empty denom",
			fields{
				Price: sdk.Coins{},
			},
			args{
				s: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one price and empty denom",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				s: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one price and one denom",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				s: "one",
			},
			sdk.NewInt64Coin("one", 1),
			true,
		},
		{
			"1one price and two denom",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				s: "two",
			},
			sdk.Coin{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Address:   tt.fields.Address,
				Provider:  tt.fields.Provider,
				Price:     tt.fields.Price,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				StatusAt:  tt.fields.StatusAt,
			}
			got, got1 := n.PriceForDenom(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceForDenom() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PriceForDenom() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_Validate(t *testing.T) {
	type fields struct {
		Address   string
		Provider  string
		Price     sdk.Coins
		RemoteURL string
		Status    hubtypes.Status
		StatusAt  time.Time
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
				Address: "invalid",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				Address: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"30 bytes address",
			fields{
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			true,
		},
		{
			"empty provider and nil price",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    nil,
			},
			true,
		},
		{
			"non-empty provider and non-nil price",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{},
			},
			true,
		},
		{
			"invalid prefix provider",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes provider",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			true,
		},
		{
			"20 bytes provider",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			true,
		},
		{
			"30 bytes provider",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			true,
		},
		{
			"empty price",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				Address:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"empty remote_url",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "",
			},
			true,
		},
		{
			"length 72 remote_url",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url",
			},
			true,
		},
		{
			"non-empty remote_url port",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
			},
			true,
		},
		{
			"unknown status",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusUnknown,
			},
			true,
		},
		{
			"active status",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusActive,
			},
			true,
		},
		{
			"inactive pending status",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusInactive,
			},
			true,
		},
		{
			"zero status_at",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusInactive,
				StatusAt:  time.Time{},
			},
			true,
		},
		{
			"now status_at",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusInactive,
				StatusAt:  time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Address:   tt.fields.Address,
				Provider:  tt.fields.Provider,
				Price:     tt.fields.Price,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				StatusAt:  tt.fields.StatusAt,
			}
			if err := n.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
