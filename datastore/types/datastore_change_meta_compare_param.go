// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

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

	copied.SetStructureVersion(dataStoreChangeMetaCompareParam.StructureVersion())

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

	if dataStoreChangeMetaCompareParam.StructureVersion() != other.StructureVersion() {
		return false
	}

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

// String returns a string representation of the struct
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) String() string {
	return dataStoreChangeMetaCompareParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaCompareParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreChangeMetaCompareParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sComparisonFlag: %d,\n", indentationValues, dataStoreChangeMetaCompareParam.ComparisonFlag))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, dataStoreChangeMetaCompareParam.Name))

	if dataStoreChangeMetaCompareParam.Permission != nil {
		b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.Permission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPermission: nil,\n", indentationValues))
	}

	if dataStoreChangeMetaCompareParam.DelPermission != nil {
		b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStoreChangeMetaCompareParam.DelPermission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDelPermission: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, dataStoreChangeMetaCompareParam.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x,\n", indentationValues, dataStoreChangeMetaCompareParam.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %v,\n", indentationValues, dataStoreChangeMetaCompareParam.Tags))
	b.WriteString(fmt.Sprintf("%sReferredCnt: %d,\n", indentationValues, dataStoreChangeMetaCompareParam.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreChangeMetaCompareParam.DataType))
	b.WriteString(fmt.Sprintf("%sStatus: %d\n", indentationValues, dataStoreChangeMetaCompareParam.Status))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaCompareParam returns a new DataStoreChangeMetaCompareParam
func NewDataStoreChangeMetaCompareParam() *DataStoreChangeMetaCompareParam {
	return &DataStoreChangeMetaCompareParam{
		ComparisonFlag: 0,
		Name:           "",
		Permission:     NewDataStorePermission(),
		DelPermission:  NewDataStorePermission(),
		Period:         0,
		MetaBinary:     make([]byte, 0),
		Tags:           make([]string, 0),
		ReferredCnt:    0,
		DataType:       0,
		Status:         0,
	}
}
