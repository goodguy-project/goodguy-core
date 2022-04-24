package handler

import (
	"context"
	"encoding/json"
	"github.com/goodguy-project/goodguy-core/core/constant"
	"github.com/goodguy-project/goodguy-core/core/web/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/util"
	"github.com/goodguy-project/goodguy-core/util/conf"
)

func doEmailConf(emailConf []*idl.EmailConf) error {
	s := conf.Viper().GetString(constant.EmailConfigName)
	var d []*idl.EmailConf
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		return err
	}
	for _, c := range emailConf {
		if c.Op == idl.Op_Op_Add {
			d = append(d, c)
		} else if c.Op == idl.Op_Op_Update {
			for i, x := range d {
				if x.Email == c.Email {
					d[i] = c
					break
				}
			}
		} else if c.Op == idl.Op_Op_Delete {
			for i, x := range d {
				if x.Email == c.Email {
					y := d[0:i]
					d = append(y, d[i+1:]...)
					break
				}
			}
		}
	}
	conf.Viper().Set(constant.EmailConfigName, util.Json(d))
	return nil
}

func AdminSet(ctx context.Context, req *idl.AdminSetRequest) (*idl.AdminSetResponse, error) {
	sid, ok := token.Auth(ctx)
	if !ok || sid != "admin" {
		return new(idl.AdminSetResponse), status.Error(codes.Unauthenticated, "auth failed")
	}
	if len(req.EmailConf) > 0 {
		if err := doEmailConf(req.EmailConf); err != nil {
			return new(idl.AdminSetResponse), err
		}
	}
	if req.GetOpenRegister() != nil {
		conf.Viper().Set(constant.OpenRegisterConfigName, req.GetOpenRegister().GetValue())
	}
	return new(idl.AdminSetResponse), nil
}

func AdminGet(ctx context.Context, req *idl.AdminGetRequest) (*idl.AdminGetResponse, error) {
	sid, ok := token.Auth(ctx)
	if !ok || sid != "admin" {
		return new(idl.AdminGetResponse), status.Error(codes.Unauthenticated, "auth failed")
	}
	var err error
	resp := new(idl.AdminGetResponse)
	emailConf := conf.Viper().GetString(constant.EmailConfigName)
	err = json.Unmarshal([]byte(emailConf), &resp.EmailConf)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func CommonGet(ctx context.Context, req *idl.CommonGetRequest) (*idl.CommonGetResponse, error) {
	resp := new(idl.CommonGetResponse)
	resp.OpenRegister = conf.Viper().GetBool(constant.OpenRegisterConfigName)
	return resp, nil
}
