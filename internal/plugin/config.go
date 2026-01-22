package plugin

import (
	"flag"
	"fmt"

	"github.com/einride/protoc-gen-go-aip-test/internal/util"
)

// APIMode represents the protobuf API mode for generated code.
type APIMode int

const (
	// APIModeOpen generates code for the Open Struct API (default).
	APIModeOpen APIMode = iota
	// APIModeOpaque generates code for the Opaque API.
	APIModeOpaque
)

func (m APIMode) String() string {
	switch m {
	case APIModeOpen:
		return "API_OPEN"
	case APIModeOpaque:
		return "API_OPAQUE"
	default:
		return ""
	}
}

func (m *APIMode) Set(s string) error {
	switch s {
	case "API_OPAQUE":
		*m = APIModeOpaque
	case "API_OPEN":
		*m = APIModeOpen
	default:
		return fmt.Errorf("invalid api_mode '%s'. Must be either 'API_OPEN' or 'API_OPAQUE'", s)
	}
	return nil
}

// Config for the protoc-gen-go-aip-cli plugin.
type Config struct {
	// APIMode specifies which protobuf API style to generate code for.
	APIMode APIMode
}

// AddToFlagSet adds the config to a pflag.FlagSet.
func (c *Config) AddToFlagSet(flags *flag.FlagSet) {
	flags.Var(
		&c.APIMode,
		"api_mode",
		"the protobuf api mode. Allowed values are 'API_OPEN' and 'API_OPAQUE'",
	)
}

func (c Config) toUtilAPIMode() util.APIMode {
	switch c.APIMode {
	case APIModeOpen:
		return util.APIModeOpen
	case APIModeOpaque:
		return util.APIModeOpaque
	}
	return util.APIModeOpen
}
