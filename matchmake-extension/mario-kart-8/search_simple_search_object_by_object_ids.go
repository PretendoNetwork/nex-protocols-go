// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchSimpleSearchObjectByObjectIDs sets the SearchSimpleSearchObjectByObjectIDs handler function
func (protocol *Protocol) SearchSimpleSearchObjectByObjectIDs(handler func(err error, packet nex.PacketInterface, callID uint32, objectIDs []uint32) uint32) {
	protocol.searchSimpleSearchObjectByObjectIDsHandler = handler
}

func (protocol *Protocol) handleSearchSimpleSearchObjectByObjectIDs(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.searchSimpleSearchObjectByObjectIDsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::SearchSimpleSearchObjectByObjectIDs not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	objectIDs, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.searchSimpleSearchObjectByObjectIDsHandler(fmt.Errorf("Failed to read objectIDs from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.searchSimpleSearchObjectByObjectIDsHandler(nil, packet, callID, objectIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
