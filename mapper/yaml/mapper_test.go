package yaml

import (
	"errors"
	"fmt"
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

		mapper := Mapper[test]{}
		conf, err := mapper.Map()
		fmt.Printf("%+v", err)
		errrr := errors.New("no such file or directory")
		if errors.As(err, errrr) {
			fmt.Printf("HELLO")
		}
		fmt.Println(conf)
	})
}
