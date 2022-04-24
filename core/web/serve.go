package web

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/goodguy-project/goodguy-core/core/web/handler"
	"github.com/goodguy-project/goodguy-core/idl"
	"github.com/goodguy-project/goodguy-core/util/conf"
)

type server struct {
	idl.UnimplementedGoodguyWebServiceServer
}

func (s *server) GetMember(ctx context.Context, req *idl.GetMemberRequest) (*idl.GetMemberResponse, error) {
	return handler.GetMember(ctx, req)
}

func (s *server) Login(ctx context.Context, req *idl.LoginRequest) (*idl.LoginResponse, error) {
	return handler.Login(ctx, req)
}

func (s *server) Register(ctx context.Context, req *idl.RegisterRequest) (*idl.RegisterResponse, error) {
	return handler.Register(ctx, req)
}

func (s *server) UpdateMember(ctx context.Context, req *idl.UpdateMemberRequest) (*idl.UpdateMemberResponse, error) {
	return handler.UpdateMember(ctx, req)
}

func (s *server) AdminSet(ctx context.Context, req *idl.AdminSetRequest) (*idl.AdminSetResponse, error) {
	return handler.AdminSet(ctx, req)
}

func (s *server) AdminGet(ctx context.Context, req *idl.AdminGetRequest) (*idl.AdminGetResponse, error) {
	return handler.AdminGet(ctx, req)
}

func (s *server) CommonGet(ctx context.Context, req *idl.CommonGetRequest) (*idl.CommonGetResponse, error) {
	return handler.CommonGet(ctx, req)
}

func (s *server) CheckToken(ctx context.Context, req *idl.CheckTokenRequest) (*idl.CheckTokenResponse, error) {
	return handler.CheckToken(ctx, req)
}

func (s *server) GetOfficialMember(context.Context, *idl.GetOfficialMemberRequest) (*idl.GetOfficialMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOfficialMember not implemented")
}

func Serve() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.Viper().GetString("web.host"), conf.Viper().GetInt("web.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("server listening at %v", listen.Addr())
	idl.RegisterGoodguyWebServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
