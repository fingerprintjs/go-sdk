package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk"
	sdk "github.com/fingerprintjs/go-sdk/sdk"
	"github.com/stretchr/testify/assert"
)

func TestGetEvent(t *testing.T) {
	t.Run("Returns event", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.Event]("../test/mocks/get_event_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")

			apiKey := r.Header.Get("Authorization")
			assert.Equal(t, apiKey, "Bearer api_key")

			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(mockResponse)

			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		event, _, err := client.GetEvent(context.Background(), "123")

		assert.Nil(t, err)
		assert.NotNil(t, event)
		assert.Equal(t, *event, mockResponse)

	})

	//t.Run("Returns event with errors in all signals", func(t *testing.T) {
	//	mockResponse := GetMockResponse[sdk.Event]("../test/mocks/get_event_200_all_errors.json")
	//
	//	ts := httptest.NewServer(http.HandlerFunc(func(
	//		w http.ResponseWriter,
	//		r *http.Request,
	//	) {
	//		integrationInfo := r.URL.Query().Get("ii")
	//		assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
	//		assert.Equal(t, r.URL.Path, "/events/123")
	//
	//		apiKey := r.Header.Get("Auth-Api-Key")
	//		assert.Equal(t, apiKey, "api_key")
	//
	//		w.Header().Set("Content-Type", "application/json")
	//
	//		err := json.NewEncoder(w).Encode(mockResponse)
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}))
	//	defer ts.Close()
	//
	//	client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
	//
	//	event, _, err := client.GetEvent(context.Background(), "123")
	//
	//	assert.Nil(t, err)
	//	assert.NotNil(t, event)
	//	assert.Equal(t, *event, mockResponse)
	//
	//})
	//
	//t.Run("Returns event with unexpected fields", func(t *testing.T) {
	//	mockResponse := GetMockResponse[sdk.Event]("../test/mocks/get_event_200_extra_fields.json")
	//
	//	ts := httptest.NewServer(http.HandlerFunc(func(
	//		w http.ResponseWriter,
	//		r *http.Request,
	//	) {
	//		integrationInfo := r.URL.Query().Get("ii")
	//		assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
	//		assert.Equal(t, r.URL.Path, "/events/123")
	//
	//		apiKey := r.Header.Get("Auth-Api-Key")
	//		assert.Equal(t, apiKey, "api_key")
	//
	//		w.Header().Set("Content-Type", "application/json")
	//
	//		err := json.NewEncoder(w).Encode(mockResponse)
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}))
	//	defer ts.Close()
	//
	//	client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
	//
	//	event, _, err := client.GetEvent(context.Background(), "123")
	//
	//	assert.Nil(t, err)
	//	assert.NotNil(t, event)
	//	assert.Equal(t, *event, mockResponse)
	//})
	//
	//t.Run("Returns ErrorResponse on 403 error", func(t *testing.T) {
	//	mockResponse := GetMockResponse[sdk.ErrorResponse]("../test/mocks/errors/403_subscription_not_active.json")
	//
	//	ts := httptest.NewServer(http.HandlerFunc(func(
	//		w http.ResponseWriter,
	//		r *http.Request,
	//	) {
	//		integrationInfo := r.URL.Query().Get("ii")
	//		assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
	//		assert.Equal(t, r.URL.Path, "/events/123")
	//
	//		apiKey := r.Header.Get("Auth-Api-Key")
	//		assert.Equal(t, apiKey, "api_key")
	//
	//		w.Header().Set("Content-Type", "application/json")
	//
	//		w.WriteHeader(403)
	//
	//		err := json.NewEncoder(w).Encode(mockResponse)
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//	}))
	//	defer ts.Close()
	//
	//	client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
	//
	//	event, _, err := client.GetEvent(context.Background(), "123")
	//
	//	assert.Error(t, err)
	//	assert.IsType(t, err, &sdk.ErrorResponse{})
	//	assert.NotNil(t, event)
	//
	//	assert.Equal(t, err.Error(), mockResponse.Error.Message)
	//})
	//
	//t.Run("Return too many requests error in all fields", func(t *testing.T) {
	//	mockResponse := GetMockResponse[sdk.Event]("../test/mocks/get_event_200_too_many_requests_error.json")
	//
	//	ts := httptest.NewServer(http.HandlerFunc(func(
	//		w http.ResponseWriter,
	//		r *http.Request,
	//	) {
	//		integrationInfo := r.URL.Query().Get("ii")
	//		assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
	//		assert.Equal(t, r.URL.Path, "/events/123")
	//
	//		apiKey := r.Header.Get("Auth-Api-Key")
	//		assert.Equal(t, apiKey, "api_key")
	//
	//		w.Header().Set("Content-Type", "application/json")
	//
	//		err := json.NewEncoder(w).Encode(mockResponse)
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}))
	//	defer ts.Close()
	//
	//	client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
	//
	//	event, _, err := client.GetEvent(context.Background(), "123")
	//
	//	assert.Nil(t, err)
	//	assert.NotNil(t, event)
	//	assert.Equal(t, *event, mockResponse)
	//})

	//t.Run("Returns botd error", func(t *testing.T) {
	//	mockResponse := GetMockResponse[sdk.Event]("../test/mocks/get_event_200_botd_failed_error.json")
	//
	//	ts := httptest.NewServer(http.HandlerFunc(func(
	//		w http.ResponseWriter,
	//		r *http.Request,
	//	) {
	//		integrationInfo := r.URL.Query().Get("ii")
	//		assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
	//		assert.Equal(t, r.URL.Path, "/events/123")
	//
	//		apiKey := r.Header.Get("Auth-Api-Key")
	//		assert.Equal(t, apiKey, "api_key")
	//
	//		w.Header().Set("Content-Type", "application/json")
	//
	//		err := json.NewEncoder(w).Encode(mockResponse)
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}))
	//	defer ts.Close()
	//
	//	client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
	//
	//	res, _, err := client.GetEvent(context.Background(), "123")
	//
	//	assert.Nil(t, err)
	//	assert.NotNil(t, res)
	//	assert.Equal(t, res, mockResponse)
	//	assert.Equal(t, res.Bot.Error.Code, sdk.ErrorCodeFailed)
	//
	//})

	//	t.Run("Returns identification error", func(t *testing.T) {
	//		mockResponse := GetMockResponse[sdk.Event]("../test/mocks/get_event_200_identification_failed_error.json")
	//
	//		ts := httptest.NewServer(http.HandlerFunc(func(
	//			w http.ResponseWriter,
	//			r *http.Request,
	//		) {
	//			integrationInfo := r.URL.Query().Get("ii")
	//			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
	//			assert.Equal(t, r.URL.Path, "/events/123")
	//
	//			apiKey := r.Header.Get("Auth-Api-Key")
	//			assert.Equal(t, apiKey, "api_key")
	//
	//			w.Header().Set("Content-Type", "application/json")
	//
	//			err := json.NewEncoder(w).Encode(mockResponse)
	//
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//		}))
	//		defer ts.Close()
	//
	//		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
	//
	//		event, _, err := client.GetEvent(context.Background(), "123")
	//
	//		assert.Nil(t, err)
	//		assert.NotNil(t, event)
	//		assert.Equal(t, *event, mockResponse)
	//		assert.Equal(t, err.Error(), sdk.ErrorCodeFailed)
	//
	//	})
	//
	//	t.Run("Returns response when JSON field type is not matching response schema", func(t *testing.T) {
	//		// Changed fields: products.identification.data.incognito bool -> int
	//		malformedResponse := `{
	//  "products": {
	//    "identification": {
	//      "data": {
	//        "visitorId": "Ibk1527CUFmcnjLwIs4A9",
	//        "requestId": "0KSh65EnVoB85JBmloQK",
	//
	//        "incognito": 1,
	//
	//        "linkedId": "somelinkedId",
	//        "tag": {},
	//        "time": "2019-05-21T16:40:13Z",
	//        "timestamp": 1582299576512,
	//        "url": "https://www.example.com/login",
	//        "ip": "61.127.217.15",
	//        "ipLocation": {
	//          "accuracyRadius": 10,
	//          "latitude": 49.982,
	//          "longitude": 36.2566,
	//          "postalCode": "61202",
	//          "timezone": "Europe/Dusseldorf",
	//          "city": {
	//            "name": "Dusseldorf"
	//          },
	//          "continent": {
	//            "code": "EU",
	//            "name": "Europe"
	//          },
	//          "country": {
	//            "code": "DE",
	//            "name": "Germany"
	//          },
	//          "subdivisions": [
	//            {
	//              "isoCode": "63",
	//              "name": "North Rhine-Westphalia"
	//            }
	//          ]
	//        },
	//        "browserDetails": {
	//          "browserName": "Chrome",
	//          "browserMajorVersion": "74",
	//          "browserFullVersion": "74.0.3729",
	//          "os": "Windows",
	//          "osVersion": "7",
	//          "device": "Other",
	//          "userAgent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64) ...."
	//        },
	//        "confidence": {
	//          "score": 0.97,
	//          "revision": "v1.1"
	//        },
	//        "visitorFound": true,
	//        "firstSeenAt": {
	//          "global": "2022-03-16T11:26:45.362Z",
	//          "subscription": "2022-03-16T11:31:01.101Z"
	//        },
	//        "lastSeenAt": {
	//          "global": "2022-03-16T11:28:34.023Z",
	//          "subscription": null
	//        }
	//      }
	//    },
	//    "botd": {
	//      "data": {
	//        "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36",
	//
	//        "requestId": "1708102555327.NLOjmg",
	//
	//        "bot": {
	//          "result": "notDetected"
	//        },
	//        "url": "https://www.example.com/login",
	//        "ip": "61.127.217.15",
	//        "time": "2019-05-21T16:40:13Z"
	//      }
	//    }
	//  }
	//}
	//`
	//
	//		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//			parseErr := r.ParseForm()
	//			assert.Nil(t, parseErr)
	//
	//			apiKey := r.Header.Get("Auth-Api-Key")
	//			assert.Equal(t, apiKey, "api_key")
	//
	//			w.Header().Set("Content-Type", "application/json")
	//
	//			_, err := w.Write([]byte(malformedResponse))
	//
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//		}))
	//		defer ts.Close()
	//
	//		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
	//
	//		event, _, err := client.GetEvent(context.Background(), "request_id")
	//		assert.Error(t, err)
	//		assert.NotNil(t, event)
	//		assert.Equal(t, event.Incognito, false)
	//	})

	//	t.Run("Returns response when URL in JSON is not RFC compilant", func(t *testing.T) {
	//		// Changed fields: products.botd.data.url -> not RFC compliant URL
	//		malformedResponse := `{
	//  "products": {
	//    "identification": {
	//      "data": {
	//        "visitorId": "Ibk1527CUFmcnjLwIs4A9",
	//        "requestId": "0KSh65EnVoB85JBmloQK",
	//        "incognito": true,
	//        "linkedId": "somelinkedId",
	//        "tag": {},
	//        "time": "2019-05-21T16:40:13Z",
	//        "timestamp": 1582299576512,
	//        "url": "https://www.example.com/login",
	//        "ip": "61.127.217.15",
	//        "ipLocation": {
	//          "accuracyRadius": 10,
	//          "latitude": 49.982,
	//          "longitude": 36.2566,
	//          "postalCode": "61202",
	//          "timezone": "Europe/Dusseldorf",
	//          "city": {
	//            "name": "Dusseldorf"
	//          },
	//          "continent": {
	//            "code": "EU",
	//            "name": "Europe"
	//          },
	//          "country": {
	//            "code": "DE",
	//            "name": "Germany"
	//          },
	//          "subdivisions": [
	//            {
	//              "isoCode": "63",
	//              "name": "North Rhine-Westphalia"
	//            }
	//          ]
	//        },
	//        "browserDetails": {
	//          "browserName": "Chrome",
	//          "browserMajorVersion": "74",
	//          "browserFullVersion": "74.0.3729",
	//          "os": "Windows",
	//          "osVersion": "7",
	//          "device": "Other",
	//          "userAgent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64) ...."
	//        },
	//        "confidence": {
	//          "score": 0.97,
	//          "revision": "v1.1"
	//        },
	//        "visitorFound": true,
	//        "firstSeenAt": {
	//          "global": "2022-03-16T11:26:45.362Z",
	//          "subscription": "2022-03-16T11:31:01.101Z"
	//        },
	//        "lastSeenAt": {
	//          "global": "2022-03-16T11:28:34.023Z",
	//          "subscription": null
	//        }
	//      }
	//    },
	//    "botd": {
	//      "data": {
	//        "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36",
	//
	//        "requestId": "1708102555327.NLOjmg",
	//
	//        "bot": {
	//          "result": "notDetected"
	//        },
	//        "url": "https://www.example.com/{{{login",
	//        "ip": "61.127.217.15",
	//        "time": "2019-05-21T16:40:13Z"
	//      }
	//    }
	//  }
	//}
	//`
	//
	//		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//			parseErr := r.ParseForm()
	//			assert.Nil(t, parseErr)
	//
	//			apiKey := r.Header.Get("Auth-Api-Key")
	//			assert.Equal(t, apiKey, "api_key")
	//
	//			w.Header().Set("Content-Type", "application/json")
	//
	//			_, err := w.Write([]byte(malformedResponse))
	//
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//		}))
	//		defer ts.Close()
	//
	//		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
	//
	//		res, _, err := client.GetEvent(context.Background(), "request_id")
	//		assert.Nil(t, err)
	//		assert.NotNil(t, res)
	//		assert.Equal(t, res.Bot.Url, "https://www.example.com/{{{login")
	//	})

}
