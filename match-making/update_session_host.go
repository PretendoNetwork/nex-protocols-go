// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateSessionHost(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateSessionHost == nil {
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
		errorCode = protocol.UpdateSessionHost(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	isMigrateOwner, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.UpdateSessionHost(fmt.Errorf("Failed to read isMigrateOwner from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.UpdateSessionHost(nil, packet, callID, gid, isMigrateOwner)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
