// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetRatings sets the ResetRatings handler function
func (protocol *Protocol) ResetRatings(handler func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, transactional bool) uint32) {
	protocol.resetRatingsHandler = handler
}

func (protocol *Protocol) handleResetRatings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.resetRatingsHandler == nil {
		globals.Logger.Warning("DataStore::ResetRatings not implemented")
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
		errorCode = protocol.resetRatingsHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.resetRatingsHandler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), client, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.resetRatingsHandler(nil, client, callID, target.(*datastore_types.DataStoreRatingTarget), transactional)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
