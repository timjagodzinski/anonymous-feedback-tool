container:
	docker build -t go_crud_service .
run:
	docker run -p 8080:8080 --rm --name my-running-app go_crud_service
runInteractive:
	docker run -p 8080:8080 -it --rm --name my-running-app go_crud_service