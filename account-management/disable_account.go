// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DisableAccount sets the DisableAccount handler function
func (protocol *Protocol) DisableAccount(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32, dtUntil *nex.DateTime, strMessage string)) {
	protocol.disableAccountHandler = handler
}

func (protocol *Protocol) handleDisableAccount(packet nex.PacketInterface) {
	if protocol.disableAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::DisableAccount not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idPrincipal, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.disableAccountHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	dtUntil, err := parametersStream.ReadDateTime()
	if err != nil {
		go protocol.disableAccountHandler(fmt.Errorf("Failed to read dtUntil from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.disableAccountHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, nil, "")
		return
	}

	go protocol.disableAccountHandler(nil, client, callID, idPrincipal, dtUntil, strMessage)
}
