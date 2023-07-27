// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MatchmakeRefereePersonalRoundResult contains the results of a round
type MatchmakeRefereePersonalRoundResult struct {
	nex.Structure
	*nex.Data
	PID                     uint32
	PersonalRoundResultFlag uint32
	RoundWinLoss            uint32
	RatingValueChange       int32
	Buffer                  []byte
}

// Bytes encodes the MatchmakeRefereePersonalRoundResult and returns a byte array
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(matchmakeRefereePersonalRoundResult.PID)
	stream.WriteUInt32LE(matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag)
	stream.WriteUInt32LE(matchmakeRefereePersonalRoundResult.RoundWinLoss)
	stream.WriteInt32LE(matchmakeRefereePersonalRoundResult.RatingValueChange)
	stream.WriteQBuffer(matchmakeRefereePersonalRoundResult.Buffer)

	return stream.Bytes()
}

// ExtractFromStream extracts a MatchmakeRefereePersonalRoundResult structure from a stream
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeRefereePersonalRoundResult.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.PID. %s", err.Error())
	}

	matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.PersonalRoundResultFlag. %s", err.Error())
	}

	matchmakeRefereePersonalRoundResult.RoundWinLoss, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereePersonalRoundResult.RoundWinLoss. %s", err.Error())
	}

	matchmakeRefereePersonalRoundResult.RatingValueChange, err = stream.ReadInt32LE()
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
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) Copy() nex.StructureInterface {
	copied := NewMatchmakeRefereePersonalRoundResult()

	copied.Data = matchmakeRefereePersonalRoundResult.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.PID = matchmakeRefereePersonalRoundResult.PID
	copied.PersonalRoundResultFlag = matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag
	copied.RoundWinLoss = matchmakeRefereePersonalRoundResult.RoundWinLoss
	copied.RatingValueChange = matchmakeRefereePersonalRoundResult.RatingValueChange
	copied.Buffer = make([]byte, len(matchmakeRefereePersonalRoundResult.Buffer))
	copy(copied.Buffer, matchmakeRefereePersonalRoundResult.Buffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereePersonalRoundResult *MatchmakeRefereePersonalRoundResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeRefereePersonalRoundResult)

	if !matchmakeRefereePersonalRoundResult.ParentType().Equals(other.ParentType()) {
		return false
	}

	if matchmakeRefereePersonalRoundResult.PID != other.PID {
		return false
	}

	if matchmakeRefereePersonalRoundResult.PersonalRoundResultFlag != other.PersonalRoundResultFlag {
		return false
	}

	if matchmakeRefereePersonalRoundResult.RoundWinLoss != other.RoundWinLoss {
		return false
	}

	if matchmakeRefereePersonalRoundResult.RatingValueChange != other.RatingValueChange {
		return false
	}

	if !bytes.Equal(matchmakeRefereePersonalRoundResult.Buffer, other.Buffer) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeRefereePersonalRoundResult.StructureVersion()))
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
