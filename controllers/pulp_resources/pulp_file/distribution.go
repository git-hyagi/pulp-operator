package pulp_file

import (
	"context"
	"fmt"
	"log/slog"

	ctrl "sigs.k8s.io/controller-runtime"

	pulpv1 "github.com/pulp/pulp-operator/apis/repo-manager.pulpproject.org/v1"
	bindings "github.com/pulp/pulp-operator/bindings/release"
	"github.com/pulp/pulp-operator/controllers/pulp_resources/utils"
)

func ManageFileDistribution(ctx context.Context, pulpResource *pulpv1.PulpResource, clientBinding *bindings.APIClient) (*ctrl.Result, error) {
	pluginName := "File"
	resourceType := "Distribution"
	distributionBasePath := pulpResource.Spec.File.Distribution.BasePath

	distributionName := utils.GetResource(pulpResource, pluginName, resourceType)
	if distributionName == nil {
		return nil, nil
	}

	// check if resource already exists
	distributionListRequest := clientBinding.DistributionsFileAPI.DistributionsFileFileList(ctx).Name(*distributionName)
	distributionList, _, err := distributionListRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("List %s %s request failed! %v\n", pluginName, resourceType, err)
		slog.Error(errMessage)
		return &ctrl.Result{}, err
	}
	if len(distributionList.Results) > 0 {
		// pending implement this!
		/* 		if !utils.OwnedByOperator(response, owner) {
			slog.Error(pluginName + " " + resourceType + " " + resourceName + " already exists, but is not managed by Pulp operator!")
			return
		} */
		distributionFoundMsg := fmt.Sprintf("%s %s %s already exists!", pluginName, resourceType, *distributionName)
		slog.Info(distributionFoundMsg)
		return nil, nil
	}

	newDistribution := bindings.NewFileFileDistribution(distributionBasePath, *distributionName)
	distributionCreateRequest := clientBinding.DistributionsFileAPI.DistributionsFileFileCreate(ctx).FileFileDistribution(*newDistribution)
	_, _, err = distributionCreateRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("Create %s %s request failed! %v\n", pluginName, resourceType, err)
		slog.Error(errMessage)
		return &ctrl.Result{}, err
	}

	// WIP!!! missing labels in distribution
	/* 	bindings.Task
	   	bindings.TasksAPITasksListRequest.PulpHrefIn([]string{newDistributionFile.Task})

	   	setLabelRequest := clientBinding.DistributionsFileAPI.DistributionsFileFileSetLabel(ctx, *newDistributionFile.PulpHref).SetLabel(utils.OperatorLabels(*pulpResource))
	   	_, _, err = setLabelRequest.Execute()
	   	if err != nil {
	   		errMessage := fmt.Sprintf("Failed to set label in %s %s %s! %v\n", *distributionName, pluginName, resourceType, err)
	   		slog.Error(errMessage)
	   		return &ctrl.Result{}, err
	   	} */

	/*
		 	// temporarily set the labels to create pulp resoruces
			clientBinding.RepositoriesFileAPI.RepositoriesFileFileSetLabel(ctx, respoitoriesList)
			set_operator_label_to_resource(operatorLabels, &resourceValue)
			if _, err := make_request("POST", pulpClient.Address+endpoint, pulpClient.Username, pulpClient.Password, resourceValue.Interface()); err != nil {
				log.Error(err, "Failed to create "+pluginName+" "+resourceType+" "+resourceName)
				continue
			}
			// unset the labels to not populate status field with it
			unset_operator_label_to_resource(&resourceValue)
			log.Info(pluginName + " " + resourceType + " " + resourceName + " created!")
	*/
	return nil, nil
}
