package service_test

import (
	"testing"

	"github.com/owncloud/ocis/v2/ocis-pkg/registry"
	mRegistry "go-micro.dev/v4/registry"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func init() {
	registry.Configure("memory")
	r := registry.GetRegistry()
	service := registry.BuildGRPCService("com.owncloud.api.gateway", "", "", "")
	service.Nodes = []*mRegistry.Node{{
		Address: "any",
	}}

	_ = r.Register(service)
}

func TestNotifications(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Notification Suite")
}
