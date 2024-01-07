// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreChangeMetaParamV1 is a data structure used by the DataStore protocol
type DataStoreChangeMetaParamV1 struct {
	types.Structure
	DataID         *types.PrimitiveU64
	ModifiesFlag   *types.PrimitiveU32
	Name           *types.String
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         *types.PrimitiveU16
	MetaBinary     *types.QBuffer
	Tags           *types.List[*types.String]
	UpdatePassword *types.PrimitiveU64
}

// ExtractFrom extracts the DataStoreChangeMetaParamV1 from the given readable
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreChangeMetaParamV1.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreChangeMetaParamV1 header. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.DataID. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.ModifiesFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.ModifiesFlag. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Name. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Permission. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.DelPermission. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Period. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.MetaBinary. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Tags. %s", err.Error())
	}

	err = dataStoreChangeMetaParamV1.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.UpdatePassword. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreChangeMetaParamV1 to the given writable
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreChangeMetaParamV1.DataID.WriteTo(contentWritable)
	dataStoreChangeMetaParamV1.ModifiesFlag.WriteTo(contentWritable)
	dataStoreChangeMetaParamV1.Name.WriteTo(contentWritable)
	dataStoreChangeMetaParamV1.Permission.WriteTo(contentWritable)
	dataStoreChangeMetaParamV1.DelPermission.WriteTo(contentWritable)
	dataStoreChangeMetaParamV1.Period.WriteTo(contentWritable)
	dataStoreChangeMetaParamV1.MetaBinary.WriteTo(contentWritable)
	dataStoreChangeMetaParamV1.Tags.WriteTo(contentWritable)
	dataStoreChangeMetaParamV1.UpdatePassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreChangeMetaParamV1.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreChangeMetaParamV1
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) Copy() types.RVType {
	copied := NewDataStoreChangeMetaParamV1()

	copied.StructureVersion = dataStoreChangeMetaParamV1.StructureVersion

	copied.DataID = dataStoreChangeMetaParamV1.DataID.Copy().(*types.PrimitiveU64)
	copied.ModifiesFlag = dataStoreChangeMetaParamV1.ModifiesFlag.Copy().(*types.PrimitiveU32)
	copied.Name = dataStoreChangeMetaParamV1.Name.Copy().(*types.String)
	copied.Permission = dataStoreChangeMetaParamV1.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaParamV1.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaParamV1.Period.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dataStoreChangeMetaParamV1.MetaBinary.Copy().(*types.QBuffer)

	copied.Tags = dataStoreChangeMetaParamV1.Tags.Copy().(*types.List[*types.String])

	copied.UpdatePassword = dataStoreChangeMetaParamV1.UpdatePassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreChangeMetaParamV1); !ok {
		return false
	}

	other := o.(*DataStoreChangeMetaParamV1)

	if dataStoreChangeMetaParamV1.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreChangeMetaParamV1.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreChangeMetaParamV1.ModifiesFlag.Equals(other.ModifiesFlag) {
		return false
	}

	if !dataStoreChangeMetaParamV1.Name.Equals(other.Name) {
		return false
	}

	if !dataStoreChangeMetaParamV1.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStoreChangeMetaParamV1.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dataStoreChangeMetaParamV1.Period.Equals(other.Period) {
		return false
	}

	if !dataStoreChangeMetaParamV1.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dataStoreChangeMetaParamV1.Tags.Equals(other.Tags) {
		return false
	}

	return dataStoreChangeMetaParamV1.UpdatePassword == other.UpdatePassword
}

// String returns a string representation of the struct
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) String() string {
	return dataStoreChangeMetaParamV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaParamV1{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreChangeMetaParamV1.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreChangeMetaParamV1.DataID))
	b.WriteString(fmt.Sprintf("%sModifiesFlag: %s,\n", indentationValues, dataStoreChangeMetaParamV1.ModifiesFlag))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dataStoreChangeMetaParamV1.Name))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStoreChangeMetaParamV1.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStoreChangeMetaParamV1.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dataStoreChangeMetaParamV1.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dataStoreChangeMetaParamV1.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dataStoreChangeMetaParamV1.Tags))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s\n", indentationValues, dataStoreChangeMetaParamV1.UpdatePassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaParamV1 returns a new DataStoreChangeMetaParamV1
func NewDataStoreChangeMetaParamV1() *DataStoreChangeMetaParamV1 {
	dataStoreChangeMetaParamV1 := &DataStoreChangeMetaParamV1{
		DataID:         types.NewPrimitiveU64(0),
		ModifiesFlag:   types.NewPrimitiveU32(0),
		Name:           types.NewString(""),
		Permission:     NewDataStorePermission(),
		DelPermission:  NewDataStorePermission(),
		Period:         types.NewPrimitiveU16(0),
		MetaBinary:     types.NewQBuffer(nil),
		Tags:           types.NewList[*types.String](),
		UpdatePassword: types.NewPrimitiveU64(0),
	}

	dataStoreChangeMetaParamV1.Tags.Type = types.NewString("")

	return dataStoreChangeMetaParamV1
}
