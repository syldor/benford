'use strict';

var app = angular.module('app', [
  'ngRoute', 'app.directives', 'ui.bootstrap', 'ngCookies', 'app.services'
]).
config([
    '$routeProvider', function($routeProvider) {
  return $routeProvider.when('/', {
    redirectTo: '/benford'
  })
  .when('/benford', {
    templateUrl: 'components/benford/benford.html',
    controller: 'BenfordCtrl'
  })
  .otherwise({redirectTo: '/benford'});
}]);

angular.module('app.directives', []);

angular.module('app').controller('MasterCtrl', 
    function MasterCtrl($scope, $cookieStore) {
  $scope.toggle = true;
});

