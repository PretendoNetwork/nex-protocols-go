// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSearchSimpleSearchObjectByObjectIDs(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.SearchSimpleSearchObjectByObjectIDs == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::SearchSimpleSearchObjectByObjectIDs not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	objectIDs := types.NewList[*types.PrimitiveU32]()
	objectIDs.Type = types.NewPrimitiveU32(0)
	err = objectIDs.ExtractFrom(parametersStream)
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
