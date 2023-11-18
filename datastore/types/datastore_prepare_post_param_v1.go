// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePreparePostParamV1 is a data structure used by the DataStore protocol
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

	dataStorePreparePostParamV1.Permission, err = nex.StreamReadStructure(stream, NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Permission. %s", err.Error())
	}

	dataStorePreparePostParamV1.DelPermission, err = nex.StreamReadStructure(stream, NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.DelPermission. %s", err.Error())
	}

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

	ratingInitParams, err := nex.StreamReadListStructure(stream, NewDataStoreRatingInitParamWithSlot())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.RatingInitParams. %s", err.Error())
	}

	dataStorePreparePostParamV1.RatingInitParams = ratingInitParams

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

	copied.SetStructureVersion(dataStorePreparePostParamV1.StructureVersion())

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

	if dataStorePreparePostParamV1.StructureVersion() != other.StructureVersion() {
		return false
	}

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

// String returns a string representation of the struct
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) String() string {
	return dataStorePreparePostParamV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostParamV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePreparePostParamV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStorePreparePostParamV1.Size))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, dataStorePreparePostParamV1.Name))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStorePreparePostParamV1.DataType))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x,\n", indentationValues, dataStorePreparePostParamV1.MetaBinary))

	if dataStorePreparePostParamV1.Permission != nil {
		b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStorePreparePostParamV1.Permission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPermission: nil,\n", indentationValues))
	}

	if dataStorePreparePostParamV1.DelPermission != nil {
		b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStorePreparePostParamV1.DelPermission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDelPermission: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sFlag: %d,\n", indentationValues, dataStorePreparePostParamV1.Flag))
	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, dataStorePreparePostParamV1.Period))
	b.WriteString(fmt.Sprintf("%sReferDataID: %d,\n", indentationValues, dataStorePreparePostParamV1.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %v,\n", indentationValues, dataStorePreparePostParamV1.Tags))

	if len(dataStorePreparePostParamV1.RatingInitParams) == 0 {
		b.WriteString(fmt.Sprintf("%sRatingInitParams: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRatingInitParams: [\n", indentationValues))

		for i := 0; i < len(dataStorePreparePostParamV1.RatingInitParams); i++ {
			str := dataStorePreparePostParamV1.RatingInitParams[i].FormatToString(indentationLevel + 2)
			if i == len(dataStorePreparePostParamV1.RatingInitParams)-1 {
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

// NewDataStorePreparePostParamV1 returns a new DataStorePreparePostParamV1
func NewDataStorePreparePostParamV1() *DataStorePreparePostParamV1 {
	return &DataStorePreparePostParamV1{
		Size:             0,
		Name:             "",
		DataType:         0,
		MetaBinary:       make([]byte, 0),
		Permission:       NewDataStorePermission(),
		DelPermission:    NewDataStorePermission(),
		Flag:             0,
		Period:           0,
		ReferDataID:      0,
		Tags:             make([]string, 0),
		RatingInitParams: make([]*DataStoreRatingInitParamWithSlot, 0),
	}
}
