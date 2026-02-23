package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
)

func TestSearchEvents(t *testing.T) {
	t.Run("Search with just limit", func(t *testing.T) {
		mockResponse := GetMockResponse[fingerprint.EventSearch]("mocks/get_event_search_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))

			assertAuthorizationHeader(t, r, "api_key")

			assert.Equal(t, "/events", r.URL.Path)
			assert.Equal(t, "2", r.URL.Query().Get("limit"))
			assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 2)

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(mockResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		eventSearch, _, err := client.SearchEvents(client.NewSearchEventsRequest(context.Background()).Limit(2))
		assert.Nil(t, err)
		assert.NotNil(t, eventSearch)
		assert.Equal(t, mockResponse, *eventSearch)
	})

	t.Run("Search with partial params", func(t *testing.T) {
		mockResponse := GetMockResponse[fingerprint.EventSearch]("mocks/get_event_search_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			integrationInfo := query.Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))

			assertAuthorizationHeader(t, r, "api_key")

			assert.Equal(t, "/events", r.URL.Path)
			assert.Equal(t, "10", query.Get("limit"), "limit")
			assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 7)
			assert.Equal(t, "", query.Get("suspect"), "suspect")
			assert.False(t, query.Has("suspect"), "has suspect")
			assert.Equal(t, "", query.Get("bot"), "bot")
			assert.False(t, query.Has("bot"), "has bot")
			assert.Equal(t, "10", query.Get("end"), "end")
			assert.Equal(t, "5", query.Get("start"), "start")
			assert.Equal(t, "", query.Get("ip_address"), "ip_address")
			assert.False(t, query.Has("ip_address"), "has ip_address")
			assert.Equal(t, "linked_id", query.Get("linked_id"), "linked_id")
			assert.Equal(t, "false", query.Get("reverse"), "reverse")
			assert.Equal(t, "XIkiQhRyp7edU9SA0jBb", query.Get("visitor_id"), "visitor_id")

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(mockResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		var end int64 = 10
		var start int64 = 5
		reverse := false
		linkedId := "linked_id"
		visitorId := "XIkiQhRyp7edU9SA0jBb"
		opts := client.NewSearchEventsRequest(context.Background()).
			End(end).
			Start(start).
			LinkedId(linkedId).
			Reverse(reverse).
			VisitorId(visitorId)

		eventSearch, _, err := client.SearchEvents(opts)
		assert.Nil(t, err)
		assert.NotNil(t, eventSearch)
		assert.Equal(t, mockResponse, *eventSearch)
	})

	t.Run("Search with all params", func(t *testing.T) {
		mockResponse := GetMockResponse[fingerprint.EventSearch]("mocks/get_event_search_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			integrationInfo := query.Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))

			assertAuthorizationHeader(t, r, "api_key")

			assert.Equal(t, "/events", r.URL.Path)

			assert.Equal(t, "2", query.Get("limit"), "limit")
			assert.Equal(t, "XIkiQhRyp7edU9SA0jBb", query.Get("visitor_id"), "visitorId")
			assert.Equal(t, "good", query.Get("bot"), "bot")
			assert.Equal(t, "127.0.0.1", query.Get("ip_address"), "ipAddress")
			assert.Equal(t, "ASN 20", query.Get("asn"), "asn")
			assert.Equal(t, "linked_id", query.Get("linked_id"), "linkedId")
			assert.Equal(t, "https://example.com", query.Get("url"), "url")
			assert.Equal(t, "bundleId", query.Get("bundle_id"), "bundleId")
			assert.Equal(t, "com.example.app", query.Get("package_name"), "packageName")
			assert.Equal(t, "https://example.com", query.Get("origin"), "origin")
			assert.Equal(t, "5", query.Get("start"), "start")
			assert.Equal(t, "10", query.Get("end"), "end")
			assert.Equal(t, "testSdkVersion", query.Get("sdk_version"), "sdkVersion")
			assert.Equal(t, string(fingerprint.SearchEventsSdkPlatformJs), query.Get("sdk_platform"), "sdkPlatform")
			assert.Equal(t, []string{"env1", "env2"}, r.URL.Query()["environment"], "environment")
			assert.Equal(t, "testProximityId", query.Get("proximity_id"), "proximityId")
			assert.Equal(t, "10", query.Get("total_hits"), "totalHits")
			assert.Equal(t, "85.5", query.Get("min_suspect_score"), "minSuspectScore")

			assert.Equal(t, "false", query.Get("reverse"), "reverse")
			assert.True(t, query.Has("reverse"), "has reverse")

			assert.Equal(t, "true", query.Get("suspect"), "suspect")
			assert.True(t, query.Has("suspect"), "has suspect")

			assert.Equal(t, "false", query.Get("vpn"), "vpn")
			assert.True(t, query.Has("vpn"), "has vpn")

			assert.Equal(t, "true", query.Get("virtual_machine"), "virtualMachine")
			assert.True(t, query.Has("virtual_machine"), "has virtualMachine")

			assert.Equal(t, "true", query.Get("tampering"), "tampering")
			assert.True(t, query.Has("tampering"), "has tampering")

			assert.Equal(t, "false", query.Get("anti_detect_browser"), "antiDetectBrowser")
			assert.True(t, query.Has("anti_detect_browser"), "has antiDetectBrowser")

			assert.Equal(t, "true", query.Get("incognito"), "incognito")
			assert.True(t, query.Has("incognito"), "has incognito")

			assert.Equal(t, "false", query.Get("privacy_settings"), "privacySettings")
			assert.True(t, query.Has("privacy_settings"), "has privacySettings")

			assert.Equal(t, "false", query.Get("jailbroken"), "jailbroken")
			assert.True(t, query.Has("jailbroken"), "has jailbroken")

			assert.Equal(t, "false", query.Get("frida"), "frida")
			assert.True(t, query.Has("frida"), "has frida")

			assert.Equal(t, "false", query.Get("factory_reset"), "factoryReset")
			assert.True(t, query.Has("factory_reset"), "has factoryReset")

			assert.Equal(t, "false", query.Get("cloned_app"), "clonedApp")
			assert.True(t, query.Has("cloned_app"), "has clonedApp")

			assert.Equal(t, "false", query.Get("emulator"), "emulator")
			assert.True(t, query.Has("emulator"), "has emulator")

			assert.Equal(t, "false", query.Get("root_apps"), "rootApps")
			assert.True(t, query.Has("root_apps"), "has rootApps")

			assert.Equal(t, "high", query.Get("vpn_confidence"), "vpnConfidence")
			assert.True(t, query.Has("vpn_confidence"), "has vpnConfidence")

			assert.Equal(t, "false", query.Get("developer_tools"), "developerTools")
			assert.True(t, query.Has("developer_tools"), "has developerTools")

			assert.Equal(t, "true", query.Get("location_spoofing"), "locationSpoofing")
			assert.True(t, query.Has("location_spoofing"), "has locationSpoofing")

			assert.Equal(t, "false", query.Get("mitm_attack"), "mitmAttack")
			assert.True(t, query.Has("mitm_attack"), "has mitmAttack")

			assert.Equal(t, "false", query.Get("proxy"), "proxy")
			assert.True(t, query.Has("proxy"), "has proxy")

			assert.Equal(t, "false", query.Get("tor_node"), "torNode")
			assert.True(t, query.Has("tor_node"), "has torNode")

			assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 40, "expected all parameters in query")

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(mockResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		var (
			limit             int32    = 2
			visitorId         string   = "XIkiQhRyp7edU9SA0jBb"
			bot                        = fingerprint.SearchEventsBotGood
			ipAddress         string   = "127.0.0.1"
			asn               string   = "ASN 20"
			linkedId          string   = "linked_id"
			url               string   = "https://example.com"
			bundleId          string   = "bundleId"
			packageName       string   = "com.example.app"
			origin            string   = "https://example.com"
			start             int64    = 5
			end               int64    = 10
			reverse           bool     = false
			suspect           bool     = true
			vpn               bool     = false
			virtualMachine    bool     = true
			tampering         bool     = true
			antiDetectBrowser bool     = false
			incognito         bool     = true
			privacySettings   bool     = false
			jailbroken        bool     = false
			frida             bool     = false
			factoryReset      bool     = false
			clonedApp         bool     = false
			emulator          bool     = false
			rootApps          bool     = false
			vpnConfidence              = fingerprint.SearchEventsVpnConfidenceHigh
			minSuspectScore   float32  = 85.5
			developerTools    bool     = false
			locationSpoofing  bool     = true
			mitmAttack        bool     = false
			proxy             bool     = false
			sdkVersion        string   = "testSdkVersion"
			sdkPlatform                = fingerprint.SearchEventsSdkPlatformJs
			environment       []string = []string{"env1", "env2"}
			proximityId       string   = "testProximityId"
			totalHits         int64    = 10
			torNode           bool     = false
		)

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		opts := client.NewSearchEventsRequest(context.Background()).
			AntiDetectBrowser(antiDetectBrowser).
			Asn(asn).
			Bot(bot).
			BundleId(bundleId).
			ClonedApp(clonedApp).
			DeveloperTools(developerTools).
			Emulator(emulator).
			End(end).
			Environment(environment).
			FactoryReset(factoryReset).
			Frida(frida).
			Incognito(incognito).
			IpAddress(ipAddress).
			Jailbroken(jailbroken).
			Limit(limit).
			LinkedId(linkedId).
			LocationSpoofing(locationSpoofing).
			MinSuspectScore(minSuspectScore).
			MitmAttack(mitmAttack).
			Origin(origin).
			PackageName(packageName).
			PrivacySettings(privacySettings).
			ProximityId(proximityId).
			Proxy(proxy).
			Reverse(reverse).
			RootApps(rootApps).
			SdkPlatform(sdkPlatform).
			SdkVersion(sdkVersion).
			Start(start).
			Suspect(suspect).
			Tampering(tampering).
			TorNode(torNode).
			TotalHits(totalHits).
			Url(url).
			VirtualMachine(virtualMachine).
			VisitorId(visitorId).
			Vpn(vpn).
			VpnConfidence(vpnConfidence)

		res, _, err := client.SearchEvents(opts)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockResponse, *res)
	})

	t.Run("ErrorResponse", func(t *testing.T) {
		testCases := []struct {
			StatusCode       int
			MockResponsePath string
		}{
			{400, "mocks/errors/400_ip_address_invalid.json"},
			{403, "mocks/errors/403_feature_not_enabled.json"},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("Returns ErrorResponse on %d", testCase.StatusCode), func(t *testing.T) {
				mockResponse := GetMockResponse[fingerprint.ErrorResponse](testCase.MockResponsePath)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					integrationInfo := r.URL.Query().Get("ii")
					assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))

					assertAuthorizationHeader(t, r, "api_key")

					assert.Equal(t, "/events", r.URL.Path)
					assert.Equal(t, "10", r.URL.Query().Get("limit"))
					assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 2)

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(testCase.StatusCode)
					err := json.NewEncoder(w).Encode(mockResponse)
					if err != nil {
						log.Fatal(err)
					}
				}))
				defer ts.Close()

				client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

				eventSearch, res, err := client.SearchEvents(client.NewSearchEventsRequest(context.Background()))

				assert.Nil(t, eventSearch)
				assertErrorResponse(t, testCase.StatusCode, mockResponse, res, err)
			})
		}
	})
}
