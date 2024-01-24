package ini

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Ini struct {
	Data map[string]map[string]string
}

func Load(iniFile string, iniFiles ...string) (*Ini, error) {
	ini := &Ini{
		Data: make(map[string]map[string]string),
	}

	if err := ini.parseDataSource(iniFile); err != nil {
		return nil, err
	}

	for _, file := range iniFiles {
		if err := ini.parseDataSource(file); err != nil {
			return nil, err
		}
	}
	return ini, nil
}

func (ini *Ini) parseDataSource(iniFile string) error {
	f, err := os.Open(iniFile)
	if err != nil {
		return err
	}
	defer f.Close()

	var section string
	r := bufio.NewReader(f)
	m := make(map[string]string)

	for {
		byts, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		s := strings.TrimSpace(string(byts))
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "#") || strings.HasPrefix(s, ";") {
			continue
		}

		if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
			i := strings.Index(s, "[")
			j := strings.LastIndex(s, "]")
			section = strings.TrimSpace(s[i+1 : j])
			m = make(map[string]string)
			continue
		}

		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		key := strings.TrimSpace(s[:index])
		if key == "" {
			continue
		}

		value := strings.TrimSpace(s[index+1:])
		if value == "" {
			continue
		}

		pos := strings.Index(value, "\t#")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, " #")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, " ;")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, "\t//")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, " //")
		if pos > -1 {
			value = value[0:pos]
		}

		value = strings.TrimSpace(value)

		if _, ok := ini.Data[section]; ok {
			ini.Data[section][key] = value
		} else {
			m[key] = value
			ini.Data[section] = m
		}
	}
	return nil
}

func (ini *Ini) Read(section, key string) string {
	v, ok := ini.Data[section][key]
	if !ok {
		return ""
	}
	return v
}
