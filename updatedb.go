// Copyright (C) 2024-2026 Murilo Gomes <profmugomes.com.br>
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.profmugomes.com.br

package main

import (
	c "profmugomes/miantivirus/controls"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/profmugomes/mgdialogbox/v2"
	"github.com/profmugomes/mgrun/v2"
	"github.com/profmugomes/mgsmartflow/v2"
)

func showUpdateDB(app fyne.App) {
	c.LoadTranslations()
	window := app.NewWindow(c.T("Update Database"))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(400, 400))

	flow := mgsmartflow.New()

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
		s := mgrun.New("pkexec sh -c 'killall freshclam;freshclam'")
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
			mgdialogbox.NewAlert(app, c.T("Update Database"), err.Error(), true, "Ok", nil)
		}

		fyne.Do(func() {
			lblInfo.SetText(c.T("Finish"))
		})
	}()
	window.SetContent(flow.Container)
	window.Show()
}
