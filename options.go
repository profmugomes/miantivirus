// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.bluice.com.br

package main

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/bluiceoficial/blucolumnview"
	"github.com/bluiceoficial/bludialogbox"
	"github.com/bluiceoficial/blunumericentry"
	"github.com/bluiceoficial/blusettings"
	"github.com/bluiceoficial/blusmartflow"

	c "bluiceoficial/miantivirus/controls"
)

func SelectRowsToStrings(rows []blucolumnview.SelectRow) []string {
	out := make([]string, 0)

	for _, r := range rows {
		out = append(out, r.Data...)
	}

	return out
}

func showOptions(app fyne.App) {
	c.LoadTranslations()

	window := app.NewWindow(c.T("Options"))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(500, 600))

	bluconfig, _ := blusettings.Load("miantivirus", true)

	flowGeral := blusmartflow.New()

	lblEscanear := widget.NewLabel(c.T("Scan"))
	lblEscanear.TextStyle = fyne.TextStyle{Bold: true}
	chkOptions := widget.NewCheckGroup([]string{
		c.T("1) Detect PUA (Potentially Unwanted Application)"),
		c.T("2) Heuristic scanning"),
		c.T("3) Scan hidden files"),
		c.T("4) Check symbolic files"),
		c.T("5) Check symbolic folders"),
		c.T("6) Scan compressed files"),
		c.T("7) Scan Email"),
	}, nil)

	flowGeral.AddRow(lblEscanear)
	flowGeral.AddRow(chkOptions)

	lblTamanho := widget.NewLabel(c.T("File Size for Scanning"))
	lblTamanho.TextStyle = fyne.TextStyle{Bold: true}
	ctnTamanho, txtTamanho := blunumericentry.NewBluNumericEntryWithButtons(0, 10000000, 0)

	flowGeral.AddRow(lblTamanho)
	flowGeral.AddRow(ctnTamanho)

	flowIgnorar := blusmartflow.New()

	cvIgnorarPastas := blucolumnview.NewColumnView([]string{c.T("Ignore Folders")}, []float32{38, 400}, true)
	btnIgnorarPastasAdd := widget.NewButton(c.T("Add"), func() {
		bludialogbox.NewSelectDirectory(app, c.T("Select Directory"), true, func(s []string) {
			if len(s) > 0 {
				for _, pathname := range s {
					cvIgnorarPastas.AddRow([]string{pathname})
				}
			}
		})
	})

	btnIgnorarPastasRemove := widget.NewButton(c.T("Remove"), func() {
		cvIgnorarPastas.RemoveSelected()
	})

	flowIgnorar.AddColumn(btnIgnorarPastasAdd, btnIgnorarPastasRemove)
	flowIgnorar.AddRow(cvIgnorarPastas)
	flowIgnorar.Resize(cvIgnorarPastas, window.Canvas().Size().Width, 157)

	separator1 := widget.NewSeparator()
	flowIgnorar.AddRow(separator1)

	cvIgnorarArquivos := blucolumnview.NewColumnView([]string{c.T("Ignore Files")}, []float32{38, 400}, true)
	btnIgnorarArquivosAdd := widget.NewButton(c.T("Add"), func() {
		bludialogbox.NewOpenFile(app, c.T("Open File"), []string{}, true, func(s []string) {
			if len(s) > 0 {
				for _, filename := range s {
					cvIgnorarArquivos.AddRow([]string{filename})
				}
			}
		})
	})

	btnIgnorarArquivosRemove := widget.NewButton(c.T("Remove"), func() {
		cvIgnorarArquivos.RemoveSelected()
	})
	flowIgnorar.AddColumn(btnIgnorarArquivosAdd, btnIgnorarArquivosRemove)
	flowIgnorar.AddRow(cvIgnorarArquivos)
	flowIgnorar.Resize(cvIgnorarArquivos, window.Canvas().Size().Width, 157)

	tabs := container.NewAppTabs(
		container.NewTabItem(c.T("General"), flowGeral.Container),
		container.NewTabItem(c.T("Ignore"), flowIgnorar.Container),
	)

	// Load
	options := bluconfig.GetStringSlice("options", []string{})

	var value []string
	for _, row := range options {
		if slices.Contains(chkOptions.Options, row) {
			value = append(value, row)
		}
	}

	chkOptions.SetSelected(value)
	txtTamanho.SetValue(bluconfig.GetInt("filesize", 0))

	ignorefolders := bluconfig.GetStringSlice("ignorefolders", []string{})

	for _, row := range ignorefolders {
		if row != "" {
			cvIgnorarPastas.AddRow([]string{row})
		}
	}

	ignorefiles := bluconfig.GetStringSlice("ignorefiles", []string{})

	for _, row := range ignorefiles {
		if row != "" {
			cvIgnorarArquivos.AddRow([]string{row})
		}
	}

	flowIgnorar.AddRow(widget.NewLabel(" "))

	// Save
	btnSave := widget.NewButton(c.T("Save"), func() {
		bluconfig.SetStringSlice("options", chkOptions.Selected)
		bluconfig.SetInt("filesize", txtTamanho.GetValue())

		bluconfig.SetStringSlice("ignorefolders", SelectRowsToStrings(cvIgnorarPastas.ListAll()))
		bluconfig.SetStringSlice("ignorefiles", SelectRowsToStrings(cvIgnorarArquivos.ListAll()))

		bluconfig.Save()
		window.Close()
	})

	window.SetContent(container.NewVBox(
		tabs,
		container.NewHBox(
			layout.NewSpacer(),
			btnSave,
			layout.NewSpacer(),
		),
	))
	window.Show()
}
