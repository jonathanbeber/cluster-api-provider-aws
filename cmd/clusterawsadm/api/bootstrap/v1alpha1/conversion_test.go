package v1alpha1

import (
	"testing"

	"github.com/onsi/gomega"

	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/api/bootstrap/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
)

func TestFuzzyConversion(t *testing.T) {
	g := gomega.NewWithT(t)
	scheme := runtime.NewScheme()
	g.Expect(AddToScheme(scheme)).To(gomega.Succeed())
	g.Expect(v1beta1.AddToScheme(scheme)).To(gomega.Succeed())

	t.Run("for AWSIAMConfiguration", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Scheme: scheme,
		Hub:    &v1beta1.AWSIAMConfiguration{},
		Spoke:  &AWSIAMConfiguration{},
	}))
}
