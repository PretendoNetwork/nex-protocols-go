package datastore_super_smash_bros_4_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreSearchReplayParam struct {
	nex.Structure
	Mode        uint8
	Style       uint8
	Fighter     uint8
	ResultRange *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreSearchReplayParam structure from a stream
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreSearchReplayParam.Mode, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Mode. %s", err.Error())
	}
	dataStoreSearchReplayParam.Style, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Style. %s", err.Error())
	}
	dataStoreSearchReplayParam.Fighter, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Fighter. %s", err.Error())
	}

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.ResultRange. %s", err.Error())
	}

	dataStoreSearchReplayParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreSearchReplayParam and returns a byte array
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStoreSearchReplayParam.Mode)
	stream.WriteUInt8(dataStoreSearchReplayParam.Style)
	stream.WriteUInt8(dataStoreSearchReplayParam.Fighter)
	stream.WriteStructure(dataStoreSearchReplayParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSearchReplayParam
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchReplayParam()

	copied.Mode = dataStoreSearchReplayParam.Mode
	copied.Style = dataStoreSearchReplayParam.Style
	copied.Fighter = dataStoreSearchReplayParam.Fighter
	copied.ResultRange = dataStoreSearchReplayParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchReplayParam)

	if dataStoreSearchReplayParam.Mode != other.Mode {
		return false
	}

	if dataStoreSearchReplayParam.Style != other.Style {
		return false
	}

	if dataStoreSearchReplayParam.Fighter != other.Fighter {
		return false
	}

	if !dataStoreSearchReplayParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// NewDataStoreSearchReplayParam returns a new DataStoreSearchReplayParam
func NewDataStoreSearchReplayParam() *DataStoreSearchReplayParam {
	return &DataStoreSearchReplayParam{}
}
