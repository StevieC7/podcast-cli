package persistence

import "github.com/StevieC7/podcast-cli/internal/parsing"

func CheckIfShowFolderExists(showNameAsPath parsing.ValidPath) (bool, string)
func CreateShowFolder(showNameAsPath parsing.ValidPath) (string, error)

func CheckIfEpisodeFileExists(showNameAsPath, episodeNameAsPath parsing.ValidPath) (bool, string)
func SaveEpisodeFile(showNameAsPath, episodeNameAsPath parsing.ValidPath) (string, error)
