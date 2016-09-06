package conf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Conf struct {
	conf_map map[string]map[string]string
}

func NewConf() *Conf {
	return &Conf{make(map[string]map[string]string)}
}

func (conf *Conf) Load(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("file %v is cannot open", filename))
	}
	rd := bufio.NewReader(f)
	section_key := ""
	line := ""
	for {
		line, err = rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			line = strings.TrimPrefix(line, "[")
			if !strings.HasSuffix(line, "]") {
				panic(fmt.Errorf("line %v section err", line))
			}
			section_key = strings.TrimSuffix(line, "]")
		} else {
			strs := strings.Split(line, "=")
			if len(strs) != 2 {
				fmt.Println(line)
				panic(fmt.Errorf("line %v dont right", line))
			}
			key := strs[0]
			value := strs[1]
			if _, ok := conf.conf_map[section_key]; !ok {
				conf.conf_map[section_key] = make(map[string]string)
			}
			conf.conf_map[section_key][key] = value
		}
	}
}

func (conf *Conf) GetValue(section_key, key string) string {
	if _, ok := conf.conf_map[section_key]; !ok {
		return ""
	}
	if _, ok := conf.conf_map[section_key][key]; !ok {
		return ""
	}
	return conf.conf_map[section_key][key]
}
