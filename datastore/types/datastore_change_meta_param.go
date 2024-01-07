// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreChangeMetaParam is sent in the ChangeMeta method
type DataStoreChangeMetaParam struct {
	types.Structure
	DataID            *types.PrimitiveU64
	ModifiesFlag      *types.PrimitiveU32
	Name              *types.String
	Permission        *DataStorePermission
	DelPermission     *DataStorePermission
	Period            *types.PrimitiveU16
	MetaBinary        *types.QBuffer
	Tags              *types.List[*types.String]
	UpdatePassword    *types.PrimitiveU64
	ReferredCnt       *types.PrimitiveU32
	DataType          *types.PrimitiveU16
	Status            *types.PrimitiveU8
	CompareParam      *DataStoreChangeMetaCompareParam
	PersistenceTarget *DataStorePersistenceTarget
}

// WriteTo writes the DataStoreChangeMetaParam to the given writable
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreChangeMetaParam.DataID.WriteTo(contentWritable)
	dataStoreChangeMetaParam.ModifiesFlag.WriteTo(contentWritable)
	dataStoreChangeMetaParam.Name.WriteTo(contentWritable)
	dataStoreChangeMetaParam.Permission.WriteTo(contentWritable)
	dataStoreChangeMetaParam.DelPermission.WriteTo(contentWritable)
	dataStoreChangeMetaParam.Period.WriteTo(contentWritable)
	dataStoreChangeMetaParam.MetaBinary.WriteTo(contentWritable)
	dataStoreChangeMetaParam.Tags.WriteTo(contentWritable)
	dataStoreChangeMetaParam.ReferredCnt.WriteTo(contentWritable)
	dataStoreChangeMetaParam.DataType.WriteTo(contentWritable)
	dataStoreChangeMetaParam.Status.WriteTo(contentWritable)
	dataStoreChangeMetaParam.CompareParam.WriteTo(contentWritable)

	if dataStoreChangeMetaParam.StructureVersion >= 1 {
		dataStoreChangeMetaParam.PersistenceTarget.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dataStoreChangeMetaParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreChangeMetaParam from the given readable
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreChangeMetaParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreChangeMetaParam header. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DataID. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.ModifiesFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.ModifiesFlag. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Name. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Permission. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DelPermission. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Period. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.MetaBinary. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Tags. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.UpdatePassword. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.ReferredCnt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.ReferredCnt. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.DataType. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.Status.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.Status. %s", err.Error())
	}

	err = dataStoreChangeMetaParam.CompareParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.CompareParam. %s", err.Error())
	}

	if dataStoreChangeMetaParam.StructureVersion >= 1 {
		err = dataStoreChangeMetaParam.PersistenceTarget.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreChangeMetaParam.PersistenceTarget. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStoreChangeMetaParam
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) Copy() types.RVType {
	copied := NewDataStoreChangeMetaParam()

	copied.StructureVersion = dataStoreChangeMetaParam.StructureVersion

	copied.DataID = dataStoreChangeMetaParam.DataID.Copy().(*types.PrimitiveU64)
	copied.ModifiesFlag = dataStoreChangeMetaParam.ModifiesFlag.Copy().(*types.PrimitiveU32)
	copied.Name = dataStoreChangeMetaParam.Name.Copy().(*types.String)
	copied.Permission = dataStoreChangeMetaParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaParam.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaParam.Period.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dataStoreChangeMetaParam.MetaBinary.Copy().(*types.QBuffer)

	copied.Tags = dataStoreChangeMetaParam.Tags.Copy().(*types.List[*types.String])

	copied.UpdatePassword = dataStoreChangeMetaParam.UpdatePassword.Copy().(*types.PrimitiveU64)
	copied.ReferredCnt = dataStoreChangeMetaParam.ReferredCnt.Copy().(*types.PrimitiveU32)
	copied.DataType = dataStoreChangeMetaParam.DataType.Copy().(*types.PrimitiveU16)
	copied.Status = dataStoreChangeMetaParam.Status.Copy().(*types.PrimitiveU8)
	copied.CompareParam = dataStoreChangeMetaParam.CompareParam.Copy().(*DataStoreChangeMetaCompareParam)

	copied.PersistenceTarget = dataStoreChangeMetaParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangeMetaParam); !ok {
		return false
	}

	other := o.(*DataStoreChangeMetaParam)

	if dataStoreChangeMetaParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreChangeMetaParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreChangeMetaParam.ModifiesFlag.Equals(other.ModifiesFlag) {
		return false
	}

	if !dataStoreChangeMetaParam.Name.Equals(other.Name) {
		return false
	}

	if dataStoreChangeMetaParam.Permission.Equals(other.Permission) {
		return false
	}

	if dataStoreChangeMetaParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dataStoreChangeMetaParam.Period.Equals(other.Period) {
		return false
	}

	if !dataStoreChangeMetaParam.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dataStoreChangeMetaParam.Tags.Equals(other.Tags) {
		return false
	}

	if !dataStoreChangeMetaParam.UpdatePassword.Equals(other.UpdatePassword) {
		return false
	}

	if !dataStoreChangeMetaParam.ReferredCnt.Equals(other.ReferredCnt) {
		return false
	}

	if !dataStoreChangeMetaParam.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStoreChangeMetaParam.Status.Equals(other.Status) {
		return false
	}

	if dataStoreChangeMetaParam.CompareParam.Equals(other.CompareParam) {
		return false
	}

	if !dataStoreChangeMetaParam.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreChangeMetaParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreChangeMetaParam.DataID))
	b.WriteString(fmt.Sprintf("%sModifiesFlag: %s,\n", indentationValues, dataStoreChangeMetaParam.ModifiesFlag))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dataStoreChangeMetaParam.Name))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStoreChangeMetaParam.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStoreChangeMetaParam.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dataStoreChangeMetaParam.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dataStoreChangeMetaParam.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dataStoreChangeMetaParam.Tags))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s,\n", indentationValues, dataStoreChangeMetaParam.UpdatePassword))
	b.WriteString(fmt.Sprintf("%sReferredCnt: %s,\n", indentationValues, dataStoreChangeMetaParam.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dataStoreChangeMetaParam.DataType))
	b.WriteString(fmt.Sprintf("%sStatus: %s,\n", indentationValues, dataStoreChangeMetaParam.Status))
	b.WriteString(fmt.Sprintf("%sCompareParam: %s,\n", indentationValues, dataStoreChangeMetaParam.CompareParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s\n", indentationValues, dataStoreChangeMetaParam.PersistenceTarget.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaParam returns a new DataStoreChangeMetaParam
func NewDataStoreChangeMetaParam() *DataStoreChangeMetaParam {
	dataStoreChangeMetaParam := &DataStoreChangeMetaParam{
		DataID:            types.NewPrimitiveU64(0),
		ModifiesFlag:      types.NewPrimitiveU32(0),
		Name:              types.NewString(""),
		Permission:        NewDataStorePermission(),
		DelPermission:     NewDataStorePermission(),
		Period:            types.NewPrimitiveU16(0),
		MetaBinary:        types.NewQBuffer(nil),
		Tags:              types.NewList[*types.String](),
		UpdatePassword:    types.NewPrimitiveU64(0),
		ReferredCnt:       types.NewPrimitiveU32(0),
		DataType:          types.NewPrimitiveU16(0),
		Status:            types.NewPrimitiveU8(0),
		CompareParam:      NewDataStoreChangeMetaCompareParam(),
		PersistenceTarget: NewDataStorePersistenceTarget(),
	}

	dataStoreChangeMetaParam.Tags.Type = types.NewString("")

	return dataStoreChangeMetaParam
}
