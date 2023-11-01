// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetObjectInfos sets the GetObjectInfos handler function
func (protocol *Protocol) GetObjectInfos(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs uint64) uint32) {
	protocol.getObjectInfosHandler = handler
}

func (protocol *Protocol) handleGetObjectInfos(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getObjectInfosHandler == nil {
		globals.Logger.Warning("DataStore::GetObjectInfos not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getObjectInfosHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getObjectInfosHandler(nil, packet, callID, dataIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
