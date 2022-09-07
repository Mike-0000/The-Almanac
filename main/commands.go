package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
	"os"
)

type HelpCommand struct{}

var (
	_ ken.DmCapable    = (*HelpCommand)(nil)
	_ ken.SlashCommand = (*HelpCommand)(nil)
)

func (c *HelpCommand) IsDmCapable() bool {
	return true
}

func (c *HelpCommand) Name() string {
	return "help"
}
func (c *HelpCommand) Description() string {
	return "Help Command"
}
func (c *HelpCommand) Version() string {
	return "1.0.0"
}
func (c *HelpCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}
func (c *HelpCommand) Run(ctx ken.Context) (err error) {
	fmt.Println("Helped")
	ctx.RespondEmbed(&discordgo.MessageEmbed{
		Description: "The Almanac Help Page" +
			"\nInteract with an AI in your text chat/\n" +
			"**Commands**\n\n" +
			"**Help Command:**  `/help` \nSends this message into chat\n\n" +
			"**Personality Help Page**  `/phelp`\n`/phelp`   -   Sends the help page for interacting with AI personalities like Bob and Evil Bob \n\n" +
			"**Generation:**  `/generate`\n`/generate <prompt>`   -   Generates a short continuation of the given prompt\n\n" +
			"**Question Response:**  `/question`\n`/question <prompt>`   -   Answers a single question provided by the user\n\n\n" +
			"**Summary Response:**  `/summarize`\n`/summarize <prompt>`   -   Summarizes a piece of text provided by the user\n\n\n" +
			//"**Patreon Only Commands:**\n\n" +
			//"**Longer Response:**  `/generate-long`\n`/generate-long <prompt>`   -   Generates a longer continuation of the given prompt\n\n" +
			//"**Bulk Response:**  `/generate-bulk`\n`/generate-bulk <prompt>`   -   Generates a very long continuation of the given prompt\n\n" +
			"\n  [Invite bot to your server!](https://top.gg/bot/879320798035345438)  |  [Patreon](https://www.patreon.com/user?u=60681312)  |  [Support Server](https://discord.gg/PsSuqG7ypM) | [Vote Here!](https://top.gg/bot/879320798035345438)",
	})
	return
}

type pHelpCommand struct{}

var (
	_ ken.SlashCommand = (*pHelpCommand)(nil)
	_ ken.DmCapable    = (*pHelpCommand)(nil)
)

func (c *pHelpCommand) IsDmCapable() bool {
	return true
}

func (c *pHelpCommand) Name() string {
	return "phelp"
}
func (c *pHelpCommand) Description() string {
	return "pHelp Command"
}
func (c *pHelpCommand) Version() string {
	return "1.0.0"
}
func (c *pHelpCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}
func (c *pHelpCommand) Run(ctx ken.Context) (err error) {
	fmt.Println("pHelped")
	ctx.RespondEmbed(&discordgo.MessageEmbed{
		Description: "The Almanac Personality Help Page" +
			"\nInteract with unique AI personalities in your text chat/\n" +
			"**Commands**\n\n" +
			"Personality interaction contains a memory that allows the AI to remember the previous conversation/prompts. The memory system is currently set up to create a new personality for each discord channel.\n\n" +
			"**Personality Help Page**  `/phelp`\n`/phelp`   -   Sends this message into chat.\n\n" +
			"**Talk to Bob:**  `/Bob`\n`/Bob <prompt>`   -   Converse with an AI assistant named Bob.\n\n" +
			"**Wipe Bob's Memory:**  `/wipe-bob`\n`/wipe-bob`   -   Resets the memory back to the original state.\n\n" +
			//"**Register Bob in a channel:**  `/register-bob`\n`/register-bob`   -   Forces Bob to respond to every message sent into this channel.\n\n" +
			//"**Unregister:**  `/unregister`\n`/unregister`   -   Unregister all personalities from a channel\n\n" +
			//"\n**Patreon Only Commands:**\n\n" +
			"**Talk to Evil Bob:**  `/EvilBob`\n`/EvilBob <prompt>`   -   Converse with an AI that wants to see humanity destroyed.\n\n" +
			"**Wipe Evil Bob's Memory:**  `/wipe-evilbob`\n`/wipe-evilbob`   -   Resets the memory back to the original state.\n\n" +
			"**More Coming Soon** \n\n Reach out to Mike\\0#0001 for suggestions" +
			"\n\n  [Invite bot to your server!](https://top.gg/bot/879320798035345438)  |  [Patreon](https://www.patreon.com/user?u=60681312)  |  [Support Server](https://discord.gg/PsSuqG7ypM) | [Vote Here!](https://top.gg/bot/879320798035345438)",
	})
	return
}

