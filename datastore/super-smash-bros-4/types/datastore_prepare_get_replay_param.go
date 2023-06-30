// Package datastore_super_smash_bros_4_types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package datastore_super_smash_bros_4_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePrepareGetReplayParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStorePrepareGetReplayParam struct {
	nex.Structure
	ReplayID  uint64
	ExtraData []string
}

// ExtractFromStream extracts a DataStorePrepareGetReplayParam structure from a stream
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePrepareGetReplayParam.ReplayID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetReplayParam.ReplayID. %s", err.Error())
	}

	dataStorePrepareGetReplayParam.ExtraData, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetReplayParam.ExtraData. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePrepareGetReplayParam and returns a byte array
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStorePrepareGetReplayParam.ReplayID)
	stream.WriteListString(dataStorePrepareGetReplayParam.ExtraData)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePrepareGetReplayParam
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStorePrepareGetReplayParam()

	copied.ReplayID = dataStorePrepareGetReplayParam.ReplayID
	copied.ExtraData = make([]string, len(dataStorePrepareGetReplayParam.ExtraData))

	copy(copied.ExtraData, dataStorePrepareGetReplayParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePrepareGetReplayParam)

	if dataStorePrepareGetReplayParam.ReplayID != other.ReplayID {
		return false
	}

	if len(dataStorePrepareGetReplayParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePrepareGetReplayParam.ExtraData); i++ {
		if dataStorePrepareGetReplayParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) String() string {
	return dataStorePrepareGetReplayParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareGetReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePrepareGetReplayParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReplayID: %d,\n", indentationValues, dataStorePrepareGetReplayParam.ReplayID))
	b.WriteString(fmt.Sprintf("%sExtraData: %v\n", indentationValues, dataStorePrepareGetReplayParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetReplayParam returns a new DataStorePrepareGetReplayParam
func NewDataStorePrepareGetReplayParam() *DataStorePrepareGetReplayParam {
	return &DataStorePrepareGetReplayParam{}
}
