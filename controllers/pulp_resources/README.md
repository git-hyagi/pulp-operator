
### Custom Resources

* [PulpResource](#pulpresource)

### Sub Resources

* [PulpAnsibleCollectionRemote](#pulpansiblecollectionremote)
* [PulpAnsibleDistribution](#pulpansibledistribution)
* [PulpAnsibleGitRemote](#pulpansiblegitremote)
* [PulpAnsiblePlugin](#pulpansibleplugin)
* [PulpAnsibleRepository](#pulpansiblerepository)
* [PulpContainerDistribution](#pulpcontainerdistribution)
* [PulpContainerPlugin](#pulpcontainerplugin)
* [PulpContainerRemote](#pulpcontainerremote)
* [PulpCoreDistribution](#pulpcoredistribution)
* [PulpCoreRemote](#pulpcoreremote)
* [PulpCoreRepository](#pulpcorerepository)
* [PulpFileDistribution](#pulpfiledistribution)
* [PulpFilePlugin](#pulpfileplugin)
* [PulpFileRepository](#pulpfilerepository)
* [PulpRPMDistribution](#pulprpmdistribution)
* [PulpRPMPlugin](#pulprpmplugin)
* [PulpRPMRemote](#pulprpmremote)
* [PulpRPMRepository](#pulprpmrepository)
* [PulpResourceList](#pulpresourcelist)
* [PulpResourceSpec](#pulpresourcespec)
* [PulpResourceStatus](#pulpresourcestatus)
* [PulpSync](#pulpsync)

#### PulpAnsibleCollectionRemote



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| requirements_file |  | string | false |
| auth_url |  | string | false |
| token |  | string | false |
| sync_dependencies |  | bool | false |
| signed_only |  | bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpAnsibleDistribution



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| repository_version |  | int | false |

[Back to Custom Resources](#custom-resources)

#### PulpAnsibleGitRemote



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata_only |  | bool | false |
| git_ref |  | string | false |

[Back to Custom Resources](#custom-resources)

#### PulpAnsiblePlugin



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| repository |  | *[PulpAnsibleRepository](#pulpansiblerepository) | false |
| distribution |  | *[PulpAnsibleDistribution](#pulpansibledistribution) | false |
| role_remote |  | *[PulpAnsibleRoleRemote](#pulpansibleroleremote) | false |
| collection_remote |  | *[PulpAnsibleCollectionRemote](#pulpansiblecollectionremote) | false |
| git_remote |  | *[PulpAnsibleGitRemote](#pulpansiblegitremote) | false |
| sync |  | *bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpAnsibleRepository



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| gpgkey |  | string | false |
| private |  | bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpContainerDistribution



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| version |  | int | false |
| private |  | bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpContainerPlugin



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| repository |  | *[PulpContainerRepository](#pulpcontainerrepository) | false |
| distribution |  | *[PulpContainerDistribution](#pulpcontainerdistribution) | false |
| remote |  | *[PulpContainerRemote](#pulpcontainerremote) | false |
| sync |  | *bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpContainerRemote



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| upstream_name |  | string | false |
| include_tags |  | []string | false |
| exclude_tags |  | []string | false |
| sigstore |  | string | false |

[Back to Custom Resources](#custom-resources)

#### PulpCoreDistribution



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| name |  | string | true |
| base_path |  | string | false |
| repository |  | string | false |
| content-guard |  | string | false |
| pulp_labels |  | map[string]string | false |

[Back to Custom Resources](#custom-resources)

#### PulpCoreRemote



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| name |  | string | true |
| url |  | string | false |
| ca_cert |  | string | false |
| client_cert |  | string | false |
| client_key |  | string | false |
| tls_validation |  | bool | false |
| proxy_url |  | string | false |
| proxy_username |  | string | false |
| proxy_password |  | string | false |
| username |  | string | false |
| password |  | string | false |
| connection_timeout |  | int | false |
| download_concurrency |  | int | false |
| rate_limit |  | int | false |
| sock_connect_timeout |  | int | false |
| sock_read_timeout |  | int | false |
| max_retries |  | int | false |
| policy |  | string | false |
| total_timeout |  | string | false |
| headers |  | []map[string]string | false |
| pulp_labels |  | map[string]string | false |

[Back to Custom Resources](#custom-resources)

#### PulpCoreRepository



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| description |  | string | false |
| remote |  | string | false |
| retain_repo_versions |  | int | false |
| pulp_labels |  | map[string]string | false |
| name |  | string | true |

[Back to Custom Resources](#custom-resources)

#### PulpFileDistribution



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| publication |  | string | false |

[Back to Custom Resources](#custom-resources)

#### PulpFilePlugin



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| repository |  | *[PulpFileRepository](#pulpfilerepository) | false |
| distribution |  | *[PulpFileDistribution](#pulpfiledistribution) | false |
| remote |  | *[PulpFileRemote](#pulpfileremote) | false |
| sync |  | *bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpFileRepository



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| manifest |  | string | false |
| autopublish |  | bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpRPMDistribution



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| publication |  | string | false |
| generate-repo-config |  | bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpRPMPlugin



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| repository |  | *[PulpRPMRepository](#pulprpmrepository) | false |
| distribution |  | *[PulpRPMDistribution](#pulprpmdistribution) | false |
| remote |  | *[PulpRPMRemote](#pulprpmremote) | false |
| sync |  | *bool | false |

[Back to Custom Resources](#custom-resources)

#### PulpRPMRemote



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| sles_auth_token |  | string | false |

[Back to Custom Resources](#custom-resources)

#### PulpRPMRepository



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| retain-package-vesions |  | int | false |
| metadata-signing-service |  | string | false |
| checksum-type |  | string | false |
| manifest |  | string | false |
| autopublish |  | bool | false |
| repo-config |  | string | false |

[Back to Custom Resources](#custom-resources)

#### PulpResource

PulpResource is the Schema for the pulpresource API

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | metav1.ObjectMeta | false |
| spec |  | [PulpResourceSpec](#pulpresourcespec) | false |
| status |  | *[PulpResourceStatus](#pulpresourcestatus) | false |

[Back to Custom Resources](#custom-resources)

#### PulpResourceList

PulpResourceList contains a list of PulpResource

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | metav1.ListMeta | false |
| items |  | [][PulpResource](#pulpresource) | true |

[Back to Custom Resources](#custom-resources)

#### PulpResourceSpec

PulpResourceSpec defines the desired state of PulpResource

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| client_config |  | *string | false |
| ansible |  | *[PulpAnsiblePlugin](#pulpansibleplugin) | false |
| container |  | *[PulpContainerPlugin](#pulpcontainerplugin) | false |
| file |  | *[PulpFilePlugin](#pulpfileplugin) | false |
| rpm |  | *[PulpRPMPlugin](#pulprpmplugin) | false |

[Back to Custom Resources](#custom-resources)

#### PulpResourceStatus

PulpResourceStatus defines the observed state of PulpResource

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| ansible |  | *[PulpAnsiblePlugin](#pulpansibleplugin) | false |
| container |  | *[PulpContainerPlugin](#pulpcontainerplugin) | false |
| file |  | *[PulpFilePlugin](#pulpfileplugin) | false |
| rpm |  | *[PulpRPMPlugin](#pulprpmplugin) | false |

[Back to Custom Resources](#custom-resources)

#### PulpSync



| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| remote |  | string | false |
| mirror |  | bool | false |
| signed_only |  | bool | false |

[Back to Custom Resources](#custom-resources)
