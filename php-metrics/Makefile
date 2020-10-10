COMPOSER_DIR ?= ${HOME}/.composer

require:
	docker run --rm \
	    --volume `pwd`:/app/ \
	    --volume ${COMPOSER_DIR}:/composer \
	    composer:1.10 require $(package) --ignore-platform-reqs
.PHONY: require

vendor:
	docker run --rm \
	    --volume `pwd`:/app/ \
	    --volume ${COMPOSER_DIR}:/composer \
	    composer:1.10 install --ignore-platform-reqs
.PHONY: vendor