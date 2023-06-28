package datastore_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreChangeMetaCompareParam is sent in the ChangeMeta method
type DataStoreChangeMetaCompareParam struct {
	nex.Structure
	ComparisonFlag uint32
	Name           string
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         uint16
	MetaBinary     []byte
	Tags           []string
	ReferredCnt    uint32
	DataType       uint16
	Status         uint8
}

// ExtractFromStream extracts a DataStoreChangeMetaCompareParam structure from a stream
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreChangeMetaCompareParam.ComparisonFlag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.ComparisonFlag. %s", err.Error())
	}

	dataStoreChangeMetaCompareParam.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Name. %s", err.Error())
	}

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Permission. %s", err.Error())
	}

	dataStoreChangeMetaCompareParam.Permission = permission.(*DataStorePermission)
	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.DelPermission. %s", err.Error())
	}

	dataStoreChangeMetaCompareParam.DelPermission = delPermission.(*DataStorePermission)
	dataStoreChangeMetaCompareParam.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Period. %s", err.Error())
	}

	dataStoreChangeMetaCompareParam.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.MetaBinary. %s", err.Error())
	}

	dataStoreChangeMetaCompareParam.Tags, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Tags. %s", err.Error())
	}

	dataStoreChangeMetaCompareParam.ReferredCnt, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.ReferredCnt. %s", err.Error())
	}

	dataStoreChangeMetaCompareParam.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.DataType. %s", err.Error())
	}

	dataStoreChangeMetaCompareParam.Status, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaCompareParam.Status. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangeMetaCompareParam
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) Copy() nex.StructureInterface {
	copied := NewDataStoreChangeMetaCompareParam()

	copied.ComparisonFlag = dataStoreChangeMetaCompareParam.ComparisonFlag
	copied.Name = dataStoreChangeMetaCompareParam.Name
	copied.Permission = dataStoreChangeMetaCompareParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaCompareParam.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaCompareParam.Period
	copied.MetaBinary = make([]byte, len(dataStoreChangeMetaCompareParam.MetaBinary))

	copy(copied.MetaBinary, dataStoreChangeMetaCompareParam.MetaBinary)

	copied.Tags = make([]string, len(dataStoreChangeMetaCompareParam.Tags))

	copy(copied.Tags, dataStoreChangeMetaCompareParam.Tags)

	copied.ReferredCnt = dataStoreChangeMetaCompareParam.ReferredCnt
	copied.DataType = dataStoreChangeMetaCompareParam.DataType
	copied.Status = dataStoreChangeMetaCompareParam.Status

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangeMetaCompareParam)

	if dataStoreChangeMetaCompareParam.ComparisonFlag != other.ComparisonFlag {
		return false
	}

	if dataStoreChangeMetaCompareParam.Name != other.Name {
		return false
	}

	if dataStoreChangeMetaCompareParam.Permission.Equals(other.Permission) {
		return false
	}

	if dataStoreChangeMetaCompareParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStoreChangeMetaCompareParam.Period != other.Period {
		return false
	}

	if !bytes.Equal(dataStoreChangeMetaCompareParam.MetaBinary, other.MetaBinary) {
		return false
	}

	if len(dataStoreChangeMetaCompareParam.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreChangeMetaCompareParam.Tags); i++ {
		if dataStoreChangeMetaCompareParam.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if dataStoreChangeMetaCompareParam.ReferredCnt != other.ReferredCnt {
		return false
	}

	if dataStoreChangeMetaCompareParam.DataType != other.DataType {
		return false
	}

	if dataStoreChangeMetaCompareParam.Status != other.Status {
		return false
	}

	return true
}

// NewDataStoreChangeMetaCompareParam returns a new DataStoreChangeMetaCompareParam
func NewDataStoreChangeMetaCompareParam() *DataStoreChangeMetaCompareParam {
	return &DataStoreChangeMetaCompareParam{}
}
