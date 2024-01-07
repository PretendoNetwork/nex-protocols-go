// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreRatingInitParam is sent in the PreparePostObject method
type DataStoreRatingInitParam struct {
	types.Structure
	Flag           *types.PrimitiveU8
	InternalFlag   *types.PrimitiveU8
	LockType       *types.PrimitiveU8
	InitialValue   *types.PrimitiveS64
	RangeMin       *types.PrimitiveS32
	RangeMax       *types.PrimitiveS32
	PeriodHour     *types.PrimitiveS8
	PeriodDuration *types.PrimitiveS16
}

// WriteTo writes the DataStoreRatingInitParam to the given writable
func (dataStoreRatingInitParam *DataStoreRatingInitParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreRatingInitParam.Flag.WriteTo(contentWritable)
	dataStoreRatingInitParam.InternalFlag.WriteTo(contentWritable)
	dataStoreRatingInitParam.LockType.WriteTo(contentWritable)
	dataStoreRatingInitParam.InitialValue.WriteTo(contentWritable)
	dataStoreRatingInitParam.RangeMin.WriteTo(contentWritable)
	dataStoreRatingInitParam.RangeMax.WriteTo(contentWritable)
	dataStoreRatingInitParam.PeriodHour.WriteTo(contentWritable)
	dataStoreRatingInitParam.PeriodDuration.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreRatingInitParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRatingInitParam from the given readable
func (dataStoreRatingInitParam *DataStoreRatingInitParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreRatingInitParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreRatingInitParam header. %s", err.Error())
	}

	err = dataStoreRatingInitParam.Flag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.Flag. %s", err.Error())
	}

	err = dataStoreRatingInitParam.InternalFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.InternalFlag. %s", err.Error())
	}

	err = dataStoreRatingInitParam.LockType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.LockType. %s", err.Error())
	}

	err = dataStoreRatingInitParam.InitialValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.InitialValue. %s", err.Error())
	}

	err = dataStoreRatingInitParam.RangeMin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.RangeMin. %s", err.Error())
	}

	err = dataStoreRatingInitParam.RangeMax.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.RangeMax. %s", err.Error())
	}

	err = dataStoreRatingInitParam.PeriodHour.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.PeriodHour. %s", err.Error())
	}

	err = dataStoreRatingInitParam.PeriodDuration.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInitParam.PeriodDuration. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInitParam
func (dataStoreRatingInitParam *DataStoreRatingInitParam) Copy() types.RVType {
	copied := NewDataStoreRatingInitParam()

	copied.StructureVersion = dataStoreRatingInitParam.StructureVersion

	copied.Flag = dataStoreRatingInitParam.Flag.Copy().(*types.PrimitiveU8)
	copied.InternalFlag = dataStoreRatingInitParam.InternalFlag.Copy().(*types.PrimitiveU8)
	copied.LockType = dataStoreRatingInitParam.LockType.Copy().(*types.PrimitiveU8)
	copied.InitialValue = dataStoreRatingInitParam.InitialValue.Copy().(*types.PrimitiveS64)
	copied.RangeMin = dataStoreRatingInitParam.RangeMin.Copy().(*types.PrimitiveS32)
	copied.RangeMax = dataStoreRatingInitParam.RangeMax.Copy().(*types.PrimitiveS32)
	copied.PeriodHour = dataStoreRatingInitParam.PeriodHour.Copy().(*types.PrimitiveS8)
	copied.PeriodDuration = dataStoreRatingInitParam.PeriodDuration.Copy().(*types.PrimitiveS16)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInitParam *DataStoreRatingInitParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRatingInitParam); !ok {
		return false
	}

	other := o.(*DataStoreRatingInitParam)

	if dataStoreRatingInitParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreRatingInitParam.Flag.Equals(other.Flag) {
		return false
	}

	if !dataStoreRatingInitParam.InternalFlag.Equals(other.InternalFlag) {
		return false
	}

	if !dataStoreRatingInitParam.LockType.Equals(other.LockType) {
		return false
	}

	if !dataStoreRatingInitParam.InitialValue.Equals(other.InitialValue) {
		return false
	}

	if !dataStoreRatingInitParam.RangeMin.Equals(other.RangeMin) {
		return false
	}

	if !dataStoreRatingInitParam.RangeMax.Equals(other.RangeMax) {
		return false
	}

	if !dataStoreRatingInitParam.PeriodHour.Equals(other.PeriodHour) {
		return false
	}

	if !dataStoreRatingInitParam.PeriodDuration.Equals(other.PeriodDuration) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreRatingInitParam *DataStoreRatingInitParam) String() string {
	return dataStoreRatingInitParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRatingInitParam *DataStoreRatingInitParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInitParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreRatingInitParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sFlag: %s,\n", indentationValues, dataStoreRatingInitParam.Flag))
	b.WriteString(fmt.Sprintf("%sInternalFlag: %s,\n", indentationValues, dataStoreRatingInitParam.InternalFlag))
	b.WriteString(fmt.Sprintf("%sLockType: %s,\n", indentationValues, dataStoreRatingInitParam.LockType))
	b.WriteString(fmt.Sprintf("%sInitialValue: %s,\n", indentationValues, dataStoreRatingInitParam.InitialValue))
	b.WriteString(fmt.Sprintf("%sRangeMin: %s,\n", indentationValues, dataStoreRatingInitParam.RangeMin))
	b.WriteString(fmt.Sprintf("%sRangeMax: %s,\n", indentationValues, dataStoreRatingInitParam.RangeMax))
	b.WriteString(fmt.Sprintf("%sPeriodHour: %s,\n", indentationValues, dataStoreRatingInitParam.PeriodHour))
	b.WriteString(fmt.Sprintf("%sPeriodDuration: %s\n", indentationValues, dataStoreRatingInitParam.PeriodDuration))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInitParam returns a new DataStoreRatingInitParam
func NewDataStoreRatingInitParam() *DataStoreRatingInitParam {
	return &DataStoreRatingInitParam{
		Flag:           types.NewPrimitiveU8(0),
		InternalFlag:   types.NewPrimitiveU8(0),
		LockType:       types.NewPrimitiveU8(0),
		InitialValue:   types.NewPrimitiveS64(0),
		RangeMin:       types.NewPrimitiveS32(0),
		RangeMax:       types.NewPrimitiveS32(0),
		PeriodHour:     types.NewPrimitiveS8(0),
		PeriodDuration: types.NewPrimitiveS16(0),
	}
}
