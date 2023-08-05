// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStringSettings sets the GetStringSettings handler function
func (protocol *Protocol) GetStringSettings(handler func(err error, client *nex.Client, callID uint32, stringSettingIndex uint32) uint32) {
	protocol.getStringSettingsHandler = handler
}

func (protocol *Protocol) handleGetStringSettings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getStringSettingsHandler == nil {
		globals.Logger.Warning("Utility::GetStringSettings not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	stringSettingIndex, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getStringSettingsHandler(fmt.Errorf("Failed to read stringSettingIndex from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getStringSettingsHandler(nil, client, callID, stringSettingIndex)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
