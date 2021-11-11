package types

import "github.com/near/borsh-go"

// TransactionStatusEnum enumerates transaction statuses
type TransactionStatusEnum borsh.Enum

// Existing transaction statuses
const (
	SuccessStatus TransactionStatusEnum = iota
	RevertStatus
	OutOfGas
	OutOfFund
	OutOfOffset
	CallTooDeep
)

// TransactionStatus describes status of a transaction
type TransactionStatus struct {
	Enum        TransactionStatusEnum `borsh_enum:"true"`
	Success     []uint8
	Revert      []uint8
	OutOfGas    struct{}
	OutOfFund   struct{}
	OutOfOffset struct{}
	CallTooDeep struct{}
}
