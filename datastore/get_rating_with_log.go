// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRatingWithLog sets the GetRatingWithLog handler function
func (protocol *Protocol) GetRatingWithLog(handler func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64)) {
	protocol.getRatingWithLogHandler = handler
}

func (protocol *Protocol) handleGetRatingWithLog(packet nex.PacketInterface) {
	if protocol.getRatingWithLogHandler == nil {
		globals.Logger.Warning("DataStore::GetRatingWithLog not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		go protocol.getRatingWithLogHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	accessPassword, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getRatingWithLogHandler(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.getRatingWithLogHandler(nil, client, callID, target.(*datastore_types.DataStoreRatingTarget), accessPassword)
}