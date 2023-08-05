// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPasswordInfos sets the GetPasswordInfos handler function
func (protocol *Protocol) GetPasswordInfos(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64) uint32) {
	protocol.getPasswordInfosHandler = handler
}

func (protocol *Protocol) handleGetPasswordInfos(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPasswordInfosHandler == nil {
		globals.Logger.Warning("DataStore::GetPasswordInfos not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.getPasswordInfosHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPasswordInfosHandler(nil, client, callID, dataIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
