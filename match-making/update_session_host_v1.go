// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateSessionHostV1 sets the UpdateSessionHostV1 handler function
func (protocol *MatchMakingProtocol) UpdateSessionHostV1(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	protocol.UpdateSessionHostV1Handler = handler
}

func (protocol *MatchMakingProtocol) handleUpdateSessionHostV1(packet nex.PacketInterface) {
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

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.UpdateSessionHostV1Handler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
	}

	go protocol.UpdateSessionHostV1Handler(nil, client, callID, gid)
}
