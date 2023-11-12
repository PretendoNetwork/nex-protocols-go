// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateSessionHostV1 sets the UpdateSessionHostV1 handler function
func (protocol *Protocol) UpdateSessionHostV1(handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32) {
	protocol.updateSessionHostV1Handler = handler
}

func (protocol *Protocol) handleUpdateSessionHostV1(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateSessionHostV1Handler == nil {
		fmt.Println("[Warning] MatchMaking::UpdateSessionHostV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.updateSessionHostV1Handler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateSessionHostV1Handler(nil, packet, callID, gid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
