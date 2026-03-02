package fingerprint

import (
	"github.com/fingerprintjs/go-sdk/v8/internal/openapi"
)

type BotInfo = openapi.BotInfo

var NewBotInfo = openapi.NewBotInfo

var NewBotInfoWithDefaults = openapi.NewBotInfoWithDefaults

type NullableBotInfo = openapi.NullableBotInfo

var NewNullableBotInfo = openapi.NewNullableBotInfo

type BotResult = openapi.BotResult

const (
	BotResultBad          = openapi.BotResultBad
	BotResultGood         = openapi.BotResultGood
	BotResultNot_detected = openapi.BotResultNot_detected
)

var AllowedBotResultEnumValues = openapi.AllowedBotResultEnumValues

var NewBotResultFromValue = openapi.NewBotResultFromValue

type NullableBotResult = openapi.NullableBotResult

var NewNullableBotResult = openapi.NewNullableBotResult

type BrowserDetails = openapi.BrowserDetails

var NewBrowserDetails = openapi.NewBrowserDetails

var NewBrowserDetailsWithDefaults = openapi.NewBrowserDetailsWithDefaults

type NullableBrowserDetails = openapi.NullableBrowserDetails

var NewNullableBrowserDetails = openapi.NewNullableBrowserDetails

type Canvas = openapi.Canvas

var NewCanvas = openapi.NewCanvas

var NewCanvasWithDefaults = openapi.NewCanvasWithDefaults

type NullableCanvas = openapi.NullableCanvas

var NewNullableCanvas = openapi.NewNullableCanvas

type Emoji = openapi.Emoji

var NewEmoji = openapi.NewEmoji

var NewEmojiWithDefaults = openapi.NewEmojiWithDefaults

type NullableEmoji = openapi.NullableEmoji

var NewNullableEmoji = openapi.NewNullableEmoji

type Error = openapi.Error

var NewError = openapi.NewError

var NewErrorWithDefaults = openapi.NewErrorWithDefaults

type NullableError = openapi.NullableError

var NewNullableError = openapi.NewNullableError

type ErrorCode = openapi.ErrorCode

const (
	ErrorCodeRequest_cannot_be_parsed = openapi.ErrorCodeRequest_cannot_be_parsed
	ErrorCodeSecret_api_key_required  = openapi.ErrorCodeSecret_api_key_required
	ErrorCodeSecret_api_key_not_found = openapi.ErrorCodeSecret_api_key_not_found
	ErrorCodePublic_api_key_required  = openapi.ErrorCodePublic_api_key_required
	ErrorCodePublic_api_key_not_found = openapi.ErrorCodePublic_api_key_not_found
	ErrorCodeSubscription_not_active  = openapi.ErrorCodeSubscription_not_active
	ErrorCodeWrong_region             = openapi.ErrorCodeWrong_region
	ErrorCodeFeature_not_enabled      = openapi.ErrorCodeFeature_not_enabled
	ErrorCodeRequest_not_found        = openapi.ErrorCodeRequest_not_found
	ErrorCodeVisitor_not_found        = openapi.ErrorCodeVisitor_not_found
	ErrorCodeToo_many_requests        = openapi.ErrorCodeToo_many_requests
	ErrorCodeState_not_ready          = openapi.ErrorCodeState_not_ready
	ErrorCodeFailed                   = openapi.ErrorCodeFailed
	ErrorCodeEvent_not_found          = openapi.ErrorCodeEvent_not_found
	ErrorCodeMissing_module           = openapi.ErrorCodeMissing_module
	ErrorCodePayload_too_large        = openapi.ErrorCodePayload_too_large
	ErrorCodeService_unavailable      = openapi.ErrorCodeService_unavailable
	ErrorCodeRuleset_not_found        = openapi.ErrorCodeRuleset_not_found
)

var AllowedErrorCodeEnumValues = openapi.AllowedErrorCodeEnumValues

var NewErrorCodeFromValue = openapi.NewErrorCodeFromValue

type NullableErrorCode = openapi.NullableErrorCode

var NewNullableErrorCode = openapi.NewNullableErrorCode

