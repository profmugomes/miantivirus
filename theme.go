// Copyright (c) 2024-2026 Murilo Gomes <profmugomes.com.br>. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myDarkTheme struct{}

func (m myDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{28, 28, 28, 255} // Fundo preto
	case theme.ColorNameForeground:
		return color.White
	default:
		return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

func (m myDarkTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (m myDarkTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m myDarkTheme) Size(n fyne.ThemeSizeName) float32 {
	if n == theme.SizeNameText {
		return 16
	}
	return theme.DefaultTheme().Size(n)
}
