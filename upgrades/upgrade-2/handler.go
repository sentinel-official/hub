package upgrade2

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	custommintkeeper "github.com/sentinel-official/hub/x/mint/keeper"
	customminttypes "github.com/sentinel-official/hub/x/mint/types"
)

func Handler(
	setStoreLoader func(baseapp.StoreLoader),
	accountKeeper authkeeper.AccountKeeper,
	upgradeKeeper upgradekeeper.Keeper,
	customMintKeeper custommintkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	info, err := upgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if info.Name == Name && !upgradeKeeper.IsSkipHeight(info.Height) {
		upgrades := &storetypes.StoreUpgrades{
			Added: []string{customminttypes.ModuleName},
		}

		setStoreLoader(
			upgradetypes.UpgradeStoreLoader(
				info.Height,
				upgrades,
			),
		)
	}

	return func(ctx sdk.Context, _ upgradetypes.Plan) {
		if err := setInflations(ctx, customMintKeeper); err != nil {
			panic(err)
		}

		if err := updateVestingAccounts(ctx, accountKeeper); err != nil {
			panic(err)
		}
	}
}

func setInflations(ctx sdk.Context, k custommintkeeper.Keeper) error {
	var (
		inflations = []customminttypes.Inflation{
			{
				Max:        sdk.NewDecWithPrec(49, 2),
				Min:        sdk.NewDecWithPrec(43, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2021, 9, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(43, 2),
				Min:        sdk.NewDecWithPrec(37, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2022, 3, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(37, 2),
				Min:        sdk.NewDecWithPrec(31, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2022, 9, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(31, 2),
				Min:        sdk.NewDecWithPrec(25, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2023, 3, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(25, 2),
				Min:        sdk.NewDecWithPrec(19, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2023, 9, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(19, 2),
				Min:        sdk.NewDecWithPrec(13, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2024, 3, 27, 0, 0, 0, 0, time.UTC),
			},
		}
	)

	for _, inflation := range inflations {
		if err := inflation.Validate(); err != nil {
			return err
		}
	}
	for _, inflation := range inflations {
		k.SetInflation(ctx, inflation)
	}

	return nil
}

func updateVestingAccounts(ctx sdk.Context, accountKeeper authkeeper.AccountKeeper) error {
	var (
		foundation = "sent1vv8kmwrs24j5emzw8dp7k8satgea62l7knegd7"
		investors  = []string{
			"sent10lqg96mxv26k4lfcagytj7qql65nd068kpa43q",
			"sent10ven70nt35qr3yjgcl7psustkmrfu35qdnxa6z",
			"sent10y4eael9fpgm7pyarc7tzdfhe5n2fu4920sxel",
			"sent13ggn65j009s3n37ekl5whlw80qh6vhmwsqe7uz",
			"sent14am4gt4j3xa3746nc0vxx0ksej8yn8tx924exg",
			"sent15gf8mpdxmkkfxl2a4mffe2l0vn08jaseztj6c5",
			"sent15ptpn8qqjrgv69ayk26ap33rmkeysz7qux803h",
			"sent17980ady5mgq9tcuyj6cwfvqt2duk7wyxdy26cf",
			"sent189lj5prumzunezczcay0yy4tp88md6crmn6s9j",
			"sent18dsg7kc3f56c0fnelqzhag4hk6snjv0rtxzspe",
			"sent1ajzdnrefxnu9d9v5eh3ucn99de96zxx6muzw5a",
			"sent1au4sj50sddjqaj3mg9t62dgzk7rw34p26tpg6r",
			"sent1exwnrcjqc8d776v8sxevt47hvzgehlqct39elz",
			"sent1ey2k5azqkshjnsugrn2xtjfdh7m0wrqehugjxr",
			"sent1fa86yfpdkvj88a8p95pp3fmwsttcyg8wmy9vzm",
			"sent1gvvfs5raxje8l60mg96shgny88390wyz3qmcfr",
			"sent1lkjg98vpqrnr3m0n0dkdkhsmylymdt4r68u0zj",
			"sent1lml57rkh0kmqan6cnj20d66f8madwc6e8397qu",
			"sent1mc8m5vz62e7ymx20alr45swlmk0ta823k2xxz0",
			"sent1n85kmtr94q54u0plcpuxt5lrzz4j09l69m9z3r",
			"sent1pg6y9kn3jhuyzkhw5nn7jnh4mfmdzdnmpzn62a",
			"sent1q8xzh6qgww5upzmwvkztmm4x0zf3n8275mp6vq",
			"sent1r2el34ydppd2ugrlt9cstv0cl34yqwalwg7zm9",
			"sent1reeywehc76ndcd3gkzaz3yasdk7nfxc06undq8",
			"sent1s88cq7g7geukqccwrw7t6kaq8csq00na3x708t",
			"sent1t4lq5fy20tane82vfyqlc700g0q8e3d9a9kkmq",
			"sent1u2z7t0wcv0w9huctaat4jqed3kl790ueuvwkp8",
			"sent1uddk4mfqq3uyu2yzhym8spheqatqs30fzh5tmr",
			"sent1ueuncjn7devskqjjweyjhg6fvqngrzrnp8vq8n",
			"sent1uysdw03gz84glddm82cgddy7jza6aeqsq2rru9",
			"sent1vrt9zgl3u6yz55hssy8a2zu8u0c822u6n3xl29",
			"sent1wcszkf3psqqrt4llzt0y3lahsypa5s4d0r80nf",
			"sent1xp9dg7vejefvhx8nr8uedkagnk9k5muq7znrhk",
			"sent1zmqe22g8k4jnqq82urfayftq4xzkkvf3rxmhet",
			"sent1zx0sc9x8j6enwm7405yyk9xmwkzx7da0unu2kw",
			"sent1zyjlu0s8tl3hvugfyar8evsyux8xt70rv8027j",
		}
	)

	for _, s := range investors {
		account, err := getContinuousVestingAccount(ctx, accountKeeper, s)
		if err != nil {
			return err
		}

		account.StartTime = time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC).Unix()
		account.BaseVestingAccount.EndTime = time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC).Unix()
		accountKeeper.SetAccount(ctx, account)
	}

	account, err := getPeriodicVestingAccount(ctx, accountKeeper, foundation)
	if err != nil {
		return err
	}

	var (
		lengths = []int64{
			2678400, 2592000, 2678400, 2592000, 2678400, 2678400, 8294400, 2678400, 2419200, 2678400,
			2592000, 2678400, 2592000, 2678400, 2678400, 2592000, 2678400, 2592000, 2678400, 2678400,
			2419200, 2678400, 2592000, 2678400, 2592000, 2678400, 2678400, 2592000, 2678400, 2592000,
			2678400, 2678400, 2505600, 2678400, 2592000, 2678400, 2592000, 2678400, 2678400, 2592000,
			2678400, 2592000, 2678400, 2678400, 2419200, 2678400, 2592000, 2678400, 2592000, 2678400,
			2678400, 2592000, 2678400, 2592000, 2678400, 2678400, 2419200, 2678400, 2592000, 2678400,
		}
		amounts = []int64{
			110000000, 110000000, 55670430, 55670430, 55670430, 81536090, 56536090, 56536090, 81536090, 56536090,
			56536090, 129136090, 78440650, 78440650, 103440650, 78440650, 78440650, 102574990, 77574990, 77574990,
			102574990, 77574990, 47600000, 93900000, 68900000, 68900000, 98900000, 73900000, 73900000, 73900000,
			73900000, 73900000, 73900000, 73900000, 73900000, 136500000, 136500000, 136500000, 136500000, 136500000,
			136500000, 136500000, 136500000, 136500000, 136500000, 136500000, 136500000, 176999999, 176999999, 176999999,
			176999999, 176999999, 176999999, 176999999, 176999999, 176999999, 176999999, 176999999, 176999999, 35000000,
		}
	)

	account.VestingPeriods = make(vestingtypes.Periods, 0, len(lengths))
	account.BaseVestingAccount.EndTime = account.StartTime

	for i := 0; i < len(lengths); i++ {
		account.VestingPeriods = append(
			account.VestingPeriods,
			vestingtypes.Period{
				Length: lengths[i],
				Amount: sdk.NewCoins(sdk.NewCoin("udvpn", sdk.NewInt(amounts[i]).MulRaw(1e6))),
			},
		)
		account.EndTime += lengths[i]
	}

	accountKeeper.SetAccount(ctx, account)

	return nil
}

func getAccount(ctx sdk.Context, k authkeeper.AccountKeeper, s string) (authtypes.AccountI, error) {
	address, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		return nil, err
	}

	account := k.GetAccount(ctx, address)
	if account == nil {
		return nil, fmt.Errorf("account for address %s does not exist", s)
	}

	return account, nil
}

func getContinuousVestingAccount(ctx sdk.Context, k authkeeper.AccountKeeper, s string) (*vestingtypes.ContinuousVestingAccount, error) {
	base, err := getAccount(ctx, k, s)
	if err != nil {
		return nil, err
	}

	account, ok := base.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		return nil, fmt.Errorf("account for address %s is not a countinuous vesting type", s)
	}

	return account, nil
}

func getPeriodicVestingAccount(ctx sdk.Context, k authkeeper.AccountKeeper, s string) (*vestingtypes.PeriodicVestingAccount, error) {
	base, err := getAccount(ctx, k, s)
	if err != nil {
		return nil, err
	}

	account, ok := base.(*vestingtypes.PeriodicVestingAccount)
	if !ok {
		return nil, fmt.Errorf("account for address %s is not a periodic vesting type", s)
	}

	return account, nil
}
