var app = angular.module('myApp', []);

app.controller('myCtrl', function ($scope, $http) {
    $scope.master = {};

    $scope.reset = function () {
        document.getElementById("myForm").reset();
        $scope.contato = angular.copy($scope.master);
    }

    $scope.showtabela = function () {
        $http.get("/contato")
            .then(function (response) {
                $scope.contatos = response.data;
            });
    }

    $scope.enviar = function (contato) {
        if (typeof contato.id == "undefined" || contato.id == 0) {
            metodo = "POST";
            msg = " adicionado";
        } else {
            metodo = "PUT";
            msg = " alterado";
        }
        $http({
            method: metodo,
            url: "/contato",
            data: contato
        }).then(function mySuccess(response) {
            document.getElementById("alert").innerHTML = response.data;
            $scope.showtabela();
            $scope.reset();
        }, function myError(response) {
            alert(response.statusText);
        });
    }

    $scope.editar = function (id) {
        $http.get("/contato/" + id)
            .then(function (response) {
                $scope.contato = response.data;
            });
    }

    $scope.remover = function (id) {
        $http({
            method: "DELETE",
            url: "/contato/" + id
        }).then(function mySuccess(response) {
            document.getElementById("alert").innerHTML = response.data;
            document.getElementById("linha" + id).setAttribute("class", "ng-hide")
        }, function myError(response) {
            alert(response.statusText);
        });
    }

    $scope.showtabela()
});