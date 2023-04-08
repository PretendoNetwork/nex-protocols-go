package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObjects sets the CompletePostObjects handler function
func (protocol *DataStoreProtocol) CompletePostObjects(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64)) {
	protocol.CompletePostObjectsHandler = handler
}

func (protocol *DataStoreProtocol) HandleCompletePostObjects(packet nex.PacketInterface) {
	if protocol.CompletePostObjectsHandler == nil {
		globals.Logger.Warning("DataStore::CompletePostObjects not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs := parametersStream.ReadListUInt64LE()

	go protocol.CompletePostObjectsHandler(nil, client, callID, dataIDs)
}
