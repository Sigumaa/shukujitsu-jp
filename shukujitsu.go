// Package shukujitsu は内閣府が提供している祝日一覧 CSV ファイルを取得・解析します。
package shukujitsu

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

// Entry は祝日一日分の情報を保持する構造体です。
type Entry struct {
	YMD  string
	Name string
}

const csvURL = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

// AllEntries は内閣府ウェブサイトから祝日 CSV を取得して Entry スライスに変換します。
func AllEntries() ([]Entry, error) {
	resp, err := http.Get(csvURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get csv: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read csv: %w", err)
	}
	records, err := csv.NewReader(transform.NewReader(bytes.NewReader(body), japanese.ShiftJIS.NewDecoder())).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse csv: %w", err)
	}
	var entries []Entry
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) != 2 {
			return nil, fmt.Errorf("invalid record: %v", record)
		}
		entries = append(entries, Entry{YMD: record[0], Name: record[1]})
	}
	return entries, nil
}
