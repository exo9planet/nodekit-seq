// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"

	"github.com/AnomalyFi/hypersdk/crypto"
	"github.com/AnomalyFi/nodekit-seq/genesis"
	"github.com/AnomalyFi/nodekit-seq/orderbook"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/trace"
)

type Controller interface {
	Genesis() *genesis.Genesis
	Tracer() trace.Tracer
	GetTransaction(context.Context, ids.ID) (bool, int64, bool, uint64, error)
	GetAssetFromState(context.Context, ids.ID) (bool, []byte, uint64, crypto.PublicKey, bool, error)
	GetBalanceFromState(context.Context, crypto.PublicKey, ids.ID) (uint64, error)
	Orders(pair string, limit int) []*orderbook.Order
	GetLoanFromState(context.Context, ids.ID, ids.ID) (uint64, error)
}
