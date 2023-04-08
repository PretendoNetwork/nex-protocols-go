package matchmake_extension

import (
	"encoding/hex"
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

// AutoMatchmakeWithSearchCriteria_Postpone sets the AutoMatchmakeWithSearchCriteria_Postpone handler function
func (protocol *MatchmakeExtensionProtocol) AutoMatchmakeWithSearchCriteria_Postpone(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *match_making.MatchmakeSession, message string)) {
	protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleAutoMatchmakeWithSearchCriteria_Postpone(packet nex.PacketInterface) {
	if protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithSearchCriteria_PostponeHandler not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	globals.Logger.Info(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	criteriaCount := int(parametersStream.ReadUInt32LE())
	for i := 0; i < criteriaCount; i++ {
		_, _ = parametersStream.ReadStructure(match_making.NewMatchmakeSessionSearchCriteria())
	}
	dataHolderType, err := parametersStream.ReadString()

	if err != nil {
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	globals.Logger.Info(dataHolderType)

	if dataHolderType != "MatchmakeSession" {
		err := errors.New("[MatchmakeExtension::AutoMatchmakeWithSearchCriteria_Postpone] Data holder name does not match")
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	if (parametersStream.ByteCapacity() - parametersStream.ByteOffset()) < 8 {
		err := errors.New("[MatchmakeExtension::AutoMatchmakeWithSearchCriteria_Postpone] Data holder missing lengths")
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	parametersStream.SeekByte(4, true) // Skip length including next buffer length field
	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, protocol.Server)

	gatheringStructureInterface, err := dataHolderContentStream.ReadStructure(match_making.NewGathering())
	if err != nil {
		globals.Logger.Error(err.Error())
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	matchmakeSessionStructureInterface, err := dataHolderContentStream.ReadStructure(match_making.NewMatchmakeSession())
	if err != nil {
		globals.Logger.Error(err.Error())
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}
	matchmakeSession := matchmakeSessionStructureInterface.(*match_making.MatchmakeSession)
	matchmakeSession.Gathering = gatheringStructureInterface.(*match_making.Gathering)

	message, err := parametersStream.ReadString()
	if err != nil {
		go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	go protocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(nil, client, callID, matchmakeSession, message)
}
