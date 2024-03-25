package yaml

import (
	"config-mapper/mapper"
	"github.com/go-yaml/yaml"
	"os"
	"strings"
)

const (
	defaultConfigDirectoryPath = "../config"
	defaultFileName            = "config"
	defaultProfileDelimiter    = "-"
	extension                  = "yaml"
	shortenExtension           = "yml"
)

type Mapper[ConfigFileType any] struct {
	phase            string
	directoryPath    string
	fileName         string
	profileDelimiter string
}

func NewDefaultYamlConfigurationMapper[ConfigFileType any]() Mapper[ConfigFileType] {
	return Mapper[ConfigFileType]{
		phase:            "",
		directoryPath:    defaultConfigDirectoryPath,
		fileName:         defaultFileName,
		profileDelimiter: defaultProfileDelimiter,
	}
}

func (m Mapper[ConfigFileType]) Map() (ConfigFileType, error) {
	var emptyConf ConfigFileType
	var config ConfigFileType
	fullPath := mapper.FullFilePath(m.directoryPath, m.fileName, "", "", extension)
	file, err := readYamlFile(fullPath)
	if err != nil {
		return emptyConf, err
	}
	if err = yaml.Unmarshal(file, &config); err != nil {
		return emptyConf, err
	}
	if m.phase == "" {
		return config, nil
	}
	// 중복 코드 제거 필요
	fullPath = mapper.FullFilePath(m.directoryPath, m.fileName, m.profileDelimiter, m.phase, extension)
	file, err = readYamlFile(fullPath)
	if err != nil {
		return emptyConf, err
	}
	if err = yaml.Unmarshal(file, &config); err != nil {
		return emptyConf, err
	}
	return config, nil
}

func readYamlFile(fullPath string) ([]byte, error) {
	file, err := os.ReadFile(fullPath)
	if err == nil {
		return file, nil
	}
	var shortenExtensionFullPath string
	if os.IsNotExist(err) {
		shortenExtensionFullPath = strings.Replace(fullPath, extension, shortenExtension, 1)
	}
	return os.ReadFile(shortenExtensionFullPath)
}
