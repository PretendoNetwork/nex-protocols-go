package friends_wiiu

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAndGetAllInformation sets the UpdateAndGetAllInformation handler function
func (protocol *FriendsWiiUProtocol) UpdateAndGetAllInformation(handler func(err error, client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)) {
	protocol.UpdateAndGetAllInformationHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	if protocol.UpdateAndGetAllInformationHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateAndGetAllInformation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nnaInfoStructureInterface, err := parametersStream.ReadStructure(NewNNAInfo())
	if err != nil {
		go protocol.UpdateAndGetAllInformationHandler(err, client, callID, nil, nil, nil)
		return
	}

	presenceStructureInterface, err := parametersStream.ReadStructure(NewNintendoPresenceV2())
	if err != nil {
		go protocol.UpdateAndGetAllInformationHandler(err, client, callID, nil, nil, nil)
		return
	}

	nnaInfo := nnaInfoStructureInterface.(*NNAInfo)
	presence := presenceStructureInterface.(*NintendoPresenceV2)
	birthday := nex.NewDateTime(parametersStream.ReadUInt64LE())

	go protocol.UpdateAndGetAllInformationHandler(nil, client, callID, nnaInfo, presence, birthday)
}
