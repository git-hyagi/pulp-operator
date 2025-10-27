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

func ManageFileRemote(ctx context.Context, pulpResource *pulpv1.PulpResource, clientBinding *bindings.APIClient) (*ctrl.Result, error) {
	pluginName := "File"
	resourceType := "Remote"
	remoteURL := pulpResource.Spec.File.Remote.URL

	remoteName := utils.GetResource(pulpResource, pluginName, resourceType)
	if remoteName == nil {
		return nil, nil
	}

	// check if resource already exists
	remoteListRequest := clientBinding.RemotesFileAPI.RemotesFileFileList(ctx).Name(*remoteName)
	remoteList, _, err := remoteListRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("List %s %s request failed! %v\n", pluginName, resourceType, err)
		slog.Error(errMessage)
		return &ctrl.Result{}, err
	}
	if len(remoteList.Results) > 0 {
		// pending implement this!
		/* 		if !utils.OwnedByOperator(response, owner) {
			slog.Error(pluginName + " " + resourceType + " " + resourceName + " already exists, but is not managed by Pulp operator!")
			return
		} */
		remoteFoundMsg := fmt.Sprintf("%s %s %s already exists!", pluginName, resourceType, *remoteName)
		slog.Info(remoteFoundMsg)
		return nil, nil
	}

	newRemote := bindings.NewFileFileRemote(*remoteName, remoteURL)
	remoteCreateRequest := clientBinding.RemotesFileAPI.RemotesFileFileCreate(ctx).FileFileRemote(*newRemote)
	newRemoteFile, _, err := remoteCreateRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("Create %s %s request failed! %v\n", pluginName, resourceType, err)
		slog.Error(errMessage)
		return &ctrl.Result{}, err
	}

	setLabelRequest := clientBinding.RemotesFileAPI.RemotesFileFileSetLabel(ctx, *newRemoteFile.PulpHref).SetLabel(utils.OperatorLabels(*pulpResource))
	_, _, err = setLabelRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("Failed to set label in %s %s %s! %v\n", *remoteName, pluginName, resourceType, err)
		slog.Error(errMessage)
		return &ctrl.Result{}, err
	}

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
