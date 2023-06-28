package friends_wiiu_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// MiiV2 contains data about a Mii
type MiiV2 struct {
	nex.Structure
	Name     string
	Unknown1 uint8
	Unknown2 uint8
	Data     []byte
	Datetime *nex.DateTime
}

// Bytes encodes the MiiV2 and returns a byte array
func (mii *MiiV2) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(mii.Name)
	stream.WriteUInt8(mii.Unknown1)
	stream.WriteUInt8(mii.Unknown2)
	stream.WriteBuffer(mii.Data)
	stream.WriteDateTime(mii.Datetime)

	return stream.Bytes()
}

// ExtractFromStream extracts a MiiV2 structure from a stream
func (mii *MiiV2) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	mii.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Name. %s", err.Error())
	}

	mii.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown1. %s", err.Error())
	}

	mii.Unknown2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown2. %s", err.Error())
	}

	mii.Data, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Data. %s", err.Error())
	}

	mii.Datetime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Datetime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiV2
func (mii *MiiV2) Copy() nex.StructureInterface {
	copied := NewMiiV2()

	copied.Name = mii.Name
	copied.Unknown1 = mii.Unknown1
	copied.Unknown2 = mii.Unknown2
	copied.Data = make([]byte, len(mii.Data))

	copy(copied.Data, mii.Data)

	copied.Datetime = mii.Datetime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (mii *MiiV2) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MiiV2)

	if mii.Name != other.Name {
		return false
	}

	if mii.Unknown1 != other.Unknown1 {
		return false
	}

	if mii.Unknown2 != other.Unknown2 {
		return false
	}

	if !bytes.Equal(mii.Data, other.Data) {
		return false
	}

	if !mii.Datetime.Equals(other.Datetime) {
		return false
	}

	return true
}

// NewMiiV2 returns a new MiiV2
func NewMiiV2() *MiiV2 {
	return &MiiV2{}
}
