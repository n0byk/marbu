complete -W "\`grep -oE '^[a-zA-Z0-9_.-]+:([^=]|$)' ?akefile | sed 's/[^a-zA-Z0-9_.-]*$//'\`" make

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

up: ## start project
	docker-compose -f ./build/marbu_ecosystem.yml up --build

up_infra: ## up_infra
	docker-compose  -f ./build/marbu_ecosystem.yml up marbu_postgres

down: ##  down project
	docker-compose -f ./build/marbu_ecosystem.yml down

stop: ## stop project
	docker-compose -f ./build/marbu_ecosystem.yml stop