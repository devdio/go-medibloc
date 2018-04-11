package core

import (
	"encoding/json"

	"github.com/medibloc/go-medibloc/common"
)

type RegisterWriterPayload struct {
	// Writer to register
	Writer common.Address
}

type RemoveWriterPayload struct {
	// Writer to remove
	Writer common.Address
}

// NewRegisterWriterPayload generates a RegisterWriterPayload value
func NewRegisterWriterPayload(writer common.Address) *RegisterWriterPayload {
	return &RegisterWriterPayload{Writer: writer}
}

// NewRemoveWriterPayload generates a RemoveWriterPayload value
func NewRemoveWriterPayload(writer common.Address) *RemoveWriterPayload {
	return &RemoveWriterPayload{Writer: writer}
}

// BytesToRegisterWriterPayload converts bytes to RegisterWriterPayload struct
func BytesToRegisterWriterPayload(b []byte) (*RegisterWriterPayload, error) {
	payload := &RegisterWriterPayload{}
	if err := json.Unmarshal(b, payload); err != nil {
		return nil, ErrInvalidTxPayload
	}
	return payload, nil
}

// ToBytes returns marshalled RegisterWriterPayload
func (payload *RegisterWriterPayload) ToBytes() ([]byte, error) {
	return json.Marshal(payload)
}

// BytesToRemoveWriterPayload converts bytes to RemoveWriterPayload struct
func BytesToRemoveWriterPayload(b []byte) (*RemoveWriterPayload, error) {
	payload := &RemoveWriterPayload{}
	if err := json.Unmarshal(b, payload); err != nil {
		return nil, ErrInvalidTxPayload
	}
	return payload, nil
}

// ToBytes returns marshalled RemoveWriterPayload
func (payload *RemoveWriterPayload) ToBytes() ([]byte, error) {
	return json.Marshal(payload)
}
