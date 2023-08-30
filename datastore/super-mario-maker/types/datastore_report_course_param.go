// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReportCourseParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreReportCourseParam struct {
	nex.Structure
	DataID         uint64
	MiiName        string
	ReportCategory uint8
	ReportReason   string
}

// ExtractFromStream extracts a DataStoreReportCourseParam structure from a stream
func (dataStoreReportCourseParam *DataStoreReportCourseParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReportCourseParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.DataID from stream. %s", err.Error())
	}

	dataStoreReportCourseParam.MiiName, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.MiiName from stream. %s", err.Error())
	}

	dataStoreReportCourseParam.ReportCategory, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.ReportCategory from stream. %s", err.Error())
	}

	dataStoreReportCourseParam.ReportReason, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.ReportReason from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReportCourseParam and returns a byte array
func (dataStoreReportCourseParam *DataStoreReportCourseParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreReportCourseParam.DataID)
	stream.WriteString(dataStoreReportCourseParam.MiiName)
	stream.WriteUInt8(dataStoreReportCourseParam.ReportCategory)
	stream.WriteString(dataStoreReportCourseParam.ReportReason)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReportCourseParam
func (dataStoreReportCourseParam *DataStoreReportCourseParam) Copy() nex.StructureInterface {
	copied := NewDataStoreReportCourseParam()

	copied.SetStructureVersion(dataStoreReportCourseParam.StructureVersion())

	copied.DataID = dataStoreReportCourseParam.DataID
	copied.MiiName = dataStoreReportCourseParam.MiiName
	copied.ReportCategory = dataStoreReportCourseParam.ReportCategory
	copied.ReportReason = dataStoreReportCourseParam.ReportReason

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReportCourseParam *DataStoreReportCourseParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReportCourseParam)

	if dataStoreReportCourseParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreReportCourseParam.DataID != other.DataID {
		return false
	}

	if dataStoreReportCourseParam.MiiName != other.MiiName {
		return false
	}

	if dataStoreReportCourseParam.ReportCategory != other.ReportCategory {
		return false
	}

	if dataStoreReportCourseParam.ReportReason != other.ReportReason {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReportCourseParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreReportCourseParam.DataID))
	b.WriteString(fmt.Sprintf("%sMiiName: %q,\n", indentationValues, dataStoreReportCourseParam.MiiName))
	b.WriteString(fmt.Sprintf("%sReportCategory: %d,\n", indentationValues, dataStoreReportCourseParam.ReportCategory))
	b.WriteString(fmt.Sprintf("%sReportReason: %q,\n", indentationValues, dataStoreReportCourseParam.ReportReason))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReportCourseParam returns a new DataStoreReportCourseParam
func NewDataStoreReportCourseParam() *DataStoreReportCourseParam {
	return &DataStoreReportCourseParam{}
}
