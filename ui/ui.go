package ui

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/cjbrigato/d4-bnet-mitm/bgspacket"
	"github.com/cjbrigato/d4-bnet-mitm/log"
	"github.com/cjbrigato/d4-bnet-mitm/services"
	"github.com/cjbrigato/d4-bnet-mitm/ui/tview"
	"github.com/gdamore/tcell/v2"
	"github.com/pterm/pterm"
)

var (
	Pages      = tview.NewPages()
	App        = tview.NewApplication()
	PacketList = &tview.List{
		Box:                tview.NewBox(),
		ShowSecondaryText:  false,
		WrapAround:         false,
		HighlightFullLine:  true,
		MainTextStyle:      tcell.StyleDefault.Foreground(tview.Styles.PrimaryTextColor),
		SecondaryTextStyle: tcell.StyleDefault.Foreground(tview.Styles.TertiaryTextColor),
		ShortcutStyle:      tcell.StyleDefault.Foreground(tview.Styles.SecondaryTextColor),
		SelectedStyle:      tcell.StyleDefault.Foreground(tview.Styles.PrimitiveBackgroundColor).Background(tview.Styles.PrimaryTextColor),
	}
	LogList     = tview.NewList()
	TreeLogList = tview.NewTreeView()
	TreeRoot    = tview.NewTreeNode("root")
	TvPackets   = []*tview.TextView{}
	Info        = tview.NewTextView()
	PacketInfo  = tview.NewTextView()
)

/*
const banner = ` ____  _  _   _                _                  _ _
|  _ \| || | | |__  _ __   ___| |_      _ __ ___ (_) |_ _ __ ___
| | | | || |_| '_ \| '_ \ / _ \ __|____| '_ ` + "`" + ` _ \| | __| '_ ` + "`" + ` _ \
| |_| |__   _| |_) | | | |  __/ ||_____| | | | | | | |_| | | | | |
|____/   |_| |_.__/|_| |_|\___|\__|    |_| |_| |_|_|\__|_| |_| |_|
 [ TLS: using embedded X509 pair ]   +----Ready ----------------  `
*/

const banner = ` ____  _  _   _                _                  _ _             
 |  _ \| || | | |__  _ __   ___| |_      _ __ ___ (_) |_ _ __ ___  
 | | | | || |_| '_ \| '_ \ / _ \ __|____| '_ ` + "`" + ` _ \| | __| '_ ` + "`" + ` _ \ 
 | |_| |__   _| |_) | | | |  __/ ||_____| | | | | | | |_| | | | | |
 |____/   |_| |_.__/|_| |_|\___|\__|    |_| |_| |_|_|\__|_| |_| |_|`

func init() {
	Pages = tview.NewPages()
	PacketList = PacketList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		UpdatePacketInfo(int64(index))
	})
	LogList = tview.NewList().AddItem("Switch to packet list", "Press to switch to packet list", 'p', func() {
		Pages.SwitchToPage("Packets")
	}).
		AddItem("Quit", "Press to exit", 'q', func() {
			App.Stop()
		})
	TreeLogList = tview.NewTreeView().SetRoot(TreeRoot).SetCurrentNode(TreeRoot).SetGraphics(false).SetTopLevel(1).SetPrefixes([]string{"[red]* ", "[darkcyan]- ", "[darkmagenta]- "})
	TvPackets = []*tview.TextView{}
	Info = tview.NewTextView()
	PacketInfo = tview.NewTextView().SetDynamicColors(true).SetDoneFunc(func(key tcell.Key) {
		App.SetFocus(PacketList)
	})
}

var SelectedPacket int64 = -1

func UpdatePacketInfo(selectable int64) {
	if SelectedPacket == selectable {
		SelectedPacket = -1
	} else {
		SelectedPacket = selectable
		//App.SetFocus(PacketInfo)
	}
	if SelectedPacket != -1 {
		PacketInfo.SetText(bgspacket.PacketHistory[SelectedPacket].Packet.String)
	} else {
		PacketInfo.SetText("")
	}
}

