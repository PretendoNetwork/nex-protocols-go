package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePreference sets the UpdatePreference handler function
func (protocol *Friends3DSProtocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, publicMode bool, showGame bool, showPlayedGame bool)) {
	protocol.UpdatePreferenceHandler = handler
}

func (protocol *Friends3DSProtocol) HandleUpdatePreference(packet nex.PacketInterface) {
	if protocol.UpdatePreferenceHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePreference not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	publicMode := (parametersStream.ReadUInt8() == 1)
	showGame := (parametersStream.ReadUInt8() == 1)
	showPlayedGame := (parametersStream.ReadUInt8() == 1)

	go protocol.UpdatePreferenceHandler(nil, client, callID, publicMode, showGame, showPlayedGame)
}
