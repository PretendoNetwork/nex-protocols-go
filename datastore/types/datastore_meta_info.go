// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreMetaInfo contains DataStore meta information
type DataStoreMetaInfo struct {
	types.Structure
	DataID        *types.PrimitiveU64
	OwnerID       *types.PID
	Size          *types.PrimitiveU32
	DataType      *types.PrimitiveU16
	Name          *types.String
	MetaBinary    *types.QBuffer
	Permission    *DataStorePermission
	DelPermission *DataStorePermission
	CreatedTime   *types.DateTime
	UpdatedTime   *types.DateTime
	Period        *types.PrimitiveU16
	Status        *types.PrimitiveU8
	ReferredCnt   *types.PrimitiveU32
	ReferDataID   *types.PrimitiveU32
	Flag          *types.PrimitiveU32
	ReferredTime  *types.DateTime
	ExpireTime    *types.DateTime
	Tags          *types.List[*types.String]
	Ratings       *types.List[*DataStoreRatingInfoWithSlot]
}

// ExtractFrom extracts the DataStoreMetaInfo from the given readable
func (dataStoreMetaInfo *DataStoreMetaInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreMetaInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreMetaInfo header. %s", err.Error())
	}

	err = dataStoreMetaInfo.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DataID. %s", err.Error())
	}

	err = dataStoreMetaInfo.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.OwnerID. %s", err.Error())
	}

	err = dataStoreMetaInfo.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Size. %s", err.Error())
	}

	err = dataStoreMetaInfo.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Name. %s", err.Error())
	}

	err = dataStoreMetaInfo.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DataType. %s", err.Error())
	}

	err = dataStoreMetaInfo.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.MetaBinary. %s", err.Error())
	}

	err = dataStoreMetaInfo.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Permission. %s", err.Error())
	}

	err = dataStoreMetaInfo.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DelPermission. %s", err.Error())
	}

	err = dataStoreMetaInfo.CreatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.CreatedTime. %s", err.Error())
	}

	err = dataStoreMetaInfo.UpdatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.UpdatedTime. %s", err.Error())
	}

	err = dataStoreMetaInfo.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Period. %s", err.Error())
	}

	err = dataStoreMetaInfo.Status.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Status. %s", err.Error())
	}

	err = dataStoreMetaInfo.ReferredCnt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferredCnt. %s", err.Error())
	}

	err = dataStoreMetaInfo.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferDataID. %s", err.Error())
	}

	err = dataStoreMetaInfo.Flag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Flag. %s", err.Error())
	}

	err = dataStoreMetaInfo.ReferredTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferredTime. %s", err.Error())
	}

	err = dataStoreMetaInfo.ExpireTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ExpireTime. %s", err.Error())
	}

	err = dataStoreMetaInfo.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Tags. %s", err.Error())
	}

	err = dataStoreMetaInfo.Ratings.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Ratings. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreMetaInfo to the given writable
