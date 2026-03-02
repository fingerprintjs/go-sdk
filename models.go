package fingerprint

import (
	"github.com/fingerprintjs/go-sdk/v8/internal"
)

type BotInfo = internal.BotInfo

var NewBotInfo = internal.NewBotInfo

var NewBotInfoWithDefaults = internal.NewBotInfoWithDefaults

type NullableBotInfo = internal.NullableBotInfo

var NewNullableBotInfo = internal.NewNullableBotInfo

type BotResult = internal.BotResult

const (
	BotResultBad          = internal.BotResultBad
	BotResultGood         = internal.BotResultGood
	BotResultNot_detected = internal.BotResultNot_detected
)

var AllowedBotResultEnumValues = internal.AllowedBotResultEnumValues

var NewBotResultFromValue = internal.NewBotResultFromValue

type NullableBotResult = internal.NullableBotResult

var NewNullableBotResult = internal.NewNullableBotResult

type BrowserDetails = internal.BrowserDetails

var NewBrowserDetails = internal.NewBrowserDetails

var NewBrowserDetailsWithDefaults = internal.NewBrowserDetailsWithDefaults

type NullableBrowserDetails = internal.NullableBrowserDetails

var NewNullableBrowserDetails = internal.NewNullableBrowserDetails

type Canvas = internal.Canvas

var NewCanvas = internal.NewCanvas

var NewCanvasWithDefaults = internal.NewCanvasWithDefaults

type NullableCanvas = internal.NullableCanvas

var NewNullableCanvas = internal.NewNullableCanvas

type Emoji = internal.Emoji

var NewEmoji = internal.NewEmoji

var NewEmojiWithDefaults = internal.NewEmojiWithDefaults

type NullableEmoji = internal.NullableEmoji

var NewNullableEmoji = internal.NewNullableEmoji

type Error = internal.Error

var NewError = internal.NewError

var NewErrorWithDefaults = internal.NewErrorWithDefaults

type NullableError = internal.NullableError

var NewNullableError = internal.NewNullableError

type ErrorCode = internal.ErrorCode

const (
	ErrorCodeRequest_cannot_be_parsed = internal.ErrorCodeRequest_cannot_be_parsed
	ErrorCodeSecret_api_key_required  = internal.ErrorCodeSecret_api_key_required
	ErrorCodeSecret_api_key_not_found = internal.ErrorCodeSecret_api_key_not_found
	ErrorCodePublic_api_key_required  = internal.ErrorCodePublic_api_key_required
	ErrorCodePublic_api_key_not_found = internal.ErrorCodePublic_api_key_not_found
	ErrorCodeSubscription_not_active  = internal.ErrorCodeSubscription_not_active
	ErrorCodeWrong_region             = internal.ErrorCodeWrong_region
	ErrorCodeFeature_not_enabled      = internal.ErrorCodeFeature_not_enabled
	ErrorCodeRequest_not_found        = internal.ErrorCodeRequest_not_found
	ErrorCodeVisitor_not_found        = internal.ErrorCodeVisitor_not_found
	ErrorCodeToo_many_requests        = internal.ErrorCodeToo_many_requests
	ErrorCodeState_not_ready          = internal.ErrorCodeState_not_ready
	ErrorCodeFailed                   = internal.ErrorCodeFailed
	ErrorCodeEvent_not_found          = internal.ErrorCodeEvent_not_found
	ErrorCodeMissing_module           = internal.ErrorCodeMissing_module
	ErrorCodePayload_too_large        = internal.ErrorCodePayload_too_large
	ErrorCodeService_unavailable      = internal.ErrorCodeService_unavailable
	ErrorCodeRuleset_not_found        = internal.ErrorCodeRuleset_not_found
)

var AllowedErrorCodeEnumValues = internal.AllowedErrorCodeEnumValues

var NewErrorCodeFromValue = internal.NewErrorCodeFromValue

type NullableErrorCode = internal.NullableErrorCode

