package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func getTodayURL() string {
	now := time.Now()

	dateStr := now.Format("02_01_2006")
	yearStr := now.Format("2006")

	result := fmt.Sprintf("http://cdn.iz.ru/sites/default/files/pdf/%s/%s.pdf", yearStr, dateStr)

	return result
}

func getCurYear() string {
	now := time.Now()
	return strconv.Itoa(now.Year())
}

func getCurMonth() string {
	now := time.Now()
	return now.Month().String()
}

func getWeekOfMonth() string {
	now := time.Now()
	day := now.Day()

	week := (day-1)/7 + 1

	return fmt.Sprintf("Week_%v", week)
}

func getTodayFileName() string {
	now := time.Now()

	dateStr := now.Format("02_01_2006")
	result := fmt.Sprintf("%s.pdf", dateStr)

	return result
}

func main() {
	url := getTodayURL()

	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, "Izvestia_Newspaper", getCurYear(), getCurMonth(), getWeekOfMonth())

	os.MkdirAll(dir, os.ModePerm)

	filePath := filepath.Join(dir, getTodayFileName())

	err := downloadFile(url, filePath)
	if err != nil {
		panic(err)
	}
}
