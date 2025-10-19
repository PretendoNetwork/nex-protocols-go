// Package types implements all the types used by the Rating protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RatingStats is a type within the Rating protocol
type RatingStats struct {
	types.Structure
	PrincipalID types.PID
	UniqueID    types.UInt64
	Flags       types.UInt32
	Category    types.UInt32
	ReportData  types.QBuffer
	Values      types.List[types.Float]
}

// WriteTo writes the RatingStats to the given writable
func (rs RatingStats) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rs.PrincipalID.WriteTo(contentWritable)
	rs.UniqueID.WriteTo(contentWritable)
	rs.Flags.WriteTo(contentWritable)
	rs.Category.WriteTo(contentWritable)
	rs.ReportData.WriteTo(contentWritable)
	rs.Values.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RatingStats from the given readable
func (rs *RatingStats) ExtractFrom(readable types.Readable) error {
	var err error

	err = rs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingStats header. %s", err.Error())
	}

	err = rs.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingStats.PrincipalID. %s", err.Error())
	}

	err = rs.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingStats.UniqueID. %s", err.Error())
	}

	err = rs.Flags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingStats.Flags. %s", err.Error())
	}

	err = rs.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingStats.Category. %s", err.Error())
	}

	err = rs.ReportData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingStats.ReportData. %s", err.Error())
	}

	err = rs.Values.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingStats.Values. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RatingStats
func (rs RatingStats) Copy() types.RVType {
	copied := NewRatingStats()

	copied.StructureVersion = rs.StructureVersion
	copied.PrincipalID = rs.PrincipalID.Copy().(types.PID)
	copied.UniqueID = rs.UniqueID.Copy().(types.UInt64)
	copied.Flags = rs.Flags.Copy().(types.UInt32)
	copied.Category = rs.Category.Copy().(types.UInt32)
	copied.ReportData = rs.ReportData.Copy().(types.QBuffer)
	copied.Values = rs.Values.Copy().(types.List[types.Float])

	return copied
}

// Equals checks if the given RatingStats contains the same data as the current RatingStats
func (rs RatingStats) Equals(o types.RVType) bool {
	if _, ok := o.(RatingStats); !ok {
		return false
	}

	other := o.(RatingStats)

	if rs.StructureVersion != other.StructureVersion {
		return false
	}

	if !rs.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !rs.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !rs.Flags.Equals(other.Flags) {
		return false
	}

	if !rs.Category.Equals(other.Category) {
		return false
	}

	if !rs.ReportData.Equals(other.ReportData) {
		return false
	}

	return rs.Values.Equals(other.Values)
}

// CopyRef copies the current value of the RatingStats
// and returns a pointer to the new copy
func (rs RatingStats) CopyRef() types.RVTypePtr {
	copied := rs.Copy().(RatingStats)
	return &copied
}

// Deref takes a pointer to the RatingStats
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rs *RatingStats) Deref() types.RVType {
	return *rs
}

// String returns the string representation of the RatingStats
func (rs RatingStats) String() string {
	return rs.FormatToString(0)
}

// FormatToString pretty-prints the RatingStats using the provided indentation level
func (rs RatingStats) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RatingStats{\n")
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, rs.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, rs.UniqueID))
	b.WriteString(fmt.Sprintf("%sFlags: %s,\n", indentationValues, rs.Flags))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rs.Category))
	b.WriteString(fmt.Sprintf("%sReportData: %s,\n", indentationValues, rs.ReportData))
	b.WriteString(fmt.Sprintf("%sValues: %s,\n", indentationValues, rs.Values))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRatingStats returns a new RatingStats
func NewRatingStats() RatingStats {
	return RatingStats{
		PrincipalID: types.NewPID(0),
		UniqueID:    types.NewUInt64(0),
		Flags:       types.NewUInt32(0),
		Category:    types.NewUInt32(0),
		ReportData:  types.NewQBuffer(nil),
		Values:      types.NewList[types.Float](),
	}

}
