// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SendInvitation sets the SendInvitation handler function
func (protocol *Protocol) SendInvitation(handler func(err error, client *nex.Client, callID uint32, pids []uint32)) {
	protocol.sendInvitationHandler = handler
}

func (protocol *Protocol) handleSendInvitation(packet nex.PacketInterface) {
	if protocol.sendInvitationHandler == nil {
		globals.Logger.Warning("Friends3DS::SendInvitation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.sendInvitationHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.sendInvitationHandler(nil, client, callID, pids)
}
