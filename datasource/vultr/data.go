package vultr

import (
	"errors"
	"os"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	APIKey string `mapstructure:"api_key" required:"true"`
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct{}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}

	var errs *packersdk.MultiError

	if d.config.APIKey == "" {
		d.config.APIKey = os.Getenv("VULTR_API_KEY")
	}

	if d.config.APIKey == "" {
		errs = packersdk.MultiErrorAppend(err, errors.New("APIKey is required"))
	}

	if errs != nil && len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {}
