// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/datastore/constants"
)

// DataStoreRatingInitParam is a type within the DataStore protocol
type DataStoreRatingInitParam struct {
	types.Structure
	Flag           constants.RatingFlag
	InternalFlag   constants.RatingInternalFlag
	LockType       constants.RatingLockType
	InitialValue   types.Int64
	RangeMin       types.Int32
	RangeMax       types.Int32
	PeriodHour     types.Int8
	PeriodDuration constants.RatingLockPeriodDay
}

// WriteTo writes the DataStoreRatingInitParam to the given writable
func (dsrip DataStoreRatingInitParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrip.Flag.WriteTo(contentWritable)
	dsrip.InternalFlag.WriteTo(contentWritable)
	dsrip.LockType.WriteTo(contentWritable)
	dsrip.InitialValue.WriteTo(contentWritable)
	dsrip.RangeMin.WriteTo(contentWritable)
	dsrip.RangeMax.WriteTo(contentWritable)
	dsrip.PeriodHour.WriteTo(contentWritable)
	dsrip.PeriodDuration.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingInitParam from the given readable
func (dsrip *DataStoreRatingInitParam) ExtractFrom(readable types.Readable) error {
	if err := dsrip.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam header. %s", err.Error())
	}

	if err := dsrip.Flag.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.Flag. %s", err.Error())
	}

	if err := dsrip.InternalFlag.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.InternalFlag. %s", err.Error())
	}

	if err := dsrip.LockType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.LockType. %s", err.Error())
	}

	if err := dsrip.InitialValue.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.InitialValue. %s", err.Error())
	}

	if err := dsrip.RangeMin.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.RangeMin. %s", err.Error())
	}

	if err := dsrip.RangeMax.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.RangeMax. %s", err.Error())
	}

	if err := dsrip.PeriodHour.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.PeriodHour. %s", err.Error())
	}

	if err := dsrip.PeriodDuration.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.PeriodDuration. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInitParam
func (dsrip DataStoreRatingInitParam) Copy() types.RVType {
	copied := NewDataStoreRatingInitParam()

	copied.StructureVersion = dsrip.StructureVersion
	copied.Flag = dsrip.Flag
	copied.InternalFlag = dsrip.InternalFlag
	copied.LockType = dsrip.LockType
	copied.InitialValue = dsrip.InitialValue.Copy().(types.Int64)
	copied.RangeMin = dsrip.RangeMin.Copy().(types.Int32)
	copied.RangeMax = dsrip.RangeMax.Copy().(types.Int32)
	copied.PeriodHour = dsrip.PeriodHour.Copy().(types.Int8)
	copied.PeriodDuration = dsrip.PeriodDuration

	return copied
}

// Equals checks if the given DataStoreRatingInitParam contains the same data as the current DataStoreRatingInitParam
func (dsrip DataStoreRatingInitParam) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreRatingInitParam); !ok {
		return false
	}

	other := o.(DataStoreRatingInitParam)

	if dsrip.StructureVersion != other.StructureVersion {
		return false
	}

	if dsrip.Flag != other.Flag {
		return false
	}

	if dsrip.InternalFlag != other.InternalFlag {
		return false
	}

	if dsrip.LockType != other.LockType {
		return false
	}

	if !dsrip.InitialValue.Equals(other.InitialValue) {
		return false
	}

	if !dsrip.RangeMin.Equals(other.RangeMin) {
		return false
	}

	if !dsrip.RangeMax.Equals(other.RangeMax) {
		return false
	}

	if !dsrip.PeriodHour.Equals(other.PeriodHour) {
		return false
	}

	return dsrip.PeriodDuration == other.PeriodDuration
}

// CopyRef copies the current value of the DataStoreRatingInitParam
// and returns a pointer to the new copy
func (dsrip DataStoreRatingInitParam) CopyRef() types.RVTypePtr {
	copied := dsrip.Copy().(DataStoreRatingInitParam)
	return &copied
}

// Deref takes a pointer to the DataStoreRatingInitParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dsrip *DataStoreRatingInitParam) Deref() types.RVType {
	return *dsrip
}

// String returns the string representation of the DataStoreRatingInitParam
func (dsrip DataStoreRatingInitParam) String() string {
	return dsrip.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRatingInitParam using the provided indentation level
func (dsrip DataStoreRatingInitParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInitParam{\n")
	b.WriteString(fmt.Sprintf("%sFlag: %s,\n", indentationValues, dsrip.Flag))
	b.WriteString(fmt.Sprintf("%sInternalFlag: %s,\n", indentationValues, dsrip.InternalFlag))
	b.WriteString(fmt.Sprintf("%sLockType: %s,\n", indentationValues, dsrip.LockType))
	b.WriteString(fmt.Sprintf("%sInitialValue: %s,\n", indentationValues, dsrip.InitialValue))
	b.WriteString(fmt.Sprintf("%sRangeMin: %s,\n", indentationValues, dsrip.RangeMin))
	b.WriteString(fmt.Sprintf("%sRangeMax: %s,\n", indentationValues, dsrip.RangeMax))
	b.WriteString(fmt.Sprintf("%sPeriodHour: %s,\n", indentationValues, dsrip.PeriodHour))
	b.WriteString(fmt.Sprintf("%sPeriodDuration: %s,\n", indentationValues, dsrip.PeriodDuration))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInitParam returns a new DataStoreRatingInitParam
func NewDataStoreRatingInitParam() DataStoreRatingInitParam {
	return DataStoreRatingInitParam{
		Flag:           constants.RatingFlag(0),         // TODO - What is the real default?
		InternalFlag:   constants.RatingInternalFlag(0), // TODO - What is the real default?
		LockType:       constants.RatingLockNone,
		InitialValue:   types.NewInt64(0),
		RangeMin:       types.NewInt32(0),
		RangeMax:       types.NewInt32(0),
		PeriodHour:     types.NewInt8(0),
		PeriodDuration: constants.RatingLockPeriodDay(0), // TODO - What is the real default?
	}

}
