package datastore_super_smash_bros_4_types

import (
	"fmt"

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

// NewDataStoreReplayMetaInfo returns a new DataStoreReplayMetaInfo
func NewDataStoreReplayMetaInfo() *DataStoreReplayMetaInfo {
	return &DataStoreReplayMetaInfo{}
}
