// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MatchmakeRefereePersonalRoundResult contains the results of a round
type MatchmakeRefereePersonalRoundResult struct {
	types.Structure
	*types.Data
	PID                     *types.PID
	PersonalRoundResultFlag *types.PrimitiveU32
	RoundWinLoss            *types.PrimitiveU32
	RatingValueChange       *types.PrimitiveS32
	Buffer                  []byte
}

// WriteTo writes the MatchmakeRefereePersonalRoundResult to the given writable
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakeRefereePersonalRoundResult.PID.WriteTo(contentWritable)
	matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag.WriteTo(contentWritable)
	matchmakeRefereePersonalRoundResult.RoundWinLoss.WriteTo(contentWritable)
	matchmakeRefereePersonalRoundResult.RatingValueChange.WriteTo(contentWritable)
	stream.WriteQBuffer(matchmakeRefereePersonalRoundResult.Buffer)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereePersonalRoundResult from the given readable
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = matchmakeRefereePersonalRoundResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MatchmakeRefereePersonalRoundResult header. %s", err.Error())
	}

	err = matchmakeRefereePersonalRoundResult.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.PID. %s", err.Error())
	}

	err = matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.PersonalRoundResultFlag. %s", err.Error())
	}

	err = matchmakeRefereePersonalRoundResult.RoundWinLoss.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.RoundWinLoss. %s", err.Error())
	}

	err = matchmakeRefereePersonalRoundResult.RatingValueChange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.RatingValueChange. %s", err.Error())
	}

	matchmakeRefereePersonalRoundResult.Buffer, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.Buffer. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereePersonalRoundResult
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) Copy() types.RVType {
	copied := NewMatchmakeRefereePersonalRoundResult()

	copied.StructureVersion = matchmakeRefereePersonalRoundResult.StructureVersion

	copied.Data = matchmakeRefereePersonalRoundResult.Data.Copy().(*types.Data)

	copied.PID = matchmakeRefereePersonalRoundResult.PID.Copy()
	copied.PersonalRoundResultFlag = matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag
	copied.RoundWinLoss = matchmakeRefereePersonalRoundResult.RoundWinLoss
	copied.RatingValueChange = matchmakeRefereePersonalRoundResult.RatingValueChange
	copied.Buffer = make([]byte, len(matchmakeRefereePersonalRoundResult.Buffer))
	copy(copied.Buffer, matchmakeRefereePersonalRoundResult.Buffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereePersonalRoundResult); !ok {
		return false
	}

	other := o.(*MatchmakeRefereePersonalRoundResult)

	if matchmakeRefereePersonalRoundResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !matchmakeRefereePersonalRoundResult.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !matchmakeRefereePersonalRoundResult.PID.Equals(other.PID) {
		return false
	}

	if !matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag.Equals(other.PersonalRoundResultFlag) {
		return false
	}

	if !matchmakeRefereePersonalRoundResult.RoundWinLoss.Equals(other.RoundWinLoss) {
		return false
	}

	if !matchmakeRefereePersonalRoundResult.RatingValueChange.Equals(other.RatingValueChange) {
		return false
	}

	if !matchmakeRefereePersonalRoundResult.Buffer.Equals(other.Buffer) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) String() string {
	return matchmakeRefereePersonalRoundResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereePersonalRoundResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeRefereePersonalRoundResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, matchmakeRefereePersonalRoundResult.PID))
	b.WriteString(fmt.Sprintf("%sPersonalRoundResultFlag: %d,\n", indentationValues, matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag))
	b.WriteString(fmt.Sprintf("%sRoundWinLoss: %d,\n", indentationValues, matchmakeRefereePersonalRoundResult.RoundWinLoss))
	b.WriteString(fmt.Sprintf("%sRatingValueChange: %d,\n", indentationValues, matchmakeRefereePersonalRoundResult.RatingValueChange))
	b.WriteString(fmt.Sprintf("%sBuffer: %x\n", indentationValues, matchmakeRefereePersonalRoundResult.Buffer))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereePersonalRoundResult returns a new MatchmakeRefereePersonalRoundResult
func NewMatchmakeRefereePersonalRoundResult() *MatchmakeRefereePersonalRoundResult {
	return &MatchmakeRefereePersonalRoundResult{}
}