type Bob struct{}

func (c *Bob) IsDmCapable() bool {
	return true
}

var (
	_ ken.SlashCommand = (*Bob)(nil)
	_ ken.DmCapable    = (*Bob)(nil)
)

func (c *Bob) Name() string {
	return "bob"
}
func (c *Bob) Description() string {
	return "Talk to Bob!"
}
func (c *Bob) Version() string {
	return "1.0.0"
}

func (c *Bob) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "prompt",
			Description: "Prompt",
			Required:    true,
		},
	}
}

func (c *Bob) Run(ctx ken.Context) (err error) {
	ctx.Defer()
	username := ctx.User().Username
	channelID := ctx.GetEvent().ChannelID

	path := "main/log/" + "Bob" + channelID + ".txt"
	createFile(path, bobString)

	prompt := ctx.Options().GetByName("prompt").StringValue()
	fmt.Println("bob command enacted")
	var top_p float32 = 0.96
	var temp float32 = 0.9
	length := 75

	str := generatePersonality(length, temp, top_p, prompt, "Bob", channelID, username)
	ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: str,
		},
	})
	return
}

type Generate struct{}

func (c *Generate) IsDmCapable() bool {
	return true
}

var (
	_ ken.SlashCommand = (*Generate)(nil)
	_ ken.DmCapable    = (*Generate)(nil)
)

func (c *Generate) Name() string {
	return "textgenerate"
}
func (c *Generate) Description() string {
	return "AI Text Continuation"
}
func (c *Generate) Version() string {
	return "1.0.0"
}

var prompt string

func (c *Generate) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "prompt",
			Description: "Prompt",
			Required:    true,
		},
	}
}

func (c *Generate) Run(ctx ken.Context) (err error) {
	ctx.Defer()

	prompt = ctx.Options().GetByName("prompt").StringValue()
	fmt.Println("normal command enacted")
	var top_p float32 = 0.97
	var temp float32 = 0.6
	length := 65

	str := generate(length, temp, top_p, prompt, false, false)
	ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: str,
		},
	})
	return
}

type Question struct{}

func (c *Question) IsDmCapable() bool {
	return true
}

var (
	_ ken.SlashCommand = (*Question)(nil)
	_ ken.DmCapable    = (*Question)(nil)
)

func (c *Question) Name() string {
	return "question"
}
func (c *Question) Description() string {
	return "AI Question Answering"
}
func (c *Question) Version() string {
	return "1.0.0"
}

func (c *Question) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "prompt",
			Description: "Prompt",
			Required:    true,
		},
	}
}

func (c *Question) Run(ctx ken.Context) (err error) {
	ctx.Defer()

	prompt = ctx.Options().GetByName("prompt").StringValue()
	fmt.Println("question command enacted")
	var top_p float32 = 0.97
	var temp float32 = 0.25
	length := 65

	str := generate(length, temp, top_p, prompt, true, false)
	ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: str,
		},
	})
	return
}

type EvilBob struct{}

func (c *EvilBob) IsDmCapable() bool {
	return true
}

var (
	_ ken.SlashCommand = (*EvilBob)(nil)
	_ ken.DmCapable    = (*EvilBob)(nil)
)

