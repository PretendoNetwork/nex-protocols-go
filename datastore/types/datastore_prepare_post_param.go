// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePreparePostParam is sent in the PreparePostObject method
type DataStorePreparePostParam struct {
	types.Structure
	Size                 *types.PrimitiveU32
	Name                 *types.String
	DataType             *types.PrimitiveU16
	MetaBinary           *types.QBuffer
	Permission           *DataStorePermission
	DelPermission        *DataStorePermission
	Flag                 *types.PrimitiveU32
	Period               *types.PrimitiveU16
	ReferDataID          *types.PrimitiveU32
	Tags                 *types.List[*types.String]
	RatingInitParams     *types.List[*DataStoreRatingInitParamWithSlot]
	PersistenceInitParam *DataStorePersistenceInitParam
	ExtraData            *types.List[*types.String] // NEX 3.5.0+
}

// WriteTo writes the DataStorePreparePostParam to the given writable
func (dataStorePreparePostParam *DataStorePreparePostParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	contentWritable := writable.CopyNew()

	dataStorePreparePostParam.Size.WriteTo(contentWritable)
	dataStorePreparePostParam.Name.WriteTo(contentWritable)
	dataStorePreparePostParam.DataType.WriteTo(contentWritable)
	dataStorePreparePostParam.MetaBinary.WriteTo(contentWritable)
	dataStorePreparePostParam.Permission.WriteTo(contentWritable)
	dataStorePreparePostParam.DelPermission.WriteTo(contentWritable)
	dataStorePreparePostParam.Flag.WriteTo(contentWritable)
	dataStorePreparePostParam.Period.WriteTo(contentWritable)
	dataStorePreparePostParam.ReferDataID.WriteTo(contentWritable)
	dataStorePreparePostParam.Tags.WriteTo(contentWritable)
	dataStorePreparePostParam.RatingInitParams.WriteTo(contentWritable)
	dataStorePreparePostParam.PersistenceInitParam.WriteTo(contentWritable)

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		dataStorePreparePostParam.ExtraData.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dataStorePreparePostParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePreparePostParam from the given readable
func (dataStorePreparePostParam *DataStorePreparePostParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if err = dataStorePreparePostParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePreparePostParam header. %s", err.Error())
	}

	err = dataStorePreparePostParam.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Size. %s", err.Error())
	}

	err = dataStorePreparePostParam.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Name. %s", err.Error())
	}

	err = dataStorePreparePostParam.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.DataType. %s", err.Error())
	}

	err = dataStorePreparePostParam.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.MetaBinary. %s", err.Error())
	}

	err = dataStorePreparePostParam.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Permission. %s", err.Error())
	}

	err = dataStorePreparePostParam.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.DelPermission. %s", err.Error())
	}

	err = dataStorePreparePostParam.Flag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Flag. %s", err.Error())
	}

	err = dataStorePreparePostParam.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Period. %s", err.Error())
	}

	err = dataStorePreparePostParam.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.ReferDataID. %s", err.Error())
	}

	err = dataStorePreparePostParam.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Tags. %s", err.Error())
	}

	err = dataStorePreparePostParam.RatingInitParams.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.RatingInitParams. %s", err.Error())
	}

	err = dataStorePreparePostParam.PersistenceInitParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.PersistenceInitParam. %s", err.Error())
	}

	if datastoreVersion.GreaterOrEqual("3.5.0") {
	err = 	dataStorePreparePostParam.ExtraData.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePreparePostParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStorePreparePostParam
func (dataStorePreparePostParam *DataStorePreparePostParam) Copy() types.RVType {
	copied := NewDataStorePreparePostParam()

	copied.StructureVersion = dataStorePreparePostParam.StructureVersion

	copied.Size = dataStorePreparePostParam.Size.Copy().(*types.PrimitiveU32)
	copied.Name = dataStorePreparePostParam.Name.Copy().(*types.String)
	copied.DataType = dataStorePreparePostParam.DataType.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dataStorePreparePostParam.MetaBinary.Copy().(*types.QBuffer)
	copied.Permission = dataStorePreparePostParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStorePreparePostParam.DelPermission.Copy().(*DataStorePermission)
	copied.Flag = dataStorePreparePostParam.Flag.Copy().(*types.PrimitiveU32)
	copied.Period = dataStorePreparePostParam.Period.Copy().(*types.PrimitiveU16)
	copied.ReferDataID = dataStorePreparePostParam.ReferDataID.Copy().(*types.PrimitiveU32)
	copied.Tags = dataStorePreparePostParam.Tags.Copy().(*types.List[*types.String])
	copied.RatingInitParams = dataStorePreparePostParam.RatingInitParams.Copy().(*types.List[*DataStoreRatingInitParamWithSlot])
	copied.PersistenceInitParam = dataStorePreparePostParam.PersistenceInitParam.Copy().(*DataStorePersistenceInitParam)
	copied.ExtraData = dataStorePreparePostParam.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostParam *DataStorePreparePostParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostParam); !ok {
		return false
	}

	other := o.(*DataStorePreparePostParam)

	if dataStorePreparePostParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePreparePostParam.Size.Equals(other.Size) {
		return false
	}

	if !dataStorePreparePostParam.Name.Equals(other.Name) {
		return false
	}

	if !dataStorePreparePostParam.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStorePreparePostParam.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dataStorePreparePostParam.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStorePreparePostParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dataStorePreparePostParam.Flag.Equals(other.Flag) {
		return false
	}

	if !dataStorePreparePostParam.Period.Equals(other.Period) {
		return false
	}

	if !dataStorePreparePostParam.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dataStorePreparePostParam.Tags.Equals(other.Tags) {
		return false
	}

	if !dataStorePreparePostParam.RatingInitParams.Equals(other.RatingInitParams) {
		return false
	}

	if !dataStorePreparePostParam.PersistenceInitParam.Equals(other.PersistenceInitParam) {
		return false
	}

	if !dataStorePreparePostParam.ExtraData.Equals(other.ExtraData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePreparePostParam *DataStorePreparePostParam) String() string {
	return dataStorePreparePostParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePreparePostParam *DataStorePreparePostParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePreparePostParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStorePreparePostParam.Size))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dataStorePreparePostParam.Name))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dataStorePreparePostParam.DataType))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dataStorePreparePostParam.MetaBinary))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStorePreparePostParam.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStorePreparePostParam.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sFlag: %s,\n", indentationValues, dataStorePreparePostParam.Flag))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dataStorePreparePostParam.Period))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dataStorePreparePostParam.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dataStorePreparePostParam.Tags))
	b.WriteString(fmt.Sprintf("%sRatingInitParams: %s,\n", indentationValues, dataStorePreparePostParam.RatingInitParams))
	b.WriteString(fmt.Sprintf("%sPersistenceInitParam: %s,\n", indentationValues, dataStorePreparePostParam.PersistenceInitParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, dataStorePreparePostParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostParam returns a new DataStorePreparePostParam
func NewDataStorePreparePostParam() *DataStorePreparePostParam {
	dataStorePreparePostParam := &DataStorePreparePostParam{
		Size:                 types.NewPrimitiveU32(0),
		Name:                 types.NewString(""),
		DataType:             types.NewPrimitiveU16(0),
		MetaBinary:           types.NewQBuffer(nil),
		Permission:           NewDataStorePermission(),
		DelPermission:        NewDataStorePermission(),
		Flag:                 types.NewPrimitiveU32(0),
		Period:               types.NewPrimitiveU16(0),
		ReferDataID:          types.NewPrimitiveU32(0),
		Tags:                 types.NewList[*types.String](),
		RatingInitParams:     types.NewList[*DataStoreRatingInitParamWithSlot](),
		PersistenceInitParam: NewDataStorePersistenceInitParam(),
		ExtraData:            types.NewList[*types.String](),
	}

	dataStorePreparePostParam.Tags.Type = types.NewString("")
	dataStorePreparePostParam.RatingInitParams.Type = NewDataStoreRatingInitParamWithSlot()
	dataStorePreparePostParam.ExtraData.Type = types.NewString("")

	return dataStorePreparePostParam
}