func (dataStoreMetaInfo *DataStoreMetaInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreMetaInfo.DataID.WriteTo(contentWritable)
	dataStoreMetaInfo.OwnerID.WriteTo(contentWritable)
	dataStoreMetaInfo.Size.WriteTo(contentWritable)
	dataStoreMetaInfo.Name.WriteTo(contentWritable)
	dataStoreMetaInfo.DataType.WriteTo(contentWritable)
	dataStoreMetaInfo.MetaBinary.WriteTo(contentWritable)
	dataStoreMetaInfo.Permission.WriteTo(contentWritable)
	dataStoreMetaInfo.DelPermission.WriteTo(contentWritable)
	dataStoreMetaInfo.CreatedTime.WriteTo(contentWritable)
	dataStoreMetaInfo.UpdatedTime.WriteTo(contentWritable)
	dataStoreMetaInfo.Period.WriteTo(contentWritable)
	dataStoreMetaInfo.Status.WriteTo(contentWritable)
	dataStoreMetaInfo.ReferredCnt.WriteTo(contentWritable)
	dataStoreMetaInfo.ReferDataID.WriteTo(contentWritable)
	dataStoreMetaInfo.Flag.WriteTo(contentWritable)
	dataStoreMetaInfo.ReferredTime.WriteTo(contentWritable)
	dataStoreMetaInfo.ExpireTime.WriteTo(contentWritable)
	dataStoreMetaInfo.Tags.WriteTo(contentWritable)
	dataStoreMetaInfo.Ratings.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreMetaInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreMetaInfo
func (dataStoreMetaInfo *DataStoreMetaInfo) Copy() types.RVType {
	copied := NewDataStoreMetaInfo()

	copied.StructureVersion = dataStoreMetaInfo.StructureVersion

	copied.DataID = dataStoreMetaInfo.DataID.Copy().(*types.PrimitiveU64)
	copied.OwnerID = dataStoreMetaInfo.OwnerID.Copy().(*types.PID)
	copied.Size = dataStoreMetaInfo.Size.Copy().(*types.PrimitiveU32)
	copied.DataType = dataStoreMetaInfo.DataType.Copy().(*types.PrimitiveU16)
	copied.Name = dataStoreMetaInfo.Name.Copy().(*types.String)
	copied.MetaBinary = dataStoreMetaInfo.MetaBinary.Copy().(*types.QBuffer)

	copied.Permission = dataStoreMetaInfo.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreMetaInfo.DelPermission.Copy().(*DataStorePermission)
	copied.CreatedTime = dataStoreMetaInfo.CreatedTime.Copy().(*types.DateTime)
	copied.UpdatedTime = dataStoreMetaInfo.UpdatedTime.Copy().(*types.DateTime)
	copied.Period = dataStoreMetaInfo.Period.Copy().(*types.PrimitiveU16)
	copied.Status = dataStoreMetaInfo.Status.Copy().(*types.PrimitiveU8)
	copied.ReferredCnt = dataStoreMetaInfo.ReferredCnt.Copy().(*types.PrimitiveU32)
	copied.ReferDataID = dataStoreMetaInfo.ReferDataID.Copy().(*types.PrimitiveU32)
	copied.Flag = dataStoreMetaInfo.Flag.Copy().(*types.PrimitiveU32)
	copied.ReferredTime = dataStoreMetaInfo.ReferredTime.Copy().(*types.DateTime)
	copied.ExpireTime = dataStoreMetaInfo.ExpireTime.Copy().(*types.DateTime)
	copied.Tags = dataStoreMetaInfo.Tags.Copy().(*types.List[*types.String])

	copied.Ratings = dataStoreMetaInfo.Ratings.Copy().(*types.List[*DataStoreRatingInfoWithSlot])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreMetaInfo *DataStoreMetaInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreMetaInfo); !ok {
		return false
	}

	other := o.(*DataStoreMetaInfo)

	if dataStoreMetaInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreMetaInfo.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreMetaInfo.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dataStoreMetaInfo.Size.Equals(other.Size) {
		return false
	}

	if !dataStoreMetaInfo.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStoreMetaInfo.Name.Equals(other.Name) {
		return false
	}

	if !dataStoreMetaInfo.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dataStoreMetaInfo.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStoreMetaInfo.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dataStoreMetaInfo.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !dataStoreMetaInfo.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	if !dataStoreMetaInfo.Period.Equals(other.Period) {
		return false
	}

	if !dataStoreMetaInfo.Status.Equals(other.Status) {
		return false
	}

	if !dataStoreMetaInfo.ReferredCnt.Equals(other.ReferredCnt) {
		return false
	}

	if !dataStoreMetaInfo.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dataStoreMetaInfo.Flag.Equals(other.Flag) {
		return false
	}

	if !dataStoreMetaInfo.ReferredTime.Equals(other.ReferredTime) {
		return false
	}

	if !dataStoreMetaInfo.ExpireTime.Equals(other.ExpireTime) {
		return false
	}

	if !dataStoreMetaInfo.Tags.Equals(other.Tags) {
		return false
	}

	if !dataStoreMetaInfo.Ratings.Equals(other.Ratings) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreMetaInfo *DataStoreMetaInfo) String() string {
	return dataStoreMetaInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreMetaInfo *DataStoreMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreMetaInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreMetaInfo.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dataStoreMetaInfo.OwnerID))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStoreMetaInfo.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dataStoreMetaInfo.DataType))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dataStoreMetaInfo.Name))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dataStoreMetaInfo.MetaBinary))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStoreMetaInfo.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStoreMetaInfo.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCreatedTime: %s,\n", indentationValues, dataStoreMetaInfo.CreatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedTime: %s,\n", indentationValues, dataStoreMetaInfo.UpdatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dataStoreMetaInfo.Period))
	b.WriteString(fmt.Sprintf("%sStatus: %s,\n", indentationValues, dataStoreMetaInfo.Status))
	b.WriteString(fmt.Sprintf("%sReferredCnt: %s,\n", indentationValues, dataStoreMetaInfo.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dataStoreMetaInfo.ReferDataID))
	b.WriteString(fmt.Sprintf("%sFlag: %s,\n", indentationValues, dataStoreMetaInfo.Flag))
	b.WriteString(fmt.Sprintf("%sReferredTime: %s,\n", indentationValues, dataStoreMetaInfo.ReferredTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sExpireTime: %s,\n", indentationValues, dataStoreMetaInfo.ExpireTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dataStoreMetaInfo.Tags))
	b.WriteString(fmt.Sprintf("%sRatings: %s,\n", indentationValues, dataStoreMetaInfo.Ratings))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// FilterPropertiesByResultOption zeroes out certain struct properties based on the input flags
func (dataStoreMetaInfo *DataStoreMetaInfo) FilterPropertiesByResultOption(resultOption *types.PrimitiveU8) {
	// * This is kind of backwards
	// *
	// * This method assumes all struct data exists
	// * by default. This is done in order to simplify
	// * database calls by just querying for all fields
	// * at once. Therefore, instead of the ResultOption
	// * flags being used to conditionally ADD properties,
	// * it's used to conditionally REMOVE them

	if resultOption&0x1 == 0 {
		dataStoreMetaInfo.Tags = types.NewList[*types.String]()
		dataStoreMetaInfo.Tags.Type = types.NewString("")
	}

	if resultOption&0x2 == 0 {
		dataStoreMetaInfo.Ratings = types.NewList[*DataStoreRatingInfoWithSlot]()
		dataStoreMetaInfo.Ratings.Type = NewDataStoreRatingInfoWithSlot()
	}

	if resultOption&0x4 == 0 {
		dataStoreMetaInfo.MetaBinary = types.NewQBuffer(nil)
	}
}

// NewDataStoreMetaInfo returns a new DataStoreMetaInfo
func NewDataStoreMetaInfo() *DataStoreMetaInfo {
	dataStoreMetaInfo := &DataStoreMetaInfo{
		DataID:        types.NewPrimitiveU64(0),
		OwnerID:       types.NewPID(0),
		Size:          types.NewPrimitiveU32(0),
		DataType:      types.NewPrimitiveU16(0),
		Name:          types.NewString(""),
		MetaBinary:    types.NewQBuffer(nil),
		Permission:    NewDataStorePermission(),
		DelPermission: NewDataStorePermission(),
		CreatedTime:   types.NewDateTime(0),
		UpdatedTime:   types.NewDateTime(0),
		Period:        types.NewPrimitiveU16(0),
		Status:        types.NewPrimitiveU8(0),
		ReferredCnt:   types.NewPrimitiveU32(0),
		ReferDataID:   types.NewPrimitiveU32(0),
		Flag:          types.NewPrimitiveU32(0),
		ReferredTime:  types.NewDateTime(0),
		ExpireTime:    types.NewDateTime(0),
		Tags:          types.NewList[*types.String](),
		Ratings:       types.NewList[*DataStoreRatingInfoWithSlot](),
	}

	dataStoreMetaInfo.Tags.Type = types.NewString("")
	dataStoreMetaInfo.Ratings.Type = NewDataStoreRatingInfoWithSlot()

	return dataStoreMetaInfo
}
