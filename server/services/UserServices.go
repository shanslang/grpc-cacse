package services

import (
	"context"
	"fmt"
	"io"
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

// 客户端流模式
func (this *UserService) GetUserScoreByClientStream(in UserService_GetUserScoreByClientStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := in.Recv()
		if err == io.EOF { // 服务端接收完数据后回数据给客户端
			return in.SendAndClose(&UserScoreResponse{Users: users})
		}
		if err != nil {
			return err
		}

		for _, user := range req.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
	}
}

func (this *UserService) GetUserScoreByCSStream(in UserService_GetUserScoreByCSStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := in.Recv()
		if err == io.EOF { // 服务端接收完数据
			return nil
		}
		if err != nil {
			return err
		}

		for _, user := range req.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
		err = in.Send(&UserScoreResponse{Users: users})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (this *UserService) mustEmbedUnimplementedUserServiceServer() {

}
