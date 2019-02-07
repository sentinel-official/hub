package types

import (
	"reflect"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestMsgInitSession(t *testing.T) {
	type fields struct {
		From         csdkTypes.AccAddress
		NodeID       NodeID
		AmountToLock csdkTypes.Coin
	}
	tests := []struct {
		name   string
		fields fields
		want   csdkTypes.Error
	}{
		{"InValid test case : NodeID", fields{From: ClientAddress1, NodeID: NewNodeID("new-node//1"), AmountToLock: Coin}, ErrorInvalidField("id")},
		{"InValid test case : From", fields{From: nil, NodeID: NewNodeID("new-node/1"), AmountToLock: Coin}, ErrorInvalidField("from")},
		{"Valid test case ", fields{From: ClientAddress2, NodeID: NewNodeID("new-node/1"), AmountToLock: Coin}, nil},
		{"InValid test case : Amount To Lock ", fields{From: ClientAddress2, NodeID: NewNodeID("new-node/1"), AmountToLock: CoinZero}, ErrorInvalidField("amount_to_lock")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgInitSession{
				From:         tt.fields.From,
				NodeID:       tt.fields.NodeID,
				AmountToLock: tt.fields.AmountToLock,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgInitSession.ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUpdateSessionBandwidth(t *testing.T) {
	type fields struct {
		From          csdkTypes.AccAddress
		SessionID     SessionID
		Bandwidth     sdkTypes.Bandwidth
		ClientSign    []byte
		NodeOwnerSign []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   csdkTypes.Error
	}{
		{"InValid test case : SessionID", fields{From: ClientAddress1, SessionID: StatusActive,
			Bandwidth:  sdkTypes.Bandwidth{},
			ClientSign: nil, NodeOwnerSign: nil}, ErrorInvalidField("session_id")},

		{"InValid test case : Bandwidth", fields{From: ClientAddress1, SessionID: NewSessionID("session/1"),
			Bandwidth:  sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadNeg},
			ClientSign: nil, NodeOwnerSign: nil}, ErrorInvalidField("bandwidth")},

		{"InValid test case : ClientSign", fields{From: ClientAddress1, SessionID: NewSessionID("session/2"),
			Bandwidth:  sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos},
			ClientSign: nil, NodeOwnerSign: nil}, ErrorInvalidField("client_sign")},

		{"Valid test case ", fields{From: ClientAddress1, SessionID: NewSessionID("session/1"),
			Bandwidth:  sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos},
			ClientSign: []byte("clientSignBytes"), NodeOwnerSign: []byte("clientSignBytes")}, nil},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgUpdateSessionBandwidth{
				From:          tt.fields.From,
				SessionID:     tt.fields.SessionID,
				Bandwidth:     tt.fields.Bandwidth,
				ClientSign:    tt.fields.ClientSign,
				NodeOwnerSign: tt.fields.NodeOwnerSign,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgUpdateSessionBandwidth.ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}
