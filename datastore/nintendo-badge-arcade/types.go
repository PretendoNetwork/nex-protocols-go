package datastore_nintendo_badge_arcade

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

type DataStoreGetMetaByOwnerIDParam struct {
	nex.Structure
	OwnerIDs     []uint32
	DataTypes    []uint16
	ResultOption uint8
	ResultRange  *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreGetMetaByOwnerIDParam structure from a stream
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetMetaByOwnerIDParam.OwnerIDs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.OwnerIDs. %s", err.Error())
	}

	dataStoreGetMetaByOwnerIDParam.DataTypes, err = stream.ReadListUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.DataTypes. %s", err.Error())
	}

	dataStoreGetMetaByOwnerIDParam.ResultOption, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultOption. %s", err.Error())
	}

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultRange. %s", err.Error())
	}

	dataStoreGetMetaByOwnerIDParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreGetMetaByOwnerIDParam and returns a byte array
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt32LE(dataStoreGetMetaByOwnerIDParam.OwnerIDs)
	stream.WriteListUInt16LE(dataStoreGetMetaByOwnerIDParam.DataTypes)
	stream.WriteUInt8(dataStoreGetMetaByOwnerIDParam.ResultOption)
	stream.WriteStructure(dataStoreGetMetaByOwnerIDParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetMetaByOwnerIDParam
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetMetaByOwnerIDParam()

	copied.OwnerIDs = make([]uint32, len(dataStoreGetMetaByOwnerIDParam.OwnerIDs))

	copy(copied.OwnerIDs, dataStoreGetMetaByOwnerIDParam.OwnerIDs)

	copied.DataTypes = make([]uint16, len(dataStoreGetMetaByOwnerIDParam.DataTypes))

	copy(copied.DataTypes, dataStoreGetMetaByOwnerIDParam.DataTypes)

	copied.ResultOption = dataStoreGetMetaByOwnerIDParam.ResultOption
	copied.ResultRange = dataStoreGetMetaByOwnerIDParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetMetaByOwnerIDParam)

	if len(dataStoreGetMetaByOwnerIDParam.OwnerIDs) != len(other.OwnerIDs) {
		return false
	}

	for i := 0; i < len(dataStoreGetMetaByOwnerIDParam.OwnerIDs); i++ {
		if dataStoreGetMetaByOwnerIDParam.OwnerIDs[i] != other.OwnerIDs[i] {
			return false
		}
	}

	if len(dataStoreGetMetaByOwnerIDParam.DataTypes) != len(other.DataTypes) {
		return false
	}

	for i := 0; i < len(dataStoreGetMetaByOwnerIDParam.DataTypes); i++ {
		if dataStoreGetMetaByOwnerIDParam.DataTypes[i] != other.DataTypes[i] {
			return false
		}
	}

	if dataStoreGetMetaByOwnerIDParam.ResultOption != other.ResultOption {
		return false
	}

	if !dataStoreGetMetaByOwnerIDParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// NewDataStoreGetMetaByOwnerIDParam returns a new DataStoreGetMetaByOwnerIDParam
func NewDataStoreGetMetaByOwnerIDParam() *DataStoreGetMetaByOwnerIDParam {
	return &DataStoreGetMetaByOwnerIDParam{}
}
