# default to running tests

SHELL=/bin/bash -o pipefail

# where to mount the sojanode tree
BASEDIR=/sojanode

in_progress_marker=data/tests-in-progress.txt
failure_marker=data/tests_failed.txt

.PHONY: all
all: 
	rm -rf data/ ../ganachedb ../__pycache__ .depend/integrationenv
	$(MAKE) data
	touch ${in_progress_marker}
	if ! $(MAKE) test; then touch ${failure_marker}; fi
	$(MAKE) pulldata
	if [ -f ${failure_marker} ]; then echo tests failed; exit 1; else touch data/tests_passed; fi

.depend data: 
	mkdir -p $@

.depend/up: | .depend
	vagrant up
	touch $@

.depend/setup: .depend/up
	vagrant ssh -c "cd ${BASEDIR} && bash test/integration/setup-linux-environment.sh"
	touch $@

.depend/ganache: .depend/setup | data
	vagrant ssh -c "cd ${BASEDIR} && bash -x test/integration/ganache_start.sh" 2>&1 | tee data/start-ganache
	touch $@

.depend/integrationenv: .depend/setup | data
	vagrant ssh -c "cd ${BASEDIR} && bash -x test/integration/start-integration-env.sh" 2>&1 | tee data/start-integration-env.txt
	touch $@

.PHONY: test-whitelistedvalidators
test-whitelistedvalidators: .depend/integrationenv
	vagrant ssh -c "cd ${BASEDIR} && bash -x test/integration/execute_integration_tests_whitelisted_validators.sh" 2>&1 | tee data/$@.sh

.PHONY: test-main
test-main: .depend/integrationenv
	for i in execute_integration_tests_against_any_chain execute_integration_tests_against_test_chain execute_integration_tests_whitelisted_validators; \
      do \
        vagrant ssh -c "cd ${BASEDIR} && . test/integration/vagrantenv.sh; cd \$${TEST_INTEGRATION_DIR}; bash $${i}.sh" 2>&1 | tee data/$${i}.sh; \
      done

.PHONY: test
test: test-main test-whitelistedvalidators
	rm ${in_progress_marker}

# convenience targets

up: .depend/up

reload:
	vagrant reload

clean:
	vagrant destroy -f
	rm -rf .depend

pulldata: data
	vagrant ssh -c "bash /sojanode/test/integration/assemble_data.sh"
	tar cfz datafiles.tar.gz data/*
	cp datafiles.tar.gz datafiles.$(shell date +%m-%d-%H-%M-%S).tar.gz
	
ssh:
	vagrant ssh -c "cd /sojanode/test/integration && . vagrantenv.sh && bash -i"

truffle:
	vagrant ssh -c "cd /sojanode && . test/integration/vagrantenv.sh && cd /sojanode/smart-contracts && truffle console --network develop --verbose-rpc"

.PHONY: sojadocker
sojadocker:
	cd .. && docker build -t sojadocker -f Dockerfile.devenv .

# start a docker container and keep it running
sojadocker-start:
	nohup docker run -p 7546:7545 --rm --name sojadocker -v $(shell pwd)/../../..:/sojanode sojadocker:latest bash -c "tail -f /dev/null" > /dev/null 2>&1 &

sojadocker-sh:
	docker exec -ti sojadocker bash