// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRatings sets the GetRatings handler function
func (protocol *Protocol) GetRatings(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64, accessPassword uint64) uint32) {
	protocol.getRatingsHandler = handler
}

func (protocol *Protocol) handleGetRatings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRatingsHandler == nil {
		globals.Logger.Warning("DataStore::GetRatings not implemented")
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
		errorCode = protocol.getRatingsHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	accessPassword, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getRatingsHandler(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), client, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRatingsHandler(nil, client, callID, dataIDs, accessPassword)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
