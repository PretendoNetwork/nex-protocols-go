package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePresence sets the UpdatePresence handler function
func (protocol *Friends3DSProtocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *NintendoPresence, showGame bool)) {
	protocol.UpdatePresenceHandler = handler
}

func (protocol *Friends3DSProtocol) HandleUpdatePresence(packet nex.PacketInterface) {
	if protocol.UpdatePresenceHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePresence not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nintendoPresence, err := parametersStream.ReadStructure(NewNintendoPresence())
	if err != nil {
		go protocol.UpdatePresenceHandler(err, client, callID, nil, false)
		return
	}

	showGame := (parametersStream.ReadUInt8() == 1)

	go protocol.UpdatePresenceHandler(nil, client, callID, nintendoPresence.(*NintendoPresence), showGame)
}
