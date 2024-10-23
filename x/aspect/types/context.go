package types

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/artela-network/artela-evm/vm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type storeContext struct {
	cosmosCtx    sdk.Context
	storeService store.KVStoreService
	gas          uint64

	chargeGas bool
}

func (s *storeContext) clone() StoreContext {
	return &storeContext{
		cosmosCtx:    s.cosmosCtx,
		storeService: s.storeService,
		gas:          s.gas,
	}
}

func (s *storeContext) Logger() log.Logger {
	return s.cosmosCtx.Logger()
}

func (s *storeContext) CosmosContext() sdk.Context {
	return s.cosmosCtx
}

func (s *storeContext) StoreService() store.KVStoreService {
	return s.StoreService()
}

func (s *storeContext) Gas() uint64 {
	return s.gas
}

func (s *storeContext) UpdateGas(gas uint64) {
	s.gas = gas
}

func (s *storeContext) ChargeGas() bool {
	return s.chargeGas
}

func (s *storeContext) ConsumeGas(gas uint64) error {
	if s.gas < gas {
		s.gas = 0
		return vm.ErrOutOfGas
	}
	s.gas -= gas
	return nil
}

type StoreContext interface {
	CosmosContext() sdk.Context
	Gas() uint64
	ConsumeGas(gas uint64) error
	UpdateGas(gas uint64)
	Logger() log.Logger
	ChargeGas() bool

	clone() StoreContext
}

func NewStoreContext(ctx sdk.Context, storeService store.KVStoreService, gas uint64) StoreContext {
	return &storeContext{
		cosmosCtx:    ctx,
		storeService: storeService,
		gas:          gas,
		chargeGas:    true,
	}
}

func NewGasFreeStoreContext(ctx sdk.Context, storeService store.KVStoreService) StoreContext {
	return &storeContext{
		cosmosCtx:    ctx,
		storeService: storeService,
		chargeGas:    false,
	}
}

type AccountStoreContext struct {
	StoreContext
	Account common.Address
}

func (a *AccountStoreContext) Clone() AccountStoreContext {
	return AccountStoreContext{
		StoreContext: a.StoreContext.clone(),
		Account:      a.Account,
	}
}

type AspectStoreContext struct {
	StoreContext
	AspectID common.Address
}

func (a *AspectStoreContext) Clone() AspectStoreContext {
	return AspectStoreContext{
		StoreContext: a.StoreContext.clone(),
		AspectID:     a.AspectID,
	}
}