var NewNullableErrorCode = internal.NewNullableErrorCode

type ErrorResponse = internal.ErrorResponse

var NewErrorResponse = internal.NewErrorResponse

var NewErrorResponseWithDefaults = internal.NewErrorResponseWithDefaults

type NullableErrorResponse = internal.NullableErrorResponse

var NewNullableErrorResponse = internal.NewNullableErrorResponse

type Event = internal.Event

var NewEvent = internal.NewEvent

var NewEventWithDefaults = internal.NewEventWithDefaults

type NullableEvent = internal.NullableEvent

var NewNullableEvent = internal.NewNullableEvent

type EventRuleAction = internal.EventRuleAction

type NullableEventRuleAction = internal.NullableEventRuleAction

var NewNullableEventRuleAction = internal.NewNullableEventRuleAction

type EventRuleActionAllow = internal.EventRuleActionAllow

var NewEventRuleActionAllow = internal.NewEventRuleActionAllow

var NewEventRuleActionAllowWithDefaults = internal.NewEventRuleActionAllowWithDefaults

type NullableEventRuleActionAllow = internal.NullableEventRuleActionAllow

var NewNullableEventRuleActionAllow = internal.NewNullableEventRuleActionAllow

type EventRuleActionBlock = internal.EventRuleActionBlock

var NewEventRuleActionBlock = internal.NewEventRuleActionBlock

var NewEventRuleActionBlockWithDefaults = internal.NewEventRuleActionBlockWithDefaults

type NullableEventRuleActionBlock = internal.NullableEventRuleActionBlock

var NewNullableEventRuleActionBlock = internal.NewNullableEventRuleActionBlock

type EventSearch = internal.EventSearch

var NewEventSearch = internal.NewEventSearch

var NewEventSearchWithDefaults = internal.NewEventSearchWithDefaults

type NullableEventSearch = internal.NullableEventSearch

var NewNullableEventSearch = internal.NewNullableEventSearch

type EventUpdate = internal.EventUpdate

var NewEventUpdate = internal.NewEventUpdate

var NewEventUpdateWithDefaults = internal.NewEventUpdateWithDefaults

type NullableEventUpdate = internal.NullableEventUpdate

var NewNullableEventUpdate = internal.NewNullableEventUpdate

type FontPreferences = internal.FontPreferences

var NewFontPreferences = internal.NewFontPreferences

var NewFontPreferencesWithDefaults = internal.NewFontPreferencesWithDefaults

type NullableFontPreferences = internal.NullableFontPreferences

var NewNullableFontPreferences = internal.NewNullableFontPreferences

type Geolocation = internal.Geolocation

var NewGeolocation = internal.NewGeolocation

var NewGeolocationWithDefaults = internal.NewGeolocationWithDefaults

type NullableGeolocation = internal.NullableGeolocation

var NewNullableGeolocation = internal.NewNullableGeolocation

type GeolocationSubdivisionsInner = internal.GeolocationSubdivisionsInner

var NewGeolocationSubdivisionsInner = internal.NewGeolocationSubdivisionsInner

var NewGeolocationSubdivisionsInnerWithDefaults = internal.NewGeolocationSubdivisionsInnerWithDefaults

type NullableGeolocationSubdivisionsInner = internal.NullableGeolocationSubdivisionsInner

var NewNullableGeolocationSubdivisionsInner = internal.NewNullableGeolocationSubdivisionsInner

type IPBlockList = internal.IPBlockList

var NewIPBlockList = internal.NewIPBlockList

var NewIPBlockListWithDefaults = internal.NewIPBlockListWithDefaults

type NullableIPBlockList = internal.NullableIPBlockList

var NewNullableIPBlockList = internal.NewNullableIPBlockList

type IPInfo = internal.IPInfo

var NewIPInfo = internal.NewIPInfo

var NewIPInfoWithDefaults = internal.NewIPInfoWithDefaults

type NullableIPInfo = internal.NullableIPInfo

var NewNullableIPInfo = internal.NewNullableIPInfo

