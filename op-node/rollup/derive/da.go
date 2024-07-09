package derive

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"os"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/log"
)

// DerivationVersionCelestia is a byte marker for celestia references submitted
// to the batch inbox address as calldata.
// Mnemonic 0xce = celestia
// version 0xce references are encoded as:
// [8]byte block height ++ [32]byte commitment
// in little-endian encoding.
// see: https://github.com/rollkit/celestia-da/blob/1f2df375fd2fcc59e425a50f7eb950daa5382ef0/celestia.go#L141-L160
const DerivationVersionCelestia = 0xce

var daClient *rollup.DAClient
var NodeDaNamespace = []byte{}

func init() {
	daRpc := os.Getenv("OP_NODE_DA_RPC")
	if daRpc == "" {
		panic("da rpc is nil")
	}
	daAuthToken := os.Getenv("OP_BATCHER_DA_AUTH_TOKEN")
	if daAuthToken == "" {
		panic("da auth token is nil")
	}
	daNamespace := os.Getenv("OP_NODE_DA_NAMESPACE")
	if daNamespace == "" {
		panic("da namespace is nil")
	}

	b, err := hexutil.Decode(daNamespace)
	if err != nil {
		fmt.Printf("op node namespace decode err")
		panic(err)
	}

	NodeDaNamespace = b

	daClient, err = rollup.NewDAClient(daRpc, daAuthToken)
	if err != nil {
		log.Error("celestia: unable to create DA client", "rpc", daRpc, "err", err)
		panic(err)
	}

	fmt.Printf("rollup OP_NODE_DA_RPC: %v\n", daRpc)
	fmt.Printf("rollup OP_BATCHER_DA_AUTH_TOKEN: %v\n", daAuthToken)
	fmt.Printf("rollup OP_NODE_DA_NAMESPACE: %v\n", daNamespace)
}
