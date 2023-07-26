// Package utility implements the Utility NEX protocol
package utility

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetIntegerSettings sets the GetIntegerSettings handler function
func (protocol *UtilityProtocol) GetIntegerSettings(handler func(err error, client *nex.Client, callID uint32, integerSettingIndex uint32)) {
	protocol.getIntegerSettingsHandler = handler
}

func (protocol *UtilityProtocol) handleGetIntegerSettings(packet nex.PacketInterface) {
	if protocol.getIntegerSettingsHandler == nil {
		globals.Logger.Warning("Utility::GetIntegerSettings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	integerSettingIndex, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getIntegerSettingsHandler(fmt.Errorf("Failed to read integerSettingIndex from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getIntegerSettingsHandler(nil, client, callID, integerSettingIndex)
}
