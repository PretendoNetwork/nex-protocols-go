// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReqGetAdditionalMeta is a data structure used by the DataStore protocol
type DataStoreReqGetAdditionalMeta struct {
	nex.Structure
	OwnerID    *nex.PID
	DataType   uint16
	Version    uint16
	MetaBinary []byte
}

// ExtractFromStream extracts a DataStoreReqGetAdditionalMeta structure from a stream
func (d *DataStoreReqGetAdditionalMeta) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	d.OwnerID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.OwnerID. %s", err.Error())
	}

	d.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.DataType. %s", err.Error())
	}

	d.Version, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.Version. %s", err.Error())
	}

	d.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetAdditionalMeta.MetaBinary. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReqGetAdditionalMeta and returns a byte array
func (d *DataStoreReqGetAdditionalMeta) Bytes(stream *nex.StreamOut) []byte {
	stream.WritePID(d.OwnerID)
	stream.WriteUInt16LE(d.DataType)
	stream.WriteUInt16LE(d.Version)
	stream.WriteQBuffer(d.MetaBinary)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetAdditionalMeta
func (d *DataStoreReqGetAdditionalMeta) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetAdditionalMeta()

	copied.SetStructureVersion(d.StructureVersion())

	copied.OwnerID = d.OwnerID.Copy()
	copied.DataType = d.DataType
	copied.Version = d.Version
	copied.MetaBinary = make([]byte, len(d.MetaBinary))

	copy(copied.MetaBinary, d.MetaBinary)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (d *DataStoreReqGetAdditionalMeta) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetAdditionalMeta)

	if d.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !d.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if d.DataType != other.DataType {
		return false
	}

	if d.Version != other.Version {
		return false
	}

	if !bytes.Equal(d.MetaBinary, other.MetaBinary) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (d *DataStoreReqGetAdditionalMeta) String() string {
	return d.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (d *DataStoreReqGetAdditionalMeta) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetAdditionalMeta{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, d.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, d.OwnerID))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, d.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %d,\n", indentationValues, d.Version))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x\n", indentationValues, d.MetaBinary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetAdditionalMeta returns a new DataStoreReqGetAdditionalMeta
func NewDataStoreReqGetAdditionalMeta() *DataStoreReqGetAdditionalMeta {
	return &DataStoreReqGetAdditionalMeta{
		OwnerID:    nex.NewPID[uint32](0),
		DataType:   0,
		Version:    0,
		MetaBinary: make([]byte, 0),
	}
}
