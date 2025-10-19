// Package protocol implements the Subscription protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetActivePlayerSubscriptionData(packet nex.PacketInterface) {
	if protocol.GetActivePlayerSubscriptionData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SubscriptionProtocol::GetActivePlayerSubscriptionData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var unknown1 types.UInt32
	var unknown2 types.UInt32
	var unknown3 types.UInt32

	var err error

	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetActivePlayerSubscriptionData(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetActivePlayerSubscriptionData(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown3.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetActivePlayerSubscriptionData(fmt.Errorf("Failed to read unknown3 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetActivePlayerSubscriptionData(nil, packet, callID, unknown1, unknown2, unknown3)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
