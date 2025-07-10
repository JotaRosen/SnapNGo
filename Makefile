

build:
	if [ -a SnapNGo ]; then rm SnapNGo; fi;
	go build -o SnapNGo
	@cat ascii_art.txt
	@printf "\n\n"
	@cat help.txt



