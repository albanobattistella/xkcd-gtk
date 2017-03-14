package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rkoesters/xkcd"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
)

// Window is the main application window.
type Window struct {
	comic      *xkcd.Comic
	win        *gtk.ApplicationWindow
	hdr        *gtk.HeaderBar
	previous   *gtk.Button
	next       *gtk.Button
	rand       *gtk.Button
	img        *gtk.Image
	properties *PropertiesDialog
}

// New creates a new XKCD viewer window.
func NewWindow(app *Application) (*Window, error) {
	var err error

	w := new(Window)

	w.win, err = gtk.ApplicationWindowNew(app.GtkApp)
	if err != nil {
		return nil, err
	}
	w.win.SetDefaultSize(1000, 800)
	w.win.Connect("destroy", w.Close)

	// Create HeaderBar
	w.hdr, err = gtk.HeaderBarNew()
	if err != nil {
		return nil, err
	}
	w.hdr.SetTitle("XKCD Viewer")
	w.hdr.SetShowCloseButton(true)

	navBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	navBoxStyleContext, err := navBox.GetStyleContext()
	if err != nil {
		return nil, err
	}
	navBoxStyleContext.AddClass("linked")

	w.previous, err = gtk.ButtonNewFromIconName("go-previous-symbolic", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	w.previous.Connect("clicked", w.PreviousComic)
	navBox.Add(w.previous)

	w.next, err = gtk.ButtonNewFromIconName("go-next-symbolic", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	w.next.Connect("clicked", w.NextComic)
	navBox.Add(w.next)

	w.hdr.PackStart(navBox)

	w.rand, err = gtk.ButtonNewFromIconName("media-playlist-shuffle-symbolic", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	w.rand.Connect("clicked", w.RandomComic)
	w.hdr.PackStart(w.rand)

	// Create the menu
	menuBtn, err := gtk.MenuButtonNew()
	if err != nil {
		return nil, err
	}
	cogImg, err := gtk.ImageNewFromIconName("open-menu", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	menuBtn.SetImage(cogImg)

	menu, err := gtk.MenuNew()
	if err != nil {
		return nil, err
	}

	menuGotoNewest, err := gtk.MenuItemNewWithLabel("Go to Newest Comic")
	if err != nil {
		return nil, err
	}
	menuGotoNewest.Connect("activate", w.GotoNewest)
	menu.Add(menuGotoNewest)
	menuGoto, err := gtk.MenuItemNewWithLabel("Go to...")
	if err != nil {
		return nil, err
	}
	menuGoto.Connect("activate", w.ShowGoto)
	menu.Add(menuGoto)
	menuProp, err := gtk.MenuItemNewWithLabel("Properties")
	if err != nil {
		return nil, err
	}
	menuProp.Connect("activate", w.ShowProperties)
	menu.Add(menuProp)
	menuNewWindow, err := gtk.MenuItemNewWithLabel("New Window")
	if err != nil {
		return nil, err
	}
	menuNewWindow.Connect("activate", app.Activate)
	menu.Add(menuNewWindow)
	menuSep, err := gtk.SeparatorMenuItemNew()
	if err != nil {
		return nil, err
	}
	menu.Add(menuSep)
	menuAbout, err := gtk.MenuItemNewWithLabel("About")
	if err != nil {
		return nil, err
	}
	menuAbout.Connect("activate", ShowAboutDialog)
	menu.Add(menuAbout)
	menuBtn.SetPopup(menu)
	menu.ShowAll()

	w.hdr.PackEnd(menuBtn)

	searchBtn, err := gtk.ButtonNewFromIconName("edit-find", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	w.hdr.PackEnd(searchBtn)

	w.hdr.ShowAll()
	w.win.SetTitlebar(w.hdr)

	// Create main part of window.
	scwin, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	scwin.SetSizeRequest(400, 300)

	w.img, err = gtk.ImageNew()
	if err != nil {
		return nil, err
	}
	scwin.Add(w.img)
	scwin.ShowAll()
	w.win.Add(scwin)

	return w, nil
}

// PreviousComic sets the current comic to the previous comic.
func (w *Window) PreviousComic() {
	w.SetComic(w.comic.Num - 1)
}

// NextComic sets the current comic to the next comic.
func (w *Window) NextComic() {
	w.SetComic(w.comic.Num + 1)
}

// RandomComic sets the current comic to a random comic.
func (w *Window) RandomComic() {
	newestComic, err := GetNewestComicInfo()
	if err != nil {
		log.Print(err)
	}
	if newestComic.Num <= 0 {
		w.SetComic(newestComic.Num)
	} else {
		w.SetComic(rand.Intn(newestComic.Num) + 1)
	}
}

// SetComic sets the current comic to the given comic.
func (w *Window) SetComic(n int) {
	// Make it clear that we are loading a comic.
	w.hdr.SetTitle("Loading comic...")
	w.hdr.SetSubtitle(strconv.Itoa(n))
	w.previous.SetSensitive(false)
	w.next.SetSensitive(false)
	w.rand.SetSensitive(false)

	go func() {
		var err error
		w.comic, err = GetComicInfo(n)
		if err != nil {
			log.Printf("error downloading comic info: %v", n)
		}

		err = DownloadComicImage(n)
		if err != nil {
			log.Printf("error downloading comic image: %v", n)
		}

		// Add the DisplayComic function to the event loop so our UI
		// gets updated with the new comic.
		glib.IdleAdd(w.DisplayComic)
	}()
}

// DisplayComic updates the UI to show the contents of w.comic
func (w *Window) DisplayComic() {
	w.hdr.SetTitle(w.comic.Title)
	w.hdr.SetSubtitle(strconv.Itoa(w.comic.Num))
	w.img.SetFromFile(getComicImagePath(w.comic.Num))
	w.img.SetTooltipText(w.comic.Alt)

	// Enable/disable previous button.
	if w.comic.Num > 1 {
		w.previous.SetSensitive(true)
	} else {
		w.previous.SetSensitive(false)
	}

	// Enable/disable next button.
	newest, err := GetNewestComicInfo()
	if err != nil {
		log.Print(err)
	}
	if w.comic.Num < newest.Num {
		w.next.SetSensitive(true)
	} else {
		w.next.SetSensitive(false)
	}

	// Enable random button.
	w.rand.SetSensitive(true)

	if w.properties != nil {
		w.properties.Update()
	}
}

func (w *Window) ShowProperties() {
	var err error
	if w.properties == nil {
		w.properties, err = NewPropertiesDialog(w)
		if err != nil {
			log.Print(err)
			return
		}
	}
	w.properties.Present()
}

func (w *Window) ShowGoto() {
	gt, err := NewGoto(w)
	if err != nil {
		log.Print(err)
		return
	}
	defer gt.dialog.Close()
	status := gt.dialog.Run()
	if status == 1 {
		input, err := gt.entry.GetText()
		if err != nil {
			log.Print(err)
			return
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Print(err)
			return
		}
		w.SetComic(number)
	}
}

func (w *Window) GotoNewest() {
	// Make it clear that we are checking for a new comic.
	w.hdr.SetTitle("Checking for new comic...")
	w.previous.SetSensitive(false)
	w.next.SetSensitive(false)
	w.rand.SetSensitive(false)

	newestComic, err := GetNewestComicInfo()
	if err != nil {
		log.Print(err)
	}
	w.SetComic(newestComic.Num)
}

func (w *Window) Close() {
	if w.properties != nil {
		w.properties.dialog.Close()
	}

	// Remember what comic we were viewing.
	err := ioutil.WriteFile(filepath.Join(cacheDir(), "latest"), []byte(strconv.Itoa(w.comic.Num)), 0666)
	if err != nil {
		log.Print("error writing latest comic file")
	}
}
