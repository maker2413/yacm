package profiles

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Profiles struct {
	profiles map[string]SystemProfile
}

func (p *Profiles) GetProfiles() error {
	files, err := os.ReadDir(viper.GetString("yacm_profiles_dir"))
	if err != nil {
		return err
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yml") || strings.HasSuffix(f.Name(), ".yaml") {
			p.profiles[f.Name()] = SystemProfile{}
		}
	}

	return nil
}

func (p *Profiles) Exist() bool {
	err := p.GetProfiles()
	if err != nil {
		return false
	}

	return len(p.profiles) > 0
}
