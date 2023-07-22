// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcceptInvitation sets the AcceptInvitation handler function
func (protocol *MatchMakingProtocol) AcceptInvitation(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string)) {
	protocol.acceptInvitationHandler = handler
}

func (protocol *MatchMakingProtocol) handleAcceptInvitation(packet nex.PacketInterface) {
	if protocol.acceptInvitationHandler == nil {
		globals.Logger.Warning("MatchMaking::AcceptInvitation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.acceptInvitationHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, "")
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.acceptInvitationHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, "")
	}

	go protocol.acceptInvitationHandler(nil, client, callID, idGathering, strMessage)
}
