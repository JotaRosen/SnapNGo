

build:
	@if [ -a SnapNGo ]; then rm SnapNGo; fi;
	@go build -o SnapNGo
	@cat makeHelp/ascii_art.txt
	@printf "\n"
	@cat makeHelp/help.txt



