package types

// RawAddress is a 20-bytes long account address in Ethereum
type RawAddress [20]uint8

// RawU256 is a big-endian large integer type
type RawU256 [32]uint8

// RawH256 is an unformatted binary data of fixed length
type RawH256 [32]uint8

// AccountID is a string
type AccountID string

// AccountBalance is a pair address:balance
type AccountBalance struct {
	Address RawAddress
	Balance RawU256
}
