package datastore_super_smash_bros_4_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

type DataStorePreparePostReplayParam struct {
	nex.Structure
	Size          uint32
	Mode          uint8
	Style         uint8
	Rule          uint8
	Stage         uint8
	ReplayType    uint8
	CompetitionID uint64
	Score         int32
	Players       []*DataStoreReplayPlayer
	Winners       []uint32
	KeyVersion    uint16
	ExtraData     []string
}

// ExtractFromStream extracts a DataStorePreparePostReplayParam structure from a stream
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePreparePostReplayParam.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Size. %s", err.Error())
	}

	dataStorePreparePostReplayParam.Mode, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Mode. %s", err.Error())
	}

	dataStorePreparePostReplayParam.Style, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Style. %s", err.Error())
	}

	dataStorePreparePostReplayParam.Rule, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Rule. %s", err.Error())
	}

	dataStorePreparePostReplayParam.Stage, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Stage. %s", err.Error())
	}

	dataStorePreparePostReplayParam.ReplayType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.ReplayType. %s", err.Error())
	}

	dataStorePreparePostReplayParam.CompetitionID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.CompetitionID. %s", err.Error())
	}

	dataStorePreparePostReplayParam.Score, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Score. %s", err.Error())
	}

	players, err := stream.ReadListStructure(NewDataStoreReplayPlayer())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Players. %s", err.Error())
	}

	dataStorePreparePostReplayParam.Players = players.([]*DataStoreReplayPlayer)

	dataStorePreparePostReplayParam.Winners, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Winners. %s", err.Error())
	}

	dataStorePreparePostReplayParam.KeyVersion, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.KeyVersion. %s", err.Error())
	}

	dataStorePreparePostReplayParam.ExtraData, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.ExtraData. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePreparePostReplayParam and returns a byte array
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStorePreparePostReplayParam.Size)
	stream.WriteUInt8(dataStorePreparePostReplayParam.Mode)
	stream.WriteUInt8(dataStorePreparePostReplayParam.Style)
	stream.WriteUInt8(dataStorePreparePostReplayParam.Rule)
	stream.WriteUInt8(dataStorePreparePostReplayParam.Stage)
	stream.WriteUInt8(dataStorePreparePostReplayParam.ReplayType)
	stream.WriteUInt64LE(dataStorePreparePostReplayParam.CompetitionID)
	stream.WriteInt32LE(dataStorePreparePostReplayParam.Score)
	stream.WriteListStructure(dataStorePreparePostReplayParam.Players)
	stream.WriteListUInt32LE(dataStorePreparePostReplayParam.Winners)
	stream.WriteUInt16LE(dataStorePreparePostReplayParam.KeyVersion)
	stream.WriteListString(dataStorePreparePostReplayParam.ExtraData)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePreparePostReplayParam
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStorePreparePostReplayParam()

	copied.Size = dataStorePreparePostReplayParam.Size
	copied.Mode = dataStorePreparePostReplayParam.Mode
	copied.Style = dataStorePreparePostReplayParam.Style
	copied.Rule = dataStorePreparePostReplayParam.Rule
	copied.Stage = dataStorePreparePostReplayParam.Stage
	copied.ReplayType = dataStorePreparePostReplayParam.ReplayType
	copied.CompetitionID = dataStorePreparePostReplayParam.CompetitionID
	copied.Score = dataStorePreparePostReplayParam.Score
	copied.Players = make([]*DataStoreReplayPlayer, len(dataStorePreparePostReplayParam.Players))

	for i := 0; i < len(dataStorePreparePostReplayParam.Players); i++ {
		copied.Players[i] = dataStorePreparePostReplayParam.Players[i].Copy().(*DataStoreReplayPlayer)
	}

	copied.Winners = make([]uint32, len(dataStorePreparePostReplayParam.Winners))

	copy(copied.Winners, dataStorePreparePostReplayParam.Winners)

	copied.KeyVersion = dataStorePreparePostReplayParam.KeyVersion
	copied.ExtraData = make([]string, len(dataStorePreparePostReplayParam.ExtraData))

	copy(copied.ExtraData, dataStorePreparePostReplayParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePreparePostReplayParam)

	if dataStorePreparePostReplayParam.Size != other.Size {
		return false
	}

	if dataStorePreparePostReplayParam.Mode != other.Mode {
		return false
	}

	if dataStorePreparePostReplayParam.Style != other.Style {
		return false
	}

	if dataStorePreparePostReplayParam.Rule != other.Rule {
		return false
	}

	if dataStorePreparePostReplayParam.Stage != other.Stage {
		return false
	}

	if dataStorePreparePostReplayParam.ReplayType != other.ReplayType {
		return false
	}

	if dataStorePreparePostReplayParam.CompetitionID != other.CompetitionID {
		return false
	}

	if dataStorePreparePostReplayParam.Score != other.Score {
		return false
	}

	if len(dataStorePreparePostReplayParam.Players) != len(other.Players) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostReplayParam.Players); i++ {
		if !dataStorePreparePostReplayParam.Players[i].Equals(other.Players[i]) {
			return false
		}
	}

	if len(dataStorePreparePostReplayParam.Winners) != len(other.Winners) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostReplayParam.Winners); i++ {
		if dataStorePreparePostReplayParam.Winners[i] != other.Winners[i] {
			return false
		}
	}

	if dataStorePreparePostReplayParam.KeyVersion != other.KeyVersion {
		return false
	}

	if len(dataStorePreparePostReplayParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostReplayParam.ExtraData); i++ {
		if dataStorePreparePostReplayParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) String() string {
	return dataStorePreparePostReplayParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePreparePostReplayParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStorePreparePostReplayParam.Size))
	b.WriteString(fmt.Sprintf("%sMode: %d,\n", indentationValues, dataStorePreparePostReplayParam.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %d,\n", indentationValues, dataStorePreparePostReplayParam.Style))
	b.WriteString(fmt.Sprintf("%sRule: %d,\n", indentationValues, dataStorePreparePostReplayParam.Rule))
	b.WriteString(fmt.Sprintf("%sStage: %d,\n", indentationValues, dataStorePreparePostReplayParam.Stage))
	b.WriteString(fmt.Sprintf("%sReplayType: %d,\n", indentationValues, dataStorePreparePostReplayParam.ReplayType))
	b.WriteString(fmt.Sprintf("%sCompetitionID: %d,\n", indentationValues, dataStorePreparePostReplayParam.CompetitionID))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, dataStorePreparePostReplayParam.Score))

	if len(dataStorePreparePostReplayParam.Players) == 0 {
		b.WriteString(fmt.Sprintf("%sPlayers: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sPlayers: [\n", indentationValues))

		for i := 0; i < len(dataStorePreparePostReplayParam.Players); i++ {
			str := dataStorePreparePostReplayParam.Players[i].FormatToString(indentationLevel + 2)
			if i == len(dataStorePreparePostReplayParam.Players)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sWinners: %v,\n", indentationValues, dataStorePreparePostReplayParam.Winners))
	b.WriteString(fmt.Sprintf("%sKeyVersion: %d,\n", indentationValues, dataStorePreparePostReplayParam.KeyVersion))
	b.WriteString(fmt.Sprintf("%sExtraData: %v\n", indentationValues, dataStorePreparePostReplayParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostReplayParam returns a new DataStorePreparePostReplayParam
func NewDataStorePreparePostReplayParam() *DataStorePreparePostReplayParam {
	return &DataStorePreparePostReplayParam{}
}
