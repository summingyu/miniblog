package user

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/summingyu/miniblog/internal/pkg/log"
	pb "github.com/summingyu/miniblog/pkg/proto/miniblog/v1"
)

func (ctrl *UserController) ListUser(ctx context.Context, r *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	log.C(ctx).Infow("ListUser function called")
	resp, err := ctrl.b.Users().List(ctx, int(*r.Offset), int(*r.Limit))
	if err != nil {
		return nil, err
	}
	users := make([]*pb.UserInfo, 0, len(resp.Users))
	for _, u := range resp.Users {
		createdAt, _ := time.Parse("2006-01-02 15:04:05", u.CreatedAt)
		updatedAt, _ := time.Parse("2006-01-02 15:04:05", u.UpdatedAt)
		users = append(users, &pb.UserInfo{
			Username:  u.Username,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Phone:     u.Phone,
			PostCount: u.PostCount,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		})
	}
	ret := &pb.ListUserResponse{
		TotalCount: resp.TotalCount,
		Users:      users,
	}
	return ret, nil
}
