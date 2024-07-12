package registry

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// MintVehicleSign(uint256 manufacturerNode,address owner,string[] attributes,string[] infos,uint256 nonce)
type MintVehicleSign struct {
	ManufacturerNode *big.Int       `json:"manufacturerNode"`
	Owner            common.Address `json:"owner"`
	Attributes       []string       `json:"attributes"`
	Infos            []string       `json:"infos"`
	Nonce            *big.Int       `json:"nonce"`
}

func (m *MintVehicleSign) Name() string { return "MintVehicleSign" }

func (m *MintVehicleSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "manufacturerNode", Type: "uint256"}, {Name: "owner", Type: "address"}, {Name: "attributes", Type: "string[]"}, {Name: "infos", Type: "string[]"}, {Name: "nonce", Type: "uint256"}}
}

func (m *MintVehicleSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"manufacturerNode": hexutil.EncodeBig(m.ManufacturerNode), "owner": m.Owner.Hex(), "attributes": anySlice(m.Attributes), "infos": anySlice(m.Infos), "nonce": hexutil.EncodeBig(m.Nonce)}
}

func (m *MintVehicleSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, m.Name(): m.Type(),
		},
		PrimaryType: m.Name(),
		Domain:      domain,
		Message:     m.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

// UnPairAftermarketDeviceSign(uint256 aftermarketDeviceNode,uint256 vehicleNode,uint256 nonce)
type UnPairAftermarketDeviceSign struct {
	AftermarketDeviceNode *big.Int `json:"aftermarketDeviceNode"`
	VehicleNode           *big.Int `json:"vehicleNode"`
	Nonce                 *big.Int `json:"nonce"`
}

func (u *UnPairAftermarketDeviceSign) Name() string { return "UnPairAftermarketDeviceSign" }

func (u *UnPairAftermarketDeviceSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "aftermarketDeviceNode", Type: "uint256"}, {Name: "vehicleNode", Type: "uint256"}, {Name: "nonce", Type: "uint256"}}
}

func (u *UnPairAftermarketDeviceSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"aftermarketDeviceNode": hexutil.EncodeBig(u.AftermarketDeviceNode), "vehicleNode": hexutil.EncodeBig(u.VehicleNode), "nonce": hexutil.EncodeBig(u.Nonce)}
}

func (u *UnPairAftermarketDeviceSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, u.Name(): u.Type(),
		},
		PrimaryType: u.Name(),
		Domain:      domain,
		Message:     u.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

// ClaimAftermarketDeviceSign(uint256 aftermarketDeviceNode,address owner,uint256 nonce)
type ClaimAftermarketDeviceSign struct {
	AftermarketDeviceNode *big.Int       `json:"aftermarketDeviceNode"`
	Owner                 common.Address `json:"owner"`
	Nonce                 *big.Int       `json:"nonce"`
}

func (c *ClaimAftermarketDeviceSign) Name() string { return "ClaimAftermarketDeviceSign" }

func (c *ClaimAftermarketDeviceSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "aftermarketDeviceNode", Type: "uint256"}, {Name: "owner", Type: "address"}, {Name: "nonce", Type: "uint256"}}
}

func (c *ClaimAftermarketDeviceSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"aftermarketDeviceNode": hexutil.EncodeBig(c.AftermarketDeviceNode), "owner": c.Owner.Hex(), "nonce": hexutil.EncodeBig(c.Nonce)}
}

func (c *ClaimAftermarketDeviceSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, c.Name(): c.Type(),
		},
		PrimaryType: c.Name(),
		Domain:      domain,
		Message:     c.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

// PairAftermarketDeviceSign(uint256 aftermarketDeviceNode,uint256 vehicleNode,uint256 nonce)
type PairAftermarketDeviceSign struct {
	AftermarketDeviceNode *big.Int `json:"aftermarketDeviceNode"`
	VehicleNode           *big.Int `json:"vehicleNode"`
	Nonce                 *big.Int `json:"nonce"`
}

func (p *PairAftermarketDeviceSign) Name() string { return "PairAftermarketDeviceSign" }

func (p *PairAftermarketDeviceSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "aftermarketDeviceNode", Type: "uint256"}, {Name: "vehicleNode", Type: "uint256"}, {Name: "nonce", Type: "uint256"}}
}

func (p *PairAftermarketDeviceSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"aftermarketDeviceNode": hexutil.EncodeBig(p.AftermarketDeviceNode), "vehicleNode": hexutil.EncodeBig(p.VehicleNode), "nonce": hexutil.EncodeBig(p.Nonce)}
}

func (p *PairAftermarketDeviceSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, p.Name(): p.Type(),
		},
		PrimaryType: p.Name(),
		Domain:      domain,
		Message:     p.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

// MintSyntheticDeviceSign(uint256 integrationNode,uint256 vehicleNode,uint256 nonce)
type MintSyntheticDeviceSign struct {
	IntegrationNode *big.Int `json:"integrationNode"`
	VehicleNode     *big.Int `json:"vehicleNode"`
	Nonce           *big.Int `json:"nonce"`
}

func (m *MintSyntheticDeviceSign) Name() string { return "MintSyntheticDeviceSign" }

func (m *MintSyntheticDeviceSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "integrationNode", Type: "uint256"}, {Name: "vehicleNode", Type: "uint256"}, {Name: "nonce", Type: "uint256"}}
}

