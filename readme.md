# JWT Cheatsheet

## Registered Claim Names

All of them are optional.

| Claim Name              | Description                                | Notes                                                    |
| ----------------------- | ------------------------------------------ | -------------------------------------------------------- |
| "iss" (Issuer)          | Who issued the token (e.g. auth.myapp.com) | Case-sensitive string.                                   |
| "sub" (Subject)         | Who the token is about (e.g., user ID)     | Case-sensitive string.                                   |
| "aud" (Audience)        | Who the token is for (e.g., app ID)        | Array of case-sensitive strings or case-sensitive string |
| "exp" (Expiration Time) | Token is invalid after this time           | Number: seconds since epoch                              |
| "nbf" (Not Before)      | Token is invalid before this time.         | Number: seconds since epoch                              |
| "iat" (Issued At)       | Token is issued at this time.              | Number: seconds since epoch                              |
| "jti" (JWT ID)          | A unique identifier for the JWT            | Can be used to prevent JWT from being replayed.          |
