// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareAttachFile sets the PrepareAttachFile handler function
func (protocol *Protocol) PrepareAttachFile(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreAttachFileParam) uint32) {
	protocol.prepareAttachFileHandler = handler
}

func (protocol *Protocol) handlePrepareAttachFile(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.prepareAttachFileHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::PrepareAttachFile not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreAttachFileParam())
	if err != nil {
		errorCode = protocol.prepareAttachFileHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.prepareAttachFileHandler(nil, packet, callID, param.(*datastore_super_mario_maker_types.DataStoreAttachFileParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
