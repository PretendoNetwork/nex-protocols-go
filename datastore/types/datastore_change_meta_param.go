// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreChangeMetaParam is sent in the ChangeMeta method
type DataStoreChangeMetaParam struct {
	nex.Structure
	DataID            uint64
	ModifiesFlag      uint32
	Name              string
	Permission        *DataStorePermission
	DelPermission     *DataStorePermission
	Period            uint16
	MetaBinary        []byte
	Tags              []string
	UpdatePassword    uint64
	ReferredCnt       uint32
	DataType          uint16
	Status            uint8
	CompareParam      *DataStoreChangeMetaCompareParam
	PersistenceTarget *DataStorePersistenceTarget
}

// ExtractFromStream extracts a DataStoreChangeMetaParam structure from a stream
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreChangeMetaParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DataID. %s", err.Error())
	}

	dataStoreChangeMetaParam.ModifiesFlag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.ModifiesFlag. %s", err.Error())
	}

	dataStoreChangeMetaParam.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Name. %s", err.Error())
	}

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Permission. %s", err.Error())
	}

	dataStoreChangeMetaParam.Permission = permission.(*DataStorePermission)
	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DelPermission. %s", err.Error())
	}

	dataStoreChangeMetaParam.DelPermission = delPermission.(*DataStorePermission)
	dataStoreChangeMetaParam.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Period. %s", err.Error())
	}

	dataStoreChangeMetaParam.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.MetaBinary. %s", err.Error())
	}

	dataStoreChangeMetaParam.Tags, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Tags. %s", err.Error())
	}

	dataStoreChangeMetaParam.UpdatePassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.UpdatePassword. %s", err.Error())
	}

	dataStoreChangeMetaParam.ReferredCnt, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.ReferredCnt. %s", err.Error())
	}

	dataStoreChangeMetaParam.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DataType. %s", err.Error())
	}

	dataStoreChangeMetaParam.Status, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Status. %s", err.Error())
	}

	compareParam, err := stream.ReadStructure(NewDataStoreChangeMetaCompareParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.CompareParam. %s", err.Error())
	}

	dataStoreChangeMetaParam.CompareParam = compareParam.(*DataStoreChangeMetaCompareParam)

	if dataStoreChangeMetaParam.StructureVersion() >= 1 {
		persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.PersistenceTarget. %s", err.Error())
		}

		dataStoreChangeMetaParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangeMetaParam
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreChangeMetaParam()

	copied.DataID = dataStoreChangeMetaParam.DataID
	copied.ModifiesFlag = dataStoreChangeMetaParam.ModifiesFlag
	copied.Name = dataStoreChangeMetaParam.Name
	copied.Permission = dataStoreChangeMetaParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaParam.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaParam.Period
	copied.MetaBinary = make([]byte, len(dataStoreChangeMetaParam.MetaBinary))

	copy(copied.MetaBinary, dataStoreChangeMetaParam.MetaBinary)

	copied.Tags = make([]string, len(dataStoreChangeMetaParam.Tags))

	copy(copied.Tags, dataStoreChangeMetaParam.Tags)

	copied.UpdatePassword = dataStoreChangeMetaParam.UpdatePassword
	copied.ReferredCnt = dataStoreChangeMetaParam.ReferredCnt
	copied.DataType = dataStoreChangeMetaParam.DataType
	copied.Status = dataStoreChangeMetaParam.Status
	copied.CompareParam = dataStoreChangeMetaParam.CompareParam.Copy().(*DataStoreChangeMetaCompareParam)

	if dataStoreChangeMetaParam.PersistenceTarget != nil {
		copied.PersistenceTarget = dataStoreChangeMetaParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangeMetaParam)

	if dataStoreChangeMetaParam.DataID != other.DataID {
		return false
	}

	if dataStoreChangeMetaParam.ModifiesFlag != other.ModifiesFlag {
		return false
	}

	if dataStoreChangeMetaParam.Name != other.Name {
		return false
	}

	if dataStoreChangeMetaParam.Permission.Equals(other.Permission) {
		return false
	}

	if dataStoreChangeMetaParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStoreChangeMetaParam.Period != other.Period {
		return false
	}

	if !bytes.Equal(dataStoreChangeMetaParam.MetaBinary, other.MetaBinary) {
		return false
	}

	if len(dataStoreChangeMetaParam.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreChangeMetaParam.Tags); i++ {
		if dataStoreChangeMetaParam.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if dataStoreChangeMetaParam.UpdatePassword != other.UpdatePassword {
		return false
	}

	if dataStoreChangeMetaParam.ReferredCnt != other.ReferredCnt {
		return false
	}

	if dataStoreChangeMetaParam.DataType != other.DataType {
		return false
	}

	if dataStoreChangeMetaParam.Status != other.Status {
		return false
	}

	if dataStoreChangeMetaParam.CompareParam.Equals(other.CompareParam) {
		return false
	}

	if dataStoreChangeMetaParam.PersistenceTarget != nil && other.PersistenceTarget == nil {
		return false
	}

	if dataStoreChangeMetaParam.PersistenceTarget == nil && other.PersistenceTarget != nil {
		return false
	}

	if dataStoreChangeMetaParam.PersistenceTarget != nil && other.PersistenceTarget != nil {
		if !dataStoreChangeMetaParam.PersistenceTarget.Equals(other.PersistenceTarget) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) String() string {
	return dataStoreChangeMetaParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreChangeMetaParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreChangeMetaParam.DataID))
	b.WriteString(fmt.Sprintf("%sModifiesFlag: %d,\n", indentationValues, dataStoreChangeMetaParam.ModifiesFlag))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, dataStoreChangeMetaParam.Name))

	if dataStoreChangeMetaParam.Permission != nil {
		b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStoreChangeMetaParam.Permission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPermission: nil,\n", indentationValues))
	}

	if dataStoreChangeMetaParam.DelPermission != nil {
		b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStoreChangeMetaParam.DelPermission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDelPermission: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, dataStoreChangeMetaParam.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x,\n", indentationValues, dataStoreChangeMetaParam.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %v,\n", indentationValues, dataStoreChangeMetaParam.Tags))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %d,\n", indentationValues, dataStoreChangeMetaParam.UpdatePassword))
	b.WriteString(fmt.Sprintf("%sReferredCnt: %d,\n", indentationValues, dataStoreChangeMetaParam.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreChangeMetaParam.DataType))
	b.WriteString(fmt.Sprintf("%sStatus: %d,\n", indentationValues, dataStoreChangeMetaParam.Status))

	if dataStoreChangeMetaParam.CompareParam != nil {
		b.WriteString(fmt.Sprintf("%sCompareParam: %s,\n", indentationValues, dataStoreChangeMetaParam.CompareParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCompareParam: nil,\n", indentationValues))
	}

	if dataStoreChangeMetaParam.PersistenceTarget != nil {
		b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s\n", indentationValues, dataStoreChangeMetaParam.PersistenceTarget.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPersistenceTarget: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaParam returns a new DataStoreChangeMetaParam
func NewDataStoreChangeMetaParam() *DataStoreChangeMetaParam {
	return &DataStoreChangeMetaParam{}
}
