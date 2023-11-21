// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetGatheringRelations(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetGatheringRelations == nil {
		globals.Logger.Warning("MatchMakingExt::GetGatheringRelations not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetGatheringRelations(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	descr, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.GetGatheringRelations(fmt.Errorf("Failed to read descr from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetGatheringRelations(nil, packet, callID, id, descr)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
