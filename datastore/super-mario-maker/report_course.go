// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportCourse sets the ReportCourse handler function
func (protocol *Protocol) ReportCourse(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreReportCourseParam) uint32) {
	protocol.reportCourseHandler = handler
}

func (protocol *Protocol) handleReportCourse(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportCourseHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ReportCourse not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreReportCourseParam())
	if err != nil {
		errorCode = protocol.reportCourseHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.reportCourseHandler(nil, packet, callID, param.(*datastore_super_mario_maker_types.DataStoreReportCourseParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
