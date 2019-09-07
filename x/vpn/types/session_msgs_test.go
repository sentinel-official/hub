package types_test

import (
	"encoding/json"
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestMsgUpdateSessionInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgUpdateSessionInfo
		want sdk.Error
	}{
		{
			"from is nil",
			types.NewMsgUpdateSessionInfo(nil, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1),
			types.ErrorInvalidField("from"),
		}, {
			"from is empty",
			types.NewMsgUpdateSessionInfo([]byte(""), hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1),
			types.ErrorInvalidField("from"),
		}, {
			"bandwidth is zero",
			types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1),
			types.ErrorInvalidField("bandwidth"),
		}, {
			"bandwidth is neg",
			types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(-500000000), sdk.NewInt(-500000000)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1),
			types.ErrorInvalidField("bandwidth"),
		}, {
			"bandwidth is zero",
			types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1),
			types.ErrorInvalidField("bandwidth"),
		}, {
			"node owner sign is empty  ",
			types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), auth.StdSignature{}, clientStdSignaturePos1),
			types.ErrorInvalidField("node_owner_signature"),
		}, {
			"client sign is empty  ",
			types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), nodeOwnerStdSignaturePos1, auth.StdSignature{}),
			types.ErrorInvalidField("client_signature"),
		}, {
			"valid ",
			types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgUpdateSessionInfo_GetSignBytes(t *testing.T) {
	msg := types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateSessionInfo_GetSigners(t *testing.T) {
	msg := types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1)
	require.Equal(t, []sdk.AccAddress{address1}, msg.GetSigners())
}

func TestMsgUpdateSessionInfo_Type(t *testing.T) {
	msg := types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1)
	require.Equal(t, "update_session_info", msg.Type())
}

func TestMsgUpdateSessionInfo_Route(t *testing.T) {
	msg := types.NewMsgUpdateSessionInfo(address1, hub.NewIDFromUInt64(1), hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), nodeOwnerStdSignaturePos1, clientStdSignaturePos1)
	require.Equal(t, types.RouterKey, msg.Route())
}
