// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.bluice.com.br

package main

import (
	c "bluiceoficial/miantivirus/controls"
	"os"

	"github.com/bluiceoficial/blurun"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/bluiceoficial/bludialogbox"
	"github.com/bluiceoficial/blusmartflow"
)

func showUpdateDB(app fyne.App) {
	c.LoadTranslations()
	window := app.NewWindow(c.T("Update Database"))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(400, 400))

	flow := blusmartflow.New()

	lblInfo := widget.NewLabel(c.T("Checking for updates..."))
	lblInfo.TextStyle = fyne.TextStyle{Bold: true}

	txtResult := widget.NewEntry()
	txtResult.MultiLine = true
	txtResult.Scroll = fyne.ScrollBoth
	txtResult.Wrapping = fyne.TextWrapBreak

	flow.AddRow(lblInfo)
	flow.AddRow(txtResult)
	flow.Resize(txtResult, window.Canvas().Size().Width, 279)

	go func() {
		//
		s := blurun.New("pkexec sh -c 'killall freshclam;freshclam'")
		pathHome, _ := os.UserHomeDir()
		s.SetDir(pathHome)
		s.AddEnv("teste", "abc")

		s.OnStderr(func(s string) {
			fyne.Do(func() {
				txtResult.Text += s + "\n"
				txtResult.Refresh()
			})
		})
		s.OnStdout(func(s string) {
			fyne.Do(func() {
				txtResult.Text += s + "\n"
				txtResult.Refresh()
			})
		})

		if err := s.Run(); err != nil {
			bludialogbox.NewAlert(app, c.T("Update Database"), err.Error(), true, "Ok")
		}

		fyne.Do(func() {
			lblInfo.SetText(c.T("Finish"))
		})
	}()
	window.SetContent(flow.Container)
	window.Show()
}