type IPInfoV4 = internal.IPInfoV4

var NewIPInfoV4 = internal.NewIPInfoV4

var NewIPInfoV4WithDefaults = internal.NewIPInfoV4WithDefaults

type NullableIPInfoV4 = internal.NullableIPInfoV4

var NewNullableIPInfoV4 = internal.NewNullableIPInfoV4

type IPInfoV6 = internal.IPInfoV6

var NewIPInfoV6 = internal.NewIPInfoV6

var NewIPInfoV6WithDefaults = internal.NewIPInfoV6WithDefaults

type NullableIPInfoV6 = internal.NullableIPInfoV6

var NewNullableIPInfoV6 = internal.NewNullableIPInfoV6

type Identification = internal.Identification

var NewIdentification = internal.NewIdentification

var NewIdentificationWithDefaults = internal.NewIdentificationWithDefaults

type NullableIdentification = internal.NullableIdentification

var NewNullableIdentification = internal.NewNullableIdentification

type IdentificationConfidence = internal.IdentificationConfidence

var NewIdentificationConfidence = internal.NewIdentificationConfidence

var NewIdentificationConfidenceWithDefaults = internal.NewIdentificationConfidenceWithDefaults

type NullableIdentificationConfidence = internal.NullableIdentificationConfidence

var NewNullableIdentificationConfidence = internal.NewNullableIdentificationConfidence

type Integration = internal.Integration

var NewIntegration = internal.NewIntegration

var NewIntegrationWithDefaults = internal.NewIntegrationWithDefaults

type NullableIntegration = internal.NullableIntegration

var NewNullableIntegration = internal.NewNullableIntegration

type IntegrationSubintegration = internal.IntegrationSubintegration

var NewIntegrationSubintegration = internal.NewIntegrationSubintegration

var NewIntegrationSubintegrationWithDefaults = internal.NewIntegrationSubintegrationWithDefaults

type NullableIntegrationSubintegration = internal.NullableIntegrationSubintegration

var NewNullableIntegrationSubintegration = internal.NewNullableIntegrationSubintegration

type PluginsInner = internal.PluginsInner

var NewPluginsInner = internal.NewPluginsInner

var NewPluginsInnerWithDefaults = internal.NewPluginsInnerWithDefaults

type NullablePluginsInner = internal.NullablePluginsInner

var NewNullablePluginsInner = internal.NewNullablePluginsInner

type PluginsInnerMimeTypesInner = internal.PluginsInnerMimeTypesInner

var NewPluginsInnerMimeTypesInner = internal.NewPluginsInnerMimeTypesInner

var NewPluginsInnerMimeTypesInnerWithDefaults = internal.NewPluginsInnerMimeTypesInnerWithDefaults

type NullablePluginsInnerMimeTypesInner = internal.NullablePluginsInnerMimeTypesInner

var NewNullablePluginsInnerMimeTypesInner = internal.NewNullablePluginsInnerMimeTypesInner

type Proximity = internal.Proximity

var NewProximity = internal.NewProximity

var NewProximityWithDefaults = internal.NewProximityWithDefaults

type NullableProximity = internal.NullableProximity

var NewNullableProximity = internal.NewNullableProximity

type ProxyConfidence = internal.ProxyConfidence

const (
	ProxyConfidenceLow    = internal.ProxyConfidenceLow
	ProxyConfidenceMedium = internal.ProxyConfidenceMedium
	ProxyConfidenceHigh   = internal.ProxyConfidenceHigh
)

var AllowedProxyConfidenceEnumValues = internal.AllowedProxyConfidenceEnumValues

var NewProxyConfidenceFromValue = internal.NewProxyConfidenceFromValue

type NullableProxyConfidence = internal.NullableProxyConfidence

var NewNullableProxyConfidence = internal.NewNullableProxyConfidence

type ProxyDetails = internal.ProxyDetails

var NewProxyDetails = internal.NewProxyDetails

