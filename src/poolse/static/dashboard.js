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

    app.controller('poolseControl', function ($scope, $rootScope, $http, $timeout, $filter, $cookies, $scope, $compile) {
        $rootScope.updateInterval = 60000;
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
                });
        };

        $scope.intervalFunction = function () {
            $timeout(function () {
                $scope.getStatus();
                $scope.intervalFunction();
            }, $rootScope.updateInterval);
        };
        $scope.getStatus();
        $scope.intervalFunction();
    });
})();