b:
	go build main.go
	mv main gpac
	cp gpac.gconf /etc/gpac.gconf
	install gpac /usr/bin
install:
	@echo "run make b to install gpac"
s:
	@echo "put the path to repo in gpac.conf (it ends with a /)"


