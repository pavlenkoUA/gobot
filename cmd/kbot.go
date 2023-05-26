package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot %s started", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Plaese check TELE_TOKEN env variable. %s", err)
			return
		}

		var (
			menu     = &telebot.ReplyMarkup{ResizeKeyboard: true}
			selector = &telebot.ReplyMarkup{}

			btnHello    = menu.Text("Hello")
			btnHelp     = menu.Text("ℹ Help")
			btnSettings = menu.Text("⚙ Settings")

			btnPrev = selector.Data("⬅", "prev", "")
			btnNext = selector.Data("➡", "next", "")
		)

		menu.Reply(
			menu.Row(btnHello),
			menu.Row(btnHelp),
			menu.Row(btnSettings),
		)
		selector.Inline(
			selector.Row(btnPrev, btnNext),
		)

		kbot.Handle("/start", func(c telebot.Context) error {
			return c.Send("Hello! Press button", menu)
		})

		kbot.Handle("/hello", func(c telebot.Context) error {
			return c.Send("Hello!")
		})

		kbot.Handle(&btnHello, func(c telebot.Context) error {
			return c.Send("You pressed button hhelo", menu)
		})

		kbot.Handle(&btnHelp, func(c telebot.Context) error {
			return c.Edit("Here is some help: ...")
		})

		kbot.Handle(&btnPrev, func(c telebot.Context) error {
			return c.Respond()
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
