package gcpnfsvolumerestore

import (
	"context"
	cloudresourcesv1beta1 "github.com/kyma-project/cloud-manager/api/cloud-control/v1beta1"
	"github.com/kyma-project/cloud-manager/api/cloud-resources/v1beta1"
	"github.com/kyma-project/cloud-manager/pkg/composed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

type loadScopeSuite struct {
	suite.Suite
	ctx context.Context
}

func (suite *loadScopeSuite) SetupTest() {
	suite.ctx = context.Background()
}

func (suite *loadScopeSuite) TestScopeNotFound() {
	fakeHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Fail(suite.T(), "unexpected request: "+r.URL.String())
		return
	}))
	obj := gcpNfsVolumeRestore.DeepCopy()
	factory, err := newTestStateFactoryWithObj(fakeHttpServer, obj)
	suite.Nil(err)

	//remove scope
	err = factory.kcpCluster.K8sClient().Delete(context.Background(), scope.DeepCopy())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//Get state object with GcpNfsVolume
	state, err := factory.newStateWith(obj)
	err, _ = loadScope(ctx, state)

	//validate expected return values
	suite.Equal(composed.StopAndForget, err)
	fromK8s := &v1beta1.GcpNfsVolumeRestore{}
	err = factory.skrCluster.K8sClient().Get(ctx,
		types.NamespacedName{Name: gcpNfsVolumeRestore.Name,
			Namespace: gcpNfsVolumeRestore.Namespace},
		fromK8s)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "", fromK8s.Status.OpIdentifier)
	assert.Equal(suite.T(), v1beta1.JobStateFailed, fromK8s.Status.State)
	assert.Equal(suite.T(), metav1.ConditionTrue, fromK8s.Status.Conditions[0].Status)
	assert.Equal(suite.T(), cloudresourcesv1beta1.ConditionTypeError, fromK8s.Status.Conditions[0].Type)
}

func (suite *loadScopeSuite) TestScopeExists() {
	fakeHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Fail(suite.T(), "unexpected request: "+r.URL.String())
		return
	}))
	defer fakeHttpServer.Close()

	obj := gcpNfsVolumeRestore.DeepCopy()
	factory, err := newTestStateFactoryWithObj(fakeHttpServer, obj)
	suite.Nil(err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//Get state object with GcpNfsVolume
	state, err := factory.newStateWith(obj)
	suite.Nil(err)
	err, _ = loadScope(ctx, state)
	assert.Nil(suite.T(), err)
}

func TestLoadScopeSuite(t *testing.T) {
	suite.Run(t, new(loadScopeSuite))
}
