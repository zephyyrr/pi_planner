angular.module("pplanner", ['ngRoute'])

.config(function($routeProvider) {
	$routeProvider
	.when('/', {
		templateUrl: 'templates/dashboard.html',
		controller: 'DashboardCtrl'
	})
	.when('/organization/:orgId', {
		templateUrl: 'templates/organization.html',
		controller: 'OrganizationCtrl'
	})
	.when('/projects', {
		templateUrl: 'templates/projects.html',
		controller: 'ProjectsCtrl'
	})
	.when('/user', {
		templateUrl: 'templates/user.html',
		controller: 'UserController'
	})
})



