// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GenerateMatchmakeSessionSystemPassword sets the GenerateMatchmakeSessionSystemPassword handler function
func (protocol *Protocol) GenerateMatchmakeSessionSystemPassword(handler func(err error, client *nex.Client, callID uint32, GID uint32) uint32) {
	protocol.generateMatchmakeSessionSystemPasswordHandler = handler
}

func (protocol *Protocol) handleGenerateMatchmakeSessionSystemPassword(packet nex.PacketInterface) {
	if protocol.generateMatchmakeSessionSystemPasswordHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GenerateMatchmakeSessionSystemPassword not implemented")
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
		go protocol.generateMatchmakeSessionSystemPasswordHandler(fmt.Errorf("Failed to read GID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.generateMatchmakeSessionSystemPasswordHandler(nil, client, callID, gid)
}
