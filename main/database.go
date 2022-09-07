package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func findChannels(channelID string) int {
	var val1 int
	db.QueryRow("SELECT COUNT(*) FROM registered_channels WHERE channel_id = '" + channelID + "'").Scan(&val1)

	return val1
}

func register(channelID string) {
	db.QueryRow("INSERT INTO registered_channels VALUES ('" + channelID + "', '" + "bob" + "')")

}
func unregister(channelID string) {
	db.QueryRow("DELETE FROM registered_channels WHERE channel_id = '" + channelID + "'")

}

func spamCatcher(m *discordgo.MessageCreate) spamReturn {
	var test spamReturn
	var err error

	var val1 int
	db.QueryRow("SELECT COUNT(*) FROM interaction WHERE time_stamp = CURRENT_DATE AND user_id = '" + m.Author.ID + "'").Scan(&val1)

	fmt.Println(val1)
	v := strconv.Itoa(val1)
	test.Value = v
	if subChecker(m) == true {
		test.TF = false

		fmt.Println(err)

		return test
	}

	if val1 > maxDailyMessage {
		test.TF = true

		return test
	}
	test.TF = false

	return test
}

func doDatabase(m *discordgo.MessageCreate) {
	db.QueryRow("INSERT INTO interaction VALUES (CURRENT_DATE, '" + m.Author.ID + "', '" + m.GuildID + "')")

}