var NewProxyDetailsWithDefaults = internal.NewProxyDetailsWithDefaults

type NullableProxyDetails = internal.NullableProxyDetails

var NewNullableProxyDetails = internal.NewNullableProxyDetails

type RawDeviceAttributes = internal.RawDeviceAttributes

var NewRawDeviceAttributes = internal.NewRawDeviceAttributes

var NewRawDeviceAttributesWithDefaults = internal.NewRawDeviceAttributesWithDefaults

type NullableRawDeviceAttributes = internal.NullableRawDeviceAttributes

var NewNullableRawDeviceAttributes = internal.NewNullableRawDeviceAttributes

type RequestHeaderModifications = internal.RequestHeaderModifications

var NewRequestHeaderModifications = internal.NewRequestHeaderModifications

var NewRequestHeaderModificationsWithDefaults = internal.NewRequestHeaderModificationsWithDefaults

type NullableRequestHeaderModifications = internal.NullableRequestHeaderModifications

var NewNullableRequestHeaderModifications = internal.NewNullableRequestHeaderModifications

type RuleActionHeaderField = internal.RuleActionHeaderField

var NewRuleActionHeaderField = internal.NewRuleActionHeaderField

var NewRuleActionHeaderFieldWithDefaults = internal.NewRuleActionHeaderFieldWithDefaults

type NullableRuleActionHeaderField = internal.NullableRuleActionHeaderField

var NewNullableRuleActionHeaderField = internal.NewNullableRuleActionHeaderField

type RuleActionType = internal.RuleActionType

const (
	RuleActionTypeAllow = internal.RuleActionTypeAllow
	RuleActionTypeBlock = internal.RuleActionTypeBlock
)

var AllowedRuleActionTypeEnumValues = internal.AllowedRuleActionTypeEnumValues

var NewRuleActionTypeFromValue = internal.NewRuleActionTypeFromValue

type NullableRuleActionType = internal.NullableRuleActionType

var NewNullableRuleActionType = internal.NewNullableRuleActionType

type SDK = internal.SDK

var NewSDK = internal.NewSDK

var NewSDKWithDefaults = internal.NewSDKWithDefaults

type NullableSDK = internal.NullableSDK

var NewNullableSDK = internal.NewNullableSDK

type SearchEventsBot = internal.SearchEventsBot

const (
	SearchEventsBotAll  = internal.SearchEventsBotAll
	SearchEventsBotGood = internal.SearchEventsBotGood
	SearchEventsBotBad  = internal.SearchEventsBotBad
	SearchEventsBotNone = internal.SearchEventsBotNone
)

var AllowedSearchEventsBotEnumValues = internal.AllowedSearchEventsBotEnumValues

var NewSearchEventsBotFromValue = internal.NewSearchEventsBotFromValue

type NullableSearchEventsBot = internal.NullableSearchEventsBot

var NewNullableSearchEventsBot = internal.NewNullableSearchEventsBot

type SearchEventsSDKPlatform = internal.SearchEventsSDKPlatform

const (
	SearchEventsSDKPlatformJs      = internal.SearchEventsSDKPlatformJs
	SearchEventsSDKPlatformAndroid = internal.SearchEventsSDKPlatformAndroid
	SearchEventsSDKPlatformIos     = internal.SearchEventsSDKPlatformIos
)

var AllowedSearchEventsSDKPlatformEnumValues = internal.AllowedSearchEventsSDKPlatformEnumValues

var NewSearchEventsSDKPlatformFromValue = internal.NewSearchEventsSDKPlatformFromValue

type NullableSearchEventsSDKPlatform = internal.NullableSearchEventsSDKPlatform

var NewNullableSearchEventsSDKPlatform = internal.NewNullableSearchEventsSDKPlatform

type SearchEventsVPNConfidence = internal.SearchEventsVPNConfidence

const (
	SearchEventsVPNConfidenceHigh   = internal.SearchEventsVPNConfidenceHigh
	SearchEventsVPNConfidenceMedium = internal.SearchEventsVPNConfidenceMedium
	SearchEventsVPNConfidenceLow    = internal.SearchEventsVPNConfidenceLow
)

