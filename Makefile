prepare:
	curl https://pre-commit.com/install-local.py | python3 -
	pre-commit install

generate: interface_mocks prepare
	go mod tidy
