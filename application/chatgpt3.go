package application

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
)

type ProxyGPT struct {
	Protocol string
	Addr     string
	Port     string
}

type chatGPT3 struct {
	client          *openai.Client
	contextMessages []openai.ChatCompletionMessage // 上下文message
}

func GenClient(token string) *chatGPT3 {
	return &chatGPT3{
		client:          openai.NewClient(token),
		contextMessages: make([]openai.ChatCompletionMessage, 0),
	}
}

func GenClientWithProxy(token string, p ProxyGPT) *chatGPT3 {
	config := openai.DefaultConfig(token)
	proxyURL, err := url.Parse(fmt.Sprintf("%s://%s:%s", p.Protocol, p.Addr, p.Port))
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}
	return &chatGPT3{
		client:          openai.NewClientWithConfig(config),
		contextMessages: make([]openai.ChatCompletionMessage, 0),
	}
}

// SendMessage send a completion message
func (gpt3 *chatGPT3) SendMessage(message string) (reply string, err error) {
	resp, err := gpt3.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

// SendMessagesWithContext send message in context
func (gpt3 *chatGPT3) SendMessagesWithContext(message string) (reply string, err error) {
	gpt3.contextMessages = append(gpt3.contextMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
	resp, err := gpt3.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: gpt3.contextMessages,
		},
	)
	if err != nil {
		return "", err
	}
	reply = resp.Choices[0].Message.Content
	gpt3.contextMessages = append(gpt3.contextMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: reply,
	})
	return reply, nil
}
