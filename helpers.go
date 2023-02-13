package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"path"
)

func closeBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		ShowError(err)
		Stop()
		return
	}
}

func closeFile(out *os.File) {
	err := out.Close()
	if err != nil {
		ShowError(err)
		Stop()
		return
	}
}

func ShowError(err error) {
	fmt.Printf("%s %s", RedBold("[ERROR]"), Red(err.Error()))
}

func ShowSuccess(message string) {
	fmt.Printf("%s %s", GreenBold("[SUCCESS]"), Green(message))
}

func Stop() {
	fmt.Print("\nPress 'Enter' to close updater...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func Red(text string) string {
	return color.RedString(text)
}

func Green(text string) string {
	return color.GreenString(text)
}

func RedBold(text string) string {
	c := color.New(color.Bold, color.FgRed)
	return c.Sprint(text)
}

func GreenBold(text string) string {
	c := color.New(color.Bold, color.FgGreen)
	return c.Sprint(text)
}

func DownloadFile(url string, dest string) {
	fileName := path.Base(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ShowError(err)
		Stop()
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ShowError(err)
		Stop()
		return
	}
	defer closeBody(resp.Body)
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", dest, fileName), os.O_CREATE|os.O_WRONLY, 0644)
	defer closeFile(f)
	if err != nil {
		ShowError(err)
		Stop()
		return
	}

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		fileName,
	)
	_, err = io.Copy(io.MultiWriter(f, bar), resp.Body)
	if err != nil {
		ShowError(err)
		Stop()
		return
	}
}

func dirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
