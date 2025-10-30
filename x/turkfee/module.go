package turkfee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type FeeDecorator struct{}

func (fd FeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	stakingKeeper := ctx.KVStore(ctx.Key("staking"))
	
	// Basit örnek: Validatör adresi ya da belirli stake sahipleri gas ödemez.
	for _, msg := range tx.GetMsgs() {
		signer := msg.GetSigners()[0]
		if isValidatorOrStaker(signer, stakingKeeper) {
			ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter()) // gas-free
		}
	}
	return next(ctx, tx, simulate)
}

func isValidatorOrStaker(addr sdk.AccAddress, store sdk.KVStore) bool {
	// Basit pseudo kontrol (ileride staking modülüyle bağlanacak)
	key := []byte("val_" + addr.String())
	return store.Has(key)
}
