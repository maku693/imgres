package resizer

import "os"

func InFile(name string) (*os.File, error) {
	file := os.Stdin
	if name != "" {
		f, err := os.Open(name)
		if err != nil {
			return nil, err
		}
		file = f
	}
	return file, nil
}

func OutFile(name string) (*os.File, error) {
	file := os.Stdout
	if name != "" {
		f, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		file = f
	}
	return file, nil
}
