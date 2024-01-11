// #  Generated from mk_svc_code DO NOT EDIT
package services
import (
// looping over : build/protos/bgs/low/pb/client/account_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/account/v1"
// looping over : build/protos/bgs/low/pb/client/api/client/v1/block_list_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/block_list/v1/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v1/block_list_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/block_list/v1/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/authentication_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/authentication/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/authentication_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/authentication/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/channel_membership_service.proto
// looping over : build/protos/bgs/low/pb/client/api/client/v2/channel_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/channel/v2"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/friends_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/friends/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/friends_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/friends/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/game_utilities_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/game_utilities/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/notification_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/notification/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/presence_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/presence/v2"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/presence_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/presence/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/publication_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/publication/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/publication_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/publication/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/report_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/report/v2"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/session_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/session/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/session_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/session/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/client/v2/voice_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/voice/v2/client"
// looping over : build/protos/bgs/low/pb/client/api/server/v2/channel_listener.proto
// looping over : build/protos/bgs/low/pb/client/api/server/v2/game_utilities_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/game_utilities/v2/server"
// looping over : build/protos/bgs/low/pb/client/authentication_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/authentication/v1"
// looping over : build/protos/bgs/low/pb/client/challenge_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/challenge/v1"
// looping over : build/protos/bgs/low/pb/client/club_membership_listener.proto
// looping over : build/protos/bgs/low/pb/client/club_membership_service.proto
// looping over : build/protos/bgs/low/pb/client/connection_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/connection/v1"
// looping over : build/protos/bgs/low/pb/client/resource_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/resources/v1"
// looping over : build/protos/bgs/low/pb/client/session_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/session/v1"
// looping over : build/protos/bgs/low/pb/client/whisper_listener.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/whisper/v1"
// looping over : build/protos/bgs/low/pb/client/whisper_service.proto
   _ "github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/whisper/v1"
)
var svc_pool = map[string][]svc_method{
"bgs.protocol.account.v1.AccountService": []svc_method{
    svc_method{method_id: 1, method_name: "ResolveAccount", request_msg: "bgs.protocol.account.v1.ResolveAccountRequest", response_msg: "bgs.protocol.account.v1.ResolveAccountResponse" },
    svc_method{method_id: 2, method_name: "Subscribe", request_msg: "bgs.protocol.account.v1.SubscriptionUpdateRequest", response_msg: "bgs.protocol.account.v1.SubscriptionUpdateResponse" },
    svc_method{method_id: 3, method_name: "Unsubscribe", request_msg: "bgs.protocol.account.v1.SubscriptionUpdateRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 4, method_name: "GetAccountState", request_msg: "bgs.protocol.account.v1.GetAccountStateRequest", response_msg: "bgs.protocol.account.v1.GetAccountStateResponse" },
    svc_method{method_id: 5, method_name: "GetGameAccountState", request_msg: "bgs.protocol.account.v1.GetGameAccountStateRequest", response_msg: "bgs.protocol.account.v1.GetGameAccountStateResponse" },
    svc_method{method_id: 6, method_name: "GetLicenses", request_msg: "bgs.protocol.account.v1.GetLicensesRequest", response_msg: "bgs.protocol.account.v1.GetLicensesResponse" },
    svc_method{method_id: 7, method_name: "GetGameTimeRemainingInfo", request_msg: "bgs.protocol.account.v1.GetGameTimeRemainingInfoRequest", response_msg: "bgs.protocol.account.v1.GetGameTimeRemainingInfoResponse" },
    svc_method{method_id: 8, method_name: "GetGameSessionInfo", request_msg: "bgs.protocol.account.v1.GetGameSessionInfoRequest", response_msg: "bgs.protocol.account.v1.GetGameSessionInfoResponse" },
    svc_method{method_id: 9, method_name: "GetCAISInfo", request_msg: "bgs.protocol.account.v1.GetCAISInfoRequest", response_msg: "bgs.protocol.account.v1.GetCAISInfoResponse" },
    svc_method{method_id: 10, method_name: "GetAuthorizedData", request_msg: "bgs.protocol.account.v1.GetAuthorizedDataRequest", response_msg: "bgs.protocol.account.v1.GetAuthorizedDataResponse" },
    svc_method{method_id: 11, method_name: "GetSignedAccountState", request_msg: "bgs.protocol.account.v1.GetSignedAccountStateRequest", response_msg: "bgs.protocol.account.v1.GetSignedAccountStateResponse" },
    svc_method{method_id: 12, method_name: "GetAccountInfo", request_msg: "bgs.protocol.account.v1.GetAccountInfoRequest", response_msg: "bgs.protocol.account.v1.GetAccountInfoResponse" },
    svc_method{method_id: 13, method_name: "GetAccountPlatformRestrictions", request_msg: "bgs.protocol.account.v1.GetAccountPlatformRestrictionsRequest", response_msg: "bgs.protocol.account.v1.GetAccountPlatformRestrictionsResponse" },
},
"bgs.protocol.account.v1.AccountListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnAccountStateUpdated", request_msg: "bgs.protocol.account.v1.AccountStateNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnGameAccountStateUpdated", request_msg: "bgs.protocol.account.v1.GameAccountStateNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnGameAccountsUpdated", request_msg: "bgs.protocol.account.v1.GameAccountNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnGameSessionUpdated", request_msg: "bgs.protocol.account.v1.GameAccountSessionNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.block_list.v1.client.BlockListListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnBlockedPlayerAdded", request_msg: "bgs.protocol.block_list.v1.client.BlockedPlayerAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnBlockedPlayerRemoved", request_msg: "bgs.protocol.block_list.v1.client.BlockedPlayerRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.block_list.v1.client.BlockListService": []svc_method{
    svc_method{method_id: 1, method_name: "Subscribe", request_msg: "bgs.protocol.block_list.v1.client.SubscribeRequest", response_msg: "bgs.protocol.block_list.v1.client.SubscribeResponse" },
    svc_method{method_id: 2, method_name: "Unsubscribe", request_msg: "bgs.protocol.block_list.v1.client.UnsubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "GetState", request_msg: "bgs.protocol.block_list.v1.client.GetStateRequest", response_msg: "bgs.protocol.block_list.v1.client.GetStateResponse" },
    svc_method{method_id: 4, method_name: "BlockPlayer", request_msg: "bgs.protocol.block_list.v1.client.BlockPlayerRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 5, method_name: "UnblockPlayer", request_msg: "bgs.protocol.block_list.v1.client.UnblockPlayerRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 6, method_name: "BlockPlayerForSession", request_msg: "bgs.protocol.block_list.v1.client.BlockPlayerRequest", response_msg: "bgs.protocol.NoData" },
},
"bgs.protocol.authentication.v2.client.AuthenticationListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnLogonComplete", request_msg: "bgs.protocol.authentication.v2.client.LogonCompleteNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnLogonQueueUpdate", request_msg: "bgs.protocol.authentication.v2.client.LogonQueueUpdateNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnLogonQueueEnd", request_msg: "bgs.protocol.NoData", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnExternalChallenge", request_msg: "bgs.protocol.authentication.v2.client.ExternalChallengeNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.authentication.v2.client.AuthenticationService": []svc_method{
    svc_method{method_id: 1, method_name: "Logon", request_msg: "bgs.protocol.authentication.v2.client.LogonRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 2, method_name: "VerifyAuthToken", request_msg: "bgs.protocol.authentication.v2.client.VerifyAuthTokenRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "GenerateAuthToken", request_msg: "bgs.protocol.authentication.v2.client.GenerateAuthTokenRequest", response_msg: "bgs.protocol.authentication.v2.client.GenerateAuthTokenResponse" },
},
"bgs.protocol.channel.v2.membership.ChannelMembershipService": []svc_method{
    svc_method{method_id: 1, method_name: "Subscribe", request_msg: "bgs.protocol.channel.v2.membership.SubscribeRequest", response_msg: "bgs.protocol.channel.v2.membership.SubscribeResponse" },
    svc_method{method_id: 2, method_name: "Unsubscribe", request_msg: "bgs.protocol.channel.v2.membership.UnsubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "GetState", request_msg: "bgs.protocol.channel.v2.membership.GetStateRequest", response_msg: "bgs.protocol.channel.v2.membership.GetStateResponse" },
},
"bgs.protocol.channel.v2.membership.ChannelMembershipListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnChannelAdded", request_msg: "bgs.protocol.channel.v2.membership.ChannelAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnChannelRemoved", request_msg: "bgs.protocol.channel.v2.membership.ChannelRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnReceivedInvitationAdded", request_msg: "bgs.protocol.channel.v2.membership.ReceivedInvitationAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnReceivedInvitationRemoved", request_msg: "bgs.protocol.channel.v2.membership.ReceivedInvitationRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.channel.v2.ChannelService": []svc_method{
    svc_method{method_id: 1, method_name: "CreateChannel", request_msg: "bgs.protocol.channel.v2.CreateChannelRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 2, method_name: "DissolveChannel", request_msg: "bgs.protocol.channel.v2.DissolveChannelRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "GetChannel", request_msg: "bgs.protocol.channel.v2.GetChannelRequest", response_msg: "bgs.protocol.channel.v2.GetChannelResponse" },
    svc_method{method_id: 4, method_name: "GetPublicChannelTypes", request_msg: "bgs.protocol.channel.v2.GetPublicChannelTypesRequest", response_msg: "bgs.protocol.channel.v2.GetPublicChannelTypesResponse" },
    svc_method{method_id: 5, method_name: "FindChannel", request_msg: "bgs.protocol.channel.v2.FindChannelRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 6, method_name: "Subscribe", request_msg: "bgs.protocol.channel.v2.SubscribeRequest", response_msg: "bgs.protocol.channel.v2.SubscribeResponse" },
    svc_method{method_id: 7, method_name: "Unsubscribe", request_msg: "bgs.protocol.channel.v2.UnsubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 8, method_name: "SetAttribute", request_msg: "bgs.protocol.channel.v2.SetAttributeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 9, method_name: "SetPrivacyLevel", request_msg: "bgs.protocol.channel.v2.SetPrivacyLevelRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 10, method_name: "SendMessage", request_msg: "bgs.protocol.channel.v2.SendMessageRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 11, method_name: "SetTypingIndicator", request_msg: "bgs.protocol.channel.v2.SetTypingIndicatorRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 12, method_name: "Join", request_msg: "bgs.protocol.channel.v2.JoinRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 13, method_name: "Leave", request_msg: "bgs.protocol.channel.v2.LeaveRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 14, method_name: "Kick", request_msg: "bgs.protocol.channel.v2.KickRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 15, method_name: "SetMemberAttribute", request_msg: "bgs.protocol.channel.v2.SetMemberAttributeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 16, method_name: "AssignRole", request_msg: "bgs.protocol.channel.v2.AssignRoleRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 17, method_name: "UnassignRole", request_msg: "bgs.protocol.channel.v2.UnassignRoleRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 18, method_name: "SendInvitation", request_msg: "bgs.protocol.channel.v2.SendInvitationRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 19, method_name: "AcceptInvitation", request_msg: "bgs.protocol.channel.v2.AcceptInvitationRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 20, method_name: "DeclineInvitation", request_msg: "bgs.protocol.channel.v2.DeclineInvitationRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 21, method_name: "RevokeInvitation", request_msg: "bgs.protocol.channel.v2.RevokeInvitationRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 22, method_name: "SendSuggestion", request_msg: "bgs.protocol.channel.v2.SendSuggestionRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 23, method_name: "GetJoinVoiceToken", request_msg: "bgs.protocol.channel.v2.GetJoinVoiceTokenRequest", response_msg: "bgs.protocol.channel.v2.GetJoinVoiceTokenResponse" },
},
"bgs.protocol.channel.v2.ChannelListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnMemberAdded", request_msg: "bgs.protocol.channel.v2.MemberAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnMemberRemoved", request_msg: "bgs.protocol.channel.v2.MemberRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnMemberAttributeChanged", request_msg: "bgs.protocol.channel.v2.MemberAttributeChangedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnMemberRoleChanged", request_msg: "bgs.protocol.channel.v2.MemberRoleChangedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "OnSendMessage", request_msg: "bgs.protocol.channel.v2.SendMessageNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 6, method_name: "OnTypingIndicator", request_msg: "bgs.protocol.channel.v2.TypingIndicatorNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 7, method_name: "OnAttributeChanged", request_msg: "bgs.protocol.channel.v2.AttributeChangedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 8, method_name: "OnPrivacyLevelChanged", request_msg: "bgs.protocol.channel.v2.PrivacyLevelChangedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 9, method_name: "OnInvitationAdded", request_msg: "bgs.protocol.channel.v2.InvitationAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 10, method_name: "OnInvitationRemoved", request_msg: "bgs.protocol.channel.v2.InvitationRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 11, method_name: "OnSuggestionAdded", request_msg: "bgs.protocol.channel.v2.SuggestionAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.friends.v2.client.FriendsListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnFriendAdded", request_msg: "bgs.protocol.friends.v2.client.FriendAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnFriendRemoved", request_msg: "bgs.protocol.friends.v2.client.FriendRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnReceivedInvitationAdded", request_msg: "bgs.protocol.friends.v2.client.ReceivedInvitationAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnReceivedInvitationRemoved", request_msg: "bgs.protocol.friends.v2.client.ReceivedInvitationRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "OnSentInvitationAdded", request_msg: "bgs.protocol.friends.v2.client.SentInvitationAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 6, method_name: "OnSentInvitationRemoved", request_msg: "bgs.protocol.friends.v2.client.SentInvitationRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 7, method_name: "OnUpdateFriendState", request_msg: "bgs.protocol.friends.v2.client.UpdateFriendStateNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.friends.v2.client.FriendsService": []svc_method{
    svc_method{method_id: 1, method_name: "Subscribe", request_msg: "bgs.protocol.friends.v2.client.SubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 2, method_name: "Unsubscribe", request_msg: "bgs.protocol.friends.v2.client.UnsubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "GetSentInvitations", request_msg: "bgs.protocol.friends.v2.client.GetSentInvitationsRequest", response_msg: "bgs.protocol.friends.v2.client.GetSentInvitationsResponse" },
    svc_method{method_id: 4, method_name: "GetReceivedInvitations", request_msg: "bgs.protocol.friends.v2.client.GetReceivedInvitationsRequest", response_msg: "bgs.protocol.friends.v2.client.GetReceivedInvitationsResponse" },
    svc_method{method_id: 5, method_name: "GetFriends", request_msg: "bgs.protocol.friends.v2.client.GetFriendsRequest", response_msg: "bgs.protocol.friends.v2.client.GetFriendsResponse" },
    svc_method{method_id: 6, method_name: "IsFriend", request_msg: "bgs.protocol.friends.v2.client.IsFriendRequest", response_msg: "bgs.protocol.friends.v2.client.IsFriendResponse" },
    svc_method{method_id: 7, method_name: "ViewFriends", request_msg: "bgs.protocol.friends.v2.client.ViewFriendsRequest", response_msg: "bgs.protocol.friends.v2.client.ViewFriendsResponse" },
    svc_method{method_id: 8, method_name: "SendInvitation", request_msg: "bgs.protocol.friends.v2.client.SendInvitationRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 9, method_name: "AcceptInvitation", request_msg: "bgs.protocol.friends.v2.client.AcceptInvitationRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 10, method_name: "RevokeInvitation", request_msg: "bgs.protocol.friends.v2.client.RevokeInvitationRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 11, method_name: "RevokeAllInvitations", request_msg: "bgs.protocol.friends.v2.client.RevokeAllInvitationsRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 12, method_name: "IgnoreInvitation", request_msg: "bgs.protocol.friends.v2.client.IgnoreInvitationRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 13, method_name: "RemoveFriend", request_msg: "bgs.protocol.friends.v2.client.RemoveFriendRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 14, method_name: "UpdateFriendState", request_msg: "bgs.protocol.friends.v2.client.UpdateFriendStateRequest", response_msg: "bgs.protocol.NoData" },
},
"bgs.protocol.game_utilities.v2.client.GameUtilitiesService": []svc_method{
    svc_method{method_id: 1, method_name: "ProcessTask", request_msg: "bgs.protocol.game_utilities.v2.client.ProcessTaskRequest", response_msg: "bgs.protocol.game_utilities.v2.client.ProcessTaskResponse" },
    svc_method{method_id: 2, method_name: "GetAllValuesForAttribute", request_msg: "bgs.protocol.game_utilities.v2.client.GetAllValuesForAttributeRequest", response_msg: "bgs.protocol.game_utilities.v2.client.GetAllValuesForAttributeResponse" },
},
"bgs.protocol.notification.v2.client.NotificationListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnNotificationReceived", request_msg: "bgs.protocol.notification.v2.client.NotificationReceivedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.presence.v2.PresenceListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnPresenceStateUpdated", request_msg: "bgs.protocol.presence.v2.PresenceStateUpdatedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnSubscribeFailure", request_msg: "bgs.protocol.presence.v2.SubscribeFailureNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.presence.v2.client.PresenceService": []svc_method{
    svc_method{method_id: 1, method_name: "BatchSubscribe", request_msg: "bgs.protocol.presence.v2.client.BatchSubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 2, method_name: "BatchUnsubscribe", request_msg: "bgs.protocol.presence.v2.client.BatchUnsubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "Query", request_msg: "bgs.protocol.presence.v2.client.QueryRequest", response_msg: "bgs.protocol.presence.v2.client.QueryResponse" },
    svc_method{method_id: 4, method_name: "Update", request_msg: "bgs.protocol.presence.v2.client.UpdateRequest", response_msg: "bgs.protocol.NoData" },
},
"bgs.protocol.publication.v2.client.PublicationListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnPublicationReceived", request_msg: "bgs.protocol.publication.v2.client.PublicationReceivedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.publication.v2.client.PublicationService": []svc_method{
    svc_method{method_id: 1, method_name: "Subscribe", request_msg: "bgs.protocol.publication.v2.client.SubscribeRequest", response_msg: "bgs.protocol.publication.v2.client.SubscribeResponse" },
    svc_method{method_id: 2, method_name: "Unsubscribe", request_msg: "bgs.protocol.publication.v2.client.UnsubscribeRequest", response_msg: "bgs.protocol.NoData" },
},
"bgs.protocol.report.v2.ReportService": []svc_method{
    svc_method{method_id: 1, method_name: "SubmitReport", request_msg: "bgs.protocol.report.v2.SubmitReportRequest", response_msg: "bgs.protocol.NoData" },
},
"bgs.protocol.session.v2.client.SessionListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnSessionCreated", request_msg: "bgs.protocol.session.v2.client.SessionCreatedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnSessionDestroyed", request_msg: "bgs.protocol.session.v2.client.SessionDestroyedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnSessionGameTimeWarning", request_msg: "bgs.protocol.session.v2.client.SessionGameTimeWarningNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnSessionQueueUpdated", request_msg: "bgs.protocol.session.v2.client.SessionQueueUpdatedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "OnSessionQueueEnd", request_msg: "bgs.protocol.session.v2.client.SessionQueueEndNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.session.v2.client.SessionService": []svc_method{
    svc_method{method_id: 1, method_name: "CreateSession", request_msg: "bgs.protocol.session.v2.client.CreateSessionRequest", response_msg: "bgs.protocol.session.v2.client.CreateSessionResponse" },
    svc_method{method_id: 2, method_name: "RestoreSession", request_msg: "bgs.protocol.session.v2.client.RestoreSessionRequest", response_msg: "bgs.protocol.session.v2.client.RestoreSessionResponse" },
    svc_method{method_id: 3, method_name: "DestroySession", request_msg: "bgs.protocol.session.v2.client.DestroySessionRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 4, method_name: "GetSignedSessionState", request_msg: "bgs.protocol.session.v2.client.GetSignedSessionStateRequest", response_msg: "bgs.protocol.session.v2.client.GetSignedSessionStateResponse" },
    svc_method{method_id: 5, method_name: "GetSessionState", request_msg: "bgs.protocol.session.v2.client.GetSessionStateRequest", response_msg: "bgs.protocol.session.v2.client.GetSessionStateResponse" },
    svc_method{method_id: 6, method_name: "MarkSessionAlive", request_msg: "bgs.protocol.session.v2.client.MarkSessionAliveRequest", response_msg: "bgs.protocol.NoData" },
},
"bgs.protocol.voice.v2.client.VoiceService": []svc_method{
    svc_method{method_id: 1, method_name: "CreateLoginCredentials", request_msg: "bgs.protocol.voice.v2.client.CreateLoginCredentialsRequest", response_msg: "bgs.protocol.voice.v2.client.CreateLoginCredentialsResponse" },
    svc_method{method_id: 2, method_name: "CreateChannelSttToken", request_msg: "bgs.protocol.voice.v2.client.CreateChannelSttTokenRequest", response_msg: "bgs.protocol.voice.v2.client.CreateChannelSttTokenResponse" },
},
"bgs.protocol.channel.v2.server.ChannelServerListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnChannelStateChanged", request_msg: "bgs.protocol.channel.v2.server.ChannelStateChangedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnMemberAdded", request_msg: "bgs.protocol.channel.v2.server.MemberAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnMemberRemoved", request_msg: "bgs.protocol.channel.v2.server.MemberRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnMemberStateChanged", request_msg: "bgs.protocol.channel.v2.server.MemberStateChangedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "OnMemberRoleChanged", request_msg: "bgs.protocol.channel.v2.server.MemberRoleChangedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.game_utilities.v2.server.GameUtilitiesListener": []svc_method{
    svc_method{method_id: 1, method_name: "ProcessClientTask", request_msg: "bgs.protocol.game_utilities.v2.server.ProcessClientTaskRequest", response_msg: "bgs.protocol.game_utilities.v2.server.ProcessClientTaskResponse" },
    svc_method{method_id: 2, method_name: "ProcessServerTask", request_msg: "bgs.protocol.game_utilities.v2.server.ProcessServerTaskRequest", response_msg: "bgs.protocol.game_utilities.v2.server.ProcessServerTaskResponse" },
    svc_method{method_id: 3, method_name: "ProcessAdminTask", request_msg: "bgs.protocol.game_utilities.v2.server.ProcessAdminTaskRequest", response_msg: "bgs.protocol.game_utilities.v2.server.ProcessAdminTaskResponse" },
},
"bgs.protocol.authentication.v1.AuthenticationListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnServerStateChange", request_msg: "bgs.protocol.authentication.v1.ServerStateChangeRequest", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnLogonComplete", request_msg: "bgs.protocol.authentication.v1.LogonResult", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnLogonUpdate", request_msg: "bgs.protocol.authentication.v1.LogonUpdateRequest", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnVersionInfoUpdated", request_msg: "bgs.protocol.authentication.v1.VersionInfoNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "OnLogonQueueUpdate", request_msg: "bgs.protocol.authentication.v1.LogonQueueUpdateRequest", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 6, method_name: "OnLogonQueueEnd", request_msg: "bgs.protocol.NoData", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.authentication.v1.AuthenticationService": []svc_method{
    svc_method{method_id: 1, method_name: "Logon", request_msg: "bgs.protocol.authentication.v1.LogonRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 2, method_name: "VerifyWebCredentials", request_msg: "bgs.protocol.authentication.v1.VerifyWebCredentialsRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "GenerateWebCredentials", request_msg: "bgs.protocol.authentication.v1.GenerateWebCredentialsRequest", response_msg: "bgs.protocol.authentication.v1.GenerateWebCredentialsResponse" },
},
"bgs.protocol.challenge.v1.ChallengeListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnExternalChallenge", request_msg: "bgs.protocol.challenge.v1.ChallengeExternalRequest", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnExternalChallengeResult", request_msg: "bgs.protocol.challenge.v1.ChallengeExternalResult", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.club.v1.membership.ClubMembershipListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnClubAdded", request_msg: "bgs.protocol.club.v1.membership.ClubAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnClubRemoved", request_msg: "bgs.protocol.club.v1.membership.ClubRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnReceivedInvitationAdded", request_msg: "bgs.protocol.club.v1.membership.ReceivedInvitationAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnReceivedInvitationRemoved", request_msg: "bgs.protocol.club.v1.membership.ReceivedInvitationRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "OnSharedSettingsChanged", request_msg: "bgs.protocol.club.v1.membership.SharedSettingsChangedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 6, method_name: "OnStreamMentionAdded", request_msg: "bgs.protocol.club.v1.membership.StreamMentionAddedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 7, method_name: "OnStreamMentionRemoved", request_msg: "bgs.protocol.club.v1.membership.StreamMentionRemovedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 8, method_name: "OnStreamMentionAdvanceViewTime", request_msg: "bgs.protocol.club.v1.membership.StreamMentionAdvanceViewTimeNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.club.v1.membership.ClubMembershipService": []svc_method{
    svc_method{method_id: 1, method_name: "Subscribe", request_msg: "bgs.protocol.club.v1.membership.SubscribeRequest", response_msg: "bgs.protocol.club.v1.membership.SubscribeResponse" },
    svc_method{method_id: 2, method_name: "Unsubscribe", request_msg: "bgs.protocol.club.v1.membership.UnsubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "GetState", request_msg: "bgs.protocol.club.v1.membership.GetStateRequest", response_msg: "bgs.protocol.club.v1.membership.GetStateResponse" },
    svc_method{method_id: 4, method_name: "UpdateClubSharedSettings", request_msg: "bgs.protocol.club.v1.membership.UpdateClubSharedSettingsRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 5, method_name: "GetStreamMentions", request_msg: "bgs.protocol.club.v1.membership.GetStreamMentionsRequest", response_msg: "bgs.protocol.club.v1.membership.GetStreamMentionsResponse" },
    svc_method{method_id: 6, method_name: "RemoveStreamMentions", request_msg: "bgs.protocol.club.v1.membership.RemoveStreamMentionsRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 7, method_name: "AdvanceStreamMentionViewTime", request_msg: "bgs.protocol.club.v1.membership.AdvanceStreamMentionViewTimeRequest", response_msg: "bgs.protocol.NoData" },
},
"bgs.protocol.connection.v1.ConnectionService": []svc_method{
    svc_method{method_id: 1, method_name: "Connect", request_msg: "bgs.protocol.connection.v1.ConnectRequest", response_msg: "bgs.protocol.connection.v1.ConnectResponse" },
    svc_method{method_id: 2, method_name: "Bind", request_msg: "bgs.protocol.connection.v1.BindRequest", response_msg: "bgs.protocol.connection.v1.BindResponse" },
    svc_method{method_id: 3, method_name: "Echo", request_msg: "bgs.protocol.connection.v1.EchoRequest", response_msg: "bgs.protocol.connection.v1.EchoResponse" },
    svc_method{method_id: 4, method_name: "ForceDisconnect", request_msg: "bgs.protocol.connection.v1.DisconnectNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "KeepAlive", request_msg: "bgs.protocol.NoData", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 6, method_name: "Encrypt", request_msg: "bgs.protocol.connection.v1.EncryptRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 7, method_name: "RequestDisconnect", request_msg: "bgs.protocol.connection.v1.DisconnectRequest", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.resources.v1.ResourcesService": []svc_method{
    svc_method{method_id: 1, method_name: "GetContentHandle", request_msg: "bgs.protocol.resources.v1.ContentHandleRequest", response_msg: "bgs.protocol.ContentHandle" },
    svc_method{method_id: 2, method_name: "GetTitleIcons", request_msg: "bgs.protocol.resources.v1.GetTitleIconsRequest", response_msg: "bgs.protocol.resources.v1.GetTitleIconsResponse" },
},
"bgs.protocol.session.v1.SessionListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnSessionCreated", request_msg: "bgs.protocol.session.v1.SessionCreatedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnSessionDestroyed", request_msg: "bgs.protocol.session.v1.SessionDestroyedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnSessionUpdated", request_msg: "bgs.protocol.session.v1.SessionUpdatedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnSessionGameTimeWarning", request_msg: "bgs.protocol.session.v1.SessionGameTimeWarningNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "OnSessionQueueUpdated", request_msg: "bgs.protocol.session.v1.SessionQueueUpdatedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 6, method_name: "OnSessionQueueEnd", request_msg: "bgs.protocol.session.v1.SessionQueueEndNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.session.v1.SessionService": []svc_method{
    svc_method{method_id: 1, method_name: "CreateSession", request_msg: "bgs.protocol.session.v1.CreateSessionRequest", response_msg: "bgs.protocol.session.v1.CreateSessionResponse" },
    svc_method{method_id: 2, method_name: "DestroySession", request_msg: "bgs.protocol.session.v1.DestroySessionRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "UpdateSession", request_msg: "bgs.protocol.session.v1.UpdateSessionRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 4, method_name: "GetSessionStateByBenefactor", request_msg: "bgs.protocol.session.v1.GetSessionStateByBenefactorRequest", response_msg: "bgs.protocol.session.v1.GetSessionStateByBenefactorResponse" },
    svc_method{method_id: 5, method_name: "MarkSessionsAlive", request_msg: "bgs.protocol.session.v1.MarkSessionsAliveRequest", response_msg: "bgs.protocol.session.v1.MarkSessionsAliveResponse" },
    svc_method{method_id: 6, method_name: "GetSessionState", request_msg: "bgs.protocol.session.v1.GetSessionStateRequest", response_msg: "bgs.protocol.session.v1.GetSessionStateResponse" },
    svc_method{method_id: 7, method_name: "GetSignedSessionState", request_msg: "bgs.protocol.session.v1.GetSignedSessionStateRequest", response_msg: "bgs.protocol.session.v1.GetSignedSessionStateResponse" },
    svc_method{method_id: 8, method_name: "RefreshSessionKey", request_msg: "bgs.protocol.session.v1.RefreshSessionKeyRequest", response_msg: "bgs.protocol.session.v1.RefreshSessionKeyResponse" },
    svc_method{method_id: 9, method_name: "GetGameSessionInfo", request_msg: "bgs.protocol.session.v1.GetGameSessionInfoRequest", response_msg: "bgs.protocol.session.v1.GetGameSessionInfoResponse" },
},
"bgs.protocol.whisper.v1.WhisperListener": []svc_method{
    svc_method{method_id: 1, method_name: "OnWhisper", request_msg: "bgs.protocol.whisper.v1.WhisperNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 2, method_name: "OnWhisperEcho", request_msg: "bgs.protocol.whisper.v1.WhisperEchoNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 3, method_name: "OnTypingIndicatorUpdate", request_msg: "bgs.protocol.whisper.v1.TypingIndicatorNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 4, method_name: "OnAdvanceViewTime", request_msg: "bgs.protocol.whisper.v1.AdvanceViewTimeNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 5, method_name: "OnWhisperUpdated", request_msg: "bgs.protocol.whisper.v1.WhisperUpdatedNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
    svc_method{method_id: 6, method_name: "OnAdvanceClearTime", request_msg: "bgs.protocol.whisper.v1.AdvanceClearTimeNotification", response_msg: "bgs.protocol.NO_RESPONSE" },
},
"bgs.protocol.whisper.v1.WhisperService": []svc_method{
    svc_method{method_id: 1, method_name: "Subscribe", request_msg: "bgs.protocol.whisper.v1.SubscribeRequest", response_msg: "bgs.protocol.whisper.v1.SubscribeResponse" },
    svc_method{method_id: 2, method_name: "Unsubscribe", request_msg: "bgs.protocol.whisper.v1.UnsubscribeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 3, method_name: "SendWhisper", request_msg: "bgs.protocol.whisper.v1.SendWhisperRequest", response_msg: "bgs.protocol.whisper.v1.SendWhisperResponse" },
    svc_method{method_id: 4, method_name: "SetTypingIndicator", request_msg: "bgs.protocol.whisper.v1.SetTypingIndicatorRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 5, method_name: "AdvanceViewTime", request_msg: "bgs.protocol.whisper.v1.AdvanceViewTimeRequest", response_msg: "bgs.protocol.NoData" },
    svc_method{method_id: 6, method_name: "GetWhisperMessages", request_msg: "bgs.protocol.whisper.v1.GetWhisperMessagesRequest", response_msg: "bgs.protocol.whisper.v1.GetWhisperMessagesResponse" },
    svc_method{method_id: 7, method_name: "AdvanceClearTime", request_msg: "bgs.protocol.whisper.v1.AdvanceClearTimeRequest", response_msg: "bgs.protocol.NoData" },
},
}
