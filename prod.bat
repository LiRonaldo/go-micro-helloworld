@echo off
start "prod1" go run  prod_main.go --server_address 127.0.0.7:8001
pause