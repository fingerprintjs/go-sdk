---
"go-sdk": patch
---

`GetEvent`, `UpdateEvent`, and `DeleteVisitorData` now validate their ID path parameter and return an error without sending a request when the value is not a valid identifier.
