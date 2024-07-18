package repo_manager_pulp

const (
	PULP_RESOURCE_FINALIZER = "repo-manager.pulpproject.org/finalizer"

	PULP_ADMIN_USER = "admin"

	PULP_STATUS_ENDPOINT         = "/pulp/api/v3/status/"
	PULP_CONTAINER_REPO_ENDPOINT = "/pulp/api/v3/repositories/container/container/"
	PULP_FILE_REPO_ENDPOINT      = "/pulp/api/v3/repositories/file/file/"
	PULP_RPM_REPO_ENDPOINT       = "/pulp/api/v3/repositories/rpm/rpm/"

	PULP_CONTAINER_DISTRIBUTION_ENDPOINT = "/pulp/api/v3/distributions/container/container/"
	PULP_FILE_DISTRIBUTION_ENDPOINT      = "/pulp/api/v3/distributions/file/file/"
	PULP_RPM_DISTRIBUTION_ENDPOINT       = "/pulp/api/v3/distributions/rpm/rpm/"

	PULP_CONTAINER_REMOTE_ENDPOINT = "/pulp/api/v3/remotes/container/container/"
	PULP_FILE_REMOTE_ENDPOINT      = "/pulp/api/v3/remotes/file/file/"
	PULP_RPM_REMOTE_ENDPOINT       = "/pulp/api/v3/remotes/rpm/rpm/"
)
