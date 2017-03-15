BUILDFLAGS=-v -tags gtk_3_18

xkcd-gtk: *.go
	go build $(BUILDFLAGS)

clean:
	go clean

home-install: xkcd-gtk
	install xkcd-gtk $$HOME/bin
	mkdir -p $$HOME/.local/share/applications
	cp com.ryankoesters.xkcd-gtk.desktop $$HOME/.local/share/applications
	mkdir -p $$HOME/.local/share/icons/hicolor/scalable/apps
	cp xkcd-gtk.svg $$HOME/.local/share/icons/hicolor/scalable/apps

home-uninstall:
	rm $$HOME/bin/xkcd-gtk \
	   $$HOME/.local/share/applications/com.ryankoesters.xkcd-gtk.desktop \
	   $$HOME/.local/share/icons/hicolor/scalable/apps/xkcd-gtk.svg

root-install: xkcd-gtk
	install xkcd-gtk /usr/local/bin
	mkdir -p /usr/share/applications
	cp com.ryankoesters.xkcd-gtk.desktop /usr/share/applications
	mkdir -p /usr/local/share/icons/hicolor/scalable/apps
	cp xkcd-gtk.svg /usr/local/share/icons/hicolor/scalable/apps

root-uninstall:
	rm /usr/local/bin/xkcd-gtk \
	   /usr/share/applications/com.ryankoesters.xkcd-gtk.desktop \
	   /usr/local/share/icons/hicolor/scalable/apps/xkcd-gtk.svg
