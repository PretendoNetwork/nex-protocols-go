// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateComment sets the UpdateComment handler function
func (protocol *Friends3DSProtocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment string)) {
	protocol.updateCommentHandler = handler
}

func (protocol *Friends3DSProtocol) handleUpdateComment(packet nex.PacketInterface) {
	if protocol.updateCommentHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateComment not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	comment, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateCommentHandler(fmt.Errorf("Failed to read comment from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.updateCommentHandler(nil, client, callID, comment)
}
