NAME=shadowsocks2
BINDIR=bin
GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s -buildid='
# The -w and -s flags reduce binary sizes by excluding unnecessary symbols and debug info
# The -buildid= flag makes builds reproducible

all: linux macos-amd64 macos-arm64 win64 win32

linux:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

macos-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

macos-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

win64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

win32:
	GOARCH=386 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe


test: test-linux test-macos-amd64 test-macos-arm64 test-win64 test-win32

test-linux:
	GOARCH=amd64 GOOS=linux go test

test-macos-amd64:
	GOARCH=amd64 GOOS=darwin go test

test-macos-arm64:
	GOARCH=arm64 GOOS=darwin go test

test-win64:
	GOARCH=amd64 GOOS=windows go test

test-win32:
	GOARCH=386 GOOS=windows go test

releases: linux macos-amd64 macos-arm64 win64 win32
	chmod +x $(BINDIR)/$(NAME)-*
	tar czf $(BINDIR)/$(NAME)-linux.tgz -C $(BINDIR) $(NAME)-linux
	gzip $(BINDIR)/$(NAME)-linux
	gzip $(BINDIR)/$(NAME)-macos-amd64
	gzip $(BINDIR)/$(NAME)-macos-arm64
	zip -m -j $(BINDIR)/$(NAME)-win32.zip $(BINDIR)/$(NAME)-win32.exe
	zip -m -j $(BINDIR)/$(NAME)-win64.zip $(BINDIR)/$(NAME)-win64.exe

clean:
	rm $(BINDIR)/*

# Remove trailing {} from the release upload url
GITHUB_UPLOAD_URL=$(shell echo $${GITHUB_RELEASE_UPLOAD_URL%\{*})

upload: releases
	curl -H "Authorization: token $(GITHUB_TOKEN)" -H "Content-Type: application/gzip" --data-binary @$(BINDIR)/$(NAME)-linux.tgz  "$(GITHUB_UPLOAD_URL)?name=$(NAME)-linux.tgz"
	curl -H "Authorization: token $(GITHUB_TOKEN)" -H "Content-Type: application/gzip" --data-binary @$(BINDIR)/$(NAME)-linux.gz  "$(GITHUB_UPLOAD_URL)?name=$(NAME)-linux.gz"
	curl -H "Authorization: token $(GITHUB_TOKEN)" -H "Content-Type: application/gzip" --data-binary @$(BINDIR)/$(NAME)-macos-amd64.gz  "$(GITHUB_UPLOAD_URL)?name=$(NAME)-macos-amd64.gz"
	curl -H "Authorization: token $(GITHUB_TOKEN)" -H "Content-Type: application/gzip" --data-binary @$(BINDIR)/$(NAME)-macos-arm64.gz  "$(GITHUB_UPLOAD_URL)?name=$(NAME)-macos-arm64.gz"
	curl -H "Authorization: token $(GITHUB_TOKEN)" -H "Content-Type: application/zip"  --data-binary @$(BINDIR)/$(NAME)-win64.zip "$(GITHUB_UPLOAD_URL)?name=$(NAME)-win64.zip"
	curl -H "Authorization: token $(GITHUB_TOKEN)" -H "Content-Type: application/zip"  --data-binary @$(BINDIR)/$(NAME)-win32.zip "$(GITHUB_UPLOAD_URL)?name=$(NAME)-win32.zip"
