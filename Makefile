#Makefile
#(c) 2018 Sam Caldwell.  See LICENSE.txt.
#
.DEFAULT=help
ROOT_DIR:=$(shell pwd -P)

help:
	@echo " "
	@echo "asymmetric-toolkit Makefile"
	@echo "---------------------------"
	@echo "(c) 2018 Sam Caldwell.  See LICENSE.txt"
	@echo " "
	@echo " This is a work in progress.  Sam should implement this...."
	@echo " "
	exit 1

setup:
	@./scripts/setup.sh

lint:
	@./scripts/linter.sh

test:
	@./scripts/run_tests.sh
