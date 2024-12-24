// Package types implements all the types used by the MatchmakeReferee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeRefereeStats is a type within the MatchmakeReferee protocol
type MatchmakeRefereeStats struct {
	types.Structure
	types.Data
	UniqueID            types.UInt64
	Category            types.UInt32
	PID                 types.PID
	RecentDisconnection types.UInt32
	RecentViolation     types.UInt32
	RecentMismatch      types.UInt32
	RecentWin           types.UInt32
	RecentLoss          types.UInt32
	RecentDraw          types.UInt32
	TotalDisconnect     types.UInt32
	TotalViolation      types.UInt32
	TotalMismatch       types.UInt32
	TotalWin            types.UInt32
	TotalLoss           types.UInt32
	TotalDraw           types.UInt32
	RatingValue         types.UInt32
}

// WriteTo writes the MatchmakeRefereeStats to the given writable
func (mrs MatchmakeRefereeStats) WriteTo(writable types.Writable) {
	mrs.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mrs.UniqueID.WriteTo(contentWritable)
	mrs.Category.WriteTo(contentWritable)
	mrs.PID.WriteTo(contentWritable)
	mrs.RecentDisconnection.WriteTo(contentWritable)
	mrs.RecentViolation.WriteTo(contentWritable)
	mrs.RecentMismatch.WriteTo(contentWritable)
	mrs.RecentWin.WriteTo(contentWritable)
	mrs.RecentLoss.WriteTo(contentWritable)
	mrs.RecentDraw.WriteTo(contentWritable)
	mrs.TotalDisconnect.WriteTo(contentWritable)
	mrs.TotalViolation.WriteTo(contentWritable)
	mrs.TotalMismatch.WriteTo(contentWritable)
	mrs.TotalWin.WriteTo(contentWritable)
	mrs.TotalLoss.WriteTo(contentWritable)
	mrs.TotalDraw.WriteTo(contentWritable)
	mrs.RatingValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mrs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeStats from the given readable
func (mrs *MatchmakeRefereeStats) ExtractFrom(readable types.Readable) error {
	var err error

	err = mrs.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.Data. %s", err.Error())
	}

	err = mrs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats header. %s", err.Error())
	}

	err = mrs.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.UniqueID. %s", err.Error())
	}

	err = mrs.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.Category. %s", err.Error())
	}

	err = mrs.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.PID. %s", err.Error())
	}

	err = mrs.RecentDisconnection.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentDisconnection. %s", err.Error())
	}

	err = mrs.RecentViolation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentViolation. %s", err.Error())
	}

	err = mrs.RecentMismatch.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentMismatch. %s", err.Error())
	}

	err = mrs.RecentWin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentWin. %s", err.Error())
	}

	err = mrs.RecentLoss.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentLoss. %s", err.Error())
	}

	err = mrs.RecentDraw.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentDraw. %s", err.Error())
	}

	err = mrs.TotalDisconnect.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalDisconnect. %s", err.Error())
	}

	err = mrs.TotalViolation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalViolation. %s", err.Error())
	}

	err = mrs.TotalMismatch.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalMismatch. %s", err.Error())
	}

	err = mrs.TotalWin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalWin. %s", err.Error())
	}

	err = mrs.TotalLoss.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalLoss. %s", err.Error())
	}

	err = mrs.TotalDraw.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalDraw. %s", err.Error())
	}

	err = mrs.RatingValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RatingValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStats
