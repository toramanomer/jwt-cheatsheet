# JWT Cheatsheet

JWTs represent a set of claims as a JSON object that is encoded in JWS and/or JWE structurtes.
A JWT is represented as a sequence of base64url-encoded values separated by period ('.') characters.

-   If the JWT is a JWS, then there are three parts:
-   If the JWT is a JWE, then there are five parts:

## JOSE Header

**JOSE Header** is a JSON object that describes the cryptographic operations applied the JWT, and optionally, additional properties of the JWT.
Depending upon whether the JWT is a JWS or JWE, the corresponding rules for the JOSE Header values apply.

| Header Parameter     | Description                                   |
| -------------------- | --------------------------------------------- |
| "typ" (Type)         | the media type of the JWT                     |
| "cty" (Content Type) | must be JWT with nested signing or encryption |

## JWT Claims

**JWT Claims Set** is a JSON object that is encoded in a JWS and/or JWE structure.
It can have zero or more name/value pairs (or members).

-   **Claim Names** are the name portion of the members and are always strings. Claim names must be unique.
-   **Claim Values** are the value portion of the members and can be of any JSON value type.

-   If it is a JWS, then the **JWT Claims Set** is the **payload** of the JWS.
-   If it is a JWE, then the **JWT Claims Set** is the **plaintext** of the JWE.

### Registered Claim Names

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
