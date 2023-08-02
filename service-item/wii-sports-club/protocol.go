// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item "github.com/PretendoNetwork/nex-protocols-go/service-item"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Service Item (Wii Sports Club) protocol
	ProtocolID = 0x77

	// MethodHello is the method ID for the Hello method
	MethodHello = 0x1

	// MethodHTTPGetRequest is the method ID for the HTTPGetRequest method
	MethodHTTPGetRequest = 0x2

	// MethodHTTPGetResponse is the method ID for the HTTPGetResponse method
	MethodHTTPGetResponse = 0x3

	// MethodPurchaseServiceItemRequest is the method ID for the PurchaseServiceItemRequest method
	MethodPurchaseServiceItemRequest = 0x4

	// MethodPurchaseServiceItemResponse is the method ID for the PurchaseServiceItemResponse method
	MethodPurchaseServiceItemResponse = 0x5

	// MethodListServiceItemRequest is the method ID for the ListServiceItemRequest method
	MethodListServiceItemRequest = 0x6

	// MethodListServiceItemResponse is the method ID for the ListServiceItemResponse method
	MethodListServiceItemResponse = 0x7

	// MethodGetBalanceRequest is the method ID for the GetBalanceRequest method
	MethodGetBalanceRequest = 0x8

	// MethodGetBalanceResponse is the method ID for the GetBalanceResponse method
	MethodGetBalanceResponse = 0x9

	// MethodGetPrepurchaseInfoRequest is the method ID for the GetPrepurchaseInfoRequest method
	MethodGetPrepurchaseInfoRequest = 0xA

	// MethodGetPrepurchaseInfoResponse is the method ID for the GetPrepurchaseInfoResponse method
	MethodGetPrepurchaseInfoResponse = 0xB

	// MethodGetServiceItemRightRequest is the method ID for the GetServiceItemRightRequest method
	MethodGetServiceItemRightRequest = 0xC

	// MethodGetServiceItemRightResponse is the method ID for the GetServiceItemRightResponse method
	MethodGetServiceItemRightResponse = 0xD

	// MethodGetPurchaseHistoryRequest is the method ID for the GetPurchaseHistoryRequest method
	MethodGetPurchaseHistoryRequest = 0xE

	// MethodGetPurchaseHistoryResponse is the method ID for the GetPurchaseHistoryResponse method
	MethodGetPurchaseHistoryResponse = 0xF

	// MethodGetNotice is the method ID for the GetNotice method
	MethodGetNotice = 0x10

	// MethodUpdateAndGetTicketInfo is the method ID for the UpdateAndGetTicketInfo method
	MethodUpdateAndGetTicketInfo = 0x11

	// MethodLoadUserInfo is the method ID for the LoadUserInfo method
	MethodLoadUserInfo = 0x12

	// MethodSaveUserInfo is the method ID for the SaveUserInfo method
	MethodSaveUserInfo = 0x13

	// MethodStartChallenge is the method ID for the StartChallenge method
	MethodStartChallenge = 0x14

	// MethodEndChallenge is the method ID for the EndChallenge method
	MethodEndChallenge = 0x15

	// MethodRequestTicketRestoration is the method ID for the RequestTicketRestoration method
	MethodRequestTicketRestoration = 0x16
)

var patchedMethods = []uint32{}

type serviceItemProtocol = service_item.Protocol

