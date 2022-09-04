
build:
	go build -o bin/mfa *.go

package:
	rm -rf mfa.app
	fyne package
	hdiutil create -volname MFA -srcfolder mfa.app -ov -format UDZO mfa.dmg





