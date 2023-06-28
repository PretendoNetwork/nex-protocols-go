package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRateObjectParam is sent in the RateObjects method
type DataStoreRateObjectParam struct {
	nex.Structure
	RatingValue    int32
	AccessPassword uint64
}

// ExtractFromStream extracts a DataStoreRateObjectParam structure from a stream
func (dataStoreRateObjectParam *DataStoreRateObjectParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRateObjectParam.RatingValue, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam.RatingValue. %s", err.Error())
	}

	dataStoreRateObjectParam.AccessPassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRateObjectParam
func (dataStoreRateObjectParam *DataStoreRateObjectParam) Copy() nex.StructureInterface {
	copied := NewDataStoreRateObjectParam()

	copied.RatingValue = dataStoreRateObjectParam.RatingValue
	copied.AccessPassword = dataStoreRateObjectParam.AccessPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRateObjectParam *DataStoreRateObjectParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRateObjectParam)

	if dataStoreRateObjectParam.RatingValue != other.RatingValue {
		return false
	}

	if dataStoreRateObjectParam.AccessPassword != other.AccessPassword {
		return false
	}

	return true
}

// NewDataStoreRateObjectParam returns a new DataStoreRateObjectParam
func NewDataStoreRateObjectParam() *DataStoreRateObjectParam {
	return &DataStoreRateObjectParam{}
}
