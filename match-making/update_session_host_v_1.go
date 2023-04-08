package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSessionURLs sets the GetSessionURLs handler function
func (protocol *MatchMakingProtocol) UpdateSessionHostV1(handler func(err error, client *nex.Client, callID uint32, gatheringId uint32)) {
	protocol.UpdateSessionHostV1Handler = handler
}

func (protocol *MatchMakingProtocol) HandleUpdateSessionHostV1(packet nex.PacketInterface) {
	if protocol.UpdateSessionHostV1Handler == nil {
		fmt.Println("[Warning] MatchMaking::UpdateSessionHostV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gatheringId := parametersStream.ReadUInt32LE()

	go protocol.UpdateSessionHostV1Handler(nil, client, callID, gatheringId)
}
