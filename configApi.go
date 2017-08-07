package main

type Configs struct {
	Configurations []Config
}

type Config struct {
	Version         string
	Exposure        int
	Gal_version     string
	Layout          string
	Related_version string
	Layout_holder   string
	Qcat            string
}

func (c1 *Configs) equals(c2 Configs) (answer bool) {
	answer = true
	for i := range c1.Configurations {
		c1.Configurations[i].equals(c2.Configurations[i])
	}
	return
}

func (c1 *Config) equals(c2 Config) (answer bool) {
	answer = c1.Version == c2.Version && c1.Exposure == c2.Exposure && c1.Gal_version == c2.Gal_version
	answer = answer && c1.Layout == c2.Layout && c1.Related_version == c2.Related_version
	answer = answer && c1.Layout_holder == c2.Layout_holder && c1.Qcat == c2.Qcat
	return
}
