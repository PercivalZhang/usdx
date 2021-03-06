package peg

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "peg" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgXrpTx:
			return handleMsgXrpTx(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to mint pXrp
func handleMsgXrpTx(ctx sdk.Context, keeper Keeper, msg MsgXrpTx) sdk.Result {

	// Request tx hash from ripple api
	txData, err := keeper.fetchXrpTransactionData(msg.TxHash)
	if err != nil {
		return err.Result()
	}
	// throw if transaction not ok
	ok := keeper.isPxrpMultisgTransaction(txData)
	if ok == false {
		return sdk.ErrInternal("Tx is not to validator multisig address").Result()
	}
	ok = keeper.hasValidMemoData(txData)
	if ok == false {
		return sdk.ErrInternal("Memo is invalid").Result()
	}

	// Mint new pxrp
	tags, err := keeper.mintPxrp(ctx, txData)
	if err != nil {
		return sdk.ErrInternal("Failed to mint pxrp").Result()
	}
	return sdk.Result{
		Tags: tags,
	}
}
