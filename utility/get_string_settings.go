// Package utility implements the Utility NEX protocol
package utility

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStringSettings sets the GetStringSettings handler function
func (protocol *UtilityProtocol) GetStringSettings(handler func(err error, client *nex.Client, callID uint32, stringSettingIndex uint32)) {
	protocol.getStringSettingsHandler = handler
}

func (protocol *UtilityProtocol) handleGetStringSettings(packet nex.PacketInterface) {
	if protocol.getStringSettingsHandler == nil {
		globals.Logger.Warning("Utility::GetStringSettings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	stringSettingIndex, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getStringSettingsHandler(fmt.Errorf("Failed to read stringSettingIndex from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getStringSettingsHandler(nil, client, callID, stringSettingIndex)
}