func (c *EvilBob) Name() string {
	return "evilbob"
}
func (c *EvilBob) Description() string {
	return "Talk to Evil Bob!"
}
func (c *EvilBob) Version() string {
	return "1.0.0"
}

func (c *EvilBob) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "prompt",
			Description: "Prompt",
			Required:    true,
		},
	}
}

func (c *EvilBob) Run(ctx ken.Context) (err error) {
	ctx.Defer()
	username := ctx.User().Username
	channelID := ctx.GetEvent().ChannelID

	path := "main/log/" + "Evil Bob" + channelID + ".txt"
	createFile(path, evilBobString)

	prompt := ctx.Options().GetByName("prompt").StringValue()
	fmt.Println("evilbob command enacted")
	var top_p float32 = 0.96
	var temp float32 = 0.91
	length := 115

	str := generatePersonality(length, temp, top_p, prompt, "Evil Bob", channelID, username)
	ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: str,
		},
	})
	return
}

type wipebob struct{}

func (c *wipebob) IsDmCapable() bool {
	return true
}

var (
	_ ken.SlashCommand = (*wipebob)(nil)
	_ ken.DmCapable    = (*wipebob)(nil)
)

func (c *wipebob) Name() string {
	return "wipe-bob"
}
func (c *wipebob) Description() string {
	return "Wipe Bob's Memory (Memory is per channel)"
}
func (c *wipebob) Version() string {
	return "1.0.0"
}

func (c *wipebob) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *wipebob) Run(ctx ken.Context) (err error) {
	channelID := ctx.GetEvent().ChannelID
	err = os.Remove("main/log/Bob" + channelID + ".txt")
	if err != nil {
		fmt.Println("Error deleting file")
		ctx.Respond(&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Bob Already Wiped",
			},
		},
		)
		return
	}
	if err == nil {
		ctx.Respond(&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Bob Wiped!",
			},
		},
		)
		return
	}
	return
}

type wipeEvilBob struct{}

func (c *wipeEvilBob) IsDmCapable() bool {
	return true
}

var (
	_ ken.SlashCommand = (*wipeEvilBob)(nil)
	_ ken.DmCapable    = (*wipeEvilBob)(nil)
)

func (c *wipeEvilBob) Name() string {
	return "wipe-evilbob"
}
func (c *wipeEvilBob) Description() string {
	return "Wipe Evil Bob's Memory (Memory is per channel)"
}
func (c *wipeEvilBob) Version() string {
	return "1.0.0"
}

func (c *wipeEvilBob) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *wipeEvilBob) Run(ctx ken.Context) (err error) {
	channelID := ctx.GetEvent().ChannelID
	err = os.Remove("main/log/Evil Bob" + channelID + ".txt")
	if err != nil {
		fmt.Println("Error deleting file")
		ctx.Respond(&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Evil Bob Already Wiped",
			},
		},
		)
		return
	}
	if err == nil {
		ctx.Respond(&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Evil Bob Wiped!",
			},
		},
		)
		return
	}
	return
}

type Summarize struct{}

func (c *Summarize) IsDmCapable() bool {
	return true
}

var (
	_ ken.SlashCommand = (*Summarize)(nil)
	_ ken.DmCapable    = (*Summarize)(nil)
)

func (c *Summarize) Name() string {
	return "summarize"
}
func (c *Summarize) Description() string {
	return "AI Summarizing"
}
func (c *Summarize) Version() string {
	return "1.0.0"
}

func (c *Summarize) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "prompt",
			Description: "Prompt",
			Required:    true,
		},
	}
}

func (c *Summarize) Run(ctx ken.Context) (err error) {
	var str string
	ctx.Defer()

	prompt = ctx.Options().GetByName("prompt").StringValue()
	fmt.Println("question command enacted")
	var top_p float32 = 0.99
	var temp float32 = 0.2
	length := 95

	str = generate(length, temp, top_p, prompt, false, true)
	ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: str,
		},
	})
	return
}
