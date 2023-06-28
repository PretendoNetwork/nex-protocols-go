package match_making_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// MatchmakeBlockListParam holds parameters for a matchmake session
type MatchmakeBlockListParam struct {
	nex.Structure
	OptionFlag uint32
}

// ExtractFromStream extracts a MatchmakeBlockListParam structure from a stream
func (matchmakeBlockListParam *MatchmakeBlockListParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeBlockListParam.OptionFlag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeBlockListParam.OptionFlag. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeBlockListParam
func (matchmakeBlockListParam *MatchmakeBlockListParam) Copy() nex.StructureInterface {
	copied := NewMatchmakeBlockListParam()

	copied.OptionFlag = matchmakeBlockListParam.OptionFlag

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeBlockListParam *MatchmakeBlockListParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeBlockListParam)

	return matchmakeBlockListParam.OptionFlag == other.OptionFlag
}

// NewMatchmakeBlockListParam returns a new MatchmakeBlockListParam
func NewMatchmakeBlockListParam() *MatchmakeBlockListParam {
	return &MatchmakeBlockListParam{}
}
