// Copyright (C) 2024-2026 Murilo Gomes <profmugomes.com.br>
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.profmugomes.com.br

package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/profmugomes/mgcolumnview/v3"
	"github.com/profmugomes/mgdialogbox/v3"
	"github.com/profmugomes/mgsmartflow/v3"

	c "profmugomes/miantivirus/controls"
)

var a fyne.App

const VERSION_APP string = "6.0.0"

func main() {
	c.LoadTranslations()

	a = app.NewWithID("br.com.profmugomes.miantivirus")
	a.SetIcon(resourceMiantivirusPng)

	window := a.NewWindow("MiAntivirus")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(659, 457))

	mnuTools := fyne.NewMenu(c.T("Tools"),
		fyne.NewMenuItem(c.T("Update Database"), func() {
			showUpdateDB()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("Options"), func() {
			showOptions()
		}),
	)

	mnuAbout := fyne.NewMenu(c.T("About"),
		fyne.NewMenuItem(c.T("Check for Updates"), func() {
			url, _ := url.Parse("https://github.com/profmugomes/miantivirus/releases")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("Support MiAntivirus"), func() {
			url, _ := url.Parse("https://github.com/sponsors/profmugomes")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("About MiAntivirus"), func() {
			showAbout()
		}),
	)
	window.SetMainMenu(fyne.NewMainMenu(mnuTools, mnuAbout))

	flow := mgsmartflow.New()

	ctnSpace1 := widget.NewLabel(" ")

	flow.AddRow(ctnSpace1)
	flow.Resize(ctnSpace1, window.Canvas().Size().Width, 7)

	lstArquivos := mgcolumnview.NewColumnView(
		[]string{c.T("Files")},
		[]float32{38, 400, 100}, true,
	)

	btnAddFile := widget.NewButton(c.T("Add File"), func() {
		mgdialogbox.NewOpenFile(a, c.T("Open Files"), []string{}, true, func(filenames []string) {
			for _, filename := range filenames {
				lstArquivos.AddRow([]string{filename})
			}
		})
	})

	btnAddFolder := widget.NewButton(c.T("Add Folder"), func() {
		mgdialogbox.NewSelectDirectory(a, c.T("Select Directory"), true, func(filenames []string) {
			for _, filename := range filenames {
				lstArquivos.AddRow([]string{filename})
			}
		})
	})

	btnRemoverLista := widget.NewButton(c.T("Remove from List"), func() {
		lstArquivos.RemoveSelected()
	})

	flow.AddColumn(btnAddFile, btnAddFolder, btnRemoverLista)
	flow.Gap(btnAddFile, 7, 17)

	flow.AddRow(lstArquivos)
	flow.Resize(lstArquivos, window.Canvas().Size().Width, 272)

	btnEscanear := widget.NewButton(c.T("Scan"), func() {
		showScan(lstArquivos.ListAll())
	})

	flow.AddColumn(
		layout.NewSpacer(),
		btnEscanear,
		layout.NewSpacer(),
	)

	window.SetContent(flow.Container)
	window.ShowAndRun()
}
