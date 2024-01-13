package cmds_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/GemAi/features"
	"github.com/GemAi/utils"

	"github.com/bwmarrin/discordgo"
)

func Ollama() *features.Cmd {
	return &features.Cmd{
		Command: &discordgo.ApplicationCommand{
			Name:        "ollama",
			Description: "Talks to ollama directly",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "prompt",
					Description: "Prompt",
					Required:    true,
				},
			},
		},
		Function: ollamaFunction,
	}
}

func ollamaFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	utils.InteractionRespond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Generating...",
		},
	})

	content, err := post(utils.Options(i)["prompt"].StringValue())
	if err != nil {
		fmt.Println("error", err.Error())
		utils.EditMessageError(s, i, err.Error())
	}
	fmt.Printf("%v/n", content)
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: content,
	})
}

type Request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type Response struct {
	Model     string `json:"model"`
	CreatedAt string `json:"create_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

func post(prompt string) (*string, error) {
	request := Request{Model: "mistral:7b-instruct", Prompt: prompt, Stream: false}

	jsonReq, err := json.Marshal(request)
	resp, err := http.Post("http://gemai-ollama-test-1:11434/api/generate", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status + "/n " + string(bodyBytes[:]))
	}

	var response Response
	json.Unmarshal(bodyBytes, &response)

	return &response.Response, nil
}
