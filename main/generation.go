package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func generatePersonality(messageLength int, temp float32, top_p float32, prompt string, name string, channelID string, username string) string {
	var Stop = []string{":", "\n"}

	//t := spamCatcher(m)
	//if t.TF == true {
	//	s.ChannelMessageSend(m.ChannelID, "You have exceeded your limit of free credits today. A Patreon subscription is required for further use today.\nSubscribe to Patreon here:  https://www.patreon.com/user?u=60681312 \n\n Cooldowns reset at Midnight UTC")
	//	s.ChannelMessageSend(loggingchannel, t.Value+"auto cooldown message sent")
	//	return
	//}

	//if banChecker(m, s) == true {
	//	return
	//}
	//if cooldownChecker(m, s) == true {
	//	return
	//}
	path := "main/log/" + name + channelID + ".txt"
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)
		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}
		// Break if error occured
		if err != nil && err != io.EOF {
			log.Fatal(err)
			break
		}
	}
	newText := string(text)
	usernameAndPrompt := username + ": " + prompt
	fmt.Fprintln(file, usernameAndPrompt)
	formattedPrompt := newText + usernameAndPrompt + "\n" + name + ":"
	str := gpt3Request(formattedPrompt, temp, top_p, messageLength, Stop)
	tempstr := str
	str = name + ":" + str
	fmt.Fprintln(file, str)
	fi, _ := os.Stat(path)
	fmt.Println(fi.Size())
	file.Close()
	discordFormatted := "**" + username + ": **" + prompt + "\n" + "**" + name + ":" + "**" + tempstr
	dg.ChannelMessageSend(loggingchannel, discordFormatted)
	return discordFormatted

	//doDatabase(m)
}

func generate(messageLength int, temp float32, top_p float32, prompt string, question bool, summary bool) string {
	//t := spamCatcher(m)
	//if t.TF == true {
	//	s.ChannelMessageSend(m.ChannelID, "You have exceeded your limit of free credits today. A Patreon subscription is required for further use today.\nSubscribe to Patreon here:  https://www.patreon.com/user?u=60681312  \n\n Cooldowns reset at Midnight UTC")
	//	s.ChannelMessageSend(loggingchannel, "auto cooldown message sent")
	//	return
	//}
	//if banChecker(m, s) == true {
	//	return
	//}
	//if cooldownChecker(m, s) == true {
	//	return
	//}
	var Stop []string
	var tempPrompt string
	if question == true {
		tempPrompt = prompt
		prompt = questionString + prompt + "\nA:"
		Stop = []string{"Q: "}
	}
	if summary == true {
		tempPrompt = prompt
		prompt = summaryString + prompt + "\n\nSummary:\n"
		Stop = []string{"Summarize:"}

	}
	if question == true {
		str := gpt3Request(prompt, temp, top_p, messageLength, Stop)
		strs1 := strings.Split(str, "Q:")
		discordFormatted := "**Q: **" + tempPrompt + "\n**A: **" + strs1[0]
		dg.ChannelMessageSend(loggingchannel, discordFormatted)
		return discordFormatted
	}
	if summary == true {
		Stop[0] = ""
		str := gpt3Request(prompt, temp, top_p, messageLength, Stop)
		discordFormatted := "**Summarize:**\n" + tempPrompt + "\n**Summary:**\n" + str
		dg.ChannelMessageSend(loggingchannel, discordFormatted)
		return discordFormatted
	}
	str := gpt3Request(prompt, temp, top_p, messageLength, Stop)
	dg.ChannelMessageSend(loggingchannel, str)
	return str
}
