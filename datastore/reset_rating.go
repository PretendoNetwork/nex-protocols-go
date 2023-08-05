// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetRating sets the ResetRating handler function
func (protocol *Protocol) ResetRating(handler func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64) uint32) {
	protocol.resetRatingHandler = handler
}

func (protocol *Protocol) handleResetRating(packet nex.PacketInterface) {
	if protocol.resetRatingHandler == nil {
		globals.Logger.Warning("DataStore::ResetRating not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		go protocol.resetRatingHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	accessPassword, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.resetRatingHandler(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.resetRatingHandler(nil, client, callID, target.(*datastore_types.DataStoreRatingTarget), accessPassword)
}
