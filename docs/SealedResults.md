# Sealed results

## **UnsealEventsResponse**
> EventResponse UnsealEventsResponse(sealed []byte, keys []DecryptionKey)

Decrypts the sealed result response with the provided keys.
### Required Parameters

| Name       | Type                | Description                                                                              | Notes |
|------------|---------------------|------------------------------------------------------------------------------------------|-------|
| **sealed** | **[]byte**          | Sealed result data. Callers must Base64-decode the sealed result received from the agent before invoking this function.                                                               |       |
| **keys**   | **[]DecryptionKey** | Decryption keys. The SDK will try to decrypt the result with each key until it succeeds. |       |
