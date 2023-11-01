// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByParticipants sets the FindByParticipants handler function
func (protocol *Protocol) FindByParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, pid []uint32) uint32) {
	protocol.findByParticipantsHandler = handler
}

func (protocol *Protocol) handleFindByParticipants(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByParticipantsHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByParticipants not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.findByParticipantsHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByParticipantsHandler(nil, packet, callID, pid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
