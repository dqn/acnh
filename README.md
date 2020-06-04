# acnh

Animal Crossing: New Horizons API

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
	"github.com/dqn/go-nso"
)

func main() {
	accessToken, _ := nso.New().Auth()
	a, err := acnh.New(accessToken)
	if err != nil {
		panic(err)
	}

	r, err := a.Users()
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Users[0].Name) // => どきゅん
}
```

You can get `accessToken` by using [gonso](https://github.com/dqn/gonso).

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
