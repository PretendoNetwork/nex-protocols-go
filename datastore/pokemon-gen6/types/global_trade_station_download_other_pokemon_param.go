// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationDownloadOtherPokemonParam holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationDownloadOtherPokemonParam struct {
	types.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// ExtractFrom extracts the GlobalTradeStationDownloadOtherPokemonParam from the given readable
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationDownloadOtherPokemonParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationDownloadOtherPokemonParam header. %s", err.Error())
	}

	err = globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadOtherPokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationDownloadOtherPokemonParam to the given writable
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationDownloadOtherPokemonParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationDownloadOtherPokemonParam
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationDownloadOtherPokemonParam()

	copied.StructureVersion = globalTradeStationDownloadOtherPokemonParam.StructureVersion

	copied.PrepareUploadKey = globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDownloadOtherPokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDownloadOtherPokemonParam)

	if globalTradeStationDownloadOtherPokemonParam.StructureVersion != other.StructureVersion {
		return false
	}

	return globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey)
}

// String returns a string representation of the struct
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) String() string {
	return globalTradeStationDownloadOtherPokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadOtherPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationDownloadOtherPokemonParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadOtherPokemonParam returns a new GlobalTradeStationDownloadOtherPokemonParam
func NewGlobalTradeStationDownloadOtherPokemonParam() *GlobalTradeStationDownloadOtherPokemonParam {
	return &GlobalTradeStationDownloadOtherPokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
	}
}
