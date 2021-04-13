package server

import (
	"flag"
	"fmt"
	"log"
	"time"
	//"encoding/json"

	"transtur_card_atm/config"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

var (
	debug = flag.Bool("d", true, "enables the debug mode")
	appWindow     *astilectron.Window = nil
)

var (
	AppName = "Transtur"
	VersionAstilectron = "0.39.0"
	VersionElectron = "7.1.10"
)

func RunGui() {
	appConfig := config.GetAppConfig()
	guiLogger := log.New(log.Writer(), log.Prefix(), log.Flags())

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
		
		Logger: guiLogger,

		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Developer Tools"),
					OnClick: handleDevTools,
				},
				{
					Label: astikit.StrPtr("FTP Test Upload"),
					OnClick: handleFtpTestUpload,
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},

		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			appWindow = ws[0
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
				Frame:           astikit.BoolPtr(appConfig.Debug),
				Fullscreen:      astikit.BoolPtr(!appConfig.Debug),
				AlwaysOnTop:     astikit.BoolPtr(!appConfig.Debug),
			},
		}},
	})
	if err != nil {
		log.Fatal(fmt.Errorf("running bootstrap failed: %w", err))

	}
}

func GetGuiWindow() *astilectron.Window {
	return appWindow
}

func WaitGuiWindow() *astilectron.Window {
	for {
		w := GetGuiWindow()
		if w != nil {
			return w
		}
		time.Sleep(1 * time.Second)
	}
}

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	log.Printf("Message: %s", m.Name)
	return
}

func handleDevTools(e astilectron.Event) (deleteListener bool) {
	GetGuiWindow().OpenDevTools()						
	return
}

func handleFtpTestUpload(e astilectron.Event) (deleteListener bool) {
	UploadFtp([]byte("transtur_card_atm test upload\n"), "ftp-test.txt")
	return
}
