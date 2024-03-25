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
	var emptyConfig ConfigFileType
	var config ConfigFileType

	if err := m.mapFileContents(&config, "", ""); err != nil {
		return emptyConfig, err
	}
	if m.phase == "" {
		return config, nil
	}
	if err := m.mapFileContents(&config, m.profileDelimiter, m.phase); err != nil {
		return emptyConfig, err
	}
	return config, nil
}

func (m Mapper[ConfigFileType]) mapFileContents(config *ConfigFileType, delimiter, phase string) error {
	fullPath := mapper.FullFilePath(m.directoryPath, m.fileName, delimiter, phase, extension)
	file, err := readYamlFile(fullPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(file, &config); err != nil {
		return err
	}
	return nil
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
