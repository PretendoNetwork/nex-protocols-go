// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetIntegerSettings sets the GetIntegerSettings handler function
func (protocol *Protocol) GetIntegerSettings(handler func(err error, client *nex.Client, callID uint32, integerSettingIndex uint32) uint32) {
	protocol.getIntegerSettingsHandler = handler
}

func (protocol *Protocol) handleGetIntegerSettings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getIntegerSettingsHandler == nil {
		globals.Logger.Warning("Utility::GetIntegerSettings not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	integerSettingIndex, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getIntegerSettingsHandler(fmt.Errorf("Failed to read integerSettingIndex from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getIntegerSettingsHandler(nil, client, callID, integerSettingIndex)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
