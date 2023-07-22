// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindInvitations sets the FindInvitations handler function
func (protocol *MatchMakingProtocol) FindInvitations(handler func(err error, client *nex.Client, callID uint32, resultRange *nex.ResultRange)) {
	protocol.findInvitationsHandler = handler
}

func (protocol *MatchMakingProtocol) handleFindInvitations(packet nex.PacketInterface) {
	if protocol.findInvitationsHandler == nil {
		globals.Logger.Warning("MatchMaking::FindInvitations not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findInvitationsHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.findInvitationsHandler(nil, client, callID, resultRange.(*nex.ResultRange))
}
