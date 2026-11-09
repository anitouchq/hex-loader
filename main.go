package main

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var version = "dev"

func main() {
	fyneApp := app.New()
	window := fyneApp.NewWindow("Hex Loader " + version)
	window.Resize(fyne.NewSize(560, 400))

	status := widget.NewMultiLineEntry()
	status.Wrapping = fyne.TextWrapWord
	status.SetText("Choose an Intel HEX (.hex / .ihex) file to validate its records and checksums.")

	openButton := widget.NewButton("Open HEX file", func() {
		fileDialog := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
			if err != nil {
				status.SetText("Could not open file: " + err.Error())
				return
			}
			if file == nil {
				return
			}
			defer file.Close()

			data, readErr := io.ReadAll(file)
			if readErr != nil {
				status.SetText("Could not read file: " + readErr.Error())
				return
			}

			summary, parseErr := ParseIntelHex(string(data))
			if parseErr != nil {
				status.SetText("Validation failed:\n" + parseErr.Error())
				return
			}

			status.SetText(fmt.Sprintf(
				"Validation successful ✅\n\nFile: %s\nRecords: %d\nData records: %d\nData bytes: %d\nExtended address base: 0x%08X\nEOF record: present",
				file.URI().Name(),
				summary.Records,
				summary.DataRecords,
				summary.DataBytes,
				summary.ExtendedAddressBase,
			))
		}, window)
		fileDialog.Show()
	})

	window.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("Intel HEX Loader", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Load and validate firmware files before using them with your programmer."),
		widget.NewLabel("Version: "+version),
		openButton,
		widget.NewSeparator(),
		status,
		widget.NewButton("Exit", func() { fyneApp.Quit() }),
	))

	window.ShowAndRun()
}
