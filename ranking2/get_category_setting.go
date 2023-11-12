// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCategorySetting sets the GetCategorySetting handler function
func (protocol *Protocol) GetCategorySetting(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32) uint32) {
	protocol.getCategorySettingHandler = handler
}

func (protocol *Protocol) handleGetCategorySetting(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCategorySettingHandler == nil {
		globals.Logger.Warning("Ranking2::GetCategorySetting not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getCategorySettingHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getCategorySettingHandler(nil, packet, callID, category)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
