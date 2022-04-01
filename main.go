package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/caarlos0/env/v6"
)

func main() {
	// Parse env variables for Twitter API keys
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%v\n", err)
	}

	// // Configure a new httpClient to pass to twitter.NewClient()
	// conf := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	// token := oauth1.NewToken(cfg.AccessToken, cfg.AccessSecret)
	// httpClient := conf.Client(oauth1.NoContext, token)

	// // Create new Twitter client
	// client := twitter.NewClient(httpClient)

	// // Verify Twitter Client Authenticated
	// user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// fmt.Printf("Account @%s (%s)\n", user.ScreenName, user.Name)

	// listenToStream(client)

	// arg1 := fmt.Sprintf("consumer-key=%s", cfg.ConsumerKey)
	// arg2 := fmt.Sprintf("consumer-secret=%s", cfg.ConsumerSecret)
	// arg3 := fmt.Sprintf("access-token=%s", cfg.AccessToken)
	// arg4 := fmt.Sprintf("access-secret=%s", cfg.AccessSecret)

	// cmd := exec.Command("twitter-upload.exe", arg1, arg2, arg3, arg4)

	cmd := exec.Command("twitter-upload")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(out.String())
}
