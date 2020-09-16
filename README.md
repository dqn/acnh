# acnh

[![build status](https://github.com/dqn/acnh/workflows/build/badge.svg)](https://github.com/dqn/acnh/actions)

Animal Crossing: New Horizons API wrapper.

## Installation

```bash
$ go get github.com/dqn/acnh
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/dqn/acnh"
	"github.com/dqn/gonso"
)

func main() {
	accessToken, _ := gonso.Auth("SESSION_TOKEN")
	a, err := acnh.New(accessToken)
	if err != nil {
		// Handle error.
	}

	r, err := a.Users()
	if err != nil {
		// Handle error.
	}

	for _, user := range r.Users {
		fmt.Println(user)
	}
}
```

You can get session token and access token by using [gonso](https://github.com/dqn/gonso).

## API

### func New

```go
func New(accessToken string) *ACNH
```

New ACNH Client.

### type ACNH

#### func Users

```go
func (a *ACNH) Users() UsersResponse
```

Get all land users.

#### func AuthToken

```go
func (a *ACNH) AuthToken(userID string) AuthTokenResponse
```

Get token for user.

#### func UserProfile

```go
func (a *ACNH) UserProfile(token, userID string) UserProfileResponse
```

Get user profile.

#### func LandProfile

```go
func (a *ACNH) LandProfile(token, landID string) LandProfileResponse
```

Get land profile.

#### func Friends

```go
func (a *ACNH) Friends(token string) FriendsResponse
```

Get friends.

#### func PresenceFriends

```go
func (a *ACNH) PresenceFriends(token string) PresenceFriendsResponse
```

Get online friends.

#### func SendMessageAll

```go
func (a *ACNH) SendMessageAll(token, message string) SendMessageResponse
```

Send message to all friends.

#### func SendMessageFriend

```go
func (a *ACNH) SendMessageFriend(token, userID, message string) SendMessageResponse
```

Send message to a friend.

#### func SendMessageLocal

```go
func (a *ACNH) SendMessageLocal(token, message string) SendMessageResponse
```

Send message local.

### Response type details

See [type.go](type.go).

## License

MIT
