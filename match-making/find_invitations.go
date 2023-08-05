// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindInvitations sets the FindInvitations handler function
func (protocol *Protocol) FindInvitations(handler func(err error, client *nex.Client, callID uint32, resultRange *nex.ResultRange) uint32) {
	protocol.findInvitationsHandler = handler
}

func (protocol *Protocol) handleFindInvitations(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findInvitationsHandler == nil {
		globals.Logger.Warning("MatchMaking::FindInvitations not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findInvitationsHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findInvitationsHandler(nil, client, callID, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
