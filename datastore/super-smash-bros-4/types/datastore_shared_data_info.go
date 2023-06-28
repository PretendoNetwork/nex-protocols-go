package datastore_super_smash_bros_4_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreSharedDataInfo struct {
	nex.Structure
	DataID      uint64
	OwnerID     uint32
	DataType    uint8
	Comment     string
	MetaBinary  []byte
	Profile     []byte
	Rating      int64
	CreatedTime *nex.DateTime
	Info        *DataStoreFileServerObjectInfo
}

// ExtractFromStream extracts a DataStoreSharedDataInfo structure from a stream
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreSharedDataInfo.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.DataID. %s", err.Error())
	}

	dataStoreSharedDataInfo.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.OwnerID. %s", err.Error())
	}

	dataStoreSharedDataInfo.DataType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.DataType. %s", err.Error())
	}

	dataStoreSharedDataInfo.Comment, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Comment. %s", err.Error())
	}

	dataStoreSharedDataInfo.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.MetaBinary. %s", err.Error())
	}

	dataStoreSharedDataInfo.Profile, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.MetaBinary. %s", err.Error())
	}

	dataStoreSharedDataInfo.Rating, err = stream.ReadInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Rating. %s", err.Error())
	}

	dataStoreSharedDataInfo.CreatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.CreatedTime. %s", err.Error())
	}

	info, err := stream.ReadStructure(NewDataStoreFileServerObjectInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Info. %s", err.Error())
	}

	dataStoreSharedDataInfo.Info = info.(*DataStoreFileServerObjectInfo)

	return nil
}

// Bytes encodes the DataStoreSharedDataInfo and returns a byte array
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreSharedDataInfo.DataID)
	stream.WriteUInt32LE(dataStoreSharedDataInfo.OwnerID)
	stream.WriteUInt8(dataStoreSharedDataInfo.DataType)
	stream.WriteString(dataStoreSharedDataInfo.Comment)
	stream.WriteQBuffer(dataStoreSharedDataInfo.MetaBinary)
	stream.WriteQBuffer(dataStoreSharedDataInfo.Profile)
	stream.WriteInt64LE(dataStoreSharedDataInfo.Rating)
	stream.WriteDateTime(dataStoreSharedDataInfo.CreatedTime)
	stream.WriteStructure(dataStoreSharedDataInfo.Info)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSharedDataInfo
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreSharedDataInfo()

	copied.DataID = dataStoreSharedDataInfo.DataID
	copied.OwnerID = dataStoreSharedDataInfo.OwnerID
	copied.DataType = dataStoreSharedDataInfo.DataType
	copied.Comment = dataStoreSharedDataInfo.Comment
	copied.MetaBinary = make([]byte, len(dataStoreSharedDataInfo.MetaBinary))

	copy(copied.MetaBinary, dataStoreSharedDataInfo.MetaBinary)

	copied.Profile = make([]byte, len(dataStoreSharedDataInfo.Profile))

	copy(copied.Profile, dataStoreSharedDataInfo.Profile)

	copied.Rating = dataStoreSharedDataInfo.Rating
	copied.CreatedTime = dataStoreSharedDataInfo.CreatedTime.Copy()
	copied.Info = dataStoreSharedDataInfo.Info.Copy().(*DataStoreFileServerObjectInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSharedDataInfo)

	if dataStoreSharedDataInfo.DataType != other.DataType {
		return false
	}

	if dataStoreSharedDataInfo.DataID != other.DataID {
		return false
	}

	if dataStoreSharedDataInfo.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreSharedDataInfo.DataType != other.DataType {
		return false
	}

	if dataStoreSharedDataInfo.Comment != other.Comment {
		return false
	}

	if !bytes.Equal(dataStoreSharedDataInfo.MetaBinary, other.MetaBinary) {
		return false
	}

	if !bytes.Equal(dataStoreSharedDataInfo.Profile, other.Profile) {
		return false
	}

	if dataStoreSharedDataInfo.Rating != other.Rating {
		return false
	}

	if !dataStoreSharedDataInfo.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !dataStoreSharedDataInfo.Info.Equals(other.Info) {
		return false
	}

	return true
}

// NewDataStoreSharedDataInfo returns a new DataStoreSharedDataInfo
func NewDataStoreSharedDataInfo() *DataStoreSharedDataInfo {
	return &DataStoreSharedDataInfo{}
}
