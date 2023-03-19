package service

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	pb "gpt/api/chat"
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
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return &pb.GetChatReply{Content: resp.Choices[0].Message.Content}, nil
}
