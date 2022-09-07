package main

import (
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
)

func subChecker(m *discordgo.MessageCreate) bool {
	//ff := false
	tf := false
	for i := 0; i < len(subs); i++ { // For individual Subscribers
		if m.Author.ID == subs[i] {
			tf = true
		}
	}
	for i := 0; i < len(servers); i++ { // For Server Members
		if m.GuildID == servers[i] {
			tf = true
		}
	}
	if tf == false {
		return false
	}
	return true

}

func banChecker(m *discordgo.MessageCreate, s *discordgo.Session) bool {
	for i := 0; i < len(bans); i++ {
		if m.Author.ID == bans[i] {
			s.ChannelMessageSend(m.ChannelID, "Your messages have been marked as spam or unconstructive, you are permanently banned")
			s.ChannelMessageSend(loggingchannel, "banned message sent")
			return true
		}
	}
	for i := 0; i < len(serverBans); i++ {
		if m.GuildID == serverBans[i] {
			s.ChannelMessageSend(m.ChannelID, "This community has been permanently from using The Almanac")
			s.ChannelMessageSend(loggingchannel, "server banned message sent")
			return true
		}
	}
	for i := 0; i < len(serverWarning); i++ {
		if m.GuildID == serverWarning[i] {
			s.ChannelMessageSend(m.ChannelID, "This community has been marked for using offensive words. This includes talk of sexual assault, or the use of offensive slurs. If this community continues to use these words the server will be permanently banned from using The Almanac. You have been warned.")
			s.ChannelMessageSend(loggingchannel, "server warning message sent")
			return false
		}
	}
	for i := 0; i < len(warning); i++ {
		if m.Author.ID == warning[i] {
			s.ChannelMessageSend(m.ChannelID, "You have been marked for using offensive words. This includes talk of sexual assault, or the use of offensive slurs. If you continue to use these words you will be permanently banned from using The Almanac. You have been warned.")
			s.ChannelMessageSend(loggingchannel, "user warning message sent")
			return false
		}
	}
	return false
}

func cooldownChecker(m *discordgo.MessageCreate, s *discordgo.Session) bool {
	for i := 0; i < len(cooldown); i++ {
		if m.Author.ID == cooldown[i] {
			s.ChannelMessageSend(m.ChannelID, "You have exceeded your limit of free credits today. A Patreon subscription is required for further use today.\nSubscribe to Patreon here:  https://www.patreon.com/user?u=60681312")
			s.ChannelMessageSend(loggingchannel, "cooldown message sent")
			return true
		}
	}

	for i := 0; i < len(serverCooldown); i++ {
		if m.GuildID == serverCooldown[i] {
			s.ChannelMessageSend(m.ChannelID, "This server has excceded the limit of free credits today. A Patreon subscription is required for further use today.\nSubscribe to Patreon here:  https://www.patreon.com/user?u=60681312")
			s.ChannelMessageSend(loggingchannel, "server cooldown message sent")
			return true
		}
	}

	return false
}

func divideString(input string) []string {
	output := strings.Split(input, "\n")
	return output
}

func createFile(pathed string, dString string) {
	// check if file exists
	var _, err = os.Stat(pathed)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(pathed)
		for _, v := range divideString(dString) {
			fmt.Fprintln(file, v)
			fmt.Println(v)
		}
		fmt.Println("File Created Successfully", pathed)
		//_, err = file.WriteString()
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		return
	}
	fmt.Println("File Already Created, No action", pathed)

}

func cleanNeuroReponse(formattedResponse string) string {
	formattedResponse = strings.Replace(formattedResponse, "map[print:[] progress:0 result:[map[generated_text:", "", -1) //Remove extra tags for formatting
	if strings.Contains(formattedResponse, "map[print:[") == true {
		if strings.Contains(formattedResponse, "generated_text:") {
			strs := strings.Split(formattedResponse, "generated_text:")
			formattedResponse1 := strings.Replace(strs[1], "]] state:COMPLETE taskType:PREDICTION]", "", -1) //Remove ending tags
			return formattedResponse1
		} else {
			return formattedResponse
		}
	} else {
		formattedResponse = strings.Replace(formattedResponse, "]] state:COMPLETE taskType:PREDICTION]", "", -1)
		return formattedResponse
	}
}

func gpt3Request(prompt string, temp float32, top_p float32, messageLength int, stop []string) string {
	var messageLengthP *int
	messageLengthP = new(int)
	*messageLengthP = messageLength
	var tempP *float32
	tempP = new(float32)
	*tempP = temp
	var top_pP *float32
	top_pP = new(float32)
	*top_pP = top_p
	if stop[0] != "" {
		resp, err := client.Completion(ctx, gpt3.CompletionRequest{
			Temperature: tempP,
			TopP:        top_pP,
			MaxTokens:   messageLengthP,
			Prompt:      []string{prompt},
			Stop:        stop,
		})
		if err != nil {
			fmt.Println(err)
		}
		return resp.Choices[0].Text
	} else {
		resp, err := client.Completion(ctx, gpt3.CompletionRequest{
			Temperature: tempP,
			TopP:        top_pP,
			MaxTokens:   messageLengthP,
			Prompt:      []string{prompt},
		})
		if err != nil {
			fmt.Println(err)
		}
		return resp.Choices[0].Text
	}

}

//func trimLastChar(s string) string {
//	r, size := utf8.DecodeLastRuneInString(s)
//	if r == utf8.RuneError && (size == 0 || size == 1) {
//		size = 0
//	}
//	return s[:len(s)-size]
//}
//
//
//
//func verify(signature, hash, publicKey string) bool {
//	decodedSig, err := hex.DecodeString(signature)
//	if err != nil {
//		return false
//	}
//	decodedPubKey, err := hex.DecodeString(publicKey)
//	if err != nil {
//		return false
//	}
//	return ed25519.Verify(decodedPubKey, []byte(hash), decodedSig)
//}
