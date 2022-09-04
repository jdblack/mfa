
build:
	go build -o bin/mfa *.go

mac_package:
	rm -rf mfa.app
	fyne package
	hdiutil create -volname MFA -srcfolder mfa.app -ov -format UDZO mfa.dmg





