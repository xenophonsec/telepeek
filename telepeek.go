package main

import (
    "fmt"
    "log"
    "os"
    "context"
	"strings"
	"bytes"
    "encoding/json"
	
    "github.com/urfave/cli/v3"
	"github.com/asmcos/requests"
)

func ensurePrefix(token string) string {
	if !strings.HasPrefix(token, "bot") {
		return "bot" + token
	}
	return token
}

func prettyPrintJSON(jsonString string) {
    var prettyJSON bytes.Buffer
    err := json.Indent(&prettyJSON, []byte(jsonString), "", "    ")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(prettyJSON.String())
}

func main() {
	telegramAPI := "https://api.telegram.org/"
    cmd := &cli.Command{
		Usage: "Telegram OSINT CLI",
        Commands: []*cli.Command{
            {
                Name:    "bot",
                Aliases: []string{"b"},
                Usage:   "data from bot token",
                Commands: []*cli.Command{
					{
                        Name:  "info",
                        Usage: "get bot info",
                        Action: func(ctx context.Context, cmd *cli.Command) error {
							botToken := cmd.Args().First()
							botToken = ensurePrefix(botToken)
							resp,err := requests.Get(telegramAPI + botToken + "/getMe")
							if err != nil{
								fmt.Println(err)
							}
							prettyPrintJSON(resp.Text())
                            return nil
                        },
                    },
					{
                        Name:  "webhook",
                        Usage: "get bot webhook info",
                        Action: func(ctx context.Context, cmd *cli.Command) error {
							botToken := cmd.Args().First()
							botToken = ensurePrefix(botToken)
							resp,err := requests.Get(telegramAPI + botToken + "/getWebhookInfo")
							if err != nil{
								fmt.Println(err)
							}
							prettyPrintJSON(resp.Text())
                            return nil
                        },
                    },
					{
                        Name:  "chat",
                        Usage: "get chat info",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "chat",
								Value: "",
								Usage: "chat id",
							},
						},
                        Action: func(ctx context.Context, cmd *cli.Command) error {
							botToken := cmd.Args().First()
							botToken = ensurePrefix(botToken)
							chatId := cmd.String("chat")
							if len(chatId) < 1 {
								fmt.Println("chat id required")
								return nil
							}
							resp,err := requests.Get(telegramAPI + botToken + "/getChat?chat_id=" + chatId)
							if err != nil{
								fmt.Println(err)
							}
							prettyPrintJSON(resp.Text())
                            return nil
                        },
                    },
					{
                        Name:  "chatadmins",
                        Usage: "get chat administrators",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "chat",
								Value: "",
								Usage: "chat id",
							},
						},
                        Action: func(ctx context.Context, cmd *cli.Command) error {
							botToken := cmd.Args().First()
							botToken = ensurePrefix(botToken)
							chatId := cmd.String("chat")
							if len(chatId) < 1 {
								fmt.Println("chat id required")
								return nil
							}
							resp,err := requests.Get(telegramAPI + botToken + "/getChatAdministrators?chat_id=" + chatId)
							if err != nil{
								fmt.Println(err)
							}
							prettyPrintJSON(resp.Text())
                            return nil
                        },
                    },
					{
                        Name:  "all",
                        Usage: "get everything available",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "chat",
								Value: "",
								Usage: "chat id",
							},
						},
                        Action: func(ctx context.Context, cmd *cli.Command) error {
							botToken := cmd.Args().First()
							botToken = ensurePrefix(botToken)
							chatId := cmd.String("chat")
							println("Bot Info")
							println("=================================================")
							resp,err := requests.Get(telegramAPI + botToken + "/getMe")
							if err != nil{
								fmt.Println(err)
							}
							prettyPrintJSON(resp.Text())
							println("=================================================")
							println("")
							println("Bot Webhook Info")
							println("=================================================")
							resp,err = requests.Get(telegramAPI + botToken + "/getWebhookInfo")
							if err != nil{
								fmt.Println(err)
							}
							prettyPrintJSON(resp.Text())
							println("=================================================")
							println("")
							if len(chatId) > 0 {
								println("Chat Info")
								println("=================================================")
								resp,err = requests.Get(telegramAPI + botToken + "/getChat?chat_id=" + chatId)
								if err != nil{
									fmt.Println(err)
								}
								prettyPrintJSON(resp.Text())
								println("=================================================")
								println("")
								println("Chat Administrators")
								println("=================================================")
								resp,err = requests.Get(telegramAPI + botToken + "/getChatAdministrators?chat_id=" + chatId)
								if err != nil{
									fmt.Println(err)
								}
								prettyPrintJSON(resp.Text())
								println("=================================================")
							}
							
                            return nil
                        },
                    },
                },
            },
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
