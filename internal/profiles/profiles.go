package profiles

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Profiles struct {
	profiles map[string]SystemProfile
}

func Init() (Profiles, error) {
	p := Profiles{}
	p.profiles = make(map[string]SystemProfile)

	err := p.loadProfiles()
	if err != nil {
		return Profiles{}, err
	}

	return p, nil
}

func (p *Profiles) loadProfiles() error {
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

func (p *Profiles) GetProfiles() map[string]SystemProfile {
	return p.profiles
}

func (p *Profiles) Exist() bool {
	err := p.loadProfiles()
	if err != nil {
		return false
	}

	return len(p.profiles) > 0
}
