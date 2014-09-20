angular.module("pplanner")

.controller('MenuCtrl', ['$scope', '$route', function($scope, $route){
	$scope.links = {"Home": "/#",
					  "Projects": "/#projects"
					};
}])

.controller('DashboardCtrl' ['$scope', 'UserService', function($scope, UserService) {
	$scope.user = UserService.
}])

.controller('OrganizationCtrl', ['$scope', '$routeParams', function($scope, $routeParams) {
	$scope.name = $routeParams['orgId'];
}])

.controller('ProjectsCtrl', function($scope, UserService) {
	$scope.owner = UserService.name()
	$scope.projects = UserService.getProjects()
})
