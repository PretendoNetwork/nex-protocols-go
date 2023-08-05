// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMultiplePublicData sets the GetMultiplePublicData handler function
func (protocol *Protocol) GetMultiplePublicData(handler func(err error, client *nex.Client, callID uint32, lstPrincipals []uint32) uint32) {
	protocol.getMultiplePublicDataHandler = handler
}

func (protocol *Protocol) handleGetMultiplePublicData(packet nex.PacketInterface) {
	if protocol.getMultiplePublicDataHandler == nil {
		globals.Logger.Warning("AccountManagement::GetMultiplePublicData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPrincipals, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getMultiplePublicDataHandler(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getMultiplePublicDataHandler(nil, client, callID, lstPrincipals)
}
