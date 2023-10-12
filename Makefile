TF_ARGS=

tf-fmt:
	terraform fmt -recursive $(TF_ARGS) tf-modules
