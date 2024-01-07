// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePreparePostParamV1 is a data structure used by the DataStore protocol
type DataStorePreparePostParamV1 struct {
	types.Structure
	Size             *types.PrimitiveU32
	Name             *types.String
	DataType         *types.PrimitiveU16
	MetaBinary       *types.QBuffer
	Permission       *DataStorePermission
	DelPermission    *DataStorePermission
	Flag             *types.PrimitiveU32
	Period           *types.PrimitiveU16
	ReferDataID      *types.PrimitiveU32
	Tags             *types.List[*types.String]
	RatingInitParams *types.List[*DataStoreRatingInitParamWithSlot]
}

// ExtractFrom extracts the DataStorePreparePostParamV1 from the given readable
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePreparePostParamV1.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePreparePostParamV1 header. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Size. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Name. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.DataType. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.MetaBinary. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Permission. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.DelPermission. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.Flag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Flag. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Period. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.ReferDataID. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Tags. %s", err.Error())
	}

	err = dataStorePreparePostParamV1.RatingInitParams.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.RatingInitParams. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePreparePostParamV1 to the given writable
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePreparePostParamV1.Size.WriteTo(contentWritable)
	dataStorePreparePostParamV1.Name.WriteTo(contentWritable)
	dataStorePreparePostParamV1.DataType.WriteTo(contentWritable)
	dataStorePreparePostParamV1.MetaBinary.WriteTo(contentWritable)
	dataStorePreparePostParamV1.Permission.WriteTo(contentWritable)
	dataStorePreparePostParamV1.DelPermission.WriteTo(contentWritable)
	dataStorePreparePostParamV1.Flag.WriteTo(contentWritable)
	dataStorePreparePostParamV1.Period.WriteTo(contentWritable)
	dataStorePreparePostParamV1.ReferDataID.WriteTo(contentWritable)
	dataStorePreparePostParamV1.Tags.WriteTo(contentWritable)
	dataStorePreparePostParamV1.RatingInitParams.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePreparePostParamV1.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePreparePostParamV1
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) Copy() types.RVType {
	copied := NewDataStorePreparePostParamV1()

	copied.StructureVersion = dataStorePreparePostParamV1.StructureVersion

	copied.Size = dataStorePreparePostParamV1.Size.Copy().(*types.PrimitiveU32)
	copied.Name = dataStorePreparePostParamV1.Name.Copy().(*types.String)
	copied.DataType = dataStorePreparePostParamV1.DataType.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dataStorePreparePostParamV1.MetaBinary.Copy().(*types.QBuffer)
	copied.Permission = dataStorePreparePostParamV1.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStorePreparePostParamV1.DelPermission.Copy().(*DataStorePermission)
	copied.Flag = dataStorePreparePostParamV1.Flag.Copy().(*types.PrimitiveU32)
	copied.Period = dataStorePreparePostParamV1.Period.Copy().(*types.PrimitiveU16)
	copied.ReferDataID = dataStorePreparePostParamV1.ReferDataID.Copy().(*types.PrimitiveU32)
	copied.Tags = dataStorePreparePostParamV1.Tags.Copy().(*types.List[*types.String])
	copied.RatingInitParams = dataStorePreparePostParamV1.RatingInitParams.Copy().(*types.List[*DataStoreRatingInitParamWithSlot])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostParamV1); !ok {
		return false
	}

	other := o.(*DataStorePreparePostParamV1)

	if dataStorePreparePostParamV1.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePreparePostParamV1.Size.Equals(other.Size) {
		return false
	}

	if !dataStorePreparePostParamV1.Name.Equals(other.Name) {
		return false
	}

	if !dataStorePreparePostParamV1.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStorePreparePostParamV1.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dataStorePreparePostParamV1.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStorePreparePostParamV1.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dataStorePreparePostParamV1.Flag.Equals(other.Flag) {
		return false
	}

	if !dataStorePreparePostParamV1.Period.Equals(other.Period) {
		return false
	}

	if !dataStorePreparePostParamV1.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dataStorePreparePostParamV1.Tags.Equals(other.Tags) {
		return false
	}

	if !dataStorePreparePostParamV1.RatingInitParams.Equals(other.RatingInitParams) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) String() string {
	return dataStorePreparePostParamV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostParamV1{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePreparePostParamV1.StructureVersion))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStorePreparePostParamV1.Size))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dataStorePreparePostParamV1.Name))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dataStorePreparePostParamV1.DataType))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dataStorePreparePostParamV1.MetaBinary))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStorePreparePostParamV1.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStorePreparePostParamV1.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sFlag: %s,\n", indentationValues, dataStorePreparePostParamV1.Flag))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dataStorePreparePostParamV1.Period))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dataStorePreparePostParamV1.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dataStorePreparePostParamV1.Tags))
	b.WriteString(fmt.Sprintf("%sRatingInitParams: %s,\n", indentationValues, dataStorePreparePostParamV1.RatingInitParams))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostParamV1 returns a new DataStorePreparePostParamV1
func NewDataStorePreparePostParamV1() *DataStorePreparePostParamV1 {
	dataStorePreparePostParamV1 := &DataStorePreparePostParamV1{
		Size:             types.NewPrimitiveU32(0),
		Name:             types.NewString(""),
		DataType:         types.NewPrimitiveU16(0),
		MetaBinary:       types.NewQBuffer(nil),
		Permission:       NewDataStorePermission(),
		DelPermission:    NewDataStorePermission(),
		Flag:             types.NewPrimitiveU32(0),
		Period:           types.NewPrimitiveU16(0),
		ReferDataID:      types.NewPrimitiveU32(0),
		Tags:             types.NewList[*types.String](),
		RatingInitParams: types.NewList[*DataStoreRatingInitParamWithSlot](),
	}

	dataStorePreparePostParamV1.Tags.Type = types.NewString("")
	dataStorePreparePostParamV1.RatingInitParams.Type = NewDataStoreRatingInitParamWithSlot()

	return dataStorePreparePostParamV1
}
