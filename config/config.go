package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// DriversDB struct
type DriversDB struct {
	Host   string
	Port   string
	DBName string
}

// Home struct
type Home struct {
	Name    string
	Address string
	Zip     string
	City    string
	Point   [2]float64
}

// Config struct
type Config struct {
	Title   string
	Drivers DriversDB `toml:"drivers-database"`
	HomeCfg Home      `toml:"home"`
}

const cfgfilename = "alertrack.toml"

// NewConfig creates new config instance
func NewConfig() (*Config, error) {
	// Content object
	var conf Config
	// Check config file exists
	_, err := os.Stat(cfgfilename)
	if err != nil {
		log.Printf("Error no config file. Error: %v", err)
		return nil, err
	}
	// Open & read config file
	fileContent, err := ioutil.ReadFile(cfgfilename)
	if err != nil {
		log.Printf("Error reading config file. Error: %v", err)
		return nil, err
	}
	// Decode file content
	_, err = toml.Decode(string(fileContent), &conf)
	if err != nil {
		log.Printf("Error parsing config file. Error: %v", err)
		return nil, err
	}
	return &conf, nil
}

// GetDriversDBConfig returns drivers db configuration
func (c *Config) GetDriversDBConfig() *DriversDB {
	return &c.Drivers
}

// GetHomeConfig returns home data location
func (c *Config) GetHomeConfig() *Home {
	return &c.HomeCfg
}
