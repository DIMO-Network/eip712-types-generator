// TODO(elffjs): A lot of these break down if you have struct types with fields
// that are also struct types. We don't have any of these yet.
package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

const eip712DomainName = "EIP712Domain"

var eip712DomainType = []apitypes.Type{
	{Name: "name", Type: "string"},
	{Name: "version", Type: "string"},
	{Name: "chainId", Type: "uint256"},
	{Name: "verifyingContract", Type: "address"},
}

type Domain struct {
	Name              string
	Version           string
	ChainID           *big.Int
	VerifyingContract common.Address
}

type TypedData interface {
	Name() string
	Type() []apitypes.Type
	Message() apitypes.TypedDataMessage
}

type Client struct {
	domain apitypes.TypedDataDomain
}

func New(domain *Domain) *Client {
	return &Client{
		domain: apitypes.TypedDataDomain{
			Name:              domain.Name,
			Version:           domain.Version,
			ChainId:           (*math.HexOrDecimal256)(domain.ChainID),
			VerifyingContract: domain.VerifyingContract.Hex(),
		},
	}
}

func (c *Client) Display(data TypedData) apitypes.TypedData {
	return apitypes.TypedData{
		Types: apitypes.Types{
			eip712DomainName: eip712DomainType,
			data.Name():      data.Type(),
		},
		PrimaryType: data.Name(),
		Domain:      c.domain,
		Message:     data.Message(),
	}
}

func (c *Client) Hash(td TypedData) (common.Hash, error) {
	b, _, err := apitypes.TypedDataAndHash(c.Display(td))
	return common.BytesToHash(b), err
}
