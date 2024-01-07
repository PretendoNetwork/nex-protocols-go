// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReplayPlayer is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreReplayPlayer struct {
	types.Structure
	Fighter     *types.PrimitiveU8
	Health      *types.PrimitiveU8
	WinningRate *types.PrimitiveU16
	Color       *types.PrimitiveU8
	Color2      *types.PrimitiveU8
	PrincipalID *types.PrimitiveU32
	Country     *types.PrimitiveU32
	Region      *types.PrimitiveU8
	Number      *types.PrimitiveU8
}

// ExtractFrom extracts the DataStoreReplayPlayer from the given readable
func (dataStoreReplayPlayer *DataStoreReplayPlayer) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreReplayPlayer.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReplayPlayer header. %s", err.Error())
	}

	err = dataStoreReplayPlayer.Fighter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Fighter. %s", err.Error())
	}

	err = dataStoreReplayPlayer.Health.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Health. %s", err.Error())
	}

	err = dataStoreReplayPlayer.WinningRate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.WinningRate. %s", err.Error())
	}

	err = dataStoreReplayPlayer.Color.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Color. %s", err.Error())
	}

	err = dataStoreReplayPlayer.Color2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Color2. %s", err.Error())
	}

	err = dataStoreReplayPlayer.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.PrincipalID. %s", err.Error())
	}

	err = dataStoreReplayPlayer.Country.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Country. %s", err.Error())
	}

	err = dataStoreReplayPlayer.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Region. %s", err.Error())
	}

	err = dataStoreReplayPlayer.Number.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayPlayer.Number. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReplayPlayer to the given writable
func (dataStoreReplayPlayer *DataStoreReplayPlayer) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreReplayPlayer.Fighter.WriteTo(contentWritable)
	dataStoreReplayPlayer.Health.WriteTo(contentWritable)
	dataStoreReplayPlayer.WinningRate.WriteTo(contentWritable)
	dataStoreReplayPlayer.Color.WriteTo(contentWritable)
	dataStoreReplayPlayer.Color2.WriteTo(contentWritable)
	dataStoreReplayPlayer.PrincipalID.WriteTo(contentWritable)
	dataStoreReplayPlayer.Country.WriteTo(contentWritable)
	dataStoreReplayPlayer.Region.WriteTo(contentWritable)
	dataStoreReplayPlayer.Number.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreReplayPlayer.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReplayPlayer
func (dataStoreReplayPlayer *DataStoreReplayPlayer) Copy() types.RVType {
	copied := NewDataStoreReplayPlayer()

	copied.StructureVersion = dataStoreReplayPlayer.StructureVersion

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
func (dataStoreReplayPlayer *DataStoreReplayPlayer) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReplayPlayer); !ok {
		return false
	}

	other := o.(*DataStoreReplayPlayer)

	if dataStoreReplayPlayer.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReplayPlayer.Fighter.Equals(other.Fighter) {
		return false
	}

	if !dataStoreReplayPlayer.Health.Equals(other.Health) {
		return false
	}

	if !dataStoreReplayPlayer.WinningRate.Equals(other.WinningRate) {
		return false
	}

	if !dataStoreReplayPlayer.Color.Equals(other.Color) {
		return false
	}

	if !dataStoreReplayPlayer.Color2.Equals(other.Color2) {
		return false
	}

	if !dataStoreReplayPlayer.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !dataStoreReplayPlayer.Country.Equals(other.Country) {
		return false
	}

	if !dataStoreReplayPlayer.Region.Equals(other.Region) {
		return false
	}

	if !dataStoreReplayPlayer.Number.Equals(other.Number) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReplayPlayer.StructureVersion))
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
