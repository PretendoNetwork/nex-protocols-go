// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MatchmakeRefereeStats contains the results of a round
type MatchmakeRefereeStats struct {
	types.Structure
	*types.Data
	UniqueID            *types.PrimitiveU64
	Category            *types.PrimitiveU32
	PID                 *types.PID
	RecentDisconnection *types.PrimitiveU32
	RecentViolation     *types.PrimitiveU32
	RecentMismatch      *types.PrimitiveU32
	RecentWin           *types.PrimitiveU32
	RecentLoss          *types.PrimitiveU32
	RecentDraw          *types.PrimitiveU32
	TotalDisconnect     *types.PrimitiveU32
	TotalViolation      *types.PrimitiveU32
	TotalMismatch       *types.PrimitiveU32
	TotalWin            *types.PrimitiveU32
	TotalLoss           *types.PrimitiveU32
	TotalDraw           *types.PrimitiveU32
	RatingValue         *types.PrimitiveU32
}

// WriteTo writes the MatchmakeRefereeStats to the given writable
func (matchmakeRefereeStats *MatchmakeRefereeStats) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakeRefereeStats.UniqueID.WriteTo(contentWritable)
	matchmakeRefereeStats.Category.WriteTo(contentWritable)
	matchmakeRefereeStats.PID.WriteTo(contentWritable)
	matchmakeRefereeStats.RecentDisconnection.WriteTo(contentWritable)
	matchmakeRefereeStats.RecentViolation.WriteTo(contentWritable)
	matchmakeRefereeStats.RecentMismatch.WriteTo(contentWritable)
	matchmakeRefereeStats.RecentWin.WriteTo(contentWritable)
	matchmakeRefereeStats.RecentLoss.WriteTo(contentWritable)
	matchmakeRefereeStats.RecentDraw.WriteTo(contentWritable)
	matchmakeRefereeStats.TotalDisconnect.WriteTo(contentWritable)
	matchmakeRefereeStats.TotalViolation.WriteTo(contentWritable)
	matchmakeRefereeStats.TotalMismatch.WriteTo(contentWritable)
	matchmakeRefereeStats.TotalWin.WriteTo(contentWritable)
	matchmakeRefereeStats.TotalLoss.WriteTo(contentWritable)
	matchmakeRefereeStats.TotalDraw.WriteTo(contentWritable)
	matchmakeRefereeStats.RatingValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	matchmakeRefereeStats.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeStats from the given readable
func (matchmakeRefereeStats *MatchmakeRefereeStats) ExtractFrom(readable types.Readable) error {
	var err error

	if err = matchmakeRefereeStats.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MatchmakeRefereeStats header. %s", err.Error())
	}

	err = matchmakeRefereeStats.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.UniqueID. %s", err.Error())
	}

	err = matchmakeRefereeStats.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.Category. %s", err.Error())
	}

	err = matchmakeRefereeStats.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.PID. %s", err.Error())
	}

	err = matchmakeRefereeStats.RecentDisconnection.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentDisconnection. %s", err.Error())
	}

	err = matchmakeRefereeStats.RecentViolation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentViolation. %s", err.Error())
	}

	err = matchmakeRefereeStats.RecentMismatch.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentMismatch. %s", err.Error())
	}

	err = matchmakeRefereeStats.RecentWin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentWin. %s", err.Error())
	}

	err = matchmakeRefereeStats.RecentLoss.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentLoss. %s", err.Error())
	}

	err = matchmakeRefereeStats.RecentDraw.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentDraw. %s", err.Error())
	}

	err = matchmakeRefereeStats.TotalDisconnect.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalDisconnect. %s", err.Error())
	}

	err = matchmakeRefereeStats.TotalViolation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalViolation. %s", err.Error())
	}

	err = matchmakeRefereeStats.TotalMismatch.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalMismatch. %s", err.Error())
	}

	err = matchmakeRefereeStats.TotalWin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalWin. %s", err.Error())
	}

	err = matchmakeRefereeStats.TotalLoss.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalLoss. %s", err.Error())
	}

	err = matchmakeRefereeStats.TotalDraw.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalDraw. %s", err.Error())
	}

	err = matchmakeRefereeStats.RatingValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RatingValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStats
