package types

import "github.com/near/borsh-go"

type TransactionStatusEnum borsh.Enum

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
