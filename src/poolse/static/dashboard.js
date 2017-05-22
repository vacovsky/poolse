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

    app.controller('poolseControl', function ($scope, $rootScope, $http, $interval, $filter, $cookies, $compile) {
        $rootScope.updateInterval = 60000;
        $scope.status = {};
        $scope.searchText = "";
        $scope.showNodesOnTable = true;

        this.prettyTime = function (uglyTime) {
            if (uglyTime !== undefined && uglyTime !== null) {
                var pt = moment(uglyTime).calendar();
                return pt;
            } else {
                return "Never";
            }
        };

        $scope.showMembers = function (id, val) {
            $scope.status.Targets[id].showMembers = val
            console.log($scope.status.Targets[id].showMembers)

        }

        $scope.shouldShowNodes = function () {
            largestNodeLength = 0;
            $scope.status.Targets.forEach(function (element) {
                if (element.nodes.length > largestNodeLength) {
                    largestNodeLength = element.length.nodes;
                }
            })
            if (largestNodeLength > 0) {
                showNodesOnTable = true;
            } else {
                showNodesOnTable = false;                
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
                $scope.shouldShowNodes();
            };

        };
        $scope.getStatus()
        $interval($scope.getStatus, $rootScope.updateInterval)
    });
})();