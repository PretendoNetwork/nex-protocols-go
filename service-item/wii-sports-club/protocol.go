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

var patchedMethods = []uint32{
	MethodHello,
	MethodHTTPGetRequest,
	MethodHTTPGetResponse,
	MethodPurchaseServiceItemRequest,
	MethodPurchaseServiceItemResponse,
	MethodListServiceItemRequest,
	MethodListServiceItemResponse,
	MethodGetBalanceRequest,
	MethodGetBalanceResponse,
	MethodGetPrepurchaseInfoRequest,
	MethodGetPrepurchaseInfoResponse,
	MethodGetServiceItemRightRequest,
	MethodGetServiceItemRightResponse,
	MethodGetPurchaseHistoryRequest,
	MethodGetPurchaseHistoryResponse,
	MethodGetNotice,
	MethodUpdateAndGetTicketInfo,
	MethodLoadUserInfo,
	MethodSaveUserInfo,
	MethodStartChallenge,
	MethodEndChallenge,
	MethodRequestTicketRestoration,
}

type serviceItemProtocol = service_item.Protocol

// Protocol stores all the RMC method handlers for the Service Item (Wii Sports Club) protocol and listens for requests
// Embeds the Service Item protocol
type Protocol struct {
	Server nex.ServerInterface
	serviceItemProtocol
	Hello                       func(err error, packet nex.PacketInterface, callID uint32, name string) uint32
	HttpGetRequest              func(err error, packet nex.PacketInterface, callID uint32, url *service_item_wii_sports_club_types.ServiceItemHTTPGetParam) uint32
	HttpGetResponse             func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	PurchaseServiceItemRequest  func(err error, packet nex.PacketInterface, callID uint32, purchaseServiceItemParam *service_item_wii_sports_club_types.ServiceItemPurchaseServiceItemParam) uint32
	PurchaseServiceItemResponse func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	ListServiceItemRequest      func(err error, packet nex.PacketInterface, callID uint32, listServiceItemParam *service_item_wii_sports_club_types.ServiceItemListServiceItemParam) uint32
	ListServiceItemResponse     func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	GetBalanceRequest           func(err error, packet nex.PacketInterface, callID uint32, getBalanceParam *service_item_wii_sports_club_types.ServiceItemGetBalanceParam) uint32
	GetBalanceResponse          func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	GetPrepurchaseInfoRequest   func(err error, packet nex.PacketInterface, callID uint32, getPrepurchaseInfoParam *service_item_wii_sports_club_types.ServiceItemGetPrepurchaseInfoParam) uint32
	GetPrepurchaseInfoResponse  func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	GetServiceItemRightRequest  func(err error, packet nex.PacketInterface, callID uint32, getServiceItemRightParam *service_item_wii_sports_club_types.ServiceItemGetServiceItemRightParam) uint32
	GetServiceItemRightResponse func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	GetPurchaseHistoryRequest   func(err error, packet nex.PacketInterface, callID uint32, getPurchaseHistoryParam *service_item_wii_sports_club_types.ServiceItemGetPurchaseHistoryParam) uint32
	GetPurchaseHistoryResponse  func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	GetNotice                   func(err error, packet nex.PacketInterface, callID uint32, getNoticeParam *service_item_wii_sports_club_types.ServiceItemGetNoticeParam) uint32
	UpdateAndGetTicketInfo      func(err error, packet nex.PacketInterface, callID uint32, forceRetrieveFromEShop bool) uint32
	LoadUserInfo                func(err error, packet nex.PacketInterface, callID uint32) uint32
	SaveUserInfo                func(err error, packet nex.PacketInterface, callID uint32, userInfo *service_item_wii_sports_club_types.ServiceItemUserInfo) uint32
	StartChallenge              func(err error, packet nex.PacketInterface, callID uint32, startChallengeParam *service_item_wii_sports_club_types.ServiceItemStartChallengeParam) uint32
	EndChallenge                func(err error, packet nex.PacketInterface, callID uint32, endChallengeParam *service_item_wii_sports_club_types.ServiceItemEndChallengeParam) uint32
	RequestTicketRestoration    func(err error, packet nex.PacketInterface, callID uint32, requestTicketRestorationParam *service_item_wii_sports_club_types.ServiceItemRequestTicketRestorationParam) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID) {
				protocol.HandlePacket(packet)
			} else {
				protocol.serviceItemProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodHello:
		go protocol.handleHello(packet)
	case MethodHTTPGetRequest:
		go protocol.handleHTTPGetRequest(packet)
	case MethodHTTPGetResponse:
		go protocol.handleHTTPGetResponse(packet)
	case MethodPurchaseServiceItemRequest:
		go protocol.handlePurchaseServiceItemRequest(packet)
	case MethodPurchaseServiceItemResponse:
		go protocol.handlePurchaseServiceItemResponse(packet)
	case MethodListServiceItemRequest:
		go protocol.handleListServiceItemRequest(packet)
	case MethodListServiceItemResponse:
		go protocol.handleListServiceItemResponse(packet)
	case MethodGetBalanceRequest:
		go protocol.handleGetBalanceRequest(packet)
	case MethodGetBalanceResponse:
		go protocol.handleGetBalanceResponse(packet)
	case MethodGetPrepurchaseInfoRequest:
		go protocol.handleGetPrepurchaseInfoRequest(packet)
	case MethodGetPrepurchaseInfoResponse:
		go protocol.handleGetPrepurchaseInfoResponse(packet)
	case MethodGetServiceItemRightRequest:
		go protocol.handleGetServiceItemRightRequest(packet)
	case MethodGetServiceItemRightResponse:
		go protocol.handleGetServiceItemRightResponse(packet)
	case MethodGetPurchaseHistoryRequest:
		go protocol.handleGetPurchaseHistoryRequest(packet)
	case MethodGetPurchaseHistoryResponse:
		go protocol.handleGetPurchaseHistoryResponse(packet)
	case MethodGetNotice:
		go protocol.handleGetNotice(packet)
	case MethodUpdateAndGetTicketInfo:
		go protocol.handleUpdateAndGetTicketInfo(packet)
	case MethodLoadUserInfo:
		go protocol.handleLoadUserInfo(packet)
	case MethodSaveUserInfo:
		go protocol.handleSaveUserInfo(packet)
	case MethodStartChallenge:
		go protocol.handleStartChallenge(packet)
	case MethodEndChallenge:
		go protocol.handleEndChallenge(packet)
	case MethodRequestTicketRestoration:
		go protocol.handleRequestTicketRestoration(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Service Item (Wii Sports Club) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new ServiceItemWiiSportsClub protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.serviceItemProtocol.Server = server

	protocol.Setup()

	return protocol
}
