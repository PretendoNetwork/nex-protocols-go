package datastore_super_smash_bros_4_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreReplayMetaInfo struct {
	nex.Structure
	ReplayID   uint64
	Size       uint32
	Mode       uint8
	Style      uint8
	Rule       uint8
	Stage      uint8
	ReplayType uint8
	Players    []*DataStoreReplayPlayer
	Winners    []uint32
}

// ExtractFromStream extracts a DataStoreReplayMetaInfo structure from a stream
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReplayMetaInfo.ReplayID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.ReplayID. %s", err.Error())
	}

	dataStoreReplayMetaInfo.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Size. %s", err.Error())
	}

	dataStoreReplayMetaInfo.Mode, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Mode. %s", err.Error())
	}

	dataStoreReplayMetaInfo.Style, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Style. %s", err.Error())
	}

	dataStoreReplayMetaInfo.Rule, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Rule. %s", err.Error())
	}

	dataStoreReplayMetaInfo.Stage, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Stage. %s", err.Error())
	}

	dataStoreReplayMetaInfo.ReplayType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.ReplayType. %s", err.Error())
	}

	players, err := stream.ReadListStructure(NewDataStoreReplayPlayer())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Players. %s", err.Error())
	}

	dataStoreReplayMetaInfo.Players = players.([]*DataStoreReplayPlayer)
	dataStoreReplayMetaInfo.Winners, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Winners. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReplayMetaInfo and returns a byte array
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreReplayMetaInfo.ReplayID)
	stream.WriteUInt32LE(dataStoreReplayMetaInfo.Size)
	stream.WriteUInt8(dataStoreReplayMetaInfo.Mode)
	stream.WriteUInt8(dataStoreReplayMetaInfo.Style)
	stream.WriteUInt8(dataStoreReplayMetaInfo.Rule)
	stream.WriteUInt8(dataStoreReplayMetaInfo.Stage)
	stream.WriteUInt8(dataStoreReplayMetaInfo.ReplayType)
	stream.WriteListStructure(dataStoreReplayMetaInfo.Players)
	stream.WriteListUInt32LE(dataStoreReplayMetaInfo.Winners)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReplayMetaInfo
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReplayMetaInfo()

	copied.ReplayID = dataStoreReplayMetaInfo.ReplayID
	copied.Size = dataStoreReplayMetaInfo.Size
	copied.Mode = dataStoreReplayMetaInfo.Mode
	copied.Style = dataStoreReplayMetaInfo.Style
	copied.Rule = dataStoreReplayMetaInfo.Rule
	copied.Stage = dataStoreReplayMetaInfo.Stage
	copied.ReplayType = dataStoreReplayMetaInfo.ReplayType
	copied.Players = make([]*DataStoreReplayPlayer, len(dataStoreReplayMetaInfo.Players))

	for i := 0; i < len(dataStoreReplayMetaInfo.Players); i++ {
		copied.Players[i] = dataStoreReplayMetaInfo.Players[i].Copy().(*DataStoreReplayPlayer)
	}

	copied.Winners = make([]uint32, len(dataStoreReplayMetaInfo.Winners))

	copy(copied.Winners, dataStoreReplayMetaInfo.Winners)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReplayMetaInfo)

	if dataStoreReplayMetaInfo.ReplayID != other.ReplayID {
		return false
	}

	if dataStoreReplayMetaInfo.Size != other.Size {
		return false
	}

	if dataStoreReplayMetaInfo.Mode != other.Mode {
		return false
	}

	if dataStoreReplayMetaInfo.Style != other.Style {
		return false
	}

	if dataStoreReplayMetaInfo.Rule != other.Rule {
		return false
	}

	if dataStoreReplayMetaInfo.Stage != other.Stage {
		return false
	}

	if dataStoreReplayMetaInfo.ReplayType != other.ReplayType {
		return false
	}

	if len(dataStoreReplayMetaInfo.Players) != len(other.Players) {
		return false
	}

	for i := 0; i < len(dataStoreReplayMetaInfo.Players); i++ {
		if !dataStoreReplayMetaInfo.Players[i].Equals(other.Players[i]) {
			return false
		}
	}

	if len(dataStoreReplayMetaInfo.Winners) != len(other.Winners) {
		return false
	}

	for i := 0; i < len(dataStoreReplayMetaInfo.Winners); i++ {
		if dataStoreReplayMetaInfo.Players[i] != other.Players[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) String() string {
	return dataStoreReplayMetaInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReplayMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReplayMetaInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReplayID: %d,\n", indentationValues, dataStoreReplayMetaInfo.ReplayID))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStoreReplayMetaInfo.Size))
	b.WriteString(fmt.Sprintf("%sMode: %d,\n", indentationValues, dataStoreReplayMetaInfo.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %d,\n", indentationValues, dataStoreReplayMetaInfo.Style))
	b.WriteString(fmt.Sprintf("%sRule: %d,\n", indentationValues, dataStoreReplayMetaInfo.Rule))
	b.WriteString(fmt.Sprintf("%sStage: %d,\n", indentationValues, dataStoreReplayMetaInfo.Stage))
	b.WriteString(fmt.Sprintf("%sReplayType: %d,\n", indentationValues, dataStoreReplayMetaInfo.ReplayType))

	if len(dataStoreReplayMetaInfo.Players) == 0 {
		b.WriteString(fmt.Sprintf("%sPlayers: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sPlayers: [\n", indentationValues))

		for i := 0; i < len(dataStoreReplayMetaInfo.Players); i++ {
			str := dataStoreReplayMetaInfo.Players[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReplayMetaInfo.Players)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s]\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sWinners: %v\n", indentationValues, dataStoreReplayMetaInfo.Winners))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReplayMetaInfo returns a new DataStoreReplayMetaInfo
func NewDataStoreReplayMetaInfo() *DataStoreReplayMetaInfo {
	return &DataStoreReplayMetaInfo{}
}
