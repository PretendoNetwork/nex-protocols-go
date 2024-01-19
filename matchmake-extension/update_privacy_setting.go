// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePrivacySetting(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdatePrivacySetting == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdatePrivacySetting not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	onlineStatus := types.NewPrimitiveBool(false)
	err = onlineStatus.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePrivacySetting(fmt.Errorf("Failed to read onlineStatus from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participationCommunity := types.NewPrimitiveBool(false)
	err = participationCommunity.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePrivacySetting(fmt.Errorf("Failed to read participationCommunity from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdatePrivacySetting(nil, packet, callID, onlineStatus, participationCommunity)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
