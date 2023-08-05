// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCourseRecord sets the GetCourseRecord handler function
func (protocol *Protocol) GetCourseRecord(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam) uint32) {
	protocol.getCourseRecordHandler = handler
}

func (protocol *Protocol) handleGetCourseRecord(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCourseRecordHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetCourseRecord not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreGetCourseRecordParam())
	if err != nil {
		errorCode = protocol.getCourseRecordHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getCourseRecordHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
