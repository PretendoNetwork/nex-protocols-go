// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadCommonData sets the UploadCommonData handler function
func (protocol *Protocol) UploadCommonData(handler func(err error, client *nex.Client, callID uint32, commonData []byte, uniqueID uint64) uint32) {
	protocol.uploadCommonDataHandler = handler
}

func (protocol *Protocol) handleUploadCommonData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.uploadCommonDataHandler == nil {
		globals.Logger.Warning("Ranking::UploadCommonData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	commonData, err := parametersStream.ReadBuffer()
	if err != nil {
		errorCode = protocol.uploadCommonDataHandler(fmt.Errorf("Failed to read commonData from parameters. %s", err.Error()), client, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.uploadCommonDataHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.uploadCommonDataHandler(nil, client, callID, commonData, uniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
