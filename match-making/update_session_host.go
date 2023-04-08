package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSessionURLs sets the GetSessionURLs handler function
func (protocol *MatchMakingProtocol) UpdateSessionHost(handler func(err error, client *nex.Client, callID uint32, gatheringId uint32)) {
	protocol.UpdateSessionHostHandler = handler
}

func (protocol *MatchMakingProtocol) HandleUpdateSessionHost(packet nex.PacketInterface) {
	if protocol.UpdateSessionHostHandler == nil {
		fmt.Println("[Warning] MatchMaking::UpdateSessionHost not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gatheringId := parametersStream.ReadUInt32LE()

	go protocol.UpdateSessionHostHandler(nil, client, callID, gatheringId)
}
