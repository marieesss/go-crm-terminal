package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config contient toute la configuration de l'application
type Config struct {
	Storage StorageConfig `mapstructure:"storage"`
}

// StorageConfig contient la configuration du stockage
type StorageConfig struct {
	Type string     `mapstructure:"type"`
	JSON JSONConfig `mapstructure:"json"`
	GORM GORMConfig `mapstructure:"gorm"`
}

// JSONConfig contient la config pour le JSONStore
type JSONConfig struct {
	FilePath string `mapstructure:"filepath"`
}

// GORMConfig contient la config pour le GORMStore
type GORMConfig struct {
	DBPath string `mapstructure:"dbpath"`
}

// LoadConfig lit le fichier config.yaml et retourne la configuration
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // Nom du fichier (sans extension)
	viper.SetConfigType("yaml")   // Type de fichier
	viper.AddConfigPath(".")      // Cherche dans le répertoire courant

	// Valeurs par défaut
	viper.SetDefault("storage.type", "memory")
	viper.SetDefault("storage.json.filepath", "contacts.json")
	viper.SetDefault("storage.gorm.dbpath", "contacts.db")

	// Lecture du fichier
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("erreur lecture config: %w", err)
	}

	// Mapping dans la struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("erreur parsing config: %w", err)
	}

	return &config, nil
}
