package types

import (
	"reflect"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestMsgRegisterNode(t *testing.T) {
	type fields struct {
		From         csdkTypes.AccAddress
		AmountToLock csdkTypes.Coin
		PricesPerGB  csdkTypes.Coins
		NetSpeed     sdkTypes.Bandwidth
		APIPort      APIPort
		EncMethod    string
		NodeType     string
		Version      string
	}
	tests := []struct {
		name   string
		fields fields
		want   csdkTypes.Error
	}{
		{"Valid test case", fields{From: NodeAddress1, AmountToLock: Coin, PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos},
			APIPort:  NewAPIPort(1000), EncMethod: "Enc-Method-1", NodeType: "NodeType-1", Version: "0.01"}, nil},

		{"InValid test case : AmountToLock", fields{From: NodeAddress1, AmountToLock: CoinNeg, PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos},
			APIPort:  NewAPIPort(1000), EncMethod: "Enc-Method-1", NodeType: "NodeType-1", Version: "0.01"}, ErrorInvalidField("amount_to_lock")},

		{"InValid test case : PricePerGB", fields{From: NodeAddress1, AmountToLock: Coin, PricesPerGB: nil,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos},
			APIPort:  NewAPIPort(10000), EncMethod: "Enc-Method-1", NodeType: "NodeType-1", Version: "0.01"}, ErrorInvalidField("prices_per_gb")},

		{"Invalid test case : APIPort", fields{From: NodeAddress1, AmountToLock: Coin, PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos},
			APIPort:  APIPort(429496729), EncMethod: "Enc-Method-1", NodeType: "NodeType-1", Version: "0.01"}, ErrorInvalidField("api_port")},

		{"InValid test case : NetSpeed", fields{From: NodeAddress1, AmountToLock: Coin, PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadNeg},
			APIPort:  NewAPIPort(1000), EncMethod: "Enc-Method-1", NodeType: "NodeType-1", Version: "0.01"}, ErrorInvalidField("net_speed")},

		{"InValid test case : NodeType", fields{From: NodeAddress1, AmountToLock: Coin, PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos},
			APIPort:  NewAPIPort(1000), EncMethod: "Enc-Method-1", NodeType: "", Version: "0.01"}, ErrorInvalidField("node_type")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgRegisterNode{
				From:         tt.fields.From,
				AmountToLock: tt.fields.AmountToLock,
				PricesPerGB:  tt.fields.PricesPerGB,
				NetSpeed:     tt.fields.NetSpeed,
				APIPort:      tt.fields.APIPort,
				EncMethod:    tt.fields.EncMethod,
				NodeType:     tt.fields.NodeType,
				Version:      tt.fields.Version,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgRegisterNode.ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUpdateNodeDetails(t *testing.T) {
	type fields struct {
		From        csdkTypes.AccAddress
		ID          NodeID
		PricesPerGB csdkTypes.Coins
		NetSpeed    sdkTypes.Bandwidth
		APIPort     APIPort
		EncMethod   string
		Version     string
	}
	tests := []struct {
		name   string
		fields fields
		want   csdkTypes.Error
	}{
		{"InValid test case : ID", fields{From: NodeAddress1, ID: NewNodeID("new-node-id-1"), PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos,},
			APIPort:  NewAPIPort(10000), EncMethod: "Enc-Method1", Version: "0.01"}, ErrorInvalidField("id")},

		{"Valid test case", fields{From: NodeAddress1, ID: NewNodeID("new-node-id/0"), PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos,},
			APIPort:  NewAPIPort(10000), EncMethod: "Enc-Method1", Version: "0.01"}, nil},

		{"InValid test case : UploadNeg", fields{From: NodeAddress1, ID: NewNodeID("new-node-/id-1"), PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadNeg, Download: DownloadPos,},
			APIPort:  NewAPIPort(10000), EncMethod: "Enc-Method1", Version: "0.01"}, ErrorInvalidField("net_speed")},

		{"InValid test case : From", fields{From: nil, ID: NewNodeID("new-node-id/0"), PricesPerGB: Coins,
			NetSpeed: sdkTypes.Bandwidth{Upload: UploadPos, Download: DownloadPos,},
			APIPort:  NewAPIPort(10000), EncMethod: "Enc-Method1", Version: "0.01"}, ErrorInvalidField("from")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgUpdateNodeDetails{
				From:        tt.fields.From,
				ID:          tt.fields.ID,
				PricesPerGB: tt.fields.PricesPerGB,
				NetSpeed:    tt.fields.NetSpeed,
				APIPort:     tt.fields.APIPort,
				EncMethod:   tt.fields.EncMethod,
				Version:     tt.fields.Version,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgUpdateNodeDetails.ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUpdateNodeStatus(t *testing.T) {
	type fields struct {
		From   csdkTypes.AccAddress
		ID     NodeID
		Status string
	}
	tests := []struct {
		name   string
		fields fields
		want   csdkTypes.Error
	}{
		{"InValid test case : Status", fields{From: NodeAddress1, ID: NewNodeID("new-node/1"), Status: "REGISTERED"}, ErrorInvalidField("status")},
		{"InValid test case : ID", fields{From: NodeAddress1, ID: NewNodeID("new-node-1"), Status: StatusActive}, ErrorInvalidField("id")},
		{"InValid test case : From", fields{From: nil, ID: NewNodeID("new-node-1"), Status: "REGISTERED"}, ErrorInvalidField("from")},
		{"Valid test case ", fields{From: nil, ID: NewNodeID("new-node-1"), Status: StatusRegistered}, ErrorInvalidField("from")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgUpdateNodeStatus{
				From:   tt.fields.From,
				ID:     tt.fields.ID,
				Status: tt.fields.Status,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgUpdateNodeStatus.ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgDeregisterNode_ValidateBasic(t *testing.T) {
	type fields struct {
		From csdkTypes.AccAddress
		ID   NodeID
	}
	tests := []struct {
		name   string
		fields fields
		want   csdkTypes.Error
	}{
		{"InValid test case : ID", fields{From: NodeAddress1, ID: NewNodeID("new-node-id-1")}, ErrorInvalidField("id")},
		{"Valid test case", fields{From: NodeAddress1, ID: NewNodeID("new-node/1")}, nil},
		{"InValid test case : From", fields{From: nil, ID: NewNodeID("new-node/1")}, ErrorInvalidField("from")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgDeregisterNode{
				From: tt.fields.From,
				ID:   tt.fields.ID,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgDeregisterNode.ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}
