package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStringSettings sets the GetStringSettings handler function
func (protocol *UtilityProtocol) GetStringSettings(handler func(err error, client *nex.Client, callID uint32, stringSettingIndex uint32)) {
	protocol.GetStringSettingsHandler = handler
}

func (protocol *UtilityProtocol) HandleGetStringSettings(packet nex.PacketInterface) {
	if protocol.GetStringSettingsHandler == nil {
		globals.Logger.Warning("Utility::GetStringSettings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	stringSettingIndex := parametersStream.ReadUInt32LE()

	go protocol.GetStringSettingsHandler(nil, client, callID, stringSettingIndex)
}
