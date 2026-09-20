// Copyright (c) 2024-2026 Murilo Gomes <profmugomes.com.br>. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

package main

import (
	"fmt"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/profmugomes/mgsmartflow/v2"
)

func showAbout(app fyne.App) {
	w := app.NewWindow("Sobre")
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

	lblLicense2 := widget.NewLabel("PolyForm Perimeter 1.0.1")

	flow.AddColumn(lblLicense1, lblLicense2)
	flow.Resize(lblLicense1, 62, 0)

	txtLicense := widget.NewRichTextFromMarkdown(`
This software is available for commercial and noncommercial use, subject to the terms of the PolyForm Perimeter License 1.0.1.

You may:

* ✔ Use the software for commercial and noncommercial purposes.
* ✔ Inspect and study the source code.
* ✔ Modify the software.
* ✔ Create derivative works based on the software.
* ✔ Redistribute the software and permitted modifications.

You may not:

* ✖ Provide a product that competes with the software.

See the full license terms at LICENSE.md.

This summary is provided for convenience only and does not replace or modify the full license terms.
	`)
	txtLicense.Wrapping = fyne.TextWrapWord

	vBoxLicense := container.NewVScroll(txtLicense)

	flow.AddRow(vBoxLicense)
	flow.Resize(vBoxLicense, w.Canvas().Size().Width, 319)

	w.SetContent(flow.Container)
	w.Show()
}