type ErrorResponse = openapi.ErrorResponse

var NewErrorResponse = openapi.NewErrorResponse

var NewErrorResponseWithDefaults = openapi.NewErrorResponseWithDefaults

type NullableErrorResponse = openapi.NullableErrorResponse

var NewNullableErrorResponse = openapi.NewNullableErrorResponse

type Event = openapi.Event

var NewEvent = openapi.NewEvent

var NewEventWithDefaults = openapi.NewEventWithDefaults

type NullableEvent = openapi.NullableEvent

var NewNullableEvent = openapi.NewNullableEvent

type EventRuleAction = openapi.EventRuleAction

type NullableEventRuleAction = openapi.NullableEventRuleAction

var NewNullableEventRuleAction = openapi.NewNullableEventRuleAction

type EventRuleActionAllow = openapi.EventRuleActionAllow

var NewEventRuleActionAllow = openapi.NewEventRuleActionAllow

var NewEventRuleActionAllowWithDefaults = openapi.NewEventRuleActionAllowWithDefaults

type NullableEventRuleActionAllow = openapi.NullableEventRuleActionAllow

var NewNullableEventRuleActionAllow = openapi.NewNullableEventRuleActionAllow

type EventRuleActionBlock = openapi.EventRuleActionBlock

var NewEventRuleActionBlock = openapi.NewEventRuleActionBlock

var NewEventRuleActionBlockWithDefaults = openapi.NewEventRuleActionBlockWithDefaults

type NullableEventRuleActionBlock = openapi.NullableEventRuleActionBlock

var NewNullableEventRuleActionBlock = openapi.NewNullableEventRuleActionBlock

type EventSearch = openapi.EventSearch

var NewEventSearch = openapi.NewEventSearch

var NewEventSearchWithDefaults = openapi.NewEventSearchWithDefaults

type NullableEventSearch = openapi.NullableEventSearch

var NewNullableEventSearch = openapi.NewNullableEventSearch

type EventUpdate = openapi.EventUpdate

var NewEventUpdate = openapi.NewEventUpdate

var NewEventUpdateWithDefaults = openapi.NewEventUpdateWithDefaults

type NullableEventUpdate = openapi.NullableEventUpdate

var NewNullableEventUpdate = openapi.NewNullableEventUpdate

type FontPreferences = openapi.FontPreferences

var NewFontPreferences = openapi.NewFontPreferences

var NewFontPreferencesWithDefaults = openapi.NewFontPreferencesWithDefaults

type NullableFontPreferences = openapi.NullableFontPreferences

var NewNullableFontPreferences = openapi.NewNullableFontPreferences

type Geolocation = openapi.Geolocation

var NewGeolocation = openapi.NewGeolocation

var NewGeolocationWithDefaults = openapi.NewGeolocationWithDefaults

type NullableGeolocation = openapi.NullableGeolocation

var NewNullableGeolocation = openapi.NewNullableGeolocation

type GeolocationSubdivisionsInner = openapi.GeolocationSubdivisionsInner

var NewGeolocationSubdivisionsInner = openapi.NewGeolocationSubdivisionsInner

var NewGeolocationSubdivisionsInnerWithDefaults = openapi.NewGeolocationSubdivisionsInnerWithDefaults

type NullableGeolocationSubdivisionsInner = openapi.NullableGeolocationSubdivisionsInner

var NewNullableGeolocationSubdivisionsInner = openapi.NewNullableGeolocationSubdivisionsInner

type IPBlockList = openapi.IPBlockList

var NewIPBlockList = openapi.NewIPBlockList

var NewIPBlockListWithDefaults = openapi.NewIPBlockListWithDefaults

type NullableIPBlockList = openapi.NullableIPBlockList

var NewNullableIPBlockList = openapi.NewNullableIPBlockList

type IPInfo = openapi.IPInfo

var NewIPInfo = openapi.NewIPInfo

var NewIPInfoWithDefaults = openapi.NewIPInfoWithDefaults

type NullableIPInfo = openapi.NullableIPInfo

var NewNullableIPInfo = openapi.NewNullableIPInfo

type IPInfoV4 = openapi.IPInfoV4

var NewIPInfoV4 = openapi.NewIPInfoV4

var NewIPInfoV4WithDefaults = openapi.NewIPInfoV4WithDefaults

type NullableIPInfoV4 = openapi.NullableIPInfoV4

var NewNullableIPInfoV4 = openapi.NewNullableIPInfoV4

type IPInfoV6 = openapi.IPInfoV6

var NewIPInfoV6 = openapi.NewIPInfoV6

var NewIPInfoV6WithDefaults = openapi.NewIPInfoV6WithDefaults

type NullableIPInfoV6 = openapi.NullableIPInfoV6

var NewNullableIPInfoV6 = openapi.NewNullableIPInfoV6

type Identification = openapi.Identification

var NewIdentification = openapi.NewIdentification

var NewIdentificationWithDefaults = openapi.NewIdentificationWithDefaults

type NullableIdentification = openapi.NullableIdentification

var NewNullableIdentification = openapi.NewNullableIdentification

type IdentificationConfidence = openapi.IdentificationConfidence

var NewIdentificationConfidence = openapi.NewIdentificationConfidence

var NewIdentificationConfidenceWithDefaults = openapi.NewIdentificationConfidenceWithDefaults

type NullableIdentificationConfidence = openapi.NullableIdentificationConfidence

var NewNullableIdentificationConfidence = openapi.NewNullableIdentificationConfidence

type Integration = openapi.Integration

var NewIntegration = openapi.NewIntegration

var NewIntegrationWithDefaults = openapi.NewIntegrationWithDefaults

type NullableIntegration = openapi.NullableIntegration

var NewNullableIntegration = openapi.NewNullableIntegration

type IntegrationSubintegration = openapi.IntegrationSubintegration

var NewIntegrationSubintegration = openapi.NewIntegrationSubintegration

var NewIntegrationSubintegrationWithDefaults = openapi.NewIntegrationSubintegrationWithDefaults

type NullableIntegrationSubintegration = openapi.NullableIntegrationSubintegration

var NewNullableIntegrationSubintegration = openapi.NewNullableIntegrationSubintegration

type PluginsInner = openapi.PluginsInner

var NewPluginsInner = openapi.NewPluginsInner

var NewPluginsInnerWithDefaults = openapi.NewPluginsInnerWithDefaults

type NullablePluginsInner = openapi.NullablePluginsInner

var NewNullablePluginsInner = openapi.NewNullablePluginsInner

type PluginsInnerMimeTypesInner = openapi.PluginsInnerMimeTypesInner

var NewPluginsInnerMimeTypesInner = openapi.NewPluginsInnerMimeTypesInner

var NewPluginsInnerMimeTypesInnerWithDefaults = openapi.NewPluginsInnerMimeTypesInnerWithDefaults

type NullablePluginsInnerMimeTypesInner = openapi.NullablePluginsInnerMimeTypesInner

var NewNullablePluginsInnerMimeTypesInner = openapi.NewNullablePluginsInnerMimeTypesInner

type Proximity = openapi.Proximity

var NewProximity = openapi.NewProximity

var NewProximityWithDefaults = openapi.NewProximityWithDefaults

type NullableProximity = openapi.NullableProximity

var NewNullableProximity = openapi.NewNullableProximity

type ProxyConfidence = openapi.ProxyConfidence

const (
	ProxyConfidenceLow    = openapi.ProxyConfidenceLow
	ProxyConfidenceMedium = openapi.ProxyConfidenceMedium
	ProxyConfidenceHigh   = openapi.ProxyConfidenceHigh
)

var AllowedProxyConfidenceEnumValues = openapi.AllowedProxyConfidenceEnumValues

var NewProxyConfidenceFromValue = openapi.NewProxyConfidenceFromValue

type NullableProxyConfidence = openapi.NullableProxyConfidence

var NewNullableProxyConfidence = openapi.NewNullableProxyConfidence

type ProxyDetails = openapi.ProxyDetails

var NewProxyDetails = openapi.NewProxyDetails

