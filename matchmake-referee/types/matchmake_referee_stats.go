// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MatchmakeRefereeStats contains the results of a round
type MatchmakeRefereeStats struct {
	nex.Structure
	*nex.Data
	UniqueID            uint64
	Category            uint32
	PID                 uint32
	RecentDisconnection uint32
	RecentViolation     uint32
	RecentMismatch      uint32
	RecentWin           uint32
	RecentLoss          uint32
	RecentDraw          uint32
	TotalDisconnect     uint32
	TotalViolation      uint32
	TotalMismatch       uint32
	TotalWin            uint32
	TotalLoss           uint32
	TotalDraw           uint32
	RatingValue         uint32
}

// Bytes encodes the MatchmakeRefereeStats and returns a byte array
func (matchmakeRefereeStats *MatchmakeRefereeStats) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(matchmakeRefereeStats.UniqueID)
	stream.WriteUInt32LE(matchmakeRefereeStats.Category)
	stream.WriteUInt32LE(matchmakeRefereeStats.PID)
	stream.WriteUInt32LE(matchmakeRefereeStats.RecentDisconnection)
	stream.WriteUInt32LE(matchmakeRefereeStats.RecentViolation)
	stream.WriteUInt32LE(matchmakeRefereeStats.RecentMismatch)
	stream.WriteUInt32LE(matchmakeRefereeStats.RecentWin)
	stream.WriteUInt32LE(matchmakeRefereeStats.RecentLoss)
	stream.WriteUInt32LE(matchmakeRefereeStats.RecentDraw)
	stream.WriteUInt32LE(matchmakeRefereeStats.TotalDisconnect)
	stream.WriteUInt32LE(matchmakeRefereeStats.TotalViolation)
	stream.WriteUInt32LE(matchmakeRefereeStats.TotalMismatch)
	stream.WriteUInt32LE(matchmakeRefereeStats.TotalWin)
	stream.WriteUInt32LE(matchmakeRefereeStats.TotalLoss)
	stream.WriteUInt32LE(matchmakeRefereeStats.TotalDraw)
	stream.WriteUInt32LE(matchmakeRefereeStats.RatingValue)

	return stream.Bytes()
}

// ExtractFromStream extracts a MatchmakeRefereeStats structure from a stream
func (matchmakeRefereeStats *MatchmakeRefereeStats) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeRefereeStats.UniqueID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.UniqueID. %s", err.Error())
	}

	matchmakeRefereeStats.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.Category. %s", err.Error())
	}

	matchmakeRefereeStats.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.PID. %s", err.Error())
	}

	matchmakeRefereeStats.RecentDisconnection, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentDisconnection. %s", err.Error())
	}

	matchmakeRefereeStats.RecentViolation, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentViolation. %s", err.Error())
	}

	matchmakeRefereeStats.RecentMismatch, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentMismatch. %s", err.Error())
	}

	matchmakeRefereeStats.RecentWin, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentWin. %s", err.Error())
	}

	matchmakeRefereeStats.RecentLoss, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentLoss. %s", err.Error())
	}

	matchmakeRefereeStats.RecentDraw, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RecentDraw. %s", err.Error())
	}

	matchmakeRefereeStats.TotalDisconnect, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalDisconnect. %s", err.Error())
	}

	matchmakeRefereeStats.TotalViolation, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalViolation. %s", err.Error())
	}

	matchmakeRefereeStats.TotalMismatch, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalMismatch. %s", err.Error())
	}

	matchmakeRefereeStats.TotalWin, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalWin. %s", err.Error())
	}

	matchmakeRefereeStats.TotalLoss, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalLoss. %s", err.Error())
	}

	matchmakeRefereeStats.TotalDraw, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.TotalDraw. %s", err.Error())
	}

	matchmakeRefereeStats.RatingValue, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStats.RatingValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStats
func (matchmakeRefereeStats *MatchmakeRefereeStats) Copy() nex.StructureInterface {
	copied := NewMatchmakeRefereeStats()

	copied.SetStructureVersion(matchmakeRefereeStats.StructureVersion())

	copied.Data = matchmakeRefereeStats.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.UniqueID = matchmakeRefereeStats.UniqueID
	copied.Category = matchmakeRefereeStats.Category
	copied.PID = matchmakeRefereeStats.PID
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
func (matchmakeRefereeStats *MatchmakeRefereeStats) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeRefereeStats)

	if matchmakeRefereeStats.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !matchmakeRefereeStats.ParentType().Equals(other.ParentType()) {
		return false
	}

	if matchmakeRefereeStats.UniqueID != other.UniqueID {
		return false
	}

	if matchmakeRefereeStats.Category != other.Category {
		return false
	}

	if matchmakeRefereeStats.PID != other.PID {
		return false
	}

	if matchmakeRefereeStats.RecentDisconnection != other.RecentDisconnection {
		return false
	}

	if matchmakeRefereeStats.RecentViolation != other.RecentViolation {
		return false
	}

	if matchmakeRefereeStats.RecentMismatch != other.RecentMismatch {
		return false
	}

	if matchmakeRefereeStats.RecentWin != other.RecentWin {
		return false
	}

	if matchmakeRefereeStats.RecentLoss != other.RecentLoss {
		return false
	}

	if matchmakeRefereeStats.RecentDraw != other.RecentDraw {
		return false
	}

	if matchmakeRefereeStats.TotalDisconnect != other.TotalDisconnect {
		return false
	}

	if matchmakeRefereeStats.TotalViolation != other.TotalViolation {
		return false
	}

	if matchmakeRefereeStats.TotalMismatch != other.TotalMismatch {
		return false
	}

	if matchmakeRefereeStats.TotalWin != other.TotalWin {
		return false
	}

	if matchmakeRefereeStats.TotalLoss != other.TotalLoss {
		return false
	}

	if matchmakeRefereeStats.TotalDraw != other.TotalDraw {
		return false
	}

	if matchmakeRefereeStats.RatingValue != other.RatingValue {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeRefereeStats.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, matchmakeRefereeStats.UniqueID))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, matchmakeRefereeStats.Category))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, matchmakeRefereeStats.PID))
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
