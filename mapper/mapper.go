package mapper

func FullFilePath(directoryPath, fileName, delimiter, phase, extension string) string {
	fullPath := directoryPath + "/" + fileName
	if phase != "" {
		fullPath += delimiter + phase
	}
	fullPath += "." + extension
	return fullPath
}
