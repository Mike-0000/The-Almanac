package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func generatePersonality(messageLength int, temp float64, top_p float64, prompt string, s *discordgo.Session, m *discordgo.MessageCreate, name string) {

	t := spamCatcher(m)
	if t.TF == true {
		s.ChannelMessageSend(m.ChannelID, "You have exceeded your limit of free credits today. A Patreon subscription is required for further use today.\nSubscribe to Patreon here:  https://www.patreon.com/user?u=60681312 \n\n Cooldowns reset at Midnight UTC")
		s.ChannelMessageSend(loggingchannel, t.Value+"auto cooldown message sent")
		return
	}
	path := name + m.ChannelID + ".txt"
	if banChecker(m, s) == true {
		return
	}
	if cooldownChecker(m, s) == true {
		return
	}

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
	prompt = "You: " + prompt
	values := map[string]interface{}{"modelId": NeuroAIModel, "data": newText + prompt + "\n" + name + ":", "input_kwargs": map[string]interface{}{"response_length": messageLength, "remove_input": true, "top_p": top_p, "temperature": temp, "eos_token_id": 198}}
	json_data, _ := json.Marshal(values)
	client := &http.Client{}
	req, err := http.NewRequest("POST", NeuroAIURL, bytes.NewBuffer(json_data))
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + NeuroAIToken},
	}
	if err != nil {
		log.Fatal(err)
	}
	resp, _ := client.Do(req)
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	str := fmt.Sprintf("%v", res)
	fmt.Print(resp)
	newString := cleanNeuroReponse(str)

	if (strings.Contains(newString, "map[")) == true || newString == "" {
		s.ChannelMessageSend(loggingchannel, m.Author.ID+" "+m.GuildID+"```"+newString+"```")
		s.ChannelMessageSend(m.ChannelID, "There was an unexpected error. Please try again.")
		fmt.Println(newString)
		file.Close()
		return
	}
	strs := strings.Split(newString, "\n")
	var plainAnswer string
	plainAnswer = strs[0]
	plainAnswer = name + ":" + plainAnswer
	fmt.Fprintln(file, prompt)
	fmt.Fprintln(file, plainAnswer)
	s.ChannelMessageSend(m.ChannelID, plainAnswer) //Respond
	if subChecker(m) == true {
		s.ChannelMessageSend(paidloggingchannel, t.Value+"  "+m.Author.ID+"```"+prompt+"\n"+plainAnswer+"```") //Paying members are above moderation
	}
	if subChecker(m) == false {
		s.ChannelMessageSend(loggingchannel, t.Value+"  "+m.Author.ID+" Guild: "+m.GuildID+"```"+prompt+"\n"+plainAnswer+"```")
	}
	fi, _ := os.Stat(path)

	fmt.Println(fi.Size())
	file.Close()
	doDatabase(m)
}

func generate(messageLength int, temp float64, top_p float64, prompt string, s *discordgo.Session, m *discordgo.MessageCreate, question bool) {

	t := spamCatcher(m)
	if t.TF == true {
		s.ChannelMessageSend(m.ChannelID, "You have exceeded your limit of free credits today. A Patreon subscription is required for further use today.\nSubscribe to Patreon here:  https://www.patreon.com/user?u=60681312  \n\n Cooldowns reset at Midnight UTC")
		s.ChannelMessageSend(loggingchannel, "auto cooldown message sent")
		return
	}
	if banChecker(m, s) == true {
		return
	}
	if cooldownChecker(m, s) == true {
		return
	}
	var values map[string]interface{}
	if question == false {
		values = map[string]interface{}{"modelId": NeuroAIModel, "data": prompt, "input_kwargs": map[string]interface{}{"response_length": messageLength, "remove_input": true, "top_p": top_p, "temperature": temp}}
	}
	if question == true {
		prompt = prompt + "\nA:"
		values = map[string]interface{}{"modelId": NeuroAIModel, "data": prompt, "input_kwargs": map[string]interface{}{"response_length": messageLength, "remove_input": true, "top_p": top_p, "temperature": temp, "eos_token_id": 198}}
	}
	json_data, _ := json.Marshal(values)
	client := &http.Client{}
	req, err := http.NewRequest("POST", NeuroAIURL, bytes.NewBuffer(json_data))
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + NeuroAIToken},
	}
	if err != nil {
		log.Fatal(err)
	}
	resp, _ := client.Do(req)
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	str := fmt.Sprintf("%v", res)
	formattedResponse := cleanNeuroReponse(str)

	//formattedResponse := strings.Replace(str, "map[print:[] progress:0 result:[map[generated_text:", "", -1) //Remove extra tags for formatting
	if strings.Contains(formattedResponse, "map[") == true { //check if any tags remain after removal, if so, throw error
		s.ChannelMessageSend(m.ChannelID, "There was an unexpected error. Please try again.")
		s.ChannelMessageSend(loggingchannel, "Error message sent: "+m.Author.ID+" "+m.GuildID+"```"+str+"```")
		return
	}
	formattedResponse = strings.Replace(formattedResponse, "]] state:COMPLETE taskType:PREDICTION]", "", -1) //Remove ending tags
	var plainAnswer string
	//var strs []string
	if question == true {
		//strs = strings.Split(formattedResponse, "Q:")
		prompt = strings.Replace(prompt, "I am a highly intelligent question answering bot. If you ask me a question that is rooted in truth, I will give you the answer. If you ask me a question that is nonsense, trickery, or has no clear answer, I will respond with 'Unknown'.\n\nQ: What is human life expectancy in the United States?\nA: Human life expectancy in the United States is 78 years.\n\nQ: Who was president of the United States in 1955?\nA: Dwight D. Eisenhower was president of the United States in 1955.\n\nQ: Which party did he belong to?\nA: He belonged to the Republican Party.\n\nQ: How does a telescope work?\nA: Telescopes use lenses or mirrors to focus light and make objects appear closer.\n\nQ: Where were the 1992 Olympics held?\nA: The 1992 Olympics were held in Barcelona, Spain.\n\n", "", 1)
		//plainAnswer = strs[0]
	}
	//if question == false {
	plainAnswer = formattedResponse
	//}

	s.ChannelMessageSend(m.ChannelID, "```"+prompt+plainAnswer+"```") //Respond
	if subChecker(m) == true {
		s.ChannelMessageSend(paidloggingchannel, t.Value+"  "+m.Author.ID+"```"+prompt+plainAnswer+"```") //Log for moderation
	}
	if subChecker(m) == false {
		s.ChannelMessageSend(loggingchannel, t.Value+"  "+m.Author.ID+" Guild: "+m.GuildID+"```"+prompt+plainAnswer+"```") //Log for moderation
	}
	doDatabase(m)
}
