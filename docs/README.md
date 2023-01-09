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
      - [Get Group](#get-group)
      - [Update Group](#update-group)
      - [Delete Group](#delete-group)
      - [List Groups](#list-groups)
      - [Get Group Member](#get-group-member)
      - [Update Group Member](#update-group-member)
      - [Remove Group Member](#remove-group-member)
      - [List Group Members](#list-group-members)
      - [Add Group Note](#add-group-note)
      - [Get Group Note](#get-group-note)
      - [Update Group Note](#update-group-note)
      - [Remove Group Note](#remove-group-note)
      - [List Group Notes](#list-group-notes)
    - [Invites](#invites)
      - [Send Invite](#send-invite)
      - [Get Invite](#get-invite)
      - [Accept Invite](#accept-invite)
      - [Deny Invite](#deny-invite)
      - [List Invites](#list-invites)
    - [Notes](#notes)
      - [Create Note](#create-note)
      - [Get Note](#get-note)
      - [Update Note](#update-note)
      - [Delete Note](#delete-note)
    - [List Notes](#list-notes)
      - [Export Note](#export-note)
    - [Blocks](#blocks)
      - [Insert Block](#insert-block)
      - [Update Block](#update-block)
      - [Delete Block](#delete-block)
    - [Conversations](#conversations)
      - [Get Conversation](#get-conversation)
      - [List Conversation](#list-conversations)
      - [Update Conversation](#update-conversation)
    - [Conversation Messages](#conversations)
      - [Send Conversation Message](#send-conversation-message)
      - [Delete Conversation Message](#delete-conversation-message)
      - [Get Conversation Message](#get-conversation-message)
      - [List Conversation Messages](#list-conversation-messages)
      - [Update Conversation Message](#update-conversation-message)

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
    "update_mask": "name"
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
- `offset=<int32>`: (Optional) Integer cursor.
- `limit=<int32>`: (Optional) Maximum number of objects returned.

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

#### Get Group

**Description:** Must be group member.

**Endpoint:** `GET /groups/:group_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.

**Response:**
```json
{
    "group": {
        "id": "string",
        "name": "string",
        "description": "string",
        "created_at": "string"
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
    "update_mask": "name,description"
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

**Description:** Must be group member.

**Endpoint:** `GET /groups`

**Tags:** `AuthRequired`

**Query:**
- `account_id=<string>`: list groups of account.
- `offset=<int32>`: (Optional) Integer cursor.
- `limit=<int32>`: (Optional) Maximum number of objects returned.

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

#### Get Group Member

**Description:** Must be group member.

**Endpoint:** `GET /groups/:group_id/members/:member_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.
- `member_id`: UUID of the account.

**Response:**
```json
{
    "member": {
        "account_id": "string",
        "role": "string",
        "created_at": "string"
    }
}
```

#### Update Group Member

**Description**: Must be group administrator.

**Endpoint:** `PATCH /groups/:group_id/members/:member_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.
- `member_id`: UUID of the account.

**Body:**
```json
{
    "member": {
        "role": "string",
    },
    "update_mask": "role",
```

**Response:**
```json
{
    "member": {
        "account_id": "string",
        "role": "string",
        "created_at": "string"
    }
}
```

#### Remove Group Member

**Description:** Must be group administrator or the authenticated user removing itself from the group.

**Endpoint:** `DELETE /groups/:group_id/members/:member_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.
- `member_id`: UUID of the account.

**Response:**
```json
{}
```

#### List Group Members

**Description:** Must be group member.

**Endpoint:** `GET /groups/:group_id/members`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.

**Query:**
- `offset=<int32>`: (Optional) Integer cursor.
- `limit=<int32>`: (Optional) Maximum number of objects returned.

**Response:**
```json
{
    "members": [
        {
            "account_id": "string",
            "role": "string",
            "created_at": "string",
        }
    ]
}
```

#### Add Group Note

**Description:** Must be group member and author of the note.

**Endpoint:** `POST /groups/:group_id/notes`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.

**Body:**
```json
{
    "group_id": "string",
    "note_id": "string",
    "title": "string",
    "author_account_id": "string",
    "folder_id": "string"
}
```

**Response:**
```json
{
    "note": {
        "note_id": "string",
        "title": "string",
        "author_account_id": "string",
        "folder_id": "string"
    }
}
```

#### Get Group Note

**Description:** Must be group member.

**Endpoint:** `GET /groups/:group_id/notes/:note_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.
- `note_id`: UUID of the group note.

**Response:**
```json
{
    "note": {
        "note_id": "string",
        "title": "string",
        "author_account_id": "string",
        "folder_id": "string"
    }
}
```

#### Update Group Note

**Description:** Must be group member. Can only update `note.title` and `note.folder_id`.

**Endpoint:** `PATCH /groups/:group_id/notes/:note_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.
- `note_id`: UUID of the group note.

**Body:**
```json
{
    "note": {
        "title": "string",
        "folder_id": "string"
    },
    "update_mask": "title,folder_id"
}
```

**Response:**
```json
{
    "note": {
        "note_id": "string",
        "title": "string",
        "author_account_id": "string",
        "folder_id": "string"
    }
}
```

#### Remove Group Note

**Description:** Must be group member, author of the note or administrator.

**Endpoint:** `DELETE /groups/:group_id/notes/:note_id`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.
- `note_id`: UUID of the group note.

**Response:**
```json
{}
```

#### List Group Notes

**Description:** Must be group member.

**Endpoint:** `GET /groups/:group_id/notes`

**Tags:** `AuthRequired`

**Path:**
- `group_id`: UUID of the group.

**Query:**
- `offset=<int32>`: (Optional) Integer cursor.
- `limit=<int32>`: (Optional) Maximum number of objects returned.
- `author_account_id=<string>`: (Optional) List only notes from that account.
- `folder_id=<string>`: (Optional) coming soon.

**Response:**
```json
{
    "notes": [
        {
            "note_id": "string",
            "title": "string",
            "author_account_id": "string",
            "folder_id": "string"
        }
    ]
}
```

### Invites

#### Send Invite

**Description:** The sender defaults to the authenticated user. Must be group member.

**Endpoint:** `POST /invites`

**Tags:** `AuthRequired`

**Body:**
```json
{
    "group_id": "string",
    "recipient_account_id": "string"
}
```

**Response:**
```json
{
    "invite": {
        "id": "string",
        "group_id": "string",
        "sender_account_id": "string",
        "recipient_account_id": "string"
    }
}
```

#### Get Invite

**Description:** Must be group administrator or sender or recipient.

**Endpoint:** `GET /invites/:invite_id`

**Tags:** `AuthRequired`

**Path:**
- `invite_id`: UUID of the invite.

**Response:**
```json
{
    "invite": {
        "id": "string",
        "group_id": "string",
        "sender_account_id": "string",
        "recipient_account_id": "string"
    }
}
```

#### Accept Invite

**Description:** Must be recipient. Accepting an invitation automatically adds the recipient to the group and deletes the invite.

**Endpoint:** `POST /invites/:invite_id/accept`

**Tags:** `AuthRequired`

**Path:**
- `invite_id`: UUID of the invite.

**Response:**
```json
{}
```

#### Deny Invite

**Description:** Must be recipient. Deletes the invitation without making the recipient join the group.

**Endpoint:** `POST /invites/:invite_id/deny`

**Tags:** `AuthRequired`

**Path:**
- `invite_id`: UUID of the invite.

**Response:**
```json
{}
```

#### List Invites

**Description:** Must be group administrator or sender or recipient.

**Endpoint:** `GET /invites`

**Tags:** `AuthRequired`

**Query:**
- `sender_account_id=<string>`: (Optional) Returns only invites from sender.
- `recipient_account_id=<string>`: (Optional) Returns only invites destined to recipient.
- `group_id=<string>`: (Optional) Returns only invites for a given group.
- `offset=<int32>`: (Optional) Integer cursor.
- `limit=<int32>`: (Optional) Maximum number of objects returned.

**Response:**
```json
{
    "invites": [
        {
            "id": "string",
            "group_id": "string",
            "sender_account_id": "string",
            "recipient_account_id": "string"
        }
    ]
}
```

### Notes

#### Create Note

**Description:** Create a note.

**Endpoint:** `POST /notes`

**Tags:** `AuthRequired`

**Body:**
```json
{
    "note": {
        "title": "string",
        "blocks": [
            {
                "type": "string"
            }
        ]
    }
}
```

**Response:**
```json
{
    "note": {
        "id": "string",
        "author_id": "string",
        "title": "string",
        "blocks": [
            {
                "id": "string",
                "type": "string"
            }
        ]
    }
}
```

#### Get Note

**Description:** Get a note and its blocks. Must be author or group member.

**Endpoint:** `GET /notes/:note_id`

**Tags:** `AuthRequired`

**Path:**
- `note_id`: UUID of the note.

**Response:**
```json
{
    "note": {
        "id": "string",
        "author_id": "string",
        "title": "string",
        "blocks": [
            {
                "id": "string",
                "type": "string"
            }
        ],
        "created_at": "string",
        "modified_at": "string"
    }
}
```

#### Update Note

**Description:** Must be author.

**Endpoint:** `PATCH /notes/:note_id`

**Tags:** `AuthRequired`

**Path:**
- `note_id`: UUID of the note.

**Body:**
```json
{
    "note": {
        "title": "string",
    },
    "update_mask": "title"
}
```

**Response:**
```json
{
    "id": "string"
}
```

#### Delete Note

**Description:** Delete a note and its blocks. Must be the owner of the note.

**Endpoint:** `DELETE /notes/:note_id`

**Tags:** `AuthRequired`

**Path:**
- `note_id`: UUID of the note.

**Response:** 
```json
{}
```

### List Notes

**Description:** Must be the author or a group member. Does not return blocks.

**Endpoint:** `GET /notes`

**Query:**
- `author_id=<string>`: Filter based on note author.

**Response:**
```json
{
    "notes": [
        {
            "id": "string",
            "author_id": "string",
            "title": "string",
            "created_at": "string",
            "modified_at": "string"
        }
    ]
}
```

#### Export Note

**Description:** Return a dowloable file in pdf or markdown format. Must be author or group member.

**Endpoint:** `GET /notes/:note_id/export`

**Path:**
- `note_id`: UUID of the note.

**Query:**
- `format=<string>`: Format of the file (either "md" or "pdf").

**Response:** File Contents

### Blocks

#### Insert Block

**Description:** Insert a block at index. Must be the owner of the note.

**Endpoint:** `POST /notes/:note_id/blocks`

**Tags:** `AuthRequired`

**Path:**
- `note_id`: UUID of the note.

**Body:**
```json
{
    "block": {
        "type": "string"
    },
    "index": "number",
}
```

**Response:**
```json
{
    "block": {
        "id": "string",
        "type": "string"
    }
}
```

#### Update Block

**Description:** Modify the contents of a block. Must be the owner of the note.

**Endpoint:** `PATCH /notes/:note_id/blocks/:block_id`

**Tags:** `AuthRequired`

**Path:**
- `note_id`: UUID of the note.
- `block_id`: UUID of the block.

**Body:**
```json
{
    "index": "number",
    "block": {
        "type": "string"
    }
}
```

**Response:**
```json
{
    "block": {
        "id": "string",
        "type": "string"
    }
}
```

#### Delete Block

**Description:** Must be the owner of the note.

**Endpoint:** `DELETE /notes/:note_id/blocks/:block_id`

**Tags:** `AuthRequired`

**Path:**
- `note_id`: UUID of the note.
- `block_id`: UUID of the block.

**Response:**
```json
{}
```

### Conversations

#### Get Conversation

**Description:** Must be group member.

**Endpoint:** `GET /conversations/:conversation_id`

**Tags:** `AuthRequired`

**Path:**
- `conversation_id`: UUID of the conversation

**Response:**
```json
{
    "conversation": {
        "id": "string",
        "group_id": "string",
        "title": "string"
    }
}
```

#### List Conversations

**Description:** Must be group member.

**Endpoint:** `GET /conversations`

**Tags:** `AuthRequired`

**Query:**
- `group_id=<string>`: UUID of the group in which you want to list the notes.

**Response:**
```json
{
    "conversations": [
        {
            "id": "string",
            "group_id": "string",
            "title": "string"
        }
    ]
}
```

#### Update Conversation

**Description:** Must be group admin. Can only update `title`.

**Endpoint:** `PATCH /conversations/:conversation_id`

**Tags:** `AuthRequired`

**Path:**
- `conversation_id`: UUID of the conversation

**Body:**
```json
{
    "conversation": {
        "title": "string",
    },
}
```

**Response**
```json
{
    "conversation": {
        "id": "string",
        "group_id": "string",
        "title": "string"
    }
}
```

### Conversation Messages

#### Send Conversation Message

**Description:** Must be group member.

**Endpoint:** `POST /conversations/:conversation_id/messages`

**Tags:** `AuthRequired`

**Path:**
- `conversation_id`: UUID of the conversation in which you want to send a message

**Reponse:**
```json
{
    "conversation_message": {
        "id": "string",
        "conversation_id": "string",
        "sender_account_id": "string",
        "content": "string",
        "created_at": "string"
    }
}
```

#### Delete Conversation Message

**Description:** Must be sender or group admin.

**Endpoint:** `DELETE /conversations/:conversation_id/messages/:message_id`

**Tags:** `AuthRequired`

**Path:**
- `conversation_id`: UUID of the conversation
- `message_id`: UUID of the message

**Reponse:**
```json
{}
```

#### Get Conversation Message

**Description:** Must be group member.

**Endpoint:** `/conversations/:conversation_id/messages/:message_id`

**Tags:** `AuthRequired`

**Path:**
- `conversation_id`: UUID of the conversation
- `message_id`: UUID of the message

**Response:**
```json
{
    "conversation_message": {
        "id": "string",
        "conversation_id": "string",
        "sender_account_id": "string",
        "content": "string",
        "created_at": "string"
    }
}
```

#### List Conversation Messages

**Description:** Must be group member. Messages are sorted in reverse chronological order.

**Endpoint:** `GET /conversations/:conversation_id/messages`

**Tags:** `AuthRequired`

**Path:**
- `conversation_id`: UUID of the conversation

**Query:**
- `offset=<int32>`: (Optional) Integer cursor.
- `limit=<int32>`: (Optional) Maximum number of objects returned.

**Response:**
```json
{
    "conversations_message": [
        {
            "id": "string",
            "conversation_id": "string",
            "sender_account_id": "string",
            "content": "string",
            "created_at": "string"
        }
    ]
}
```

#### Update Conversation Message

**Description:** Must be sender. Can only update `content`.

**Endpoint:** `PATCH /conversations/:conversation_id/messages/:message_id`

**Tags:** `AuthRequired`

**Path:**
- `conversation_id`: UUID of the conversation
- `message_id`: UUID of the message

**Body:**
```json
{
    "conversation_message": {
        "content": "string",
    },
}
```

**Response**
```json
{
    "conversation_message": {
        "id": "string",
        "conversation_id": "string",
        "sender_account_id": "string",
        "content": "string",
        "created_at": "string"
    }
}
```
