// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPublicData sets the GetPublicData handler function
func (protocol *AccountManagementProtocol) GetPublicData(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32)) {
	protocol.getPublicDataHandler = handler
}

func (protocol *AccountManagementProtocol) handleGetPublicData(packet nex.PacketInterface) {
	if protocol.getPublicDataHandler == nil {
		globals.Logger.Warning("AccountManagement::GetPublicData not implemented")
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
		go protocol.getPublicDataHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getPublicDataHandler(nil, client, callID, idPrincipal)
}
