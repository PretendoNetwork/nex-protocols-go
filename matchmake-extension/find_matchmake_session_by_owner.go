// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindMatchmakeSessionByOwner sets the FindMatchmakeSessionByOwner handler function
func (protocol *Protocol) FindMatchmakeSessionByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, id uint32, resultRange *nex.ResultRange) uint32) {
	protocol.findMatchmakeSessionByOwnerHandler = handler
}

func (protocol *Protocol) handleFindMatchmakeSessionByOwner(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findMatchmakeSessionByOwnerHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByOwner not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.findMatchmakeSessionByOwnerHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findMatchmakeSessionByOwnerHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findMatchmakeSessionByOwnerHandler(nil, packet, callID, id, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
