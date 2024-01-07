// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
)

func (protocol *Protocol) handleGetContentMulti(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetContentMulti == nil {
		globals.Logger.Warning("Subscriber::GetContentMulti not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	params, err := nex.StreamReadListStructure(parametersStream, subscriber_types.NewSubscriberGetContentParam())
	if err != nil {
		_, errorCode = protocol.GetContentMulti(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetContentMulti(nil, packet, callID, params)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
