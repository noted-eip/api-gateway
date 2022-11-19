# API Documentation

This document describes all the endpoints of the Noted API gateway and their expected fields. The API gateway works as a RESTful JSON API.

- [API Documentation](#api-documentation)
  - [Concepts](#concepts)
    - [Authentication](#authentication)
    - [Internal tokens](#internal-tokens)
    - [Authorization](#authorization)
  - [Endpoints](#endpoints)
    - [Accounts](#accounts)
      - [Create Account](#create-account)
      - [Get Account](#get-account)
      - [Update Account](#update-account)
      - [Delete Account](#delete-account)
      - [List Accounts](#list-accounts)
      - [Authenticate](#authenticate)
    - [Groups](#groups)
      - [Create Group](#create-group)
      - [Update Group](#update-group)
      - [Delete Group](#delete-group)
      - [List Groups](#list-groups)
    - [Invites](#invites)
    - [Notes](#notes)
    - [Recommendations](#recommendations)
      - [Extract Keywords](#extract-keywords)

## Concepts

### Authentication

Some endpoints of the API expect some form of authentication. Within the API, authentication is carried through JSON Web Tokens. A user can obtain a JSON Web Token by logging in using the `/authenticate` route documented below.

In calls that require authentication, the `Authorization` header is expected to be set to the following:
```
Authorization: Bearer <user_token>
```

The endpoints requiring authentication are marked with the tag `AuthRequired`.

### Internal tokens

Some endpoints cannot be accessed by regular users. They can only be called using an internal token, which only developpers have access to. These endpoints are marked with the tag `InternalToken`.

### Authorization

This API enforces authorization. For example, you cannot modify a group you're not a part of, nor can you delete someone else's account. How authorization is implemented is based on common sense and in some cases it is documented in the description of an endpoint through phrases like "Must be group administrator", "Must be account owner", etc meaning the operation will fail if the user does not meet the requirements.

## Endpoints

### Accounts

#### Create Account

**Description:** Create an account using the email/password flow.

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

#### Get Account

**Endpoint:** `GET /accounts/:account_id`

**Tags:** `AuthRequired`

**Path:**
- `account_id`: UUID of the account.

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

#### Update Account

**Description:** Must be account owner.

**Endpoint:** `PATCH /accounts/:account_id`

**Tags:** `AuthRequired`

**Path:**
- `account_id`: UUID of the account.

**Body:**
```json
{
    "account": {
        "name": "string",
    },
    "update_mask": ["name"]
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

#### Delete Account

**Description**: Delete an account and its associated resources. Must be account owner.

**Endpoint:** `DELETE /accounts/:account_id`

**Tags:** `AuthRequired`

**Path:**
- `account_id`: UUID of the account.

**Response:**
```json
{}
```

#### List Accounts

**Description:** List accounts with pagination.

**Endpoint:** `GET /accounts`

**Tags:** `InternalToken`

**Query:**
- `offset=<int32>`: integer cursor.
- `limit=<int32>`: maximum number of objects returned.

**Response:**
```json
{
    "accounts": [
        {
            "id": "string",
            "name": "string",
            "email": "string"
        }
    ]
}
```

#### Authenticate

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

### Groups

#### Create Group

**Endpoint:** `POST /groups`

**Tags:** `AuthRequired`

**Body:**
```json
{
    "name": "string",
    "description": "string"
}
```
**Response:**
```json
{
    "group": {
        "id": "string",
        "name": "string",
        "description": "string",
        "created_at": "string",
    }
}
```

#### Update Group

**Description**: Must be group administrator.

**Endpoint:** `PATCH /groups/:group_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.

**Body:**
```json
{
    "group": {
        "name": "string",
        "description": "string",
    },
    "update_mask": ["name", "description"]
}
```

**Response:**
```json
{
    "group": {
        "id": "string",
        "name": "string",
        "description": "string",
        "created_at": "string",
    }
}
```

#### Delete Group

**Description:** Delete a group and its associated resources. Must be group administrator.

**Endpoint:** `DELETE /groups/:group_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.

**Response:**
```json
{}
```

#### List Groups

**Description:** Must be groups member.

**Endpoint:** `GET /groups`

**Tags:** `AuthRequired`

**Query:**
- `offset=<int32>`: integer cursor.
- `limit=<int32>`: maximum number of objects returned.

**Response:**
```json
{
    "groups": [
        {
            "id": "string",
            "name": "string",
            "description": "string",
            "created_at": "string",
        }
    ]
}
```

### Invites

### Notes

### Recommendations

#### Extract Keywords

**Endpoint:** `POST /recommendations/keywords`

**Body:**
```json
{
    "content": "string"
}
```

**Response:**
```json
{
    "keywords": ["strings"]
}
```
