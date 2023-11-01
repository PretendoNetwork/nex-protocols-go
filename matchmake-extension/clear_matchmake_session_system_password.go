// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearMatchmakeSessionSystemPassword sets the ClearMatchmakeSessionSystemPassword handler function
func (protocol *Protocol) ClearMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32) {
	protocol.clearMatchmakeSessionSystemPasswordHandler = handler
}

func (protocol *Protocol) handleClearMatchmakeSessionSystemPassword(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.clearMatchmakeSessionSystemPasswordHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ClearMatchmakeSessionSystemPassword not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.clearMatchmakeSessionSystemPasswordHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.clearMatchmakeSessionSystemPasswordHandler(nil, packet, callID, gid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
