// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetLastConnectionStats sets the GetLastConnectionStats handler function
func (protocol *AccountManagementProtocol) GetLastConnectionStats(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32)) {
	protocol.getLastConnectionStatsHandler = handler
}

func (protocol *AccountManagementProtocol) handleGetLastConnectionStats(packet nex.PacketInterface) {
	if protocol.getLastConnectionStatsHandler == nil {
		globals.Logger.Warning("AccountManagement::GetLastConnectionStats not implemented")
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
		go protocol.getLastConnectionStatsHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getLastConnectionStatsHandler(nil, client, callID, idPrincipal)
}
