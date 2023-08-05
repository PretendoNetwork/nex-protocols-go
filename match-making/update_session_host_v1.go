// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateSessionHostV1 sets the UpdateSessionHostV1 handler function
func (protocol *Protocol) UpdateSessionHostV1(handler func(err error, client *nex.Client, callID uint32, gid uint32) uint32) {
	protocol.updateSessionHostV1Handler = handler
}

func (protocol *Protocol) handleUpdateSessionHostV1(packet nex.PacketInterface) {
	if protocol.updateSessionHostV1Handler == nil {
		fmt.Println("[Warning] MatchMaking::UpdateSessionHostV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateSessionHostV1Handler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
	}

	go protocol.updateSessionHostV1Handler(nil, client, callID, gid)
}
