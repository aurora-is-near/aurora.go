package types

// BeginBlockArgs contains parameters for the `begin_block` method
type BeginBlockArgs struct {
	Hash       RawU256
	Coinbase   RawAddress
	Timestamp  RawU256
	Number     RawU256
	Difficulty RawU256
	GasLimit   RawU256
}

// BeginChainArgs contains parameters for the `begin_chain` method
type BeginChainArgs struct {
	ChainID      RawU256
	GenesisAlloc []AccountBalance
}

// FunctionCallArgs contains parameters for the `call` method
type FunctionCallArgs struct {
	Contract RawAddress
	Input    []uint8
}

// GetChainIDArgs contains parameters for the `get_chain_id` method
type GetChainIDArgs struct {
}

// GetStorageAtArgs contains parameters for the `get_storage_at` method
type GetStorageAtArgs struct {
	Address RawAddress
	Key     RawH256
}

// MetaCallArgs contains parameters for the `meta_call` method
type MetaCallArgs struct {
	Signature       [64]uint8
	V               uint8
	Nonce           RawU256
	FeeAmount       RawU256
	FeeAddress      RawAddress
	ContractAddress RawAddress
	Value           RawU256
	MethodDef       string
	Args            []uint8
}

// NewCallArgs contains parameters for the `new` method
type NewCallArgs struct {
	ChainID            RawU256
	OwnerID            AccountID
	BridgeProverID     AccountID
	UpgradeDelayBlocks uint64
}

// InitCallArgs contains parameters for the `new_eth_connector` method
type InitCallArgs struct {
	ProverAccount       AccountID
	EthCustodianAddress string
	Metadata            FungibleTokenMetadata
}

// ViewCallArgs contains parameters for the `view` method
type ViewCallArgs struct {
	Sender  RawAddress
	Address RawAddress
	Amount  RawU256
	Input   []uint8
}
