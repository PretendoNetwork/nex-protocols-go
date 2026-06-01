// Package protocol implements the Subscription protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleReplaceTargetAndGetSubscriptionData(packet nex.PacketInterface) {
	if protocol.ReplaceTargetAndGetSubscriptionData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SubscriptionProtocol::ReplaceTargetAndGetSubscriptionData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var newTargets types.List[types.PID]

	err := newTargets.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReplaceTargetAndGetSubscriptionData(fmt.Errorf("failed to read newTargets from parameters. %s", err.Error()), packet, callID, newTargets)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReplaceTargetAndGetSubscriptionData(nil, packet, callID, newTargets)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
