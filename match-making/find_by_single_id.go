package match_making

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindBySingleID sets the FindBySingleID handler function
func (protocol *MatchMakingProtocol) FindBySingleID(handler func(err error, client *nex.Client, callID uint32, id uint32)) {
	protocol.FindBySingleIDHandler = handler
}

func (protocol *MatchMakingProtocol) HandleFindBySingleID(packet nex.PacketInterface) {
	if protocol.FindBySingleIDHandler == nil {
		globals.Logger.Warning("MatchMaking::FindBySingleID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id := parametersStream.ReadUInt32LE()

	go protocol.FindBySingleIDHandler(nil, client, callID, id)
}
