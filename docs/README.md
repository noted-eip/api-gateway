# API Documentation

This document describes all the endpoints of the Noted API gateway and their expected fields. The API gateway works as a RESTful JSON API.

## Authentication

Some endpoints of the API expect some form of authentication. Within the API, authentication is carried through JSON Web Tokens. A user can obtain a JSON Web Token by logging in using the `/authenticate` route documented below.

In calls that require authentication, the `Authorization` header is expected to be set to the following:
```
Authorization: Bearer <user_token>
```

The endpoints requiring authentication are marked with the tag `AuthRequired`.

## Endpoints

### Create account

**Endpoint:** `POST /accounts`

**Body:**
```json
{
    "name": "string",
    "email": "string",
    "password": "string"
}
```

**Response:**
```json
{
    "account": {
        "id": "string",
        "name": "string",
        "email": "string"
    }
}
```

**Errors:**
- Name is invalid
- Email already exists
- Password is too weak

### Get account

**Endpoint:** `GET /accounts/:id`

**Tags:** `AuthRequired`

**Path:**
- `id`: UUID of the account.

**Response:**
```json
{
    "account": {
        "id": "string",
        "name": "string",
        "email": "string"
    }
}
```

**Errors:**
- Not found

### Update account

**Description:** Update some fields of an account. The body expects an `update_mask` field which consist of a list of strings of all the fields that must be updated. For example, if wanting to update only the `"email"` and `"name"` the `"update_mask"` must be set to `["name", "email"]`.

**Endpoint:** `PATCH /accounts/:id`

**Tags:** `AuthRequired`

**Path:**
- `id`: UUID of the account.

**Body:**
```json
{
    "account": {
        "name": "string",
        "email": "string",
        "password": "string"
    },
    "update_mask": ["name", "email", "password"]
}
```

**Response:**
```json
{
    "account": {
        "id": "string",
        "name": "string",
        "email": "string"
    }
}
```

### Delete account

**Endpoint:** `DELETE /accounts/:id`

**Tags:** `AuthRequired`

**Path:**
- `id`: UUID of the account.

**Response:**
```json
{}
```

**Errors:**
- Not found

### Authenticate account

**Description:** Obtain a JWT to make authenticated calls to the API.

**Endpoint:** `POST /authenticate`

**Body:**
```json
{
    "email": "string",
    "password": "string"
}
```

**Response:**
```json
{
    "token": "string"
}
```
