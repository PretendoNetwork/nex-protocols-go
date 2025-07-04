// Package types implements all the types used by the legacy Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingData is a type within the Ranking protocol
type RankingData struct {
	types.Structure
	UniqueID    types.UInt32
	PrincipalID types.PID
	Order       types.UInt32
	Category    types.UInt32
	Scores      types.List[types.UInt32]
	Unknown1    types.UInt8
	Unknown2    types.UInt32
	CommonData  types.Buffer
}

// WriteTo writes the RankingData to the given writable
func (rd RankingData) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.Ranking

	contentWritable := writable.CopyNew()

	rd.PrincipalID.WriteTo(contentWritable)
	rd.UniqueID.WriteTo(contentWritable)
	rd.Order.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("2.0.0") {
		rd.Category.WriteTo(contentWritable)
	} else {
		category := types.List[types.UInt16]{types.UInt16(rd.Category)}
		category.WriteTo(contentWritable)
	}

	rd.Category.WriteTo(contentWritable)
	rd.Scores.WriteTo(contentWritable)
	rd.Unknown1.WriteTo(contentWritable)
	rd.Unknown2.WriteTo(contentWritable)
	rd.CommonData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingData from the given readable
func (rd *RankingData) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.Ranking

	var err error

	err = rd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData header. %s", err.Error())
	}

	err = rd.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.PrincipalID. %s", err.Error())
	}

	err = rd.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.UniqueID. %s", err.Error())
	}

	err = rd.Order.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Order. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("2.0.0") {
		err = rd.Category.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract RankingData.Category. %s", err.Error())
		}
	} else {
		var category types.List[types.UInt16]

		err = category.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract RankingData.Category. %s", err.Error())
		}

		if len(category) != 1 {
			return fmt.Errorf("Failed to extract RankingData.Category. Expected length of 1, got %d", len(category))
		}

		rd.Category = types.UInt32(category[0])
	}

	err = rd.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Category. %s", err.Error())
	}

	err = rd.Scores.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Scores. %s", err.Error())
	}

	err = rd.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Unknown1. %s", err.Error())
	}

	err = rd.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Unknown2. %s", err.Error())
	}

	err = rd.CommonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.CommonData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingData
func (rd RankingData) Copy() types.RVType {
	copied := NewRankingData()

	copied.StructureVersion = rd.StructureVersion
	copied.PrincipalID = rd.PrincipalID.Copy().(types.PID)
	copied.UniqueID = rd.UniqueID.Copy().(types.UInt32)
	copied.Order = rd.Order.Copy().(types.UInt32)
	copied.Category = rd.Category.Copy().(types.UInt32)
	copied.Scores = rd.Scores.Copy().(types.List[types.UInt32])
	copied.Unknown1 = rd.Unknown1.Copy().(types.UInt8)
	copied.Unknown2 = rd.Unknown2.Copy().(types.UInt32)
	copied.CommonData = rd.CommonData.Copy().(types.Buffer)

	return copied
}

// Equals checks if the given RankingData contains the same data as the current RankingData
func (rd RankingData) Equals(o types.RVType) bool {
	if _, ok := o.(RankingData); !ok {
		return false
	}

	other := o.(RankingData)

	if rd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rd.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !rd.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !rd.Order.Equals(other.Order) {
		return false
	}

	if !rd.Category.Equals(other.Category) {
		return false
	}

	if !rd.Scores.Equals(other.Scores) {
		return false
	}

	if !rd.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !rd.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return rd.CommonData.Equals(other.CommonData)
}

// CopyRef copies the current value of the RankingData
// and returns a pointer to the new copy
func (rd RankingData) CopyRef() types.RVTypePtr {
	copied := rd.Copy().(RankingData)
	return &copied
}

// Deref takes a pointer to the RankingData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rd *RankingData) Deref() types.RVType {
	return *rd
}

// String returns the string representation of the RankingData
func (rd RankingData) String() string {
	return rd.FormatToString(0)
}

// FormatToString pretty-prints the RankingData using the provided indentation level
func (rd RankingData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingData{\n")
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, rd.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, rd.UniqueID))
	b.WriteString(fmt.Sprintf("%sOrder: %s,\n", indentationValues, rd.Order))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rd.Category))
	b.WriteString(fmt.Sprintf("%sScores: %s,\n", indentationValues, rd.Scores))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, rd.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, rd.Unknown2))
	b.WriteString(fmt.Sprintf("%sCommonData: %s,\n", indentationValues, rd.CommonData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingData returns a new RankingData
func NewRankingData() RankingData {
	return RankingData{
		PrincipalID: types.NewPID(0),
		UniqueID:    types.NewUInt32(0),
		Order:       types.NewUInt32(0),
		Category:    types.NewUInt32(0),
		Scores:      types.NewList[types.UInt32](),
		Unknown1:    types.NewUInt8(0),
		Unknown2:    types.NewUInt32(0),
		CommonData:  types.NewBuffer(nil),
	}

}
