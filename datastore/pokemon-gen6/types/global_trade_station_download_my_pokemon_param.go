// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationDownloadMyPokemonParam holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationDownloadMyPokemonParam struct {
	types.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// ExtractFrom extracts the GlobalTradeStationDownloadMyPokemonParam from the given readable
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationDownloadMyPokemonParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationDownloadMyPokemonParam header. %s", err.Error())
	}

	err = globalTradeStationDownloadMyPokemonParam.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationDownloadMyPokemonParam to the given writable
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationDownloadMyPokemonParam.PrepareUploadKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationDownloadMyPokemonParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationDownloadMyPokemonParam
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationDownloadMyPokemonParam()

	copied.StructureVersion = globalTradeStationDownloadMyPokemonParam.StructureVersion

	copied.PrepareUploadKey = globalTradeStationDownloadMyPokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDownloadMyPokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDownloadMyPokemonParam)

	if globalTradeStationDownloadMyPokemonParam.StructureVersion != other.StructureVersion {
		return false
	}

	return globalTradeStationDownloadMyPokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey)
}

// String returns a string representation of the struct
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) String() string {
	return globalTradeStationDownloadMyPokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadMyPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationDownloadMyPokemonParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationDownloadMyPokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadMyPokemonParam returns a new GlobalTradeStationDownloadMyPokemonParam
func NewGlobalTradeStationDownloadMyPokemonParam() *GlobalTradeStationDownloadMyPokemonParam {
	return &GlobalTradeStationDownloadMyPokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
	}
}
