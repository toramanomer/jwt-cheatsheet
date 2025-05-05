# JWT Cheatsheet

## Registered Claim Names

All of them are optional.

| Claim Name              | Description                                          | Notes                                                    |
| ----------------------- | ---------------------------------------------------- | -------------------------------------------------------- |
| "iss" (Issuer)          | Identifies the principal that issued JWT             | Case-sensitive string.                                   |
| "sub" (Subject)         | Identifies the principles that is the subject of JWT | Case-sensitive string.                                   |
| "aud" (Audience)        | The recipients that the JWT is intended for          | Array of case-sensitive strings or case-sensitive string |
| "exp" (Expiration Time) | The time after which JWT must not be accepted.       | Number: seconds since epoch                              |
| "nbf" (Not Before)      | The time before which JWT must not be accepted.      | Number: seconds since epoch                              |
| "iat" (Issued At)       | The time JWT was issued.                             | Number: seconds since epoch                              |
| "jti" (JWT ID)          | A unique identifier for the JWT                      | Can be used to prevent JWT from being replayed.          |
