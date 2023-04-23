package friends_wiiu

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateComment sets the UpdateComment handler function
func (protocol *FriendsWiiUProtocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment *Comment)) {
	protocol.UpdateCommentHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleUpdateComment(packet nex.PacketInterface) {
	if protocol.UpdateCommentHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateComment not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	commentStructureInterface, err := parametersStream.ReadStructure(NewComment())
	if err != nil {
		go protocol.UpdateCommentHandler(err, client, callID, nil)
		return
	}

	comment := commentStructureInterface.(*Comment)

	go protocol.UpdateCommentHandler(nil, client, callID, comment)
}