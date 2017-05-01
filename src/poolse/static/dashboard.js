(function () {

    var app = angular.module('poolseDash', [
        'chart.js',
        'ngCookies',
        'ngMessages',
        'ngAnimate',
        'ui.rCalendar',
        'angularMoment'
    ]).factory('moment', function ($window) {
        return $window.moment;
    });
    app.Root = '/';

    app.controller('poolseControl', function ($scope, $rootScope, $http, $interval, $filter, $cookies, $scope, $compile) {
        $rootScope.updateInterval = 1000;//60000;
        $scope.status = {};
        $scope.searchText = "";

        this.prettyTime = function (uglyTime) {
            if (uglyTime !== undefined && uglyTime !== null) {
                // console.log(uglyTime)
                var pt = moment(uglyTime).calendar();
                return pt;
            } else {
                return "Never";
            }
        };

        $scope.getStatus = function () {
            $http.get('/status')
                .then(function (response) {
                    $scope.status = response.data;
                })
            if ($scope.status.Targets != undefined) {
                $scope.status.Targets.forEach(function (element) {
                    if (element.notes != "" && element.notes != undefined) {
                    $http.get(element.endpoint + element.notes)
                        .then(function (response) {
                            $scope.status.Targets[element.id].nodes = response.data.servers.server;
                        });
                    }
                })
                console.log($scope.status);
            };
        };
        $scope.getStatus()
        $interval($scope.getStatus, $rootScope.updateInterval)
    });
})();