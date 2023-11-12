// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateSessionHost sets the UpdateSessionHost handler function
func (protocol *Protocol) UpdateSessionHost(handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32, isMigrateOwner bool) uint32) {
	protocol.updateSessionHostHandler = handler
}

func (protocol *Protocol) handleUpdateSessionHost(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateSessionHostHandler == nil {
		fmt.Println("[Warning] MatchMaking::UpdateSessionHost not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.updateSessionHostHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	isMigrateOwner, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.updateSessionHostHandler(fmt.Errorf("Failed to read isMigrateOwner from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateSessionHostHandler(nil, packet, callID, gid, isMigrateOwner)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
