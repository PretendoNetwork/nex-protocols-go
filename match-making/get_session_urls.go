package match_making

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSessionURLs sets the GetSessionURLs handler function
func (protocol *MatchMakingProtocol) GetSessionURLs(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	protocol.GetSessionURLsHandler = handler
}

func (protocol *MatchMakingProtocol) HandleGetSessionURLs(packet nex.PacketInterface) {
	if protocol.GetSessionURLsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetSessionURLs not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid := parametersStream.ReadUInt32LE()

	go protocol.GetSessionURLsHandler(nil, client, callID, gid)
}
