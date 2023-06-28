package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreTouchObjectParam struct {
	nex.Structure
	DataID         uint64
	LockID         uint32
	AccessPassword uint64
}

// ExtractFromStream extracts a DataStoreTouchObjectParam structure from a stream
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreTouchObjectParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.DataID. %s", err.Error())
	}

	dataStoreTouchObjectParam.LockID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.LockID. %s", err.Error())
	}

	dataStoreTouchObjectParam.AccessPassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreTouchObjectParam and returns a byte array
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreTouchObjectParam.DataID)
	stream.WriteUInt32LE(dataStoreTouchObjectParam.LockID)
	stream.WriteUInt64LE(dataStoreTouchObjectParam.AccessPassword)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreTouchObjectParam
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Copy() nex.StructureInterface {
	copied := NewDataStoreTouchObjectParam()

	copied.DataID = dataStoreTouchObjectParam.DataID
	copied.LockID = dataStoreTouchObjectParam.LockID
	copied.AccessPassword = dataStoreTouchObjectParam.AccessPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreTouchObjectParam)

	if dataStoreTouchObjectParam.DataID != other.DataID {
		return false
	}

	if dataStoreTouchObjectParam.LockID != other.LockID {
		return false
	}

	if dataStoreTouchObjectParam.AccessPassword != other.AccessPassword {
		return false
	}

	return true
}

// NewDataStoreTouchObjectParam returns a new DataStoreTouchObjectParam
func NewDataStoreTouchObjectParam() *DataStoreTouchObjectParam {
	return &DataStoreTouchObjectParam{}
}