var NewProxyDetailsWithDefaults = openapi.NewProxyDetailsWithDefaults

type NullableProxyDetails = openapi.NullableProxyDetails

var NewNullableProxyDetails = openapi.NewNullableProxyDetails

type RawDeviceAttributes = openapi.RawDeviceAttributes

var NewRawDeviceAttributes = openapi.NewRawDeviceAttributes

var NewRawDeviceAttributesWithDefaults = openapi.NewRawDeviceAttributesWithDefaults

type NullableRawDeviceAttributes = openapi.NullableRawDeviceAttributes

var NewNullableRawDeviceAttributes = openapi.NewNullableRawDeviceAttributes

type RequestHeaderModifications = openapi.RequestHeaderModifications

var NewRequestHeaderModifications = openapi.NewRequestHeaderModifications

var NewRequestHeaderModificationsWithDefaults = openapi.NewRequestHeaderModificationsWithDefaults

type NullableRequestHeaderModifications = openapi.NullableRequestHeaderModifications

var NewNullableRequestHeaderModifications = openapi.NewNullableRequestHeaderModifications

type RuleActionHeaderField = openapi.RuleActionHeaderField

var NewRuleActionHeaderField = openapi.NewRuleActionHeaderField

var NewRuleActionHeaderFieldWithDefaults = openapi.NewRuleActionHeaderFieldWithDefaults

type NullableRuleActionHeaderField = openapi.NullableRuleActionHeaderField

var NewNullableRuleActionHeaderField = openapi.NewNullableRuleActionHeaderField

type RuleActionType = openapi.RuleActionType

const (
	RuleActionTypeAllow = openapi.RuleActionTypeAllow
	RuleActionTypeBlock = openapi.RuleActionTypeBlock
)

var AllowedRuleActionTypeEnumValues = openapi.AllowedRuleActionTypeEnumValues

var NewRuleActionTypeFromValue = openapi.NewRuleActionTypeFromValue

type NullableRuleActionType = openapi.NullableRuleActionType

var NewNullableRuleActionType = openapi.NewNullableRuleActionType

type SDK = openapi.SDK

var NewSDK = openapi.NewSDK

var NewSDKWithDefaults = openapi.NewSDKWithDefaults

type NullableSDK = openapi.NullableSDK

var NewNullableSDK = openapi.NewNullableSDK

type SearchEventsBot = openapi.SearchEventsBot

const (
	SearchEventsBotAll  = openapi.SearchEventsBotAll
	SearchEventsBotGood = openapi.SearchEventsBotGood
	SearchEventsBotBad  = openapi.SearchEventsBotBad
	SearchEventsBotNone = openapi.SearchEventsBotNone
)

var AllowedSearchEventsBotEnumValues = openapi.AllowedSearchEventsBotEnumValues

var NewSearchEventsBotFromValue = openapi.NewSearchEventsBotFromValue

type NullableSearchEventsBot = openapi.NullableSearchEventsBot

var NewNullableSearchEventsBot = openapi.NewNullableSearchEventsBot

type SearchEventsSDKPlatform = openapi.SearchEventsSDKPlatform

const (
	SearchEventsSDKPlatformJs      = openapi.SearchEventsSDKPlatformJs
	SearchEventsSDKPlatformAndroid = openapi.SearchEventsSDKPlatformAndroid
	SearchEventsSDKPlatformIos     = openapi.SearchEventsSDKPlatformIos
)

var AllowedSearchEventsSDKPlatformEnumValues = openapi.AllowedSearchEventsSDKPlatformEnumValues

var NewSearchEventsSDKPlatformFromValue = openapi.NewSearchEventsSDKPlatformFromValue

type NullableSearchEventsSDKPlatform = openapi.NullableSearchEventsSDKPlatform

var NewNullableSearchEventsSDKPlatform = openapi.NewNullableSearchEventsSDKPlatform

type SearchEventsVPNConfidence = openapi.SearchEventsVPNConfidence

const (
	SearchEventsVPNConfidenceHigh   = openapi.SearchEventsVPNConfidenceHigh
	SearchEventsVPNConfidenceMedium = openapi.SearchEventsVPNConfidenceMedium
	SearchEventsVPNConfidenceLow    = openapi.SearchEventsVPNConfidenceLow
)

