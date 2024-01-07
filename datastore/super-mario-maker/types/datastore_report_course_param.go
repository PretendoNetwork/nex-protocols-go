// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReportCourseParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreReportCourseParam struct {
	types.Structure
	DataID         *types.PrimitiveU64
	MiiName        *types.String
	ReportCategory *types.PrimitiveU8
	ReportReason   *types.String
}

// ExtractFrom extracts the DataStoreReportCourseParam from the given readable
func (dataStoreReportCourseParam *DataStoreReportCourseParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreReportCourseParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReportCourseParam header. %s", err.Error())
	}

	err = dataStoreReportCourseParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.DataID from stream. %s", err.Error())
	}

	err = dataStoreReportCourseParam.MiiName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.MiiName from stream. %s", err.Error())
	}

	err = dataStoreReportCourseParam.ReportCategory.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.ReportCategory from stream. %s", err.Error())
	}

	err = dataStoreReportCourseParam.ReportReason.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.ReportReason from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReportCourseParam to the given writable
func (dataStoreReportCourseParam *DataStoreReportCourseParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreReportCourseParam.DataID.WriteTo(contentWritable)
	dataStoreReportCourseParam.MiiName.WriteTo(contentWritable)
	dataStoreReportCourseParam.ReportCategory.WriteTo(contentWritable)
	dataStoreReportCourseParam.ReportReason.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreReportCourseParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReportCourseParam
func (dataStoreReportCourseParam *DataStoreReportCourseParam) Copy() types.RVType {
	copied := NewDataStoreReportCourseParam()

	copied.StructureVersion = dataStoreReportCourseParam.StructureVersion

	copied.DataID = dataStoreReportCourseParam.DataID.Copy().(*types.PrimitiveU64)
	copied.MiiName = dataStoreReportCourseParam.MiiName.Copy().(*types.String)
	copied.ReportCategory = dataStoreReportCourseParam.ReportCategory.Copy().(*types.PrimitiveU8)
	copied.ReportReason = dataStoreReportCourseParam.ReportReason.Copy().(*types.String)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReportCourseParam *DataStoreReportCourseParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReportCourseParam); !ok {
		return false
	}

	other := o.(*DataStoreReportCourseParam)

	if dataStoreReportCourseParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReportCourseParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreReportCourseParam.MiiName.Equals(other.MiiName) {
		return false
	}

	if !dataStoreReportCourseParam.ReportCategory.Equals(other.ReportCategory) {
		return false
	}

	if !dataStoreReportCourseParam.ReportReason.Equals(other.ReportReason) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReportCourseParam *DataStoreReportCourseParam) String() string {
	return dataStoreReportCourseParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReportCourseParam *DataStoreReportCourseParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReportCourseParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReportCourseParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreReportCourseParam.DataID))
	b.WriteString(fmt.Sprintf("%sMiiName: %s,\n", indentationValues, dataStoreReportCourseParam.MiiName))
	b.WriteString(fmt.Sprintf("%sReportCategory: %s,\n", indentationValues, dataStoreReportCourseParam.ReportCategory))
	b.WriteString(fmt.Sprintf("%sReportReason: %s,\n", indentationValues, dataStoreReportCourseParam.ReportReason))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReportCourseParam returns a new DataStoreReportCourseParam
func NewDataStoreReportCourseParam() *DataStoreReportCourseParam {
	return &DataStoreReportCourseParam{
		DataID:         types.NewPrimitiveU64(0),
		MiiName:        types.NewString(""),
		ReportCategory: types.NewPrimitiveU8(0),
		ReportReason:   types.NewString(""),
	}
}
