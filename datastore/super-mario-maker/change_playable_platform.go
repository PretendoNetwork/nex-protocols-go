// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangePlayablePlatform sets the ChangePlayablePlatform handler function
func (protocol *Protocol) ChangePlayablePlatform(handler func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.DataStoreChangePlayablePlatformParam) uint32) {
	protocol.changePlayablePlatformHandler = handler
}

func (protocol *Protocol) handleChangePlayablePlatform(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.changePlayablePlatformHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ChangePlayablePlatform not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewDataStoreChangePlayablePlatformParam())
	if err != nil {
		errorCode = protocol.changePlayablePlatformHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.changePlayablePlatformHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.DataStoreChangePlayablePlatformParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
