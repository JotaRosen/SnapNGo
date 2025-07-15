

build:
	@if [ -a SnapNGo ]; then rm SnapNGo; fi;
	@go build -o SnapNGo
	@cat ascii_art.txt
	@printf "\n"
	@cat help.txt



