b:
	go build main.go
	mv main gpac
	cp gpac.conf /etc/gpac.conf
	install gpac /usr/bin
s:
	@echo "put the path to repo in gpac.conf (it ends with a /)"


