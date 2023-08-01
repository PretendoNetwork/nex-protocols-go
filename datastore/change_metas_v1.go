// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangeMetasV1 sets the ChangeMetasV1 handler function
func (protocol *Protocol) ChangeMetasV1(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64, params []*datastore_types.DataStoreChangeMetaParamV1, transactional bool)) {
	protocol.changeMetasV1Handler = handler
}

func (protocol *Protocol) handleChangeMetasV1(packet nex.PacketInterface) {
	if protocol.changeMetasV1Handler == nil {
		globals.Logger.Warning("DataStore::ChangeMetasV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.changeMetasV1Handler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreChangeMetaParamV1())
	if err != nil {
		go protocol.changeMetasV1Handler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.changeMetasV1Handler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	go protocol.changeMetasV1Handler(nil, client, callID, dataIDs, params.([]*datastore_types.DataStoreChangeMetaParamV1), transactional)
}