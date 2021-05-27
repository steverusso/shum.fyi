WEBDIR = website
SRVDIR = server

WEBDST = $(SRVDIR)/public
SRVBIN = $(SRVDIR)/shumfyid

.PHONY: website server

default: server

website:
	cd $(WEBDIR) && hugo -d ../$(WEBDST)

server: website
	cd $(SRVDIR) && go build

run: server
	$(SRVBIN) -p 8080

install: server
	sudo sysrc shumfyid_enable="YES"
	sudo cp -f $(SRVBIN) /usr/local/bin/
	sudo cp -f $(SRVDIR)/contrib/shumfyid.rc.d /usr/local/etc/rc.d/shumfyid
	sudo service shumfyid restart

clean:
	rm -rf $(WEBDST) $(SRVBIN)
