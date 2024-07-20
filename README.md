### EIP-712 Types Generator

Generate Go types from an [EIP-712](https://eips.ethereum.org/EIPS/eip-712) types JSON.

```
go run ./cmd/eip-712-types-generator ../dimo-identity/utils/constants/eip712.ts
```

Currently only supports the member types that DIMO uses:

* `address`
* `string`
* `string[]`
* `uint256`