var AllowedSearchEventsVPNConfidenceEnumValues = internal.AllowedSearchEventsVPNConfidenceEnumValues

var NewSearchEventsVPNConfidenceFromValue = internal.NewSearchEventsVPNConfidenceFromValue

type NullableSearchEventsVPNConfidence = internal.NullableSearchEventsVPNConfidence

var NewNullableSearchEventsVPNConfidence = internal.NewNullableSearchEventsVPNConfidence

type SupplementaryIDHighRecall = internal.SupplementaryIDHighRecall

var NewSupplementaryIDHighRecall = internal.NewSupplementaryIDHighRecall

var NewSupplementaryIDHighRecallWithDefaults = internal.NewSupplementaryIDHighRecallWithDefaults

type NullableSupplementaryIDHighRecall = internal.NullableSupplementaryIDHighRecall

var NewNullableSupplementaryIDHighRecall = internal.NewNullableSupplementaryIDHighRecall

type TamperingDetails = internal.TamperingDetails

var NewTamperingDetails = internal.NewTamperingDetails

var NewTamperingDetailsWithDefaults = internal.NewTamperingDetailsWithDefaults

type NullableTamperingDetails = internal.NullableTamperingDetails

var NewNullableTamperingDetails = internal.NewNullableTamperingDetails

type TouchSupport = internal.TouchSupport

var NewTouchSupport = internal.NewTouchSupport

var NewTouchSupportWithDefaults = internal.NewTouchSupportWithDefaults

type NullableTouchSupport = internal.NullableTouchSupport

var NewNullableTouchSupport = internal.NewNullableTouchSupport

type VPNConfidence = internal.VPNConfidence

const (
	VPNConfidenceLow    = internal.VPNConfidenceLow
	VPNConfidenceMedium = internal.VPNConfidenceMedium
	VPNConfidenceHigh   = internal.VPNConfidenceHigh
)

var AllowedVPNConfidenceEnumValues = internal.AllowedVPNConfidenceEnumValues

var NewVPNConfidenceFromValue = internal.NewVPNConfidenceFromValue

type NullableVPNConfidence = internal.NullableVPNConfidence

var NewNullableVPNConfidence = internal.NewNullableVPNConfidence

type VPNMethods = internal.VPNMethods

var NewVPNMethods = internal.NewVPNMethods

var NewVPNMethodsWithDefaults = internal.NewVPNMethodsWithDefaults

type NullableVPNMethods = internal.NullableVPNMethods

var NewNullableVPNMethods = internal.NewNullableVPNMethods

type Velocity = internal.Velocity

var NewVelocity = internal.NewVelocity

var NewVelocityWithDefaults = internal.NewVelocityWithDefaults

type NullableVelocity = internal.NullableVelocity

var NewNullableVelocity = internal.NewNullableVelocity

type VelocityData = internal.VelocityData

var NewVelocityData = internal.NewVelocityData

var NewVelocityDataWithDefaults = internal.NewVelocityDataWithDefaults

type NullableVelocityData = internal.NullableVelocityData

var NewNullableVelocityData = internal.NewNullableVelocityData

type WebGlBasics = internal.WebGlBasics

var NewWebGlBasics = internal.NewWebGlBasics

var NewWebGlBasicsWithDefaults = internal.NewWebGlBasicsWithDefaults

type NullableWebGlBasics = internal.NullableWebGlBasics

var NewNullableWebGlBasics = internal.NewNullableWebGlBasics

type WebGlExtensions = internal.WebGlExtensions

var NewWebGlExtensions = internal.NewWebGlExtensions

var NewWebGlExtensionsWithDefaults = internal.NewWebGlExtensionsWithDefaults

type NullableWebGlExtensions = internal.NullableWebGlExtensions

var NewNullableWebGlExtensions = internal.NewNullableWebGlExtensions
