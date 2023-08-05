// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObjects sets the CompletePostObjects handler function
func (protocol *Protocol) CompletePostObjects(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64) uint32) {
	protocol.completePostObjectsHandler = handler
}

func (protocol *Protocol) handleCompletePostObjects(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completePostObjectsHandler == nil {
		globals.Logger.Warning("DataStore::CompletePostObjects not implemented")
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
		errorCode = protocol.completePostObjectsHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completePostObjectsHandler(nil, client, callID, dataIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
