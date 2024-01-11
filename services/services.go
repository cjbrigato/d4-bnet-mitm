package services

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type svc_method struct {
	method_id    uint32
	method_name  string
	request_msg  string
	response_msg string
}

func HashServiceName(name string) string {
	var hash uint32 = 0x811C9DC5
	for i := 0; i < len(name); i++ {
		hash ^= uint32(name[i])
		hash *= 0x1000193
	}
	return fmt.Sprintf("0x%08X", hash)
}

func RPCCallKind(service_id uint32) string {
	switch service_id {
	case 0:
		return "request"
	case 254:
		return "response"
	}
	return "unknown"
}

func ServiceMsg(service string, method_name string, service_id uint32, message_bytes []byte) (*protoreflect.ProtoMessage, error) {
	messageName := protoreflect.FullName(PbMessageStr(service, method_name, service_id))
	pbtype, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(messageName))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	msg := pbtype.New().Interface()
	err = proto.Unmarshal(message_bytes, msg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return (&msg), nil
}

func PbMessageint(service string, method_id uint32, service_id uint32) string {
	methods := svc_pool[service]
	for _, method := range methods {
		if method.method_id != method_id {
			continue
		}
		switch service_id {
		case 0:
			return method.request_msg
		case 254:
			return method.response_msg
		}
	}
	return ""
}

func PbMessageStr(service string, method_name string, service_id uint32) string {
	methods := svc_pool[service]
	for _, method := range methods {
		if method.method_name != method_name {
			continue
		}
		switch service_id {
		case 0:
			return method.request_msg
		case 254:
			return method.response_msg
		}
	}
	return ""
}

func Test_protos() {
	for _, j := range svc_pool {
		for _, l := range j {

			_, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(l.request_msg))
			if err != nil {
				fmt.Printf("%s proto not found !\n", l.request_msg)
			}
		}
	}
}

