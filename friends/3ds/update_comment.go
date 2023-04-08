package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateComment sets the UpdateComment handler function
func (protocol *Friends3DSProtocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment string)) {
	protocol.UpdateCommentHandler = handler
}

func (protocol *Friends3DSProtocol) HandleUpdateComment(packet nex.PacketInterface) {
	if protocol.UpdateCommentHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateComment not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	comment, _ := parametersStream.ReadString()

	go protocol.UpdateCommentHandler(nil, client, callID, comment)
}