func (mrs MatchmakeRefereeStats) Copy() types.RVType {
	copied := NewMatchmakeRefereeStats()

	copied.StructureVersion = mrs.StructureVersion
	copied.Data = mrs.Data.Copy().(types.Data)
	copied.UniqueID = mrs.UniqueID.Copy().(types.UInt64)
	copied.Category = mrs.Category.Copy().(types.UInt32)
	copied.PID = mrs.PID.Copy().(types.PID)
	copied.RecentDisconnection = mrs.RecentDisconnection.Copy().(types.UInt32)
	copied.RecentViolation = mrs.RecentViolation.Copy().(types.UInt32)
	copied.RecentMismatch = mrs.RecentMismatch.Copy().(types.UInt32)
	copied.RecentWin = mrs.RecentWin.Copy().(types.UInt32)
	copied.RecentLoss = mrs.RecentLoss.Copy().(types.UInt32)
	copied.RecentDraw = mrs.RecentDraw.Copy().(types.UInt32)
	copied.TotalDisconnect = mrs.TotalDisconnect.Copy().(types.UInt32)
	copied.TotalViolation = mrs.TotalViolation.Copy().(types.UInt32)
	copied.TotalMismatch = mrs.TotalMismatch.Copy().(types.UInt32)
	copied.TotalWin = mrs.TotalWin.Copy().(types.UInt32)
	copied.TotalLoss = mrs.TotalLoss.Copy().(types.UInt32)
	copied.TotalDraw = mrs.TotalDraw.Copy().(types.UInt32)
	copied.RatingValue = mrs.RatingValue.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given MatchmakeRefereeStats contains the same data as the current MatchmakeRefereeStats
func (mrs MatchmakeRefereeStats) Equals(o types.RVType) bool {
	if _, ok := o.(MatchmakeRefereeStats); !ok {
		return false
	}

	other := o.(MatchmakeRefereeStats)

	if mrs.StructureVersion != other.StructureVersion {
		return false
	}

	if !mrs.Data.Equals(other.Data) {
		return false
	}

	if !mrs.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !mrs.Category.Equals(other.Category) {
		return false
	}

	if !mrs.PID.Equals(other.PID) {
		return false
	}

	if !mrs.RecentDisconnection.Equals(other.RecentDisconnection) {
		return false
	}

	if !mrs.RecentViolation.Equals(other.RecentViolation) {
		return false
	}

	if !mrs.RecentMismatch.Equals(other.RecentMismatch) {
		return false
	}

	if !mrs.RecentWin.Equals(other.RecentWin) {
		return false
	}

	if !mrs.RecentLoss.Equals(other.RecentLoss) {
		return false
	}

	if !mrs.RecentDraw.Equals(other.RecentDraw) {
		return false
	}

	if !mrs.TotalDisconnect.Equals(other.TotalDisconnect) {
		return false
	}

	if !mrs.TotalViolation.Equals(other.TotalViolation) {
		return false
	}

	if !mrs.TotalMismatch.Equals(other.TotalMismatch) {
		return false
	}

	if !mrs.TotalWin.Equals(other.TotalWin) {
		return false
	}

	if !mrs.TotalLoss.Equals(other.TotalLoss) {
		return false
	}

	if !mrs.TotalDraw.Equals(other.TotalDraw) {
		return false
	}

	return mrs.RatingValue.Equals(other.RatingValue)
}

// CopyRef copies the current value of the MatchmakeRefereeStats
// and returns a pointer to the new copy
func (mrs MatchmakeRefereeStats) CopyRef() types.RVTypePtr {
	copied := mrs.Copy().(MatchmakeRefereeStats)
	return &copied
}

// Deref takes a pointer to the MatchmakeRefereeStats
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (mrs *MatchmakeRefereeStats) Deref() types.RVType {
	return *mrs
}

// String returns the string representation of the MatchmakeRefereeStats
func (mrs MatchmakeRefereeStats) String() string {
	return mrs.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeRefereeStats using the provided indentation level
func (mrs MatchmakeRefereeStats) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeStats{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mrs.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, mrs.UniqueID))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, mrs.Category))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, mrs.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRecentDisconnection: %s,\n", indentationValues, mrs.RecentDisconnection))
	b.WriteString(fmt.Sprintf("%sRecentViolation: %s,\n", indentationValues, mrs.RecentViolation))
	b.WriteString(fmt.Sprintf("%sRecentMismatch: %s,\n", indentationValues, mrs.RecentMismatch))
	b.WriteString(fmt.Sprintf("%sRecentWin: %s,\n", indentationValues, mrs.RecentWin))
	b.WriteString(fmt.Sprintf("%sRecentLoss: %s,\n", indentationValues, mrs.RecentLoss))
	b.WriteString(fmt.Sprintf("%sRecentDraw: %s,\n", indentationValues, mrs.RecentDraw))
	b.WriteString(fmt.Sprintf("%sTotalDisconnect: %s,\n", indentationValues, mrs.TotalDisconnect))
	b.WriteString(fmt.Sprintf("%sTotalViolation: %s,\n", indentationValues, mrs.TotalViolation))
	b.WriteString(fmt.Sprintf("%sTotalMismatch: %s,\n", indentationValues, mrs.TotalMismatch))
	b.WriteString(fmt.Sprintf("%sTotalWin: %s,\n", indentationValues, mrs.TotalWin))
	b.WriteString(fmt.Sprintf("%sTotalLoss: %s,\n", indentationValues, mrs.TotalLoss))
	b.WriteString(fmt.Sprintf("%sTotalDraw: %s,\n", indentationValues, mrs.TotalDraw))
	b.WriteString(fmt.Sprintf("%sRatingValue: %s,\n", indentationValues, mrs.RatingValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStats returns a new MatchmakeRefereeStats
func NewMatchmakeRefereeStats() MatchmakeRefereeStats {
	return MatchmakeRefereeStats{
		Data:                types.NewData(),
		UniqueID:            types.NewUInt64(0),
		Category:            types.NewUInt32(0),
		PID:                 types.NewPID(0),
		RecentDisconnection: types.NewUInt32(0),
		RecentViolation:     types.NewUInt32(0),
		RecentMismatch:      types.NewUInt32(0),
		RecentWin:           types.NewUInt32(0),
		RecentLoss:          types.NewUInt32(0),
		RecentDraw:          types.NewUInt32(0),
		TotalDisconnect:     types.NewUInt32(0),
		TotalViolation:      types.NewUInt32(0),
		TotalMismatch:       types.NewUInt32(0),
		TotalWin:            types.NewUInt32(0),
		TotalLoss:           types.NewUInt32(0),
		TotalDraw:           types.NewUInt32(0),
		RatingValue:         types.NewUInt32(0),
	}

}
