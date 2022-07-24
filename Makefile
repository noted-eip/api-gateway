# After cloning the repo, run init.
init:
	git submodule init

# Fetch the latest version of the protos submodule.
update-submodules:
	git submodule update --remote
