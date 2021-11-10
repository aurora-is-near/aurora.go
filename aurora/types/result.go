package types

import (
	"fmt"

	"github.com/near/borsh-go"
)

// LogEvent is for use in LegacyExecutionResult
type LogEvent struct {
	Topics []RawU256
	Data   []uint8
}

// LogEventWithAddress is for use in a latest SubmitResult
type LogEventWithAddress struct {
	Address RawAddress
	Topics  []RawU256
	Data    []uint8
}

// LegacyExecutionResult is an old version of result from the `submit` method
type LegacyExecutionResult struct {
	Status  bool
	GasUsed uint64
	Output  []uint8
	Logs    []LogEvent
}

// SubmitResultV1 is a previous version of result from the `submit` method
type SubmitResultV1 struct {
	Status  TransactionStatus
	GasUsed uint64
	Logs    []LogEvent
}

// SubmitResultV2 is the latest version of result from the `submit` method
type SubmitResultV2 struct {
	Version uint8
	Status  TransactionStatus
	GasUsed uint64
	Logs    []LogEventWithAddress
}

// SubmitResultV2Version is a current version of SubmitResultV2
const SubmitResultV2Version = 7

// SubmitResult is a generic multi-version result from the `submit` method
type SubmitResult struct {
	Legacy *LegacyExecutionResult
	V1     *SubmitResultV1
	V2     *SubmitResultV2
}

// DeserializeSubmitResult deserializes the right version of SubmitResult using borsh
func DeserializeSubmitResult(result *SubmitResult, input []byte) error {
	if len(input) == 0 {
		return fmt.Errorf("input is empty")
	}

	// Note: SubmitResultV1 and LegacyExecutionResult binary formats
	// will never start with the version byte of SubmitResultV2
	// because SubmitResultV1 starts with an enum with fewer than 7
	// variants and LegacyExecutionResult starts with a boolean.
	if input[0] == SubmitResultV2Version {
		v2 := new(SubmitResultV2)
		if err := borsh.Deserialize(result.V2, input); err != nil {
			return fmt.Errorf("can't decode SubmitResultV2: %v", err)
		}
		result.V2 = v2
		return nil
	}

	v1 := new(SubmitResultV1)
	v1Err := borsh.Deserialize(v1, input)
	if v1Err == nil {
		result.V1 = v1
		return nil
	}

	legacy := new(LegacyExecutionResult)
	legacyErr := borsh.Deserialize(legacy, input)
	if legacyErr == nil {
		result.Legacy = legacy
		return nil
	}

	return fmt.Errorf(
		"can't decode neither SubmitResultV1 (%v), nor LegacyExecutionResult (%v)",
		v1Err,
		legacyErr,
	)
}

// Serialize serializes SubmitResult using borsh
func (result *SubmitResult) Serialize() ([]byte, error) {
	switch {
	case result.Legacy != nil:
		return borsh.Serialize(*result.Legacy)
	case result.V1 != nil:
		return borsh.Serialize(*result.V1)
	case result.V2 != nil:
		return borsh.Serialize(*result.V2)
	default:
		return nil, fmt.Errorf("SubmitResult wasn't initialized")
	}
}
