// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRating sets the GetRating handler function
func (protocol *Protocol) GetRating(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64) uint32) {
	protocol.getRatingHandler = handler
}

func (protocol *Protocol) handleGetRating(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRatingHandler == nil {
		globals.Logger.Warning("DataStore::GetRating not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		errorCode = protocol.getRatingHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	accessPassword, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getRatingHandler(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRatingHandler(nil, packet, callID, target.(*datastore_types.DataStoreRatingTarget), accessPassword)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
