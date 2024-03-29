// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

// PutCommonData sets the PutCommonData handler function
func (protocol *Protocol) PutCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID uint64) uint32) {
	protocol.putCommonDataHandler = handler
}

func (protocol *Protocol) handlePutCommonData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.putCommonDataHandler == nil {
		globals.Logger.Warning("Ranking2::PutCommonData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	commonData, err := parametersStream.ReadStructure(ranking2_types.NewRanking2CommonData())
	if err != nil {
		errorCode = protocol.putCommonDataHandler(fmt.Errorf("Failed to read commonData from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.putCommonDataHandler(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.putCommonDataHandler(nil, packet, callID, commonData.(*ranking2_types.Ranking2CommonData), nexUniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
