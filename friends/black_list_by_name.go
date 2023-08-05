// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// BlackListByName sets the BlackListByName handler function
func (protocol *Protocol) BlackListByName(handler func(err error, client *nex.Client, callID uint32, strPlayerName string, uiDetails uint32) uint32) {
	protocol.blackListByNameHandler = handler
}

func (protocol *Protocol) handleBlackListByName(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.blackListByNameHandler == nil {
		globals.Logger.Warning("Friends::BlackListByName not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPlayerName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.blackListByNameHandler(fmt.Errorf("Failed to read strPlayerName from parameters. %s", err.Error()), client, callID, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.blackListByNameHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), client, callID, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.blackListByNameHandler(nil, client, callID, strPlayerName, uiDetails)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
