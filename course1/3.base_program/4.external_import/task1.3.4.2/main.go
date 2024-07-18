package main

import (
	"fmt"
	"github.com/ksrof/gocolors"
)

func main() {
	fmt.Println(ColorizeRed("Красный"))
	fmt.Println(ColorizeGreen("Зеленный"))
	fmt.Println(ColorizeBlue("Синий"))
	fmt.Println(ColorizeYellow("Желтый"))
	fmt.Println(ColorizeMagenta("Пурпурный"))
	fmt.Println(ColorizeCyan("Голубой"))
	fmt.Println(ColorizeWhite("Белый"))
	fmt.Println(ColorizeCustom("Кастомный", 200, 149, 50))
}

func ColorizeRed(a string) string {
	return gocolors.Red(a, "")
}

func ColorizeGreen(a string) string {
	return gocolors.Green(a, "")
}

func ColorizeBlue(a string) string {
	return gocolors.Blue(a, "")
}

func ColorizeYellow(a string) string {
	return gocolors.BrightYellow(a, "")
}

func ColorizeMagenta(a string) string {
	return gocolors.Magenta(a, "")
}

func ColorizeCyan(a string) string {
	return gocolors.Cyan(a, "")
}

func ColorizeWhite(a string) string {
	return gocolors.BrightWhite(a, "")
}

func ColorizeCustom(a string, r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, a)
}
