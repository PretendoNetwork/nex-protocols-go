package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetIntegerSettings sets the GetIntegerSettings handler function
func (protocol *UtilityProtocol) GetIntegerSettings(handler func(err error, client *nex.Client, callID uint32, integerSettingIndex uint32)) {
	protocol.GetIntegerSettingsHandler = handler
}

func (protocol *UtilityProtocol) HandleGetIntegerSettings(packet nex.PacketInterface) {
	if protocol.GetIntegerSettingsHandler == nil {
		globals.Logger.Warning("Utility::GetIntegerSettings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	integerSettingIndex := parametersStream.ReadUInt32LE()

	go protocol.GetIntegerSettingsHandler(nil, client, callID, integerSettingIndex)
}