var ServiceHashes = map[int]string{
	/*
	   service: bnet.protocol.connection.ConnectionService, hash: 1698982289
	   service: bnet.protocol.notification.NotificationService, hash: 213793859
	   service: bnet.protocol.notification.NotificationListener, hash: 3788189352
	   service: bnet.protocol.friends.FriendsService, hash: 2749215165
	   service: bnet.protocol.friends.FriendsNotify, hash: 1864735251
	   service: bnet.protocol.presence.PresenceService, hash: 4194801407
	   service: bnet.protocol.channel_invitation.ChannelInvitationService, hash: 2198078984
	   service: bnet.protocol.channel_invitation.ChannelInvitationNotify, hash: 4035247136
	   service: bnet.protocol.channel.Channel, hash: 3073563442
	   service: bnet.protocol.channel.ChannelSubscriber, hash: 3213656212
	   service: bnet.protocol.channel.ChannelOwner, hash: 101490829
	   service: bnet.protocol.game_utilities.GameUtilities, hash: 1069623117
	   service: bnet.protocol.game_master.GameMaster, hash: 2165092757
	   service: bnet.protocol.game_master.GameMasterSubscriber, hash: 1467132723
	   service: bnet.protocol.game_master.GameFactorySubscriber, hash: 3338259653
	   service: bnet.protocol.challenge.ChallengeService, hash: 3686756121
	   service: bnet.protocol.challenge.ChallengeNotify, hash: 3151632159
	   service: bnet.protocol.account.AccountService, hash: 1658456209
	   service: bnet.protocol.account.AccountNotify, hash: 1423956503
	   service: bnet.protocol.authentication.AuthenticationServer, hash: 233634817
	   service: bnet.protocol.authentication.AuthenticationClient, hash: 1898188341
	   service: bnet.protocol.resources.Resources, hash: 3971904954
	*/

	// From get_services_hash.py
	0x65446991: "bgs.protocol.connection.v1.ConnectionService",
	0xf4729709: "bgs.protocol.presence.v2.PresenceListener",
	0x138d200c: "bgs.protocol.presence.v2.client.PresenceService",
	0x7e525e99: "bgs.protocol.channel.v2.membership.ChannelMembershipService",
	0x18007be:  "bgs.protocol.channel.v2.membership.ChannelMembershipListener",
	0x798d39d1: "bgs.protocol.channel.v2.ChannelService",
	0x80909d73: "bgs.protocol.club.v1.ClubListener",
	0xe273de0e: "bgs.protocol.club.v1.ClubService",
	0x94b94786: "bgs.protocol.club.v1.membership.ClubMembershipService",
	0x7b37d770: "bgs.protocol.session.v2.client.SessionService",
	0xbca72135: "bgs.protocol.session.v2.client.SessionListener",
	0x2aaadbc:  "bgs.protocol.publication.v2.client.PublicationService",
	0x2362becd: "bgs.protocol.notification.v2.client.NotificationListener",
	0x9e8758d9: "bgs.protocol.publication.v2.client.PublicationListener",
	0xdbc01a89: "bgs.protocol.friends.v2.client.FriendsListener",
	0x5869be8c: "bgs.protocol.friends.v2.client.FriendsService",
	0xc12828f9: "bgs.protocol.whisper.v1.WhisperService",
	0x3fe5849e: "bgs.protocol.whisper.v1.WhisperListener",
	0x8e8f5fb0: "bgs.protocol.block_list.v1.client.BlockListService",
	0xb5dd8a75: "bgs.protocol.block_list.v1.client.BlockListListener",
	0x9da8116b: "bgs.protocol.authentication.v2.client.AuthenticationListener",
	0xc02f8216: "bgs.protocol.authentication.v2.client.AuthenticationService",
	0xdecfc01:  "bgs.protocol.authentication.v1.AuthenticationService",
	0x71240e35: "bgs.protocol.authentication.v1.AuthenticationListener",
	0xbbda171f: "bgs.protocol.challenge.v1.ChallengeListener",

	// manually recovered
	0xECBE75BA: "bgs.protocol.resources.v1.ResourcesService",
	0x5dbb51c2: "bgs.protocol.game_utilities.v2.client.GameUtilitiesService",
	0x62da0891: "bgs.protocol.account.v1.AccountService",
	0x1ae52686: "bgs.protocol.channel.v2.ChannelListener",
}

type ServiceMethods[T uint32 | uint16] map[T]string
type ServiceMethodsIds map[uint16]string
type Service[T string | uint32] struct {
	id        T
	methods   ServiceMethods[uint32]
	methods16 ServiceMethods[uint16]
}
type NamedService Service[string]
type HashedService Service[uint32]
type UnionService struct {
	hashed *Service[uint32]
	named  *Service[string]
}

func (u *UnionService) Name() string {
	return u.named.id
}

func (u *UnionService) Hash() uint32 {
	return u.hashed.id
}

func (u *UnionService) Method(id uint16) (val string, ok bool) {
	val, ok = u.named.methods16[id]
	return
}

/*func (u *UnionService) Methods() *ServiceMethods {
	var m = &ServiceMethods[uint32]{}
	*m = u.named.methods
	return m
}*/

func newUnionService[T string | uint32](s T) (UnionService, bool) {
	var u UnionService
	var i interface{} = s
	var ok bool
	switch i.(type) {
	case uint32:
		u.named, ok = findByHash(i.(uint32))
		u.hashed, _ = findByName(u.named.id)
	case string:
		u.hashed, ok = findByName(i.(string))
		u.named, _ = findByHash(u.hashed.id)
	}
	return u, ok
}

func findByName(id string) (val *Service[uint32], ok bool) {
	var s = &Service[uint32]{}
	*s, ok = names[id]
	val = s
	return
}

func findByHash(id uint32) (val *Service[string], ok bool) {
	var s = &Service[string]{}
	*s, ok = hashes[id]
	val = s
	return
}

func Get[T string | uint32](id T) (val UnionService, ok bool) {
	val, ok = newUnionService(id)
	return
}

