package service

import (
	"context"
	pb "englishAI/gpt/api/chat"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

type ChatService struct {
	pb.UnimplementedChatServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (s *ChatService) GetChat(ctx context.Context, req *pb.GetChatRequest) (*pb.GetChatReply, error) {

	client := openai.NewClient("your token")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		return &pb.GetChatReply{contect: err}, nil
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return &pb.GetChatReply{contect: resp.Choices[0].Message.Content}, nil
}
