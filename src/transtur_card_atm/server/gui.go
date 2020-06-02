package server

import (
	"flag"
	"fmt"
	"log"
	"time"
	"encoding/json"

	"transtur_card_atm/config"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

var (
	debug = flag.Bool("d", true, "enables the debug mode")
	w     *astilectron.Window = nil
)

var (
	AppName = "Transtur"
	VersionAstilectron = "0.39.0"
	VersionElectron = "7.1.10"
)

func RunGui() {
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			//AppIconDarwinPath:  "resources/icon.icns",
			//AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			DataDirectoryPath:  config.GetAppRoot(),
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Debug:  *debug,
		Logger: l,

		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(5 * time.Second)
				if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
					l.Println(fmt.Errorf("sending check.out.menu event failed: %w", err))
				}
			}()
			return nil
		},

		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#333"),
				Center:          astikit.BoolPtr(true),
				Height:          astikit.IntPtr(700),
				Width:           astikit.IntPtr(700),
			},
		}},
	})
	if err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}
}

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	log.Printf("Message: %s", m.Name)
	return
}

func SendRedState(text string) {
	SendCardState("danger", text)
}

func SendGreyState(text string) {
	SendCardState("secondary", text)
}

func SendBlueState(text string) {
	SendCardState("primary", text)
}


func SendCardState(color string, text string) {
	log.Printf("SendCardState: %s - %s", color, text)

	msg := CardMessage{
		CardStateColor: color,
		CardStateText: text,
	}
	
	msg_json, err := json.MarshalIndent(msg, "", "    ")	
	if err != nil {
		log.Printf("Failed to serialize json: %s", err.Error())
		return
	}

	log.Printf("SendCardState json: %s", string(msg_json))
	if w != nil {
		bootstrap.SendMessage(w, "update", string(msg_json))
	}
}