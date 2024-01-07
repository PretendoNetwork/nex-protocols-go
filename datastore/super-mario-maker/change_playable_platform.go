// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleChangePlayablePlatform(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ChangePlayablePlatform == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ChangePlayablePlatform not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	params, err := nex.StreamReadListStructure(parametersStream, datastore_super_mario_maker_types.NewDataStoreChangePlayablePlatformParam())
	if err != nil {
		_, errorCode = protocol.ChangePlayablePlatform(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ChangePlayablePlatform(nil, packet, callID, params)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
