# Geolocation

    ## Properties

    Name | Type | Description | Notes
    ------------ | ------------- | ------------- | -------------
        **AccuracyRadius** | Pointer to **int32** | The IP address is likely to be within this radius (in km) of the specified location. | [optional] 
**Latitude** | Pointer to **float64** |  | [optional] 
**Longitude** | Pointer to **float64** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**CityName** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CountryName** | Pointer to **string** |  | [optional] 
**ContinentCode** | Pointer to **string** |  | [optional] 
**ContinentName** | Pointer to **string** |  | [optional] 
**Subdivisions** | Pointer to [**[]GeolocationSubdivisionsInner**](GeolocationSubdivisionsInner.md) |  | [optional] 

    ## Methods

        ### NewGeolocation

        `func NewGeolocation() *Geolocation`

        NewGeolocation instantiates a new Geolocation object
        This constructor will assign default values to properties that have it defined,
        and makes sure properties required by API are set, but the set of arguments
        will change when the set of required properties is changed

        ### NewGeolocationWithDefaults

        `func NewGeolocationWithDefaults() *Geolocation`

        NewGeolocationWithDefaults instantiates a new Geolocation object
        This constructor will only assign default values to properties that have it defined,
        but it doesn't guarantee that properties required by API are set

            ### GetAccuracyRadius

            `func (o *Geolocation) GetAccuracyRadius() int32`

            GetAccuracyRadius returns the AccuracyRadius field if non-nil, zero value otherwise.

            ### GetAccuracyRadiusOk

            `func (o *Geolocation) GetAccuracyRadiusOk() (*int32, bool)`

            GetAccuracyRadiusOk returns a tuple with the AccuracyRadius field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetAccuracyRadius

            `func (o *Geolocation) SetAccuracyRadius(v int32)`

            SetAccuracyRadius sets AccuracyRadius field to given value.

                ### HasAccuracyRadius

                `func (o *Geolocation) HasAccuracyRadius() bool`

                HasAccuracyRadius returns a boolean if a field has been set.

            ### GetLatitude

            `func (o *Geolocation) GetLatitude() float64`

            GetLatitude returns the Latitude field if non-nil, zero value otherwise.

            ### GetLatitudeOk

            `func (o *Geolocation) GetLatitudeOk() (*float64, bool)`

            GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetLatitude

            `func (o *Geolocation) SetLatitude(v float64)`

            SetLatitude sets Latitude field to given value.

                ### HasLatitude

                `func (o *Geolocation) HasLatitude() bool`

                HasLatitude returns a boolean if a field has been set.

            ### GetLongitude

            `func (o *Geolocation) GetLongitude() float64`

            GetLongitude returns the Longitude field if non-nil, zero value otherwise.

            ### GetLongitudeOk

            `func (o *Geolocation) GetLongitudeOk() (*float64, bool)`

            GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetLongitude

            `func (o *Geolocation) SetLongitude(v float64)`

            SetLongitude sets Longitude field to given value.

                ### HasLongitude

                `func (o *Geolocation) HasLongitude() bool`

                HasLongitude returns a boolean if a field has been set.

            ### GetPostalCode

            `func (o *Geolocation) GetPostalCode() string`

            GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

            ### GetPostalCodeOk

            `func (o *Geolocation) GetPostalCodeOk() (*string, bool)`

            GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetPostalCode

            `func (o *Geolocation) SetPostalCode(v string)`

            SetPostalCode sets PostalCode field to given value.

                ### HasPostalCode

                `func (o *Geolocation) HasPostalCode() bool`

                HasPostalCode returns a boolean if a field has been set.

            ### GetTimezone

            `func (o *Geolocation) GetTimezone() string`

            GetTimezone returns the Timezone field if non-nil, zero value otherwise.

            ### GetTimezoneOk

            `func (o *Geolocation) GetTimezoneOk() (*string, bool)`

            GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetTimezone

            `func (o *Geolocation) SetTimezone(v string)`

            SetTimezone sets Timezone field to given value.

                ### HasTimezone

                `func (o *Geolocation) HasTimezone() bool`

                HasTimezone returns a boolean if a field has been set.

            ### GetCityName

            `func (o *Geolocation) GetCityName() string`

            GetCityName returns the CityName field if non-nil, zero value otherwise.

            ### GetCityNameOk

            `func (o *Geolocation) GetCityNameOk() (*string, bool)`

            GetCityNameOk returns a tuple with the CityName field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetCityName

            `func (o *Geolocation) SetCityName(v string)`

            SetCityName sets CityName field to given value.

                ### HasCityName

                `func (o *Geolocation) HasCityName() bool`

                HasCityName returns a boolean if a field has been set.

            ### GetCountryCode

            `func (o *Geolocation) GetCountryCode() string`

            GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

            ### GetCountryCodeOk

            `func (o *Geolocation) GetCountryCodeOk() (*string, bool)`

            GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetCountryCode

            `func (o *Geolocation) SetCountryCode(v string)`

            SetCountryCode sets CountryCode field to given value.

                ### HasCountryCode

                `func (o *Geolocation) HasCountryCode() bool`

                HasCountryCode returns a boolean if a field has been set.

            ### GetCountryName

            `func (o *Geolocation) GetCountryName() string`

            GetCountryName returns the CountryName field if non-nil, zero value otherwise.

            ### GetCountryNameOk

            `func (o *Geolocation) GetCountryNameOk() (*string, bool)`

            GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetCountryName

            `func (o *Geolocation) SetCountryName(v string)`

            SetCountryName sets CountryName field to given value.

                ### HasCountryName

                `func (o *Geolocation) HasCountryName() bool`

                HasCountryName returns a boolean if a field has been set.

            ### GetContinentCode

            `func (o *Geolocation) GetContinentCode() string`

            GetContinentCode returns the ContinentCode field if non-nil, zero value otherwise.

            ### GetContinentCodeOk

            `func (o *Geolocation) GetContinentCodeOk() (*string, bool)`

            GetContinentCodeOk returns a tuple with the ContinentCode field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetContinentCode

            `func (o *Geolocation) SetContinentCode(v string)`

            SetContinentCode sets ContinentCode field to given value.

                ### HasContinentCode

                `func (o *Geolocation) HasContinentCode() bool`

                HasContinentCode returns a boolean if a field has been set.

            ### GetContinentName

            `func (o *Geolocation) GetContinentName() string`

            GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

            ### GetContinentNameOk

            `func (o *Geolocation) GetContinentNameOk() (*string, bool)`

            GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetContinentName

            `func (o *Geolocation) SetContinentName(v string)`

            SetContinentName sets ContinentName field to given value.

                ### HasContinentName

                `func (o *Geolocation) HasContinentName() bool`

                HasContinentName returns a boolean if a field has been set.

            ### GetSubdivisions

            `func (o *Geolocation) GetSubdivisions() []GeolocationSubdivisionsInner`

            GetSubdivisions returns the Subdivisions field if non-nil, zero value otherwise.

            ### GetSubdivisionsOk

            `func (o *Geolocation) GetSubdivisionsOk() (*[]GeolocationSubdivisionsInner, bool)`

            GetSubdivisionsOk returns a tuple with the Subdivisions field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetSubdivisions

            `func (o *Geolocation) SetSubdivisions(v []GeolocationSubdivisionsInner)`

            SetSubdivisions sets Subdivisions field to given value.

                ### HasSubdivisions

                `func (o *Geolocation) HasSubdivisions() bool`

                HasSubdivisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