func (matchmakeRefereeStats *MatchmakeRefereeStats) Copy() types.RVType {
	copied := NewMatchmakeRefereeStats()

	copied.StructureVersion = matchmakeRefereeStats.StructureVersion

	copied.Data = matchmakeRefereeStats.Data.Copy().(*types.Data)

	copied.UniqueID = matchmakeRefereeStats.UniqueID
	copied.Category = matchmakeRefereeStats.Category
	copied.PID = matchmakeRefereeStats.PID.Copy()
	copied.RecentDisconnection = matchmakeRefereeStats.RecentDisconnection
	copied.RecentViolation = matchmakeRefereeStats.RecentViolation
	copied.RecentMismatch = matchmakeRefereeStats.RecentMismatch
	copied.RecentWin = matchmakeRefereeStats.RecentWin
	copied.RecentLoss = matchmakeRefereeStats.RecentLoss
	copied.RecentDraw = matchmakeRefereeStats.RecentDraw
	copied.TotalDisconnect = matchmakeRefereeStats.TotalDisconnect
	copied.TotalViolation = matchmakeRefereeStats.TotalViolation
	copied.TotalMismatch = matchmakeRefereeStats.TotalMismatch
	copied.TotalWin = matchmakeRefereeStats.TotalWin
	copied.TotalLoss = matchmakeRefereeStats.TotalLoss
	copied.TotalDraw = matchmakeRefereeStats.TotalDraw
	copied.RatingValue = matchmakeRefereeStats.RatingValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeStats *MatchmakeRefereeStats) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeStats); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeStats)

	if matchmakeRefereeStats.StructureVersion != other.StructureVersion {
		return false
	}

	if !matchmakeRefereeStats.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !matchmakeRefereeStats.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !matchmakeRefereeStats.Category.Equals(other.Category) {
		return false
	}

	if !matchmakeRefereeStats.PID.Equals(other.PID) {
		return false
	}

	if !matchmakeRefereeStats.RecentDisconnection.Equals(other.RecentDisconnection) {
		return false
	}

	if !matchmakeRefereeStats.RecentViolation.Equals(other.RecentViolation) {
		return false
	}

	if !matchmakeRefereeStats.RecentMismatch.Equals(other.RecentMismatch) {
		return false
	}

	if !matchmakeRefereeStats.RecentWin.Equals(other.RecentWin) {
		return false
	}

	if !matchmakeRefereeStats.RecentLoss.Equals(other.RecentLoss) {
		return false
	}

	if !matchmakeRefereeStats.RecentDraw.Equals(other.RecentDraw) {
		return false
	}

	if !matchmakeRefereeStats.TotalDisconnect.Equals(other.TotalDisconnect) {
		return false
	}

	if !matchmakeRefereeStats.TotalViolation.Equals(other.TotalViolation) {
		return false
	}

	if !matchmakeRefereeStats.TotalMismatch.Equals(other.TotalMismatch) {
		return false
	}

	if !matchmakeRefereeStats.TotalWin.Equals(other.TotalWin) {
		return false
	}

	if !matchmakeRefereeStats.TotalLoss.Equals(other.TotalLoss) {
		return false
	}

	if !matchmakeRefereeStats.TotalDraw.Equals(other.TotalDraw) {
		return false
	}

	if !matchmakeRefereeStats.RatingValue.Equals(other.RatingValue) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeRefereeStats *MatchmakeRefereeStats) String() string {
	return matchmakeRefereeStats.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeRefereeStats *MatchmakeRefereeStats) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeStats{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeRefereeStats.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, matchmakeRefereeStats.UniqueID))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, matchmakeRefereeStats.Category))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, matchmakeRefereeStats.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRecentDisconnection: %d,\n", indentationValues, matchmakeRefereeStats.RecentDisconnection))
	b.WriteString(fmt.Sprintf("%sRecentViolation: %d,\n", indentationValues, matchmakeRefereeStats.RecentViolation))
	b.WriteString(fmt.Sprintf("%sRecentMismatch: %d,\n", indentationValues, matchmakeRefereeStats.RecentMismatch))
	b.WriteString(fmt.Sprintf("%sRecentWin: %d,\n", indentationValues, matchmakeRefereeStats.RecentWin))
	b.WriteString(fmt.Sprintf("%sRecentLoss: %d,\n", indentationValues, matchmakeRefereeStats.RecentLoss))
	b.WriteString(fmt.Sprintf("%sRecentDraw: %d,\n", indentationValues, matchmakeRefereeStats.RecentDraw))
	b.WriteString(fmt.Sprintf("%sTotalDisconnect: %d,\n", indentationValues, matchmakeRefereeStats.TotalDisconnect))
	b.WriteString(fmt.Sprintf("%sTotalViolation: %d,\n", indentationValues, matchmakeRefereeStats.TotalViolation))
	b.WriteString(fmt.Sprintf("%sTotalMismatch: %d,\n", indentationValues, matchmakeRefereeStats.TotalMismatch))
	b.WriteString(fmt.Sprintf("%sTotalWin: %d,\n", indentationValues, matchmakeRefereeStats.TotalWin))
	b.WriteString(fmt.Sprintf("%sTotalLoss: %d,\n", indentationValues, matchmakeRefereeStats.TotalLoss))
	b.WriteString(fmt.Sprintf("%sTotalDraw: %d,\n", indentationValues, matchmakeRefereeStats.TotalDraw))
	b.WriteString(fmt.Sprintf("%sRatingValue: %d,\n", indentationValues, matchmakeRefereeStats.RatingValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStats returns a new MatchmakeRefereeStats
func NewMatchmakeRefereeStats() *MatchmakeRefereeStats {
	return &MatchmakeRefereeStats{}
}
