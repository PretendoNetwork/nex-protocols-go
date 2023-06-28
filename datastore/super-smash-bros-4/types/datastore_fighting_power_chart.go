package datastore_super_smash_bros_4_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreFightingPowerChart struct {
	nex.Structure
	UserNum uint32
	Chart   []*DataStoreFightingPowerScore
}

// ExtractFromStream extracts a DataStoreFightingPowerChart structure from a stream
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreFightingPowerChart.UserNum, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart.UserNum. %s", err.Error())
	}

	chart, err := stream.ReadListStructure(NewDataStoreFightingPowerScore())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart.Chart. %s", err.Error())
	}

	dataStoreFightingPowerChart.Chart = chart.([]*DataStoreFightingPowerScore)

	return nil
}

// Bytes encodes the DataStoreFightingPowerChart and returns a byte array
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreFightingPowerChart.UserNum)
	stream.WriteListStructure(dataStoreFightingPowerChart.Chart)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreFightingPowerChart
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Copy() nex.StructureInterface {
	copied := NewDataStoreFightingPowerChart()

	copied.UserNum = dataStoreFightingPowerChart.UserNum
	copied.Chart = make([]*DataStoreFightingPowerScore, len(dataStoreFightingPowerChart.Chart))

	for i := 0; i < len(dataStoreFightingPowerChart.Chart); i++ {
		copied.Chart[i] = dataStoreFightingPowerChart.Chart[i].Copy().(*DataStoreFightingPowerScore)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreFightingPowerChart)

	if dataStoreFightingPowerChart.UserNum != other.UserNum {
		return false
	}

	if len(dataStoreFightingPowerChart.Chart) != len(other.Chart) {
		return false
	}

	for i := 0; i < len(dataStoreFightingPowerChart.Chart); i++ {
		if !dataStoreFightingPowerChart.Chart[i].Equals(other.Chart[i]) {
			return false
		}
	}

	return true
}

// NewDataStoreFightingPowerChart returns a new DataStoreFightingPowerChart
func NewDataStoreFightingPowerChart() *DataStoreFightingPowerChart {
	return &DataStoreFightingPowerChart{}
}
