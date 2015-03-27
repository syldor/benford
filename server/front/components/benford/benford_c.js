'use strict';

function BenfordCtrl($scope, $http, $filter) {
	$scope.submit = function() {
		$http.post('/benford', $scope.data).success(function(data) {
			$scope.graph_data = [];
			for (var p = 1 ; p < 10 ; p++ ) {
				$scope.graph_data.push({
					number: p,
					count: data[p.toString()] || 0
				})
			}
			if($scope.barchart) {
				$scope.barchart.setData($scope.graph_data)
			}
			else {
		    $scope.barchart = Morris.Bar({
		      element: 'benford_graph',
		      data: $scope.graph_data,
		      barColors: ["#31c0be"],
		      xkey: ['number'],
		      ykeys: ['count'],
		      labels: ['number'],
		      hideHover: 'auto'
		    });
			}
		});
	}
};