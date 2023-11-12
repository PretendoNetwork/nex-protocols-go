// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePrivacySetting sets the UpdatePrivacySetting handler function
func (protocol *Protocol) UpdatePrivacySetting(handler func(err error, packet nex.PacketInterface, callID uint32, onlineStatus bool, participationCommunity bool) uint32) {
	protocol.updatePrivacySettingHandler = handler
}

func (protocol *Protocol) handleUpdatePrivacySetting(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updatePrivacySettingHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdatePrivacySetting not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	onlineStatus, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.updatePrivacySettingHandler(fmt.Errorf("Failed to read onlineStatus from parameters. %s", err.Error()), packet, callID, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participationCommunity, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.updatePrivacySettingHandler(fmt.Errorf("Failed to read participationCommunity from parameters. %s", err.Error()), packet, callID, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updatePrivacySettingHandler(nil, packet, callID, onlineStatus, participationCommunity)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