func init() {
	for key, val := range hashes {
		methods_16 := ServiceMethods[uint16]{}
		for yek, lav := range val.methods {
			methods_16[uint16(yek)] = lav
			hashes[key].methods16[uint16(yek)] = lav
		}
		names[val.id] = Service[uint32]{key, val.methods, val.methods16}
	}
}

var names = map[string]Service[uint32]{}
var hashes = map[uint32]Service[string]{
	0x65446991: {
		"bgs.protocol.connection.v1.ConnectionService",
		ServiceMethods[uint32]{
			0x80000001: "Connect",
			0x80000003: "Echo",
			0x80000004: "ForceDisconnect",
			0x80000005: "KeepAlive",
			0x80000007: "RequestDisconnect",
		}, ServiceMethods[uint16]{},
	},
	0xf4729709: {
		"bgs.protocol.presence.v2.PresenceListener",
		ServiceMethods[uint32]{
			0xc0000001: "OnPresenceStateUpdated",
			0xc0000002: "OnSubscribeFailure",
		}, ServiceMethods[uint16]{},
	},
	0x138d200c: {
		"bgs.protocol.presence.v2.client.PresenceService",
		ServiceMethods[uint32]{
			0x40000003: "BatchSubscribe",
			0x40000004: "BatchUnsubscribe",
			0x40000005: "Query",
			0x40000006: "Update",
		}, ServiceMethods[uint16]{},
	},
	0xe273de0e: {
		"bgs.protocol.club.v1.ClubService",
		ServiceMethods[uint32]{
			0x40000001: "Subscribe",
			0x40000002: "Unsubscribe",
			0x40000003: "Create",
			0x40000004: "Destroy",
			0x40000005: "GetDescription",
			0x40000006: "GetClubType",
			0x40000007: "UpdateClubState",
			0x40000008: "UpdateClubSettings",
			0x4000001d: "AddMember",
			0x4000001e: "Join",
			0x4000001f: "Leave",
			0x40000020: "Kick",
			0x40000021: "GetMember",
			0x40000022: "GetMembers",
			0x40000023: "UpdateMemberState",
			0x40000024: "UpdateSubscriberState",
			0x40000025: "AssignRole",
			0x40000026: "UnassignRole",
			0x40000032: "SendInvitation",
			0x40000033: "AcceptInvitation",
			0x40000034: "DeclineInvitation",
			0x40000035: "RevokeInvitation",
			0x40000036: "GetInvitation",
			0x40000037: "GetInvitations",
			0x4000003c: "SendSuggestion",
			0x4000003d: "AcceptSuggestion",
			0x4000003e: "DeclineSuggestion",
			0x4000003f: "GetSuggestion",
			0x40000040: "GetSuggestions",
			0x40000046: "CreateTicket",
			0x40000047: "DestroyTicket",
			0x40000048: "RedeemTicket",
			0x40000049: "GetTicket",
			0x4000004a: "GetTickets",
			0x40000050: "AddBan",
			0x40000051: "RemoveBan",
			0x40000052: "GetBan",
			0x40000053: "GetBans",
			0x40000064: "SubscribeStream",
			0x40000065: "UnsubscribeStream",
			0x40000066: "CreateStream",
			0x40000067: "DestroyStream",
			0x40000068: "GetStream",
			0x40000069: "GetStreams",
			0x4000006a: "UpdateStreamState",
			0x4000006b: "SetStreamFocus",
			0x4000006c: "GetStreamVoiceToken",
			0x4000006d: "KickFromStreamVoice",
			0x40000096: "CreateMessage",
			0x40000097: "DestroyMessage",
			0x40000098: "EditMessage",
			0x40000099: "SetMessagePinned",
			0x4000009a: "SetTypingIndicator",
			0x4000009b: "AdvanceStreamViewTime",
			0x4000009c: "GetStreamHistory",
			0x4000009d: "GetStreamMessage",
		}, ServiceMethods[uint16]{},
	},
	0x7e525e99: {
		"bgs.protocol.channel.v2.membership.ChannelMembershipService",
		ServiceMethods[uint32]{
			0x00000001: "Subscribe",
			0x00000002: "Unsubscribe",
			0x00000003: "GetState",
		}, ServiceMethods[uint16]{},
	},
	0x18007be: {"bgs.protocol.channel.v2.membership.ChannelMembershipListener",
		ServiceMethods[uint32]{
			0x80000001: "OnChannelAdded",
			0x80000002: "OnChannelRemoved",
			0x80000003: "OnReceivedInvitationAdded",
			0x80000004: "OnReceivedInvitationRemoved",
		}, ServiceMethods[uint16]{},
	},
	0x798d39d1: {
		"bgs.protocol.channel.v2.ChannelService",
		ServiceMethods[uint32]{
			0x00000002: "CreateChannel",
			0x00000003: "DissolveChannel",
			0x00000004: "GetChannel",
			0x00000005: "GetPublicChannelTypes",
			0x00000006: "FindChannel",
			0x0000000a: "Subscribe",
			0x0000000b: "Unsubscribe",
			0x00000015: "SetAttribute",
			0x00000016: "SetPrivacyLevel",
			0x00000017: "SendMessage",
			0x00000018: "SetTypingIndicator",
			0x0000001e: "Join",
			0x0000001f: "Leave",
			0x00000020: "Kick",
			0x00000028: "SetMemberAttribute",
			0x00000029: "AssignRole",
			0x0000002a: "UnassignRole",
			0x00000032: "SendInvitation",
			0x00000033: "AcceptInvitation",
			0x00000034: "DeclineInvitation",
			0x00000035: "RevokeInvitation",
			0x0000003c: "SendSuggestion",
			0x00000046: "GetJoinVoiceToken",
		}, ServiceMethods[uint16]{},
	},
	0x80909d73: {
		"bgs.protocol.club.v1.ClubListener",
		ServiceMethods[uint32]{
			0xc0000001: "OnSubscribe",
			0xc0000002: "OnUnsubscribe",
			0xc0000003: "OnStateChanged",
			0xc0000004: "OnSettingsChanged",
			0xc000001e: "OnMemberAdded",
			0xc000001f: "OnMemberRemoved",
			0xc0000020: "OnMemberStateChanged",
			0xc0000021: "OnSubscriberStateChanged",
			0xc0000022: "OnMemberRoleChanged",
			0xc0000032: "OnInvitationAdded",
			0xc0000033: "OnInvitationRemoved",
			0xc0000046: "OnSuggestionAdded",
			0xc0000047: "OnSuggestionRemoved",
			0xc0000064: "OnStreamAdded",
			0xc0000065: "OnStreamRemoved",
			0xc0000066: "OnStreamStateChanged",
			0xc0000096: "OnStreamMessageAdded",
			0xc0000098: "OnStreamMessageUpdated",
			0xc0000099: "OnStreamTypingIndicator",
			0xc000009a: "OnStreamUnreadIndicator",
			0xc000009b: "OnStreamAdvanceViewTime",
		}, ServiceMethods[uint16]{},
	},
	0x94b94786: {
		"bgs.protocol.club.v1.membership.ClubMembershipService",
		ServiceMethods[uint32]{
			0x40000001: "Subscribe",
			0x40000002: "Unsubscribe",
			0x40000003: "GetState",
			0x40000004: "UpdateClubSharedSettings",
			0x40000005: "GetStreamMentions",
			0x40000006: "RemoveStreamMentions",
			0x40000007: "AdvanceStreamMentionViewTime",
		}, ServiceMethods[uint16]{},
	},
	0x7b37d770: {
		"bgs.protocol.session.v2.client.SessionService",
		ServiceMethods[uint32]{
			0x40000001: "CreateSession",
			0x40000002: "RestoreSession",
			0x40000003: "DestroySession",
			0x40000004: "GetSignedSessionState",
			0x40000005: "GetSessionState",
			0x40000006: "MarkSessionAlive",
		}, ServiceMethods[uint16]{},
	},
	0xbca72135: {
		"bgs.protocol.session.v2.client.SessionListener",
		ServiceMethods[uint32]{
			0xc0000001: "OnSessionCreated",
			0xc0000002: "OnSessionDestroyed",
			0xc0000004: "OnSessionGameTimeWarning",
			0xc0000005: "OnSessionQueueUpdated",
			0xc0000006: "OnSessionQueueEnd",
		}, ServiceMethods[uint16]{},
	},
	0x2aaadbc: {
		"bgs.protocol.publication.v2.client.PublicationService",
		ServiceMethods[uint32]{
			0x40000001: "Subscribe",
			0x40000002: "Unsubscribe",
		}, ServiceMethods[uint16]{},
	},
	0x2362becd: {
		"bgs.protocol.notification.v2.client.NotificationListener",
		ServiceMethods[uint32]{
			0xc0000001: "OnNotificationReceived",
		}, ServiceMethods[uint16]{},
	},
	0x9e8758d9: {
		"bgs.protocol.publication.v2.client.PublicationListener",
		ServiceMethods[uint32]{
			0xC0000001: "OnPublicationReceived",
		}, ServiceMethods[uint16]{},
	},
	0xdbc01a89: {
		"bgs.protocol.friends.v2.client.FriendsListener",
		ServiceMethods[uint32]{
			0xC0000001: "OnFriendAdded",
			0xC0000002: "OnFriendRemoved",
			0xC0000003: "OnReceivedInvitationAdded",
			0xC0000004: "OnReceivedInvitationRemoved",
			0xC0000005: "OnSentInvitationAdded",
			0xC0000006: "OnSentInvitationRemoved",
			0xC0000007: "OnUpdateFriendState",
		}, ServiceMethods[uint16]{},
	},
	0x62da0891: {
		"bgs.protocol.account.v1.AccountService",
		ServiceMethods[uint32]{
			0x00000019: "Subscribe",
			0x0000000D: "ResolveAccount",
			0x0000001A: "Unsubscribe",
			0x0000001E: "GetAccountState",
			0x0000001F: "GetGameAccountState",
			0x00000020: "GetLicenses",
			0x00000021: "GetGameTimeRemainingInfo",
			0x00000023: "GetCAISInfo",
			0x00000025: "GetAuthorizedData",
			0x0000002C: "GetSignedAccountState",
			0x0000002D: "GetAccountInfo",
			0x0000002E: "GetAccountPlatformRestrictions",
		}, ServiceMethods[uint16]{},
	},
	0x1ae52686: {
		"bgs.protocol.channel.v2.ChannelListener",
		ServiceMethods[uint32]{
			0x80000003: "OnMemberAdded",
			0x80000004: "OnMemberRemoved",
			0x80000005: "OnMemberAttributeChanged",
			0x80000006: "OnMemberRoleChanged",
			0x8000000A: "OnSendMessage",
			0x8000000B: "OnTypingIndicator",
			0x80000010: "OnAttributeChanged",
			0x80000011: "OnPrivacyLevelChanged",
			0x80000012: "OnInvitationAdded",
			0x80000013: "OnInvitationRemoved",
			0x80000014: "OnSuggestionAdded",
		}, ServiceMethods[uint16]{},
	},
	0x3fe5849e: {
		"bgs.protocol.whisper.v1.WhisperListener",
		ServiceMethods[uint32]{
			0x80000001: "OnWhisper",
			0x80000002: "OnWhisperEcho",
			0x80000003: "OnTypingIndicatorUpdate",
			0x80000004: "OnAdvanceViewTime",
			0x80000005: "OnWhisperUpdated",
			0x80000006: "OnAdvanceClearTime",
		}, ServiceMethods[uint16]{},
	},
	0x5869be8c: {
		"bgs.protocol.friends.v2.client.FriendsService",
		ServiceMethods[uint32]{
			0x40000001: "Subscribe",
			0x40000002: "Unsubscribe",
			0x40000003: "GetSentInvitations",
			0x40000004: "GetReceivedInvitations",
			0x40000005: "GetFriends",
			0x40000006: "IsFriend",
			0x40000007: "ViewFriends",
			0x40000008: "SendInvitation",
			0x40000009: "AcceptInvitation",
			0x4000000a: "RevokeInvitation",
			0x4000000b: "RevokeAllInvitations",
			0x4000000c: "IgnoreInvitation",
			0x4000000d: "RemoveFriend",
			0x4000000e: "UpdateFriendState",
		}, ServiceMethods[uint16]{},
	},
	0x71240e35: {
		"bgs.protocol.authentication.v1.AuthenticationListener",
		ServiceMethods[uint32]{
			0x80000004: "OnServerStateChange",
			0x80000005: "OnLogonComplete",
			0x8000000a: "OnLogonUpdate",
			0x8000000b: "OnVersionInfoUpdated",
			0x8000000c: "OnLogonQueueUpdate",
			0x8000000d: "OnLogonQueueEnd",
		}, ServiceMethods[uint16]{},
	},
	0xc12828f9: {
		"bgs.protocol.whisper.v1.WhisperService",
		ServiceMethods[uint32]{
			0x40000001: "Subscribe",
			0x40000002: "Unsubscribe",
			0x40000003: "SendWhisper",
			0x40000004: "SetTypingIndicator",
			0x40000005: "AdvanceViewTime",
			0x40000006: "GetWhisperMessages",
			0x40000007: "AdvanceClearTime",
		}, ServiceMethods[uint16]{},
	},
	0x8e8f5fb0: {
		"bgs.protocol.block_list.v1.client.BlockListService",
		ServiceMethods[uint32]{
			0x40000001: "Subscribe",
			0x40000002: "Unsubscribe",
			0x40000003: "GetState",
			0x40000004: "BlockPlayer",
			0x40000005: "UnblockPlayer",
			0x40000006: "BlockPlayerForSession",
		}, ServiceMethods[uint16]{},
	},
	0xb5dd8a75: {
		"bgs.protocol.block_list.v1.client.BlockListListener",
		ServiceMethods[uint32]{
			0xc0000001: "OnBlockedPlayerAdded",
			0xc0000002: "OnBlockedPlayerRemoved",
		}, ServiceMethods[uint16]{},
	},
	0x9da8116b: {
		"bgs.protocol.authentication.v2.client.AuthenticationListener",
		ServiceMethods[uint32]{
			0xc0000001: "OnLogonComplete",
			0xc0000002: "OnLogonQueueUpdate",
			0xc0000003: "OnLogonQueueEnd",
			0xc0000004: "OnExternalChallenge",
		}, ServiceMethods[uint16]{},
	},

	0xc02f8216: {
		"bgs.protocol.authentication.v2.client.AuthenticationService",
		ServiceMethods[uint32]{
			0x40000001: "Logon",
			0x40000002: "VerifyAuthToken",
			0x40000003: "GenerateAuthToken",
		}, ServiceMethods[uint16]{},
	},

	0xdecfc01: {
		"bgs.protocol.authentication.v1.AuthenticationService",
		ServiceMethods[uint32]{
			0x00000001: "Logon",
			0x00000007: "VerifyWebCredentials",
			0x00000008: "GenerateWebCredentials",
		}, ServiceMethods[uint16]{},
	},

	0xbbda171f: {
		"bgs.protocol.challenge.v1.ChallengeListener",
		ServiceMethods[uint32]{
			0x80000003: "OnExternalChallenge",
			0x80000004: "OnExternalChallengeResult",
		}, ServiceMethods[uint16]{},
	},

	0xECBE75BA: {
		"bgs.protocol.resources.v1.ResourcesService",
		ServiceMethods[uint32]{
			0x00000001: "GetContentHandle",
			0x00000002: "GetTitleIcons",
		}, ServiceMethods[uint16]{},
	},
	0x5dbb51c2: {
		"bgs.protocol.game_utilities.v2.client.GameUtilitiesService",
		ServiceMethods[uint32]{
			0x40000001: "ProcessTask",
			0x40000002: "GetAllValuesForAttribute",
		}, ServiceMethods[uint16]{},
	},
}
