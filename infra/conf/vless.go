package conf

import (
	"encoding/json"

	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/common/uuid"
	"github.com/xtls/xray-core/proxy/vless"
	"github.com/xtls/xray-core/proxy/vless/outbound"
	"google.golang.org/protobuf/proto"
)

type VLessOutboundVnext struct {
	Address *Address          `json:"address"`
	Port    uint16            `json:"port"`
	Users   []json.RawMessage `json:"users"`
}

type VLessOutboundConfig struct {
	Vnext []*VLessOutboundVnext `json:"vnext"`
}

// Build implements Buildable
func (c *VLessOutboundConfig) Build() (proto.Message, error) {
	config := new(outbound.Config)

	if len(c.Vnext) == 0 {
		return nil, errors.New(`VLESS settings: "vnext" is empty`)
	}
	config.Vnext = make([]*protocol.ServerEndpoint, len(c.Vnext))
	for idx, rec := range c.Vnext {
		if rec.Address == nil {
			return nil, errors.New(`VLESS vnext: "address" is not set`)
		}
		if len(rec.Users) == 0 {
			return nil, errors.New(`VLESS vnext: "users" is empty`)
		}
		spec := &protocol.ServerEndpoint{
			Address: rec.Address.Build(),
			Port:    uint32(rec.Port),
			User:    make([]*protocol.User, len(rec.Users)),
		}
		for idx, rawUser := range rec.Users {
			user := new(protocol.User)
			if err := json.Unmarshal(rawUser, user); err != nil {
				return nil, errors.New(`VLESS users: invalid user`).Base(err)
			}
			account := new(vless.Account)
			if err := json.Unmarshal(rawUser, account); err != nil {
				return nil, errors.New(`VLESS users: invalid user`).Base(err)
			}

			u, err := uuid.ParseString(account.Id)
			if err != nil {
				return nil, err
			}
			account.Id = u.String()

			switch account.Flow {
			case "", vless.XRV, vless.XRV + "-udp443":
			default:
				return nil, errors.New(`VLESS users: "flow" doesn't support "` + account.Flow + `" in this version`)
			}

			if account.Encryption != "none" {
				return nil, errors.New(`VLESS users: please add/set "encryption":"none" for every user`)
			}

			user.Account = serial.ToTypedMessage(account)
			spec.User[idx] = user
		}
		config.Vnext[idx] = spec
	}

	return config, nil
}
