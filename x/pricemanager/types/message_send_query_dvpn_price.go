package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

const TypeMsgSendQueryDVPNPrice = "send_query_dvpn_price"

var _ sdk.Msg = &MsgSendQueryDVPNPrice{}

func NewMsgSendQueryDVPNPrice(creator, channelId string, poolId string, baseAssetDenom string, quoteAssetDenom string, page *query.PageRequest) *MsgSendQueryDVPNPrice {
	return &MsgSendQueryDVPNPrice{
		Creator:         creator,
		ChannelId:       channelId,
		PoolId:          poolId,
		BaseAssetDenom:  baseAssetDenom,
		QuoteAssetDenom: quoteAssetDenom,
		Pagination:      page,
	}
}

func (msg *MsgSendQueryDVPNPrice) Route() string {
	return RouterKey
}

func (msg *MsgSendQueryDVPNPrice) Type() string {
	return TypeMsgSendQueryDVPNPrice
}

func (msg *MsgSendQueryDVPNPrice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendQueryDVPNPrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendQueryDVPNPrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
