'use strict';

angular.module('app.directives').controller('TopBarCtrl',
  function($scope, $cookieStore) {

  });

angular.module('app.directives').directive('topbar', function() {
  return {
    restrict: 'E',
    replace: true,
    scope: {

    },
    templateUrl: 'components/top_bar/top_bar_d.html',
    controller: 'TopBarCtrl'
  };
});