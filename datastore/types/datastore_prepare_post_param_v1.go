package datastore_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStorePreparePostParamV1 struct {
	nex.Structure

	Size             uint32
	Name             string
	DataType         uint16
	MetaBinary       []byte
	Permission       *DataStorePermission
	DelPermission    *DataStorePermission
	Flag             uint32
	Period           uint16
	ReferDataID      uint32
	Tags             []string
	RatingInitParams []*DataStoreRatingInitParamWithSlot
}

// ExtractFromStream extracts a DataStorePreparePostParamV1 structure from a stream
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePreparePostParamV1.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Size. %s", err.Error())
	}

	dataStorePreparePostParamV1.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Name. %s", err.Error())
	}

	dataStorePreparePostParamV1.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.DataType. %s", err.Error())
	}

	dataStorePreparePostParamV1.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.MetaBinary. %s", err.Error())
	}

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Permission. %s", err.Error())
	}

	dataStorePreparePostParamV1.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.DelPermission. %s", err.Error())
	}

	dataStorePreparePostParamV1.DelPermission = delPermission.(*DataStorePermission)
	dataStorePreparePostParamV1.Flag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Flag. %s", err.Error())
	}

	dataStorePreparePostParamV1.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Period. %s", err.Error())
	}

	dataStorePreparePostParamV1.ReferDataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.ReferDataID. %s", err.Error())
	}

	dataStorePreparePostParamV1.Tags, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Tags. %s", err.Error())
	}

	ratingInitParams, err := stream.ReadListStructure(NewDataStoreRatingInitParamWithSlot())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.RatingInitParams. %s", err.Error())
	}

	dataStorePreparePostParamV1.RatingInitParams = ratingInitParams.([]*DataStoreRatingInitParamWithSlot)

	return nil
}

// Bytes encodes the DataStorePreparePostParamV1 and returns a byte array
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStorePreparePostParamV1.Size)
	stream.WriteString(dataStorePreparePostParamV1.Name)
	stream.WriteUInt16LE(dataStorePreparePostParamV1.DataType)
	stream.WriteQBuffer(dataStorePreparePostParamV1.MetaBinary)
	stream.WriteStructure(dataStorePreparePostParamV1.Permission)
	stream.WriteStructure(dataStorePreparePostParamV1.DelPermission)
	stream.WriteUInt32LE(dataStorePreparePostParamV1.Flag)
	stream.WriteUInt16LE(dataStorePreparePostParamV1.Period)
	stream.WriteUInt32LE(dataStorePreparePostParamV1.ReferDataID)
	stream.WriteListString(dataStorePreparePostParamV1.Tags)
	stream.WriteListStructure(dataStorePreparePostParamV1.RatingInitParams)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePreparePostParamV1
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) Copy() nex.StructureInterface {
	copied := NewDataStorePreparePostParamV1()

	copied.Size = dataStorePreparePostParamV1.Size
	copied.Name = dataStorePreparePostParamV1.Name
	copied.DataType = dataStorePreparePostParamV1.DataType
	copied.MetaBinary = make([]byte, len(dataStorePreparePostParamV1.MetaBinary))

	copy(copied.MetaBinary, dataStorePreparePostParamV1.MetaBinary)

	copied.Permission = dataStorePreparePostParamV1.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStorePreparePostParamV1.DelPermission.Copy().(*DataStorePermission)
	copied.Flag = dataStorePreparePostParamV1.Flag
	copied.Period = dataStorePreparePostParamV1.Period
	copied.ReferDataID = dataStorePreparePostParamV1.ReferDataID
	copied.Tags = make([]string, len(dataStorePreparePostParamV1.Tags))

	copy(copied.Tags, dataStorePreparePostParamV1.Tags)

	copied.RatingInitParams = make([]*DataStoreRatingInitParamWithSlot, len(dataStorePreparePostParamV1.RatingInitParams))

	for i := 0; i < len(dataStorePreparePostParamV1.RatingInitParams); i++ {
		copied.RatingInitParams[i] = dataStorePreparePostParamV1.RatingInitParams[i].Copy().(*DataStoreRatingInitParamWithSlot)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePreparePostParamV1)

	if dataStorePreparePostParamV1.Size != other.Size {
		return false
	}

	if dataStorePreparePostParamV1.Name != other.Name {
		return false
	}

	if dataStorePreparePostParamV1.DataType != other.DataType {
		return false
	}

	if !bytes.Equal(dataStorePreparePostParamV1.MetaBinary, other.MetaBinary) {
		return false
	}

	if !dataStorePreparePostParamV1.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStorePreparePostParamV1.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStorePreparePostParamV1.Flag != other.Flag {
		return false
	}

	if dataStorePreparePostParamV1.Period != other.Period {
		return false
	}

	if dataStorePreparePostParamV1.ReferDataID != other.ReferDataID {
		return false
	}

	if len(dataStorePreparePostParamV1.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParamV1.Tags); i++ {
		if dataStorePreparePostParamV1.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if len(dataStorePreparePostParamV1.RatingInitParams) != len(other.RatingInitParams) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParamV1.RatingInitParams); i++ {
		if dataStorePreparePostParamV1.RatingInitParams[i] != other.RatingInitParams[i] {
			return false
		}
	}

	return true
}

// NewDataStorePreparePostParamV1 returns a new DataStorePreparePostParamV1
func NewDataStorePreparePostParamV1() *DataStorePreparePostParamV1 {
	return &DataStorePreparePostParamV1{}
}
