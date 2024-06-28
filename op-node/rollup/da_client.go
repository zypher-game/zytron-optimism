package rollup

import (
	"fmt"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
)

type DAClient struct {
	Client da.DA
}

func NewDAClient(rpc, authToken string) (*DAClient, error) {
	fmt.Printf("NewDAClient rpc: %v \n", rpc)
	client, err := proxy.NewClient(rpc, authToken)
	if err != nil {
		return nil, err
	}
	return &DAClient{
		Client: client,
	}, nil
}
