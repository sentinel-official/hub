package querier

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

var (
	TestNodeParamsZero = QueryNodeParams{ID: types.TestIDZero}
	TestNodeParamsPos  = QueryNodeParams{ID: types.TestIDPos}

	TestNodeOfAddressParamsEmpty = QueryNodesOfAddressPrams{Address: types.TestAddressEmpty}
	TestNodeOfAddressParams1     = QueryNodesOfAddressPrams{Address: types.TestAddress1}
	TestNodeOfAddressParams2     = QueryNodesOfAddressPrams{Address: types.TestAddress2}

	TestSessionParamsZero = QuerySessionParams{ID: types.TestIDZero}
	TestSessionParamsPos  = QuerySessionParams{ID: types.TestIDPos}

	TestSessionOfSubscriptionPramsZero = QuerySessionOfSubscriptionPrams{types.TestIDZero, 0}
	TestSessionOfSubscriptionPramsPos  = QuerySessionOfSubscriptionPrams{types.TestIDPos, 0}

	TestSessionsOfSubscriptionPramsZero = QuerySessionsOfSubscriptionPrams{types.TestIDZero}
	TestSessionsOfSubscriptionPramsPos  = QuerySessionsOfSubscriptionPrams{types.TestIDPos}

	TestSubscriptionParamsZero = QuerySubscriptionParams{ID: types.TestIDZero}
	TestSubscriptionParamsPos  = QuerySubscriptionParams{ID: types.TestIDPos}

	TestSubscriptionsOfNodeParamsZero = QuerySubscriptionsOfNodePrams{types.TestIDZero}
	TestSubscriptionsOfNodeParamsPos  = QuerySubscriptionsOfNodePrams{types.TestIDPos}

	TestSubscriptionsOfAddressParamsEmpty = QuerySubscriptionsOfAddressParams{Address: types.TestAddressEmpty}
	TestSubscriptionsOfAddressParams1     = QuerySubscriptionsOfAddressParams{Address: types.TestAddress1}
	TestSubscriptionsOfAddressParams2     = QuerySubscriptionsOfAddressParams{Address: types.TestAddress2}

	TestSessionsCountOfSubscriptionParamsZero = QuerySessionsCountOfSubscriptionParams{types.TestIDZero}
	TestSessionsCountOfSubscriptionParamsPos  = QuerySessionsCountOfSubscriptionParams{types.TestIDPos}
)
