// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSearchSimpleSearchObjectByObjectIDs(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.SearchSimpleSearchObjectByObjectIDs == nil {
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
		_, errorCode = protocol.SearchSimpleSearchObjectByObjectIDs(fmt.Errorf("Failed to read objectIDs from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SearchSimpleSearchObjectByObjectIDs(nil, packet, callID, objectIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
