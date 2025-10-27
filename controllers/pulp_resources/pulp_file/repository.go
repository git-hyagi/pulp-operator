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

func ManageFileRepository(ctx context.Context, pulpResource *pulpv1.PulpResource, clientBinding *bindings.APIClient) (*ctrl.Result, error) {
	pluginName := "File"
	resourceType := "Repository"
	/* 	owner := pulpResource.Name + "." + pulpResource.Namespace
	   	operatorLabels := map[string]string{
	   		"owner": owner,
	   	} */

	repoName := utils.GetResource(pulpResource, pluginName, resourceType)
	if repoName == nil {
		return nil, nil
	}

	// check if resource already exists
	repoListRequest := clientBinding.RepositoriesFileAPI.RepositoriesFileFileList(ctx)
	repoFilterRequest := repoListRequest.Name(*repoName)
	repoList, _, err := repoFilterRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("List %s %s request failed! %v\n", pluginName, resourceType, err)
		slog.Error(errMessage)
		return &ctrl.Result{}, err
	}
	if len(repoList.Results) > 0 {
		// pending implement this!
		/* 		if !utils.OwnedByOperator(response, owner) {
			slog.Error(pluginName + " " + resourceType + " " + resourceName + " already exists, but is not managed by Pulp operator!")
			return
		} */
		repoFoundMsg := fmt.Sprintf("%s %s %s already exists!", pluginName, resourceType, *repoName)
		slog.Info(repoFoundMsg)
		return nil, nil
	}

	newRepo := bindings.NewFileFileRepository(*repoName)
	repoCreateRequest := clientBinding.RepositoriesFileAPI.RepositoriesFileFileCreate(ctx)
	repoCreateRequest = repoCreateRequest.FileFileRepository(*newRepo)
	newRepositoryFile, _, err := repoCreateRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("Create %s %s request failed! %v\n", pluginName, resourceType, err)
		slog.Error(errMessage)
		return &ctrl.Result{}, err
	}

	setLabelRequest := clientBinding.RepositoriesFileAPI.RepositoriesFileFileSetLabel(ctx, *newRepositoryFile.PulpHref).SetLabel(utils.OperatorLabels(*pulpResource))
	_, _, err = setLabelRequest.Execute()
	if err != nil {
		errMessage := fmt.Sprintf("Failed to set label in %s %s %s! %v\n", *repoName, pluginName, resourceType, err)
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
