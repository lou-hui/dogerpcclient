// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: This file is intended to house the RPC commands that are supported by
// a chain server.

package dogejson

import (
	"github.com/btcsuite/btcd/btcjson"
)

// NewGetBlockCmd returns a new instance which can be used to issue a getblock
// JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetBlockCmd(hash string, verbose *bool) *btcjson.GetBlockCmd {
	return &btcjson.GetBlockCmd{
		Hash:      hash,
		Verbose:   verbose,
	}
}