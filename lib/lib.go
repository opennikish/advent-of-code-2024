package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetInput(day int) ([]byte, error) {
	filename := "in.txt"

	download := func() ([]byte, error) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2024/day/%02d/input", day), nil)
		if err != nil {
			return nil, fmt.Errorf("new req: %w", err)
		}

		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: os.Getenv("AOC"),
		})

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("do req: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("unexpected resp: status = %d, body = %s", resp.StatusCode, string(b))
		}

		f, err := os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("create file: %w", err)
		}

		r := io.TeeReader(resp.Body, f)
		bs, err := io.ReadAll(r)
		if err != nil {
			return nil, fmt.Errorf("tee req body: %w", err)
		}

		return bs, nil
	}

	bs, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		return download()
	}

	return bs, nil
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