// Protocol stores all the RMC method handlers for the Service Item (Wii Sports Club) protocol and listens for requests
// Embeds the Service Item protocol
type Protocol struct {
	Server *nex.Server
	serviceItemProtocol
	helloHandler                       func(err error, client *nex.Client, callID uint32, name string)
	httpGetRequestHandler              func(err error, client *nex.Client, callID uint32, url *service_item_wii_sports_club_types.ServiceItemHTTPGetParam)
	httpGetResponseHandler             func(err error, client *nex.Client, callID uint32, requestID uint32)
	purchaseServiceItemRequestHandler  func(err error, client *nex.Client, callID uint32, purchaseServiceItemParam *service_item_wii_sports_club_types.ServiceItemPurchaseServiceItemParam)
	purchaseServiceItemResponseHandler func(err error, client *nex.Client, callID uint32, requestID uint32)
	listServiceItemRequestHandler      func(err error, client *nex.Client, callID uint32, listServiceItemParam *service_item_wii_sports_club_types.ServiceItemListServiceItemParam)
	listServiceItemResponseHandler     func(err error, client *nex.Client, callID uint32, requestID uint32)
	getBalanceRequestHandler           func(err error, client *nex.Client, callID uint32, getBalanceParam *service_item_wii_sports_club_types.ServiceItemGetBalanceParam)
	getBalanceResponseHandler          func(err error, client *nex.Client, callID uint32, requestID uint32)
	getPrepurchaseInfoRequestHandler   func(err error, client *nex.Client, callID uint32, getPrepurchaseInfoParam *service_item_wii_sports_club_types.ServiceItemGetPrepurchaseInfoParam)
	getPrepurchaseInfoResponseHandler  func(err error, client *nex.Client, callID uint32, requestID uint32)
	getServiceItemRightRequestHandler  func(err error, client *nex.Client, callID uint32, getServiceItemRightParam *service_item_wii_sports_club_types.ServiceItemGetServiceItemRightParam)
	getServiceItemRightResponseHandler func(err error, client *nex.Client, callID uint32, requestID uint32)
	getPurchaseHistoryRequestHandler   func(err error, client *nex.Client, callID uint32, getPurchaseHistoryParam *service_item_wii_sports_club_types.ServiceItemGetPurchaseHistoryParam)
	getPurchaseHistoryResponseHandler  func(err error, client *nex.Client, callID uint32, requestID uint32)
	getNoticeHandler                   func(err error, client *nex.Client, callID uint32, getNoticeParam *service_item_wii_sports_club_types.ServiceItemGetNoticeParam)
	updateAndGetTicketInfoHandler      func(err error, client *nex.Client, callID uint32, forceRetrieveFromEShop bool)
	loadUserInfoHandler                func(err error, client *nex.Client, callID uint32)
	saveUserInfoHandler                func(err error, client *nex.Client, callID uint32, userInfo *service_item_wii_sports_club_types.ServiceItemUserInfo)
	startChallengeHandler              func(err error, client *nex.Client, callID uint32, startChallengeParam *service_item_wii_sports_club_types.ServiceItemStartChallengeParam)
	endChallengeHandler                func(err error, client *nex.Client, callID uint32, endChallengeParam *service_item_wii_sports_club_types.ServiceItemEndChallengeParam)
	requestTicketRestorationHandler    func(err error, client *nex.Client, callID uint32, requestTicketRestorationParam *service_item_wii_sports_club_types.ServiceItemRequestTicketRestorationParam)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.serviceItemProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodHello:
		protocol.handleHello(packet)
	case MethodHTTPGetRequest:
		protocol.handleHTTPGetRequest(packet)
	case MethodHTTPGetResponse:
		protocol.handleHTTPGetResponse(packet)
	case MethodPurchaseServiceItemRequest:
		protocol.handlePurchaseServiceItemRequest(packet)
	case MethodPurchaseServiceItemResponse:
		protocol.handlePurchaseServiceItemResponse(packet)
	case MethodListServiceItemRequest:
		protocol.handleListServiceItemRequest(packet)
	case MethodListServiceItemResponse:
		protocol.handleListServiceItemResponse(packet)
	case MethodGetBalanceRequest:
		protocol.handleGetBalanceRequest(packet)
	case MethodGetBalanceResponse:
		protocol.handleGetBalanceResponse(packet)
	case MethodGetPrepurchaseInfoRequest:
		protocol.handleGetPrepurchaseInfoRequest(packet)
	case MethodGetPrepurchaseInfoResponse:
		protocol.handleGetPrepurchaseInfoResponse(packet)
	case MethodGetServiceItemRightRequest:
		protocol.handleGetServiceItemRightRequest(packet)
	case MethodGetServiceItemRightResponse:
		protocol.handleGetServiceItemRightResponse(packet)
	case MethodGetPurchaseHistoryRequest:
		protocol.handleGetPurchaseHistoryRequest(packet)
	case MethodGetPurchaseHistoryResponse:
		protocol.handleGetPurchaseHistoryResponse(packet)
	case MethodGetNotice:
		protocol.handleGetNotice(packet)
	case MethodUpdateAndGetTicketInfo:
		protocol.handleUpdateAndGetTicketInfo(packet)
	case MethodLoadUserInfo:
		protocol.handleLoadUserInfo(packet)
	case MethodSaveUserInfo:
		protocol.handleSaveUserInfo(packet)
	case MethodStartChallenge:
		protocol.handleStartChallenge(packet)
	case MethodEndChallenge:
		protocol.handleEndChallenge(packet)
	case MethodRequestTicketRestoration:
		protocol.handleRequestTicketRestoration(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Service Item (Wii Sports Club) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new ServiceItemWiiSportsClub protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.serviceItemProtocol.Server = server

	protocol.Setup()

	return protocol
}
