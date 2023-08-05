// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CustomSearchObject sets the CustomSearchObject handler function
func (protocol *Protocol) CustomSearchObject(handler func(err error, client *nex.Client, callID uint32, condition uint32, param *datastore_types.DataStoreSearchParam) uint32) {
	protocol.customSearchObjectHandler = handler
}

func (protocol *Protocol) handleCustomSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.customSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CustomSearchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	condition, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.customSearchObjectHandler(fmt.Errorf("Failed to read condition from parameters. %s", err.Error()), client, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		errorCode = protocol.customSearchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.customSearchObjectHandler(nil, client, callID, condition, param.(*datastore_types.DataStoreSearchParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
