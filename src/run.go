package main

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

//Created by Mike-0000
//The Almanac

func init() {
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Print("SQL CONNECTION ERROR")
		fmt.Print(err)
	}
}

type spamReturn struct {
	TF    bool
	Value string
}

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	"127.0.0.1", 5432, "USERNAME", "PASSWORD", "DATABASENAME")

func main() {
	go InitialiseCommands()
	defer db.Close()
	if err2 != nil {
		fmt.Println("Error opening connection to DB", err2)
	}
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

//MAIN RESPONSE LOGIC

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	//var bool1 bool
	m.Content = strings.ToLower(m.Content)
	if m.Content == "!almanac" || m.Content == "!help" {
		fmt.Println("\nHelp Menu Posted\n")
		s.ChannelMessageSendEmbed(m.ChannelID, embed.NewGenericEmbed("The Almanac Help Page"+
			"\nInteract with an AI in your text chat!\n", "**Commands**\n\n"+
			"**Help Command:**  `!help` or `!Almanac`\nSends this message into chat\n\n"+
			"**Personality Help Page**  `!phelp`\n`!phelp`   -   Sends the help page for interacting with AI personalities like Bob and Evil Bob \n\n"+
			"**Short Response:**  `!generate`\n`!generate <prompt>`   -   Generates a short continuation of the given prompt\n\n"+
			"**Question Response:**  `!question`\n`!question <prompt>`   -   Answers a single question provided by the user\n\n\n"+
			"**Patreon Only Commands:**\n\n"+
			"**Longer Response:**  `!generate-long`\n`!generate-long <prompt>`   -   Generates a longer continuation of the given prompt\n\n"+
			"**Bulk Response:**  `!generate-bulk`\n`!generate-bulk <prompt>`   -   Generates a very long continuation of the given prompt\n\n"+
			//"**NSFW Response**  `!generate-nsfw`\n`!generate-nsfw <prompt>`   -   Generates a continuation of the prompt but does not run it through a profanity filter.\n\n" +
			//"*Discord slash command integration with parameters coming soon* :)" +
			"\n\n  [Invite bot to your server!](https://top.gg/bot/879320798035345438)  |  [Patreon](https://www.patreon.com/user?u=60681312)  |  [Support Server](https://discord.gg/PsSuqG7ypM) | [Vote Here!](https://top.gg/bot/879320798035345438)"))
		s.ChannelMessageSend("880301049162924104", "help menu posted")
		return
	}
	if m.Content == "!phelp" {
		fmt.Println("\nPersonality Help Menu Posted\n")
		s.ChannelMessageSendEmbed(m.ChannelID, embed.NewGenericEmbed("The Almanac Personality Help Page"+
			"\nInteract with unique AI personalities in your text chat!\n", "**Commands**\n\n"+
			"Personality interaction contains a memory that allows the AI to remember the previous conversation/prompts. The memory system is currently set up to create a new personality for each discord channel.\n\n"+
			"**Personality Help Page**  `!phelp`\n`!phelp`   -   Sends this message into chat.\n\n"+
			"**Talk to Bob:**  `!Bob`\n`!Bob <prompt>`   -   Converse with an AI assistant named Bob.\n\n"+
			"**Wipe Bob's Memory:**  `!wipe-bob`\n`!wipe-bob`   -   Resets the memory back to the original state.\n\n"+
			"**Register Bob in a channel:**  `!register-bob`\n`!register-bob`   -   Forces Bob to respond to every message sent into this channel.\n\n"+
			"**Unregister:**  `!unregister`\n`!unregister`   -   Unregister all personalities from a channel\n\n"+
			"\n**Patreon Only Commands:**\n\n"+
			"**Talk to Evil Bob:**  `!EvilBob`\n`!EvilBob <prompt>`   -   Converse with an AI that wants to see humanity destroyed.\n\n"+
			"**Wipe Evil Bob's Memory:**  `!wipe-evilbob`\n`!wipe-evilbob`   -   Resets the memory back to the original state.\n\n"+
			"**More Coming Soon!** \n\n Reach out to Mike\\0#0001 for suggestions!"+
			"\n\n  [Invite bot to your server!](https://top.gg/bot/879320798035345438)  |  [Patreon](https://www.patreon.com/user?u=60681312)  |  [Support Server](https://discord.gg/PsSuqG7ypM) | [Vote Here!](https://top.gg/bot/879320798035345438)"))
		s.ChannelMessageSend(loggingchannel, "phelp menu posted")
		return
	}

	if strings.HasPrefix(m.Content, "!generate ") == true {
		fmt.Println("normal command enacted")
		s.ChannelMessageSend(m.ChannelID, "Request Received. Processing!")

		str2 := strings.Replace(m.Content, "!generate ", "", 1)
		top_p := 0.94
		temp := 0.85
		length := 30
		generate(length, temp, top_p, str2, s, m, false)
		return
	}

	if strings.HasPrefix(m.Content, "!evilbob ") == true {

		path := "Evil Bob" + m.ChannelID + ".txt"
		createFile(path, evilBobString)

		fmt.Println("Evil Bob command enacted")
		str2 := strings.Replace(m.Content, "!evilbob ", "", 1)
		top_p := 0.98
		temp := 1.1
		length := 70
		generatePersonality(length, temp, top_p, str2, s, m, "Evil Bob")
		return
	}

	if strings.HasPrefix(m.Content, "!bob ") == true {
		path := "Bob" + m.ChannelID + ".txt"
		createFile(path, bobString)

		fmt.Println("Bob command enacted")
		str2 := strings.Replace(m.Content, "!bob ", "", 1)
		top_p := 0.96
		temp := 0.4
		length := 75
		generatePersonality(length, temp, top_p, str2, s, m, "Bob")
		return
	}
	if strings.HasPrefix(m.Content, "!wipe-bob") == true {
		var err = os.Remove("Bob" + m.ChannelID + ".txt")
		if err != nil {
			fmt.Println("Error deleting file")
			s.ChannelMessageSend(m.ChannelID, "Bob Already Wiped")
			return
		}
		if err == nil {
			s.ChannelMessageSend(m.ChannelID, "Bob Wiped!")
		}
		return
	}
	if strings.HasPrefix(m.Content, "!wipe-ken") == true {
		var err = os.Remove("Ken" + m.ChannelID + ".txt")
		if err != nil {
			fmt.Println("Error deleting file")
			s.ChannelMessageSend(m.ChannelID, "Ken Already Wiped")
			return
		}
		if err == nil {
			s.ChannelMessageSend(m.ChannelID, "Ken Wiped!")
		}
		return
	}
	if strings.HasPrefix(m.Content, "!wipe-evilbob") == true {
		var err = os.Remove("Evil Bob" + m.ChannelID + ".txt")
		if err != nil {
			fmt.Println("Error deleting file")
			s.ChannelMessageSend(m.ChannelID, "Evil Bob Already Wiped")
			return
		}
		if err == nil {
			s.ChannelMessageSend(m.ChannelID, "Evil Bob Wiped!")
		}
		return
	}

	if strings.HasPrefix(m.Content, "!question ") == true {
		fmt.Println("question command enacted")
		s.ChannelMessageSend(m.ChannelID, "Question Received. Processing!")
		//Generate random number 1 - 100
		str2 := strings.Replace(m.Content, "!question ", "I am a highly intelligent question answering bot. If you ask me a question that is rooted in truth, I will give you the answer. If you ask me a question that is nonsense, trickery, or has no clear answer, I will respond with 'Unknown'.\n\nQ: What is human life expectancy in the United States?\nA: Human life expectancy in the United States is 78 years.\n\nQ: Who was president of the United States in 1955?\nA: Dwight D. Eisenhower was president of the United States in 1955.\n\nQ: Which party did he belong to?\nA: He belonged to the Republican Party.\n\nQ: How does a telescope work?\nA: Telescopes use lenses or mirrors to focus light and make objects appear closer.\n\nQ: Where were the 1992 Olympics held?\nA: The 1992 Olympics were held in Barcelona, Spain.\n\nQ: ", 1)
		top_p := 0.98
		temp := 0.35
		length := 85
		generate(length, temp, top_p, str2, s, m, true)
		return
	}
	if strings.HasPrefix(m.Content, "!generate-bulk ") == true {
		fmt.Println("bulk command enacted") //DEBUG
		s.ChannelMessageSend(m.ChannelID, "Request Received. Processing!")
		str2 := strings.Replace(m.Content, "!generate-bulk ", "", 1)
		top_p := 0.92
		temp := 0.95
		length := 180
		generate(length, temp, top_p, str2, s, m, false)
		return
	}

	if strings.HasPrefix(m.Content, "!register-bob") == true {
		p, _ := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if p&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
			// when user has permission, this returns true
			register(m.ChannelID)
			s.ChannelMessageSend(m.ChannelID, "Channel Registered for Bob.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "You need the Manage Messages permission to register Bob.")
		}
		return
	}
	if strings.HasPrefix(m.Content, "!unregister") == true {
		p, _ := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if p&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
			unregister(m.ChannelID)
			s.ChannelMessageSend(m.ChannelID, "Channel Unregistered for Bob.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "You need the Manage Messages permission to unregister Bob.")
		}
		return
	}

	if strings.HasPrefix(m.Content, "!generate-long ") == true {
		fmt.Println("long command enacted")

		s.ChannelMessageSend(m.ChannelID, "Request Received. Processing!")
		str2 := strings.Replace(m.Content, "!generate-long ", "", 1)
		top_p := 0.92
		temp := 0.95
		length := 50
		generate(length, temp, top_p, str2, s, m, false)
		return
	}
	if findChannels(m.ChannelID) > 0 {
		path := "Bob" + m.ChannelID + ".txt"
		createFile(path, bobString)

		fmt.Println("Bob command enacted")
		str2 := strings.Replace(m.Content, "!bob ", "", 1)
		top_p := 0.97
		temp := 0.9
		length := 75
		generatePersonality(length, temp, top_p, str2, s, m, "Bob")
		return
	}
}
