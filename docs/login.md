# Login

Login flow

## Table of Contents

- [Success](#success)
- [Invalid Params](#invalid-parameters)
- [Credential Invalid](#password-not-match)
- [Account Inactive](#account-inactive)


## Request

- Path : `/v1/login`
- Method : `POST`
- Payload

| Field      | Type             | Description                   |
|------------|------------------|-------------------------------|
| `email`    | string, required | Must be a valid email address |
| `password` | string, required |                               |


```shell
curl --location '{{base_url}}/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "febry@gmail.com",
    "password": "makcik"
}'
```

## Response

### Success

200 OK

```json
{
  "status": true,
  "data": {
    "token": "MY.JWT.TOKEN",
    "type": "internal",
    "expired_at": "2023-06-01T01:53:48.524466441Z"
  }
}
```

example of `jwt token`

https://jwt.io/#debugger-io?token=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlVTRVItRU1BSUwiLCJleHAiOjE2ODU2ODYwNDcsImlhdCI6MTY4NTU5OTY0NywianRpIjoiR1VJRC1TRVNTSU9OLUlEIiwibmFtZSI6IkZVTExOQU1FIiwic3ViIjoiVVNFUi1JRCIsImlzcyI6ImFwaS5rdXJzdXMtbWFzYWsuaWQiLCJhdWQiOiJ3d3cua3Vyc3VzLW1hc2FrLmlkIn0.cff342HJ8IDuYqAZVpYmKObeDqgq6WYATfRmu1oxewzVeyFyfx5PeBs7kttfzPMjphj0Eh07pWeonVU3IjiS1OzQJHBCiXr2LYNgUN8EQM2C4rSi7Mn8AXUCyYELS2eoZygh0-lEJ3dzPqMc45LQcWV21FQxfYmFKK0ukjgKFzR-sssJIVqBcN5G7M7EccXRqP2oJp1uzBGkReGCddmPoJolXmZuSSlKJVuYDwXr-YyQThKc52EbZqCreDTORpEsx0MqTDoc19rFNcJGmltPltz57YMHtX2liFcTihNBNK6gPzuzxxhSwbZ7-aDhFkf2Z4pLWcsUCVEuvVbE_AUTuA&publicKey=-----BEGIN%20PUBLIC%20KEY-----%0AMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu1SU1LfVLPHCozMxH2Mo%0A4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0%2FIzW7yWR7QkrmBL7jTKEn5u%0A%2BqKhbwKfBstIs%2BbMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyeh%0Akd3qqGElvW%2FVDL5AaWTg0nLVkjRo9z%2B40RQzuVaE8AkAFmxZzow3x%2BVJYKdjykkJ%0A0iT9wCS0DRTXu269V264Vf%2F3jvredZiKRkgwlL9xNAwxXFg0x%2FXFw005UWVRIkdg%0AcKWTjpBP2dPwVZ4WWC%2B9aGVd%2BGyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbc%0AmwIDAQAB%0A-----END%20PUBLIC%20KEY-----

`claims`
```json
{
  "email": "USER-EMAIL",
  "exp": 1685686047,
  "iat": 1685599647,
  "jti": "GUID-SESSION-ID",
  "name": "FULLNAME",
  "sub": "USER-ID",
  "iss": "api.kursus-masak.id",
  "aud": "www.kursus-masak.id"
}
```

### Invalid Parameters

400 Bad Request

```json
{
    "status": false,
    "message": "Email dan Password tidak boleh kosong",
    "error_code": "ERR"
}
```

### Password Not Match

400 Bad Request

```json
{
  "status": false,
  "message": "Email / Password tidak sesuai",
  "error_code": "ERR"
}
```

### Account Inactive

400 Bad Request

```json
{
  "status": false,
  "message": "Akun ini tidak aktif, silahkan hubungi administrator",
  "error_code": "ERR"
}
```
