// Package protocol implements the Subscription protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	subscription_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscription/types"
)

func (protocol *Protocol) handleCreateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.CreateMySubscriptionData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SubscriptionProtocol::CreateMySubscriptionData not implemented")

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
	var param subscription_types.SubscriptionData
	var unknown2 types.Bool

	var err error

	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateMySubscriptionData(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, unknown1, param, unknown2)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateMySubscriptionData(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, unknown1, param, unknown2)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateMySubscriptionData(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, unknown1, param, unknown2)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CreateMySubscriptionData(nil, packet, callID, unknown1, param, unknown2)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