func Preamble() {
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println("Starting up...")
}

func PreStart() {
	fmt.Printf(Fbanner())
	for _, entry := range log.LogEntries {
		log.PrintLogEntry(entry)
	}
	time.Sleep(3 * time.Second)
}

func PaddedMessageTypeString(content string) string {
	messageTypeHeader := content
	messageTypeHeaderLen := len(messageTypeHeader)
	paddedMessageTypeHeader := fmt.Sprintf("%s%s", messageTypeHeader, strings.Repeat(" ", services.LonguestMethodName-messageTypeHeaderLen))
	return paddedMessageTypeHeader
}

func StartUi() {
	//go UpdateLogList()

	logoBox := tview.NewTextView()
	logoBox.SetDynamicColors(true).SetChangedFunc(func() { App.Draw() })

	go func() {
		w := tview.ANSIWriter(logoBox)
		if _, err := io.Copy(w, bytes.NewReader([]byte(Fbanner()))); err != nil {
			panic(err)
		}
	}()

	Info = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetChangedFunc(func() {
			App.Draw()
		}).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			Pages.SwitchToPage(added[0])
		})

	bannerSeparator := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(false).
		SetWrap(false).SetSize(1, 0)
	//SetBackgroundColor(tcell.ColorWhite)
	//_, _, width, _ := bannerSeparator.Box.GetRect()
	separatorChar := strings.Repeat("_", 500)
	fmt.Fprintf(bannerSeparator, separatorChar)

	packetheader := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			Pages.SwitchToPage(added[0])
		})

	fmt.Fprintf(packetheader, `      [darkcyan]No.[white]     [darkcyan]Time[white]      [darkcyan]Source[white]  [darkcyan]RID[white] [darkcyan]Kind[white]      [darkcyan]%s[white]  [darkcyan]bgs.protocol.[cyan]ServiceName[white]`, PaddedMessageTypeString("MessageType"))

	Pages.AddPage("Logs", TreeLogList, true, true)
	fmt.Fprintf(Info, `%d ["%d"][darkcyan]%s[white][""]  `, 1, 0, "Logs")

	Pages.AddPage("Packets", PacketList, true, true)
	fmt.Fprintf(Info, `%d ["%d"][darkcyan]%s[white][""]  `, 2, 1, "Packets")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(logoBox, 5, 1, false).
		AddItem(bannerSeparator, 1, 1, false).
		AddItem(packetheader, 1, 1, false).
		AddItem(Pages, 0, 1, true).
		AddItem(bannerSeparator, 1, 1, false).
		AddItem(PacketInfo, 0, 1, false).
		AddItem(Info, 1, 1, false)

	if err := App.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

func Fbanner() string {
	from := pterm.NewRGB(0, 255, 255)
	to := pterm.NewRGB(255, 0, 255)
	str := banner
	strs := strings.Split(str, "")
	var fadeInfo string
	for i := 0; i < len(str); i++ {
		fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
	}
	return pterm.DefaultCenter.WithCenterEachLineSeparately().Sprintf(fadeInfo)
}

type node struct {
	text     string
	expand   bool
	selected func()
	children []*node
}

func UpdateLogList() {
	/*last_logentries_len := 0
	log.Info(nil, "POUH c'est de la bombe")
	for {
		time.Sleep(1 * time.Second)
		cur_len_logentries := len(log.LogEntries)
		if cur_len_logentries > last_logentries_len {
			for i := last_logentries_len; i < cur_len_logentries; i++ {
				App.QueueUpdateDraw(func() {
					TreeRoot.
						AddChild(tview.NewTreeNode(log.LogEntries[i].Level.String()).SetExpanded(true).
							AddChild(tview.NewTreeNode(log.PrintLogEntry(log.LogEntries[i]))).SetSelectable(true).SetExpanded(true))
				})
				last_logentries_len = cur_len_logentries
			}
		}
	}*/

}
