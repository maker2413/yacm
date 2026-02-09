package profiles

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type SystemProfiles map[string]Profile

func Init() (SystemProfiles, error) {
	var sp = make(SystemProfiles)

	err := sp.readProfiles()
	if err != nil {
		return sp, err
	}

	return sp, nil
}

func (sp SystemProfiles) readProfiles() error {
	files, err := os.ReadDir(viper.GetString("yacm_profiles_dir"))
	if err != nil {
		return err
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yml") || strings.HasSuffix(f.Name(), ".yaml") {
			sp[f.Name()] = Profile{}
		}
	}

	return nil
}

func (sp SystemProfiles) exists() bool {
	exist := len(sp)
	if exist <= 0 {
		err := sp.readProfiles()
		if err != nil {
			return false
		}

		exist = len(sp)
	}

	return exist > 0
}
