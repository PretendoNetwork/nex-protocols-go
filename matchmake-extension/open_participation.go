// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// OpenParticipation sets the OpenParticipation handler function
func (protocol *Protocol) OpenParticipation(handler func(err error, client *nex.Client, callID uint32, gid uint32) uint32) {
	protocol.openParticipationHandler = handler
}

func (protocol *Protocol) handleOpenParticipation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.openParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::OpenParticipation not implemented")
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
		errorCode = protocol.openParticipationHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.openParticipationHandler(nil, client, callID, gid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
