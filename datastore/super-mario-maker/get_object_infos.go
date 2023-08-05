// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetObjectInfos sets the GetObjectInfos handler function
func (protocol *Protocol) GetObjectInfos(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64) uint32) {
	protocol.getObjectInfosHandler = handler
}

func (protocol *Protocol) handleGetObjectInfos(packet nex.PacketInterface) {
	if protocol.getObjectInfosHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetObjectInfos not implemented")
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
		go protocol.getObjectInfosHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getObjectInfosHandler(nil, client, callID, dataIDs)
}
