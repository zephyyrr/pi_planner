angular.module("pplanner")

.factory('APIService', ['$http', function($http){
	return function(method, route, data) {
		console.log("API Request to /api" + route)
		return $http({
			method: method,
			url: "/api" + route,
			data: data
		})
	}
}])

.factory('UserService', ['APIService', function(APIService){
	var user = APIService("GET", "/user").then(setUser, errorhandler)



	return {
		current: function() {
			return user
		},
		name: name,
		getProjects: getProjects,
	};

	function name() {
		console.log(user)
		user.res
		return user.Name;
	}

	function getProjects() {
		return user.Projects;
	}

	function setUser(res) {
		console.log(user)
		user = res.data;
		console.log(user)

	function errorhandler(data, statuscode, headers, config, statusText) {
		console.log("Error requesting data: " + statuscode + " (" + statusText + ")")
	}
}])