var AllowedSearchEventsVPNConfidenceEnumValues = openapi.AllowedSearchEventsVPNConfidenceEnumValues

var NewSearchEventsVPNConfidenceFromValue = openapi.NewSearchEventsVPNConfidenceFromValue

type NullableSearchEventsVPNConfidence = openapi.NullableSearchEventsVPNConfidence

var NewNullableSearchEventsVPNConfidence = openapi.NewNullableSearchEventsVPNConfidence

type SupplementaryIDHighRecall = openapi.SupplementaryIDHighRecall

var NewSupplementaryIDHighRecall = openapi.NewSupplementaryIDHighRecall

var NewSupplementaryIDHighRecallWithDefaults = openapi.NewSupplementaryIDHighRecallWithDefaults

type NullableSupplementaryIDHighRecall = openapi.NullableSupplementaryIDHighRecall

var NewNullableSupplementaryIDHighRecall = openapi.NewNullableSupplementaryIDHighRecall

type TamperingDetails = openapi.TamperingDetails

var NewTamperingDetails = openapi.NewTamperingDetails

var NewTamperingDetailsWithDefaults = openapi.NewTamperingDetailsWithDefaults

type NullableTamperingDetails = openapi.NullableTamperingDetails

var NewNullableTamperingDetails = openapi.NewNullableTamperingDetails

type TouchSupport = openapi.TouchSupport

var NewTouchSupport = openapi.NewTouchSupport

var NewTouchSupportWithDefaults = openapi.NewTouchSupportWithDefaults

type NullableTouchSupport = openapi.NullableTouchSupport

var NewNullableTouchSupport = openapi.NewNullableTouchSupport

type VPNConfidence = openapi.VPNConfidence

const (
	VPNConfidenceLow    = openapi.VPNConfidenceLow
	VPNConfidenceMedium = openapi.VPNConfidenceMedium
	VPNConfidenceHigh   = openapi.VPNConfidenceHigh
)

var AllowedVPNConfidenceEnumValues = openapi.AllowedVPNConfidenceEnumValues

var NewVPNConfidenceFromValue = openapi.NewVPNConfidenceFromValue

type NullableVPNConfidence = openapi.NullableVPNConfidence

var NewNullableVPNConfidence = openapi.NewNullableVPNConfidence

type VPNMethods = openapi.VPNMethods

var NewVPNMethods = openapi.NewVPNMethods

var NewVPNMethodsWithDefaults = openapi.NewVPNMethodsWithDefaults

type NullableVPNMethods = openapi.NullableVPNMethods

var NewNullableVPNMethods = openapi.NewNullableVPNMethods

type Velocity = openapi.Velocity

var NewVelocity = openapi.NewVelocity

var NewVelocityWithDefaults = openapi.NewVelocityWithDefaults

type NullableVelocity = openapi.NullableVelocity

var NewNullableVelocity = openapi.NewNullableVelocity

type VelocityData = openapi.VelocityData

var NewVelocityData = openapi.NewVelocityData

var NewVelocityDataWithDefaults = openapi.NewVelocityDataWithDefaults

type NullableVelocityData = openapi.NullableVelocityData

var NewNullableVelocityData = openapi.NewNullableVelocityData

type WebGlBasics = openapi.WebGlBasics

var NewWebGlBasics = openapi.NewWebGlBasics

var NewWebGlBasicsWithDefaults = openapi.NewWebGlBasicsWithDefaults

type NullableWebGlBasics = openapi.NullableWebGlBasics

var NewNullableWebGlBasics = openapi.NewNullableWebGlBasics

type WebGlExtensions = openapi.WebGlExtensions

var NewWebGlExtensions = openapi.NewWebGlExtensions

var NewWebGlExtensionsWithDefaults = openapi.NewWebGlExtensionsWithDefaults

type NullableWebGlExtensions = openapi.NullableWebGlExtensions

var NewNullableWebGlExtensions = openapi.NewNullableWebGlExtensions
