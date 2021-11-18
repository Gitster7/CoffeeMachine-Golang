#Copy over dependencies and build docker image
#Mock CI/CD part

Copy-Item -Path "..\..\internal\" -Destination "CoffeeMachine/internal" -Recurse -Container
Copy-Item -Path "..\..\cmd\CoffeeMachineRestApiV2\" -Destination "CoffeeMachine\cmd\CoffeeMachineRestApiV2" -Recurse -Container
Copy-Item -Path "..\..\go.mod" -Destination "CoffeeMachine"
Copy-Item -Path "..\..\go.sum" -Destination "CoffeeMachine"

docker build . -t github.com/kostovic/coffeemachine:restapiv2.0 --no-cache
Remove-Item "CoffeeMachine" -Recurse