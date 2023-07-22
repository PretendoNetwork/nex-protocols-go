// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearMatchmakeSessionSystemPassword sets the ClearMatchmakeSessionSystemPassword handler function
func (protocol *MatchmakeExtensionProtocol) ClearMatchmakeSessionSystemPassword(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	protocol.clearMatchmakeSessionSystemPasswordHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleClearMatchmakeSessionSystemPassword(packet nex.PacketInterface) {
	if protocol.clearMatchmakeSessionSystemPasswordHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ClearMatchmakeSessionSystemPassword not implemented")
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
		go protocol.clearMatchmakeSessionSystemPasswordHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.clearMatchmakeSessionSystemPasswordHandler(nil, client, callID, gid)
}
