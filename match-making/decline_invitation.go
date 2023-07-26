// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeclineInvitation sets the DeclineInvitation handler function
func (protocol *Protocol) DeclineInvitation(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string)) {
	protocol.declineInvitationHandler = handler
}

func (protocol *Protocol) handleDeclineInvitation(packet nex.PacketInterface) {
	if protocol.declineInvitationHandler == nil {
		globals.Logger.Warning("MatchMaking::DeclineInvitation not implemented")
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
		go protocol.declineInvitationHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, "")
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.declineInvitationHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, "")
	}

	go protocol.declineInvitationHandler(nil, client, callID, idGathering, strMessage)
}
