// Package types implements all the types used by the Rating protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RatingRankData is a type within the Rating protocol
type RatingRankData struct {
	types.Structure
	PrincipalID types.PID
	UniqueID    types.UInt64
	Order       types.UInt32
	Category    types.UInt32
	Score       types.UInt32
	Unknown1    types.UInt32
	Unknown2    types.UInt32
	Unknown3    types.UInt32
	CommonData  types.Buffer
	UpdateTime  types.DateTime
}

// WriteTo writes the RatingRankData to the given writable
func (rrd RatingRankData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rrd.PrincipalID.WriteTo(contentWritable)
	rrd.UniqueID.WriteTo(contentWritable)
	rrd.Order.WriteTo(contentWritable)
	rrd.Category.WriteTo(contentWritable)
	rrd.Score.WriteTo(contentWritable)
	rrd.Unknown1.WriteTo(contentWritable)
	rrd.Unknown2.WriteTo(contentWritable)
	rrd.Unknown3.WriteTo(contentWritable)
	rrd.CommonData.WriteTo(contentWritable)
	rrd.UpdateTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rrd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RatingRankData from the given readable
func (rrd *RatingRankData) ExtractFrom(readable types.Readable) error {
	var err error

	err = rrd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData header. %s", err.Error())
	}

	err = rrd.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.PrincipalID. %s", err.Error())
	}

	err = rrd.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.UniqueID. %s", err.Error())
	}

	err = rrd.Order.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.Order. %s", err.Error())
	}

	err = rrd.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.Category. %s", err.Error())
	}

	err = rrd.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.Score. %s", err.Error())
	}

	err = rrd.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.Unknown1. %s", err.Error())
	}

	err = rrd.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.Unknown2. %s", err.Error())
	}

	err = rrd.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.Unknown3. %s", err.Error())
	}

	err = rrd.CommonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.CommonData. %s", err.Error())
	}

	err = rrd.UpdateTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingRankData.UpdateTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RatingRankData
func (rrd RatingRankData) Copy() types.RVType {
	copied := NewRatingRankData()

	copied.StructureVersion = rrd.StructureVersion
	copied.PrincipalID = rrd.PrincipalID.Copy().(types.PID)
	copied.UniqueID = rrd.UniqueID.Copy().(types.UInt64)
	copied.Order = rrd.Order.Copy().(types.UInt32)
	copied.Category = rrd.Category.Copy().(types.UInt32)
	copied.Score = rrd.Score.Copy().(types.UInt32)
	copied.Unknown1 = rrd.Unknown1.Copy().(types.UInt32)
	copied.Unknown2 = rrd.Unknown2.Copy().(types.UInt32)
	copied.Unknown3 = rrd.Unknown3.Copy().(types.UInt32)
	copied.CommonData = rrd.CommonData.Copy().(types.Buffer)
	copied.UpdateTime = rrd.UpdateTime.Copy().(types.DateTime)

	return copied
}

// Equals checks if the given RatingRankData contains the same data as the current RatingRankData
func (rrd RatingRankData) Equals(o types.RVType) bool {
	if _, ok := o.(RatingRankData); !ok {
		return false
	}

	other := o.(RatingRankData)

	if rrd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rrd.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !rrd.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !rrd.Order.Equals(other.Order) {
		return false
	}

	if !rrd.Category.Equals(other.Category) {
		return false
	}

	if !rrd.Score.Equals(other.Score) {
		return false
	}

	if !rrd.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !rrd.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !rrd.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !rrd.CommonData.Equals(other.CommonData) {
		return false
	}

	return rrd.UpdateTime.Equals(other.UpdateTime)
}

// CopyRef copies the current value of the RatingRankData
// and returns a pointer to the new copy
func (rrd RatingRankData) CopyRef() types.RVTypePtr {
	copied := rrd.Copy().(RatingRankData)
	return &copied
}

// Deref takes a pointer to the RatingRankData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rrd *RatingRankData) Deref() types.RVType {
	return *rrd
}

// String returns the string representation of the RatingRankData
func (rrd RatingRankData) String() string {
	return rrd.FormatToString(0)
}

// FormatToString pretty-prints the RatingRankData using the provided indentation level
func (rrd RatingRankData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RatingRankData{\n")
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, rrd.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, rrd.UniqueID))
	b.WriteString(fmt.Sprintf("%sOrder: %s,\n", indentationValues, rrd.Order))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rrd.Category))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, rrd.Score))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, rrd.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, rrd.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, rrd.Unknown3))
	b.WriteString(fmt.Sprintf("%sCommonData: %s,\n", indentationValues, rrd.CommonData))
	b.WriteString(fmt.Sprintf("%sUpdateTime: %s,\n", indentationValues, rrd.UpdateTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRatingRankData returns a new RatingRankData
func NewRatingRankData() RatingRankData {
	return RatingRankData{
		PrincipalID: types.NewPID(0),
		UniqueID:    types.NewUInt64(0),
		Order:       types.NewUInt32(0),
		Category:    types.NewUInt32(0),
		Score:       types.NewUInt32(0),
		Unknown1:    types.NewUInt32(0),
		Unknown2:    types.NewUInt32(0),
		Unknown3:    types.NewUInt32(0),
		CommonData:  types.NewBuffer(nil),
		UpdateTime:  types.NewDateTime(0),
	}

}
