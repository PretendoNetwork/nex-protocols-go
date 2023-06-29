package datastore_super_smash_bros_4_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReplayPlayer is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreReplayPlayer struct {
	nex.Structure
	Fighter     uint8
	Health      uint8
	WinningRate uint16
	Color       uint8
	Color2      uint8
	PrincipalID uint32
	Country     uint32
	Region      uint8
	Number      uint8
}

// ExtractFromStream extracts a DataStoreReplayPlayer structure from a stream
func (dataStoreReplayPlayer *DataStoreReplayPlayer) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReplayPlayer.Fighter, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Fighter. %s", err.Error())
	}

	dataStoreReplayPlayer.Health, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Health. %s", err.Error())
	}

	dataStoreReplayPlayer.WinningRate, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.WinningRate. %s", err.Error())
	}

	dataStoreReplayPlayer.Color, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Color. %s", err.Error())
	}

	dataStoreReplayPlayer.Color2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Color2. %s", err.Error())
	}

	dataStoreReplayPlayer.PrincipalID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.PrincipalID. %s", err.Error())
	}

	dataStoreReplayPlayer.Country, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Country. %s", err.Error())
	}

	dataStoreReplayPlayer.Region, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Region. %s", err.Error())
	}

	dataStoreReplayPlayer.Number, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Number. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReplayPlayer and returns a byte array
func (dataStoreReplayPlayer *DataStoreReplayPlayer) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStoreReplayPlayer.Fighter)
	stream.WriteUInt8(dataStoreReplayPlayer.Health)
	stream.WriteUInt16LE(dataStoreReplayPlayer.WinningRate)
	stream.WriteUInt8(dataStoreReplayPlayer.Color)
	stream.WriteUInt8(dataStoreReplayPlayer.Color2)
	stream.WriteUInt32LE(dataStoreReplayPlayer.PrincipalID)
	stream.WriteUInt32LE(dataStoreReplayPlayer.Country)
	stream.WriteUInt8(dataStoreReplayPlayer.Region)
	stream.WriteUInt8(dataStoreReplayPlayer.Number)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReplayPlayer
func (dataStoreReplayPlayer *DataStoreReplayPlayer) Copy() nex.StructureInterface {
	copied := NewDataStoreReplayPlayer()

	copied.Fighter = dataStoreReplayPlayer.Fighter
	copied.Health = dataStoreReplayPlayer.Health
	copied.WinningRate = dataStoreReplayPlayer.WinningRate
	copied.Color = dataStoreReplayPlayer.Color
	copied.Color2 = dataStoreReplayPlayer.Color2
	copied.PrincipalID = dataStoreReplayPlayer.PrincipalID
	copied.Country = dataStoreReplayPlayer.Country
	copied.Region = dataStoreReplayPlayer.Region
	copied.Number = dataStoreReplayPlayer.Number

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReplayPlayer *DataStoreReplayPlayer) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReplayPlayer)

	if dataStoreReplayPlayer.Fighter != other.Fighter {
		return false
	}

	if dataStoreReplayPlayer.Health != other.Health {
		return false
	}

	if dataStoreReplayPlayer.WinningRate != other.WinningRate {
		return false
	}

	if dataStoreReplayPlayer.Color != other.Color {
		return false
	}

	if dataStoreReplayPlayer.Color2 != other.Color2 {
		return false
	}

	if dataStoreReplayPlayer.PrincipalID != other.PrincipalID {
		return false
	}

	if dataStoreReplayPlayer.Country != other.Country {
		return false
	}

	if dataStoreReplayPlayer.Region != other.Region {
		return false
	}

	if dataStoreReplayPlayer.Number != other.Number {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReplayPlayer *DataStoreReplayPlayer) String() string {
	return dataStoreReplayPlayer.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReplayPlayer *DataStoreReplayPlayer) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReplayPlayer{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReplayPlayer.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sFighter: %d,\n", indentationValues, dataStoreReplayPlayer.Fighter))
	b.WriteString(fmt.Sprintf("%sHealth: %d,\n", indentationValues, dataStoreReplayPlayer.Health))
	b.WriteString(fmt.Sprintf("%sWinningRate: %d,\n", indentationValues, dataStoreReplayPlayer.WinningRate))
	b.WriteString(fmt.Sprintf("%sColor: %d,\n", indentationValues, dataStoreReplayPlayer.Color))
	b.WriteString(fmt.Sprintf("%sColor2: %d,\n", indentationValues, dataStoreReplayPlayer.Color2))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %d,\n", indentationValues, dataStoreReplayPlayer.PrincipalID))
	b.WriteString(fmt.Sprintf("%sCountry: %d,\n", indentationValues, dataStoreReplayPlayer.Country))
	b.WriteString(fmt.Sprintf("%sRegion: %d,\n", indentationValues, dataStoreReplayPlayer.Region))
	b.WriteString(fmt.Sprintf("%sNumber: %d\n", indentationValues, dataStoreReplayPlayer.Number))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReplayPlayer returns a new DataStoreReplayPlayer
func NewDataStoreReplayPlayer() *DataStoreReplayPlayer {
	return &DataStoreReplayPlayer{}
}
