package configuration

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmptyFilenameConfiguration(t *testing.T) {
	t.Parallel()

	cfg, err := Load("")
	require.Error(t, err)
	require.NotNil(t, cfg)
}

func TestEmptyConfiguration(t *testing.T) {
	t.Parallel()

	cfg, err := Load("test_data/empty_config.yaml")
	require.NoError(t, err)
	require.Nil(t, cfg)
}

func TestNotFoundConfiguration(t *testing.T) {
	t.Parallel()

	cfg, err := Load("test_data/not_found_config.yaml")
	require.Error(t, err)
	require.Nil(t, cfg)
}

func TestFullConfiguration(t *testing.T) {
	t.Parallel()

	cfg, err := Load("test_data/config.yaml")
	require.NoError(t, err)
	require.NotNil(t, cfg)
}

func TestConfigurationProperties(t *testing.T) {
	t.Parallel()

	cfg, err := Load("test_data/config.yaml")
	require.NoError(t, err)
	require.NotNil(t, cfg)

	require.Equal(t, cfg.EngineConfiguration.Type, "in_memory")

}
