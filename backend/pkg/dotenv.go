package pkg

import (
	"bufio"
	"os"
	"strings"
)

func LoadDotenv(path string) error {
	envfile, err := os.Open(path)

	if err != nil {
		return err
	}

	defer func() {
		CheckError(envfile.Close())
	}()

	scanner := bufio.NewScanner(envfile)

	if err = scanner.Err(); err != nil {
		return err
	}

	inEnv := false

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "#") && text[2:] == "Go" {
			inEnv = true
			continue
		}

		if inEnv {
			if text == "" {
				inEnv = false
				continue
			}

			line := strings.SplitN(text, "=", 2)
			os.Setenv(line[0], line[1])
		}
	}

	return nil
}
