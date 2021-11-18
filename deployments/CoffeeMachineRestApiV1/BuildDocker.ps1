#Copy over dependencies and build docker image
#Mock CI/CD part

Copy-Item -Path "..\..\internal\" -Destination "CoffeeMachine/internal" -Recurse -Container
Copy-Item -Path "..\..\cmd\CoffeeMachineRestApiV1\" -Destination "CoffeeMachine\cmd\CoffeeMachineRestApiV1" -Recurse -Container

docker build . -t github.com/kostovic/coffeemachine:restapiv1.0 --no-cache
Remove-Item "CoffeeMachine" -Recurse