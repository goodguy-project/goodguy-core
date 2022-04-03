package handler

import (
	"context"
	"encoding/json"

	"github.com/spf13/viper"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/util"
)

func doEmailConf(emailConf []*idl.EmailConf) error {
	s := viper.GetString(util.EmailConfigName)
	var d []*idl.EmailConf
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		return err
	}
	for _, conf := range emailConf {
		if conf.Op == idl.Op_Op_Add {
			d = append(d, conf)
		} else if conf.Op == idl.Op_Op_Update {
			for i, x := range d {
				if x.Email == conf.Email {
					d[i] = conf
					break
				}
			}
		} else if conf.Op == idl.Op_Op_Delete {
			for i, x := range d {
				if x.Email == conf.Email {
					y := d[0:i]
					d = append(y, d[i+1:]...)
					break
				}
			}
		}
	}
	viper.Set(util.EmailConfigName, util.JsonToString(d))
	return nil
}

func AdminSet(ctx context.Context, req *idl.AdminSetRequest) (*idl.AdminSetResponse, error) {
	if len(req.EmailConf) > 0 {
		if err := doEmailConf(req.EmailConf); err != nil {
			return new(idl.AdminSetResponse), err
		}
	}
	if req.OpenRegister != idl.Bool_Bool_Undefined {
		viper.Set(util.OpenRegisterConfigName, req.OpenRegister)
	}
	return new(idl.AdminSetResponse), nil
}

func AdminGet(ctx context.Context, req *idl.AdminGetRequest) (*idl.AdminGetResponse, error) {
	var err error
	resp := new(idl.AdminGetResponse)
	emailConf := viper.GetString(util.EmailConfigName)
	err = json.Unmarshal([]byte(emailConf), &resp.EmailConf)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func CommonGet(ctx context.Context, req *idl.CommonGetRequest) (*idl.CommonGetResponse, error) {
	resp := new(idl.CommonGetResponse)
	resp.OpenRegister = viper.GetInt(util.OpenRegisterConfigName) == int(idl.Bool_Bool_True)
	return resp, nil
}
