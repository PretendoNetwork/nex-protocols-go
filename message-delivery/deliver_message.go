package message_delivery

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeliverMessage sets the DeliverMessage handler function
func (protocol *MessageDeliveryProtocol) DeliverMessage(handler func(err error, client *nex.Client, callID uint32, oUserMessage nex.StructureInterface)) {
	protocol.DeliverMessageHandler = handler
}

func (protocol *MessageDeliveryProtocol) HandleDeliverMessage(packet nex.PacketInterface) {
	if protocol.DeliverMessageHandler == nil {
		globals.Logger.Warning("MessageDelivery::DeliverMessage not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataHolderName, err := parametersStream.ReadString()

	if err != nil {
		go protocol.DeliverMessageHandler(err, client, callID, nil)
		return
	}

	_ = parametersStream.ReadUInt32LE() // length including this field

	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go protocol.DeliverMessageHandler(err, client, callID, nil)
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, protocol.Server)

	var oUserMessage nex.StructureInterface

	if dataHolderName == "BinaryMessage" {
		oUserMessage, _ = dataHolderContentStream.ReadStructure(NewBinaryMessage())
	}

	go protocol.DeliverMessageHandler(nil, client, callID, oUserMessage)
}
