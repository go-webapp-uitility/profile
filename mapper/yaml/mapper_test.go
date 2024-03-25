package yaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapper_Map(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		type st struct {
			Top    string `yaml:"top"`
			Bottom string `yaml:"bottom"`
		}
		type str struct {
			String string `yaml:"string"`
			Number int    `yaml:"number"`
			Array  []any  `yaml:"array"`
			St     st     `yaml:"st"`
		}

		type test struct {
			Test str `yaml:"test"`
		}

		mapper := NewDefaultYamlConfigurationMapper[test]()
		mapper.directoryPath = "../test_config/"
		mapper.phase = "dev"
		conf, err := mapper.Map()
		assert.NoError(t, err)
		assert.NotEqual(t, "tmp", conf.Test.St.Top, "variable is not same with default config.")
		assert.Equal(t, "b", conf.Test.St.Bottom)

		assert.Len(t, conf.Test.Array, 3)
		assert.Equal(t, "a", conf.Test.Array[0])
		assert.Equal(t, "b", conf.Test.Array[1])
		assert.Equal(t, "c", conf.Test.Array[2])

		assert.Equal(t, 1234, conf.Test.Number)

		assert.Equal(t, "hello-dev", conf.Test.String, "variable of default config is overwritten by phased config.")
		assert.Equal(t, "tmp-Dev", conf.Test.St.Top, "variable of default config is overwritten by phased config.")
	})
}
