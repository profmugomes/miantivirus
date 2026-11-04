// Copyright (C) 2024-2026 Murilo Gomes <profmugomes.com.br>
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.profmugomes.com.br

package main

import (
	"fmt"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/profmugomes/mgsmartflow/v3"
)

func showAbout() {
	w := a.NewWindow("Sobre")
	w.Resize(fyne.NewSize(697, 359))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	flow := mgsmartflow.New()

	lblSoftware := canvas.NewText(fmt.Sprintf("MiAntivirus - Version: %s", VERSION_APP), color.Opaque)
	lblSoftware.TextSize = 18
	lblSoftware.TextStyle.Bold = true

	flow.AddRow(lblSoftware)
	flow.Move(lblSoftware, 7, 7)

	lblDesenvolvedor1 := widget.NewLabel("Desenvolvido por:")
	lblDesenvolvedor1.TextStyle = fyne.TextStyle{Bold: true}
	lblDesenvolvedor2 := widget.NewLabel("Murilo Gomes")

	flow.AddColumn(lblDesenvolvedor1, lblDesenvolvedor2)
	flow.Resize(lblDesenvolvedor1, 142, 0)
	lblSite1 := widget.NewLabel("Site:")
	lblSite1.TextStyle = fyne.TextStyle{Bold: true}

	sURL, _ := url.Parse("https://www.profmugomes.com.br")
	lblSite2 := widget.NewHyperlink("https://www.profmugomes.com.br", sURL)

	flow.AddColumn(lblSite1, lblSite2)
	flow.Resize(lblSite1, 34, 0)

	lblCopyright1 := widget.NewLabel("Copyright (c) 2024-2026 Murilo Gomes <profmugomes.com.br>. All Rights Reserved.")
	lblCopyright1.TextStyle = fyne.TextStyle{Bold: true}
	flow.AddRow(lblCopyright1)

	lblLicense1 := widget.NewLabel("License:")
	lblLicense1.TextStyle = fyne.TextStyle{Bold: true}

	lblLicense2 := widget.NewLabel("GPL-2.0-only")

	flow.AddColumn(lblLicense1, lblLicense2)
	flow.Resize(lblLicense1, 62, 0)

	txtLicense := widget.NewRichTextFromMarkdown(`
MiAntivirus is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, only version 2 of the License.

MiAntivirus is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
	`)
	txtLicense.Wrapping = fyne.TextWrapWord

	vBoxLicense := container.NewVScroll(txtLicense)

	flow.AddRow(vBoxLicense)
	flow.Resize(vBoxLicense, w.Canvas().Size().Width, 319)

	w.SetContent(flow.Container)
	w.Show()
}
