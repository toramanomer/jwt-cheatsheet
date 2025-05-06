# JWT Cheatsheet

JWTs represent a set of claims as a JSON object that is encoded in JWS and/or JWE structurtes.
A JWT is represented as a sequence of base64url-encoded values separated by period ('.') characters.

-   If the JWT is a JWS, then there are three parts:

    1.  JOSE Header

        JSON object, typically consists of "alg" and "typ".

    2.  Payload

        JSON object that contains the claims.

    3.  Signature

        ```
        HMACSHA256(
            base64UrlEncode(header) + "." +
            base64UrlEncode(payload),
            secret
        )
        ```

-   If the JWT is a JWE, then there are five parts:

    ```
    base64UrlEncode(UTF8(JWE Protected Header)) + '.' +
    base64UrlEncode(JWE Encrypted Key) + '.' +
    base64UrlEncode(JWE Initialization Vector) + '.' +
    base64UrlEncode(JWE Ciphertext) + '.' +
    base64UrlEncode(JWE Authentication Tag)
    ```

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

## Best Practices

1. Perform Algorithm Verification

-   Specify a supported set of algorithms. Ensure the "alg" or "enc" header specifies the same algorithm that is used for the cryptographic operation.
-   Each key must be used with exactly one algorithm.

2. Use Appropriate Algorithms

-   Avoid all RSA-PKCS1 v1.5 encryption algorithms, preferring RSAES-OAEP
-   Ensure unique random value for every message that is signed when using ECDSA signatures.

3. Ensure Cryptographic Keys Have Sufficient Entropy

-   Human-memorizable passwords MUST NOT be directly used as the key to a keyed-MAC algorithm such as "HS256".

4. Use UTF-8

-   UTF-8 be used for encoding and decoding JSON used in Header Parameters and JWT Claims Sets.

5. Validate Issuer and Subject

-   When "iss" claim is present, the app must validate that the cryptographic keys used for cryptographic operations in the JWT belong to the issues. This can be done using jwks_uri.
-   When "sub" claim is present, the app must validate that the subject value corresponds to a valid subject.

6. Use and Validate Audience

-   If the same issuer can issue JWTs that are intended for use by more than one relying party or application, the JWT MUST contain an "aud" (audience) claim that can be used to determine whether the JWT is being used by an intended party or was substituted by an attacker at an unintended party
