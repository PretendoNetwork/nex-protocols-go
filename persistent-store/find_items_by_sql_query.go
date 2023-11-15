// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindItemsBySQLQuery(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.FindItemsBySQLQuery == nil {
		globals.Logger.Warning("PersistentStore::FindItemsBySQLQuery not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroup, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.FindItemsBySQLQuery(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, 0, "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.FindItemsBySQLQuery(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, 0, "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strQuery, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.FindItemsBySQLQuery(fmt.Errorf("Failed to read strQuery from parameters. %s", err.Error()), packet, callID, 0, "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindItemsBySQLQuery(nil, packet, callID, uiGroup, strTag, strQuery)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
