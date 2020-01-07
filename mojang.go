package mojang

import (
        "bytes"
        "encoding/json"
        "github.com/pkg/errors"
        "io/ioutil"
        "net/http"
)

// MojangPlayer struct
type MojangPlayer struct {
	UUID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Legacy bool `json:"legacy,omitempty"`
	Demo   bool `json:"demo,omitempty"`
}

// MojangPlayers type
type MojangPlayers []MojangPlayer

func MojangGetUsers(names []string) (MojangPlayers, error) {
        jsonNames, err1 := json.Marshal(names)

        if err1 != nil {
                return MojangPlayers{}, err1
	}

        resp, err2 := http.Post("https://api.mojang.com/profiles/minecraft", "application/json", bytes.NewBuffer(jsonNames))

        if err2 != nil {
                return MojangPlayers{}, err2
	}

        respData, err2 := ioutil.ReadAll(resp.Body)

        if err2 != nil {
                return MojangPlayers{}, err2
	}

        var mojResponse MojangPlayers
        err3 := json.Unmarshal(respData, &mojResponse)

        if err3 != nil {
                return MojangPlayers{}, err3
	}

        return mojResponse, nil
}

func MojangGetUser(name string) (MojangPlayer, error) {
        users, err := MojangGetUsers([]string{name})

	if err != nil {
		return MojangPlayer{}, err
	}

        if len(users) <= 0 {
                return MojangPlayer{}, errors.New("Cannot get UUID from Mojang for: " + name)
        }

        return users[0], nil
}
