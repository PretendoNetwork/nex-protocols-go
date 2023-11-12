// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindInvitations sets the FindInvitations handler function
func (protocol *Protocol) FindInvitations(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange *nex.ResultRange) uint32) {
	protocol.findInvitationsHandler = handler
}

func (protocol *Protocol) handleFindInvitations(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findInvitationsHandler == nil {
		globals.Logger.Warning("MatchMaking::FindInvitations not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findInvitationsHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findInvitationsHandler(nil, packet, callID, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
