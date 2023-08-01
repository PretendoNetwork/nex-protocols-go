// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRating sets the GetRating handler function
func (protocol *Protocol) GetRating(handler func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64)) {
	protocol.getRatingHandler = handler
}

func (protocol *Protocol) handleGetRating(packet nex.PacketInterface) {
	if protocol.getRatingHandler == nil {
		globals.Logger.Warning("DataStore::GetRating not implemented")
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
		go protocol.getRatingHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	accessPassword, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getRatingHandler(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.getRatingHandler(nil, client, callID, target.(*datastore_types.DataStoreRatingTarget), accessPassword)
}