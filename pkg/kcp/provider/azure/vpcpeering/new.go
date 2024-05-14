package vpcpeering

import (
	"context"
	"fmt"
	"github.com/kyma-project/cloud-manager/pkg/composed"
	"github.com/kyma-project/cloud-manager/pkg/kcp/vpcpeering/types"
)

func New(stateFactory StateFactory) composed.Action {
	return func(ctx context.Context, st composed.State) (error, context.Context) {
		logger := composed.LoggerFromCtx(ctx)
		state, err := stateFactory.NewState(ctx, st.(types.State), logger)

		if err != nil {
			err = fmt.Errorf("error creating new aws vpcpeering state %w", err)
			logger.Error(err, "Error")
			return composed.StopAndForget, nil
		}

		return composed.ComposeActions(
			"azureVpcPeering",
			composed.BuildSwitchAction(
				"azureVpcPeering-switch",
				// default action
				composed.ComposeActions("azureVpcPeering-non-delete",
					addFinalizer,
					createVpcPeeringConnection,
					createVpcPeeringRemote,
					updateSuccessStatus,
					composed.StopAndForgetAction),
			), // switch
			composed.StopAndForgetAction,
		)(ctx, state)
	}
}
