// Package types implements all the types used by the Ranking protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// RankingRankData holds parameters for ordering rankings
type RankingRankData struct {
	types.Structure
	PrincipalID *types.PID
	UniqueID    *types.PrimitiveU64
	Order       *types.PrimitiveU32
	Category    *types.PrimitiveU32
	Score       *types.PrimitiveU32
	Groups      []byte
	Param       *types.PrimitiveU64
	CommonData  []byte
	UpdateTime  *types.DateTime // * NEX 3.6.0+
}

// ExtractFrom extracts the RankingRankData from the given readable
func (rankingRankData *RankingRankData) ExtractFrom(readable types.Readable) error {
	nexVersion := stream.Server.LibraryVersion()

	var err error

	err = rankingRankData.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.PrincipalID from stream. %s", err.Error())
	}

	err = rankingRankData.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.UniqueID from stream. %s", err.Error())
	}

	err = rankingRankData.Order.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Order from stream. %s", err.Error())
	}

	err = rankingRankData.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Category from stream. %s", err.Error())
	}

	err = rankingRankData.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Score from stream. %s", err.Error())
	}

	rankingRankData.Groups, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Groups from stream. %s", err.Error())
	}

	err = rankingRankData.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Param from stream. %s", err.Error())
	}

	rankingRankData.CommonData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.CommonData from stream. %s", err.Error())
	}

	if nexVersion.GreaterOrEqual("3.6.0") {
	err = 	rankingRankData.UpdateTime.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract RankingRankData.UpdateTime from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the RankingRankData to the given writable
func (rankingRankData *RankingRankData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nexVersion := stream.Server.LibraryVersion()

	rankingRankData.PrincipalID.WriteTo(contentWritable)
	rankingRankData.UniqueID.WriteTo(contentWritable)
	rankingRankData.Order.WriteTo(contentWritable)
	rankingRankData.Category.WriteTo(contentWritable)
	rankingRankData.Score.WriteTo(contentWritable)
	stream.WriteBuffer(rankingRankData.Groups)
	rankingRankData.Param.WriteTo(contentWritable)
	stream.WriteBuffer(rankingRankData.CommonData)

	if nexVersion.GreaterOrEqual("4.0.0") {
		rankingRankData.UpdateTime.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of RankingRankData
func (rankingRankData *RankingRankData) Copy() types.RVType {
	copied := NewRankingRankData()

	copied.StructureVersion = rankingRankData.StructureVersion

	copied.PrincipalID = rankingRankData.PrincipalID.Copy()
	copied.UniqueID = rankingRankData.UniqueID
	copied.Order = rankingRankData.Order
	copied.Category = rankingRankData.Category
	copied.Score = rankingRankData.Score
	copied.Groups = make([]byte, len(rankingRankData.Groups))

	copy(copied.Groups, rankingRankData.Groups)

	copied.Param = rankingRankData.Param
	copied.CommonData = make([]byte, len(rankingRankData.CommonData))

	copy(copied.CommonData, rankingRankData.CommonData)

	copied.UpdateTime = rankingRankData.UpdateTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingRankData *RankingRankData) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingRankData); !ok {
		return false
	}

	other := o.(*RankingRankData)

	if rankingRankData.StructureVersion != other.StructureVersion {
		return false
	}

	if !rankingRankData.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !rankingRankData.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !rankingRankData.Order.Equals(other.Order) {
		return false
	}

	if !rankingRankData.Category.Equals(other.Category) {
		return false
	}

	if !rankingRankData.Score.Equals(other.Score) {
		return false
	}

	if !rankingRankData.Groups.Equals(other.Groups) {
		return false
	}

	if !rankingRankData.Param.Equals(other.Param) {
		return false
	}

	if !rankingRankData.CommonData.Equals(other.CommonData) {
		return false
	}

	if !rankingRankData.UpdateTime.Equals(other.UpdateTime) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (rankingRankData *RankingRankData) String() string {
	return rankingRankData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingRankData *RankingRankData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingRankData{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, rankingRankData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, rankingRankData.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, rankingRankData.UniqueID))
	b.WriteString(fmt.Sprintf("%sOrder: %d,\n", indentationValues, rankingRankData.Order))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, rankingRankData.Category))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, rankingRankData.Score))
	b.WriteString(fmt.Sprintf("%sGroups: %x,\n", indentationValues, rankingRankData.Groups))
	b.WriteString(fmt.Sprintf("%sParam: %d,\n", indentationValues, rankingRankData.Param))
	b.WriteString(fmt.Sprintf("%sCommonData: %x,\n", indentationValues, rankingRankData.CommonData))

	if rankingRankData.UpdateTime != nil {
		b.WriteString(fmt.Sprintf("%sUpdateTime: %s\n", indentationValues, rankingRankData.UpdateTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUpdateTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingRankData returns a new RankingRankData
func NewRankingRankData() *RankingRankData {
	return &RankingRankData{}
}
