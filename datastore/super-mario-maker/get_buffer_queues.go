// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBufferQueues sets the GetBufferQueues handler function
func (protocol *Protocol) GetBufferQueues(handler func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam) uint32) {
	protocol.getBufferQueuesHandler = handler
}

func (protocol *Protocol) handleGetBufferQueues(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getBufferQueuesHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetBufferQueues not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewBufferQueueParam())
	if err != nil {
		errorCode = protocol.getBufferQueuesHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getBufferQueuesHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.BufferQueueParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