func (m *MintSyntheticDeviceSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"integrationNode": hexutil.EncodeBig(m.IntegrationNode), "vehicleNode": hexutil.EncodeBig(m.VehicleNode), "nonce": hexutil.EncodeBig(m.Nonce)}
}

func (m *MintSyntheticDeviceSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, m.Name(): m.Type(),
		},
		PrimaryType: m.Name(),
		Domain:      domain,
		Message:     m.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

// MintVehicleAndSdSign(uint256 integrationNode,uint256 nonce)
type MintVehicleAndSdSign struct {
	IntegrationNode *big.Int `json:"integrationNode"`
	Nonce           *big.Int `json:"nonce"`
}

func (m *MintVehicleAndSdSign) Name() string { return "MintVehicleAndSdSign" }

func (m *MintVehicleAndSdSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "integrationNode", Type: "uint256"}, {Name: "nonce", Type: "uint256"}}
}

func (m *MintVehicleAndSdSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"integrationNode": hexutil.EncodeBig(m.IntegrationNode), "nonce": hexutil.EncodeBig(m.Nonce)}
}

func (m *MintVehicleAndSdSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, m.Name(): m.Type(),
		},
		PrimaryType: m.Name(),
		Domain:      domain,
		Message:     m.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

// MintVehicleWithDeviceDefinitionSign(uint256 manufacturerNode,address owner,string deviceDefinitionId,string[] attributes,string[] infos,uint256 nonce)
type MintVehicleWithDeviceDefinitionSign struct {
	ManufacturerNode   *big.Int       `json:"manufacturerNode"`
	Owner              common.Address `json:"owner"`
	DeviceDefinitionId string         `json:"deviceDefinitionId"`
	Attributes         []string       `json:"attributes"`
	Infos              []string       `json:"infos"`
	Nonce              *big.Int       `json:"nonce"`
}

func (m *MintVehicleWithDeviceDefinitionSign) Name() string {
	return "MintVehicleWithDeviceDefinitionSign"
}

func (m *MintVehicleWithDeviceDefinitionSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "manufacturerNode", Type: "uint256"}, {Name: "owner", Type: "address"}, {Name: "deviceDefinitionId", Type: "string"}, {Name: "attributes", Type: "string[]"}, {Name: "infos", Type: "string[]"}, {Name: "nonce", Type: "uint256"}}
}

func (m *MintVehicleWithDeviceDefinitionSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"manufacturerNode": hexutil.EncodeBig(m.ManufacturerNode), "owner": m.Owner.Hex(), "deviceDefinitionId": m.DeviceDefinitionId, "attributes": anySlice(m.Attributes), "infos": anySlice(m.Infos), "nonce": hexutil.EncodeBig(m.Nonce)}
}

func (m *MintVehicleWithDeviceDefinitionSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, m.Name(): m.Type(),
		},
		PrimaryType: m.Name(),
		Domain:      domain,
		Message:     m.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

// BurnSyntheticDeviceSign(uint256 vehicleNode,uint256 syntheticDeviceNode,uint256 nonce)
type BurnSyntheticDeviceSign struct {
	VehicleNode         *big.Int `json:"vehicleNode"`
	SyntheticDeviceNode *big.Int `json:"syntheticDeviceNode"`
	Nonce               *big.Int `json:"nonce"`
}

func (b *BurnSyntheticDeviceSign) Name() string { return "BurnSyntheticDeviceSign" }

func (b *BurnSyntheticDeviceSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "vehicleNode", Type: "uint256"}, {Name: "syntheticDeviceNode", Type: "uint256"}, {Name: "nonce", Type: "uint256"}}
}

func (b *BurnSyntheticDeviceSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"vehicleNode": hexutil.EncodeBig(b.VehicleNode), "syntheticDeviceNode": hexutil.EncodeBig(b.SyntheticDeviceNode), "nonce": hexutil.EncodeBig(b.Nonce)}
}

func (b *BurnSyntheticDeviceSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, b.Name(): b.Type(),
		},
		PrimaryType: b.Name(),
		Domain:      domain,
		Message:     b.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

// BurnVehicleSign(uint256 vehicleNode,uint256 nonce)
type BurnVehicleSign struct {
	VehicleNode *big.Int `json:"vehicleNode"`
	Nonce       *big.Int `json:"nonce"`
}

func (b *BurnVehicleSign) Name() string { return "BurnVehicleSign" }

func (b *BurnVehicleSign) Type() []apitypes.Type {
	return []apitypes.Type{{Name: "vehicleNode", Type: "uint256"}, {Name: "nonce", Type: "uint256"}}
}

func (b *BurnVehicleSign) Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{"vehicleNode": hexutil.EncodeBig(b.VehicleNode), "nonce": hexutil.EncodeBig(b.Nonce)}
}

func (b *BurnVehicleSign) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {
	td := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			}, b.Name(): b.Type(),
		},
		PrimaryType: b.Name(),
		Domain:      domain,
		Message:     b.Message()}
	hash, _, err := apitypes.TypedDataAndHash(*td)
	return hash, err
}

func anySlice[A any](v []A) []any {
	n := len(v)
	out := make([]any, n)
	for i := 0; i < n; i++ {
		out[i] = v[i]
	}

	return out
}
