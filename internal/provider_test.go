package internal

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aviatrix/provider-aviatrix/apis"
	"github.com/aviatrix/provider-aviatrix/config"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/onsi/gomega"
	"github.com/upbound/upjet/pkg/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

func TestProviderGoldenFiles(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	myscheme := runtime.NewScheme()
	err := apis.AddToScheme(myscheme)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	decode := serializer.NewCodecFactory(myscheme).
		UniversalDecoder(myscheme.PreferredVersionAllGroups()...).Decode

	type args struct {
		tfJsonFile string
		yamlFile   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "vpc",
			args: args{
				tfJsonFile: "./test-data/vpc.tf.json",
				yamlFile:   "../examples/vpc.yaml",
			},
		}, {
			name: "transit",
			args: args{
				tfJsonFile: "./test-data/transit-gateway.tf.json",
				yamlFile:   "../examples/transit-gateway.yaml",
			},
		}, {
			name: "spoke",
			args: args{
				tfJsonFile: "./test-data/spoke-gateway.tf.json",
				yamlFile:   "../examples/spoke-gateway.yaml",
			},
		}, {
			name: "attachment",
			args: args{
				tfJsonFile: "./test-data/transitattachment.tf.json",
				yamlFile:   "../examples/transitattachment.yaml",
			},
		}, {
			name: "domain",
			args: args{
				tfJsonFile: "./test-data/networkdomain.tf.json",
				yamlFile:   "../examples/networkdomain.yaml",
			},
		}, {
			name: "association",
			args: args{
				tfJsonFile: "./test-data/networkdomainassociation.tf.json",
				yamlFile:   "../examples/networkdomainassociation.yaml",
			},
		}, {
			name: "policy",
			args: args{
				tfJsonFile: "./test-data/networkdomainconnectionpolicy.tf.json",
				yamlFile:   "../examples/networkdomainconnectionpolicy.yaml",
			},
		}, {
			name: "snat",
			args: args{
				tfJsonFile: "./test-data/snat.tf.json",
				yamlFile:   "../examples/snat.yaml",
			},
		}, {
			name: "dnat",
			args: args{
				tfJsonFile: "./test-data/dnat.tf.json",
				yamlFile:   "../examples/dnat.yaml",
			},
		}, {
			name: "peering",
			args: args{
				tfJsonFile: "./test-data/gatewaypeering.tf.json",
				yamlFile:   "../examples/gatewaypeering.yaml",
			},
		}, {
			name: "account",
			args: args{
				tfJsonFile: "./test-data/partial/account.tf.json",
				yamlFile:   "../examples/account.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream, err := ioutil.ReadFile(tt.args.yamlFile)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			obj, _, err := decode(stream, nil, nil)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			res := obj.(resource.Terraformed)
			resources := writeResourceBlock(res)

			jstream, _ := ioutil.ReadFile(tt.args.tfJsonFile)
			x := map[string]any{}

			json.Unmarshal([]byte(jstream), &x)

			g := gomega.NewGomegaWithT(t)

			g.Expect(resources).To(gomega.Equal(x["resource"]))
		})
	}

}

// this function reproduces the private logic in upjet/pkg/terraform/files.go
func writeResourceBlock(res resource.Terraformed) map[string]any {
	parms, _ := res.GetParameters()

	parms["lifecycle"] = map[string]interface{}{
		"prevent_destroy": !meta.WasDeleted(res),
	}
	config.ExternalNameConfigs[res.GetTerraformResourceType()].SetIdentifierArgumentFn(parms, res.GetName())
	// parms["name"] = meta.GetExternalName(res)
	return map[string]any{
		res.GetTerraformResourceType(): map[string]any{
			res.GetName(): parms,
		},
	}
}
