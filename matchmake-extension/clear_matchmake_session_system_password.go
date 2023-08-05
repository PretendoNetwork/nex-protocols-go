// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearMatchmakeSessionSystemPassword sets the ClearMatchmakeSessionSystemPassword handler function
func (protocol *Protocol) ClearMatchmakeSessionSystemPassword(handler func(err error, client *nex.Client, callID uint32, gid uint32) uint32) {
	protocol.clearMatchmakeSessionSystemPasswordHandler = handler
}

func (protocol *Protocol) handleClearMatchmakeSessionSystemPassword(packet nex.PacketInterface) {
	if protocol.clearMatchmakeSessionSystemPasswordHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ClearMatchmakeSessionSystemPassword not implemented")
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
		go protocol.clearMatchmakeSessionSystemPasswordHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.clearMatchmakeSessionSystemPasswordHandler(nil, client, callID, gid)
}
