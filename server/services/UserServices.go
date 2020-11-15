package services

import (
	"context"
)

type UserService struct {
}

// 普通方式
func (this *UserService) GetUserScore(ctx context.Context, in *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for _, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}
	return &UserScoreResponse{Users: users}, nil
}

// 流
func (this *UserService) GetUserScoreByServerStream(in *UserScoreRequest, s UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for index, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)

		if (index+1)%2 == 0 {
			err := s.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			users = (users)[0:0] // 空切片赋值给users
		}

	}
	if len(users) > 0 {
		err := s.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *UserService) mustEmbedUnimplementedUserServiceServer() {

}
