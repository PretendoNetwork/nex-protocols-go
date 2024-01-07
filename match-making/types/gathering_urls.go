// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// GatheringURLs holds information about a matchmake gatheringURLs
type GatheringURLs struct {
	types.Structure
	GID            *types.PrimitiveU32
	LstStationURLs *types.List[*types.StationURL]
}

// ExtractFrom extracts the GatheringURLs from the given readable
func (gatheringURLs *GatheringURLs) ExtractFrom(readable types.Readable) error {
	var err error

	if err = gatheringURLs.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GatheringURLs header. %s", err.Error())
	}

	err = gatheringURLs.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringURLs.GID. %s", err.Error())
	}

	gatheringURLs.LstStationURLs, err = stream.ReadListStationURL()
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringURLs.LstStationURLs. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GatheringURLs to the given writable
func (gatheringURLs *GatheringURLs) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gatheringURLs.GID.WriteTo(contentWritable)
	stream.WriteListStationURL(gatheringURLs.LstStationURLs)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GatheringURLs
func (gatheringURLs *GatheringURLs) Copy() types.RVType {
	copied := NewGatheringURLs()

	copied.StructureVersion = gatheringURLs.StructureVersion

	copied.GID = gatheringURLs.GID
	copied.LstStationURLs = make(*types.List[*types.StationURL], len(gatheringURLs.LstStationURLs))

	for i := 0; i < len(gatheringURLs.LstStationURLs); i++ {
		copied.LstStationURLs[i] = gatheringURLs.LstStationURLs[i].Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gatheringURLs *GatheringURLs) Equals(o types.RVType) bool {
	if _, ok := o.(*GatheringURLs); !ok {
		return false
	}

	other := o.(*GatheringURLs)

	if gatheringURLs.StructureVersion != other.StructureVersion {
		return false
	}

	if !gatheringURLs.GID.Equals(other.GID) {
		return false
	}

	if len(gatheringURLs.LstStationURLs) != len(other.LstStationURLs) {
		return false
	}

	for i := 0; i < len(gatheringURLs.LstStationURLs); i++ {
		if !gatheringURLs.LstStationURLs[i].Equals(other.LstStationURLs[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (gatheringURLs *GatheringURLs) String() string {
	return gatheringURLs.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (gatheringURLs *GatheringURLs) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GatheringURLs{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, gatheringURLs.StructureVersion))
	b.WriteString(fmt.Sprintf("%sGID: %d,\n", indentationValues, gatheringURLs.GID))

	if len(gatheringURLs.LstStationURLs) == 0 {
		b.WriteString(fmt.Sprintf("%sLstStationURLs: []\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sLstStationURLs: [\n", indentationValues))

		for i := 0; i < len(gatheringURLs.LstStationURLs); i++ {
			str := gatheringURLs.LstStationURLs[i].FormatToString(indentationLevel + 2)
			if i == len(gatheringURLs.LstStationURLs)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s]\n", indentationValues))
	}
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGatheringURLs returns a new GatheringURLs
func NewGatheringURLs() *GatheringURLs {
	return &GatheringURLs{}
}
