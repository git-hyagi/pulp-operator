Architecture
============

Here is a diagram with the components deployed by Pulp Operator:


![Pulp Architecture](images/pulp_architecture.png "Pulp Architecture")

The above picture represents a common installation.

Some components, like `pulp-web`, are not mandatory and depending on how `Pulp CR` is configured
the operator will take care of configuring the other resources that depend on them.

See the [Pulpcore documentation](https://pulpproject.org/pulpcore/docs/admin/learn/architecture/) for more information of Pulp Architecture.