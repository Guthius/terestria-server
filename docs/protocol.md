# Protocol

This document describes the TCP protocol of *Mirage Nova*.

## 1. Login

The login request is send from the Client to the Server.

The server will respond back with a result code. If the result code is 0 the Client can continue to the character selection screen.

If the result code is non-zero, there was a problem with the login and the Message field should be displayed to the player.

### Request

| Field | Name        | Type   | Description            |
| ----- | ----------- | ------ | ---------------------- |
| 0     | Player Name | string | The name of the player |

### Response

| Field | Name        | Type   | Description                                          |
| ----- | ----------- | ------ | ---------------------------------------------------- |
| 0     | Result      | byte   | The login result                                     |
| 1     | Message     | string | A error message (only present if Result is non-zero) |


## 2. JoinGame

Send by the Server to the Client. 

When the client receives this command it should load the requested map and create a Player actor at the specified coordinates.

| Field | Name        | Type   | Description                                                    |
| ----- | ----------- | ------ | -------------------------------------------------------------- |
| 0     | Player ID   | int    | A unique ID of the player (this is the ID the Client controls) |
| 1     | X           | int    | The X position                                                 |
| 2     | Y           | int    | The Y position                                                 |
| 3     | Map         | string | The path of the map to load (without the .tscn extension)      |

## 3. AddPlayer

Send by the Server to the client when a player enters the map the Client is currently on.

| Field | Name        | Type   | Description               |
| ----- | ----------- | ------ | ------------------------- |
| 0     | Player ID   | int    | A unique ID of the player |
| 1     | Player Name | string | The name of the player    |
| 2     | Sprite      | string | The sprite of the texture |
| 3     | X           | int    | The X position            |
| 4     | Y           | int    | The Y position            |

## 4. RemovePlayer

Send by the Server to the Client when a player leaves the map the Client is currently on.

| Field | Name        | Type   | Description                           |
| ----- | ----------- | ------ | ------------------------------------- |
| 0     | Player ID   | int    | The unique ID of the player to remove |


## 5. MovePlayer

Send by the Client to the Server to indicate it has started moving in the specified.

If the move is allowed, the server will send a `MovePlayer` response to all other players on the same map.

### Request

| Field | Name        | Type   | Description               |
| ----- | ----------- | ------ | ------------------------- |
| 1     | Direction   | byte   | The direction to move in  |

### Response

| Field | Name        | Type   | Description               |
| ----- | ----------- | ------ | ------------------------- |
| 0     | Player ID   | int    | A unique ID of the player |
| 1     | Direction   | byte   | The direction to move in  |

## 6. SetPlayerPosition

Send by the Server to the Client to set the position of a player.

When a `MovePlayer` request is initiated by the Client, the Client begins 
movement immediately without waiting for the Server's approval. If the Server 
determines that the requested movement violates the rules or constraints of the
game, it will issue a `SetPlayerPosition` command. This command corrects the 
player's position by relocating them to a valid location.

| Field | Name        | Type   | Description               |
| ----- | ----------- | ------ | ------------------------- |
| 0     | Player ID   | int    | A unique ID of the player |
| 1     | X           | int    | The X position            |
| 2     | Y           | int    | The Y position            |
