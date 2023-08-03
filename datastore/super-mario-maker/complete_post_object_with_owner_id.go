// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObjectWithOwnerID sets the CompletePostObjectWithOwnerID handler function
func (protocol *Protocol) CompletePostObjectWithOwnerID(handler func(err error, client *nex.Client, callID uint32, ownerID uint32, param *datastore_types.DataStoreCompletePostParam)) {
	protocol.completePostObjectWithOwnerIDHandler = handler
}

func (protocol *Protocol) handleCompletePostObjectWithOwnerID(packet nex.PacketInterface) {
	if protocol.completePostObjectWithOwnerIDHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CompletePostObjectWithOwnerID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ownerID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.completePostObjectWithOwnerIDHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		go protocol.completePostObjectWithOwnerIDHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.completePostObjectWithOwnerIDHandler(nil, client, callID, ownerID, param.(*datastore_types.DataStoreCompletePostParam))
}
