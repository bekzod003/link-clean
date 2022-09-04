package bot_buttons

import "gopkg.in/telebot.v3"

var (
	// Universal markup builders.
	Menu     = &telebot.ReplyMarkup{ResizeKeyboard: true}
	Selector = &telebot.ReplyMarkup{}

	// Reply buttons.
	BtnHelp     = Menu.Text("ℹ Help")
	BtnSettings = Menu.Text("⚙ Settings")

	// Inline buttons.
	//
	// Pressing it will cause the client to
	// send the bot a callback.
	//
	// Make sure Unique stays unique as per button kind,
	// as it has to be for callback routing to work.
	//
	BtnPrev = Selector.Data("⬅", "prev")
	BtnNext = Selector.Data("➡", "next")
)
