/* jshint undef: true, unused: true */
/* global angular */
// (function () {
'use strict';

  /**
   * The main Kassa app module
   *
   * @type {angular.Module}
   */
var app = angular.module('MyApp',['ngMaterial', 'ngMessages']);

app.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('{[{').endSymbol('}]}');
});

app.controller('MainCtrl', MainCtrl);
app.controller('CategoryCtrl', CategoryCtrl);
app.controller('KassaProductCategoryCtrl', KassaProductCategoryCtrl);
app.controller('KassaSearchCtrl', KassaSearchCtrl);
app.controller('KassaProductCtrl', KassaProductCtrl);
app.controller('KassaAmountCtrl', KassaAmountCtrl);

app.directive('clock', Clock);

app.factory('loadCategoryFactory', LoadCategory);

  app.factory('loadProductsFactory', [ '$rootScope', '$http', '$log', function($rootScope, $http, $log) {
    // $log.info('loadProducts...');

    var apiUrl = $rootScope.Url,
      url = null,
      obj = {
        nav: {
          count: null,
          page: null,
          current: null,
          total: null,
          next: null,
          priv: null,
          isNext: true,
          isPriv: true
        },
        items: [],
        category: null
      };

    return {
      get: function (category_id, page) {
        if (category_id === undefined) {
          category_id = 0;
        }
        obj.category =category_id;
        if (page === undefined) {
          page = '?limit=18';
        }
        if (obj.category > 0) {
          url = apiUrl + '/api/v1/category/' + obj.category + '/products';
        } else {
          url = apiUrl + '/api/v1/products';
        }
        if (page.length === 0) {
          page = '?limit=18';
        }
        // go = obj.category + url + lim;
        $http.get(url + page)
          .success(function(data) {
            // $log.info('Load: ' + url + page + ' success!');
            obj.items = data.data;
            // $log.info(data.data);
            obj.nav.count = data.count;
            obj.nav.current = data.current;
            obj.nav.total = data.total;
            obj.nav.page = data.current + '/' + data.total;
            obj.nav.next = data.next;
            obj.nav.priv = data.priv;
            if (data.priv.length > 0) {
              obj.nav.isPriv = false;
            } else {
              obj.nav.isPriv = true;
            }
            if (data.next.length > 0) {
              obj.nav.isNext = false;
            } else {
              obj.nav.isNext = true;
            }
        });
        // $log.info(obj);
        return obj;
      // },
      // put: function (todos) {
      //   localStorage.setItem(STORAGE_ID, JSON.stringify(todos));
      }
    };
  }]);

// })();
/* jshint undef: true, unused: true */
/* global angular */
'use strict';

Clock.$inject = ['dateFilter', '$timeout'];

function Clock(dateFilter, $timeout){
    return {
        restrict: 'E',
        scope: {
            format: '@'
        },
        link: function(scope, element, attrs){
            var updateTime = function(){
                var now = Date.now();

                element.html(dateFilter(now, scope.format));
                $timeout(updateTime, now % 1000);
            };

            updateTime();
        }
    };

}
/*global angular, app */

/*
 * Line below lets us save `this` as `TC`
 * to make properties look exactly the same as in the template
 */
//jscs:disable safeContextKeyword
'use strict';

LoadCategory.$inject = ['$rootScope', '$http', '$log'];

function LoadCategory($rootScope, $http, $log) {
  var apiUrl = $rootScope.Url,
      service = {};

  service.GetAll = GetAll;

  return service;

  function GetAll() {
    return $http.get(apiUrl + '/api/v1/category').then(handleSuccess, handleError('Error getting all category'));
  }

  // private functions

  function handleSuccess(res) {
    // $log.debug('Load: ' + apiUrl + '/api/v1/category');
    // $log.debug(res.data.data);
    return res.data.data;
  }

  function handleError(error) {
    return function () {
      return { success: false, message: error };
    };
  }
    // return {
    //   get: function () {
    //     $http.get(apiUrl + '/api/v1/category')
    //       .success(function(data) {
    //         $log.debug('Load: ' + apiUrl + '/api/v1/category');
    //         // $log.info(data.data);
    //         obj =  data.data;
    //       });
    //     $log.info('loadCategoryFactory --->' + obj);
    //     return obj;
    //   }
    // }
  }
/* jshint undef: true, unused: true */
/*global angular, app */

/*
 * Line below lets us save `this` as `TC`
 * to make properties look exactly the same as in the template
 */
//jscs:disable safeContextKeyword
// (function () {
  'use strict';
  // var app = angular.module('MyApp',['ngMaterial', 'ngMessages']);

  // var app = angular
  //     .module('MyApp',['ngMaterial'])
  //     // .module('MyApp',['ngMaterial', 'ngMessages', 'material.svgAssetsCache'])
  //     // .controller('DemoCtrl', DemoCtrl);
  //     .controller('CategoryCtrl', CategoryCtrl);



  // function KassaSearchCtrl($rootScope, $scope, $http, $element, $log) {
  //   var self = this;

  //   // self.newState = newState;
  // }

  // function KassaAmountCtrl($rootScope, $scope, $http, $element, $log) {
  //   var self = this;

  //   // self.newState = newState;
  // }




  // function KassaCtrl ($timeout, $q, $log) {
  //   var self = this;

  //   self.isDisabled    = false;

  // }

DemoCtrl.$inject = ['$timeout', '$q', '$log'];

function DemoCtrl ($timeout, $q, $log) {
    var self = this;

    self.simulateQuery = false;
    self.isDisabled    = false;

    // list of `state` value/display objects
    self.states        = loadAll();
    self.querySearch   = querySearch;
    self.selectedItemChange = selectedItemChange;
    self.searchTextChange   = searchTextChange;

    self.newState = newState;

    function newState(state) {
      alert("Sorry! You'll need to create a Constituion for " + state + " first!");
    }

    // ******************************
    // Internal methods
    // ******************************

    /**
     * Search for states... use $timeout to simulate
     * remote dataservice call.
     */
    function querySearch (query) {
      var results = query ? self.states.filter( createFilterFor(query) ) : self.states,
          deferred;
      if (self.simulateQuery) {
        deferred = $q.defer();
        $timeout(function () { deferred.resolve( results ); }, Math.random() * 1000, false);
        return deferred.promise;
      } else {
        return results;
      }
    }

    function searchTextChange(text) {
      $log.info('Text changed to ' + text);
    }

    function selectedItemChange(item) {
      $log.info('Item changed to ' + JSON.stringify(item));
    }

    /**
     * Build `states` list of key/value pairs
     */
    function loadAll() {
      var allStates = 'Alabama, Alaska, Arizona, Arkansas, California, Colorado, Connecticut, Delaware,\
              Florida, Georgia, Hawaii, Idaho, Illinois, Indiana, Iowa, Kansas, Kentucky, Louisiana,\
              Maine, Maryland, Massachusetts, Michigan, Minnesota, Mississippi, Missouri, Montana,\
              Nebraska, Nevada, New Hampshire, New Jersey, New Mexico, New York, North Carolina,\
              North Dakota, Ohio, Oklahoma, Oregon, Pennsylvania, Rhode Island, South Carolina,\
              South Dakota, Tennessee, Texas, Utah, Vermont, Virginia, Washington, West Virginia,\
              Wisconsin, Wyoming';

      return allStates.split(/, +/g).map( function (state) {
        return {
          value: state.toLowerCase(),
          display: state
        };
      });
    }

    /**
     * Create filter function for a query string
     */
    function createFilterFor(query) {
      var lowercaseQuery = angular.lowercase(query);

      return function filterFn(state) {
        return (state.value.indexOf(lowercaseQuery) === 0);
      };

    }
  }

// })();
/* jshint undef: true, unused: true */
/*global angular, app */

/*
 * Line below lets us save `this` as `TC`
 * to make properties look exactly the same as in the template
 */
//jscs:disable safeContextKeyword
// (function () {
'use strict';

MainCtrl.$inject = ['$rootScope', '$log'];

function MainCtrl($rootScope, $log) {
    var self = this;

    self.productCalc = productCalc;
    $rootScope.globalLoader = true;
    // $log.info('Start MainCtrl');
    $rootScope.Url = 'http://k.bbox.kiev.ua';
    // $rootScope.Url = 'http://localhost:7000';
    $rootScope.CurrentCategory = 0;
    $rootScope.user = {};
    $rootScope.user.id = 3;
    $rootScope.user.name = 'Вейдер Д. А.';
    $rootScope.OrderCode = 'XXX-XXXXXXX';
    $rootScope.products = [];
    $rootScope.prod = [];
    $rootScope.orderAmount = 0.00;

    $rootScope.$on('rootScope.summa', function() {
      $log.debug('MainCtrl - on - rootScope.summa');
      self.productCalc();
    });

    function productCalc() {
      var amount = 0;
      angular.forEach($rootScope.products, function(val, key){
      var curAmount = 0;
      curAmount = val.price * val.count;
      amount += curAmount;
      if (curAmount != val.amount) {
        $log.error('curAmount: ' + curAmount +' != val.amount: ' + val.amount);
      }
      // if(value.Password == "thomasTheKing") {
      //    console.log("username is thomas");
      //   }
      });
      $log.debug('amount:' + amount);
      $rootScope.orderAmount = amount;
    }

    // $rootScope.$on('rootScope.CurrentCategory', function() {
    //   $log.info('MainCtrl - rootScope - rootScope.emit');
    // });
  }

// })();
"use strict";
/* jshint undef: true, unused: true */
/* global angular */

/*
 * Line below lets us save `this` as `self`
 * to make properties look exactly the same as in the template
 */
// jscs:disable safeContextKeyword
'use strict';

KassaAmountCtrl.$inject = ['$rootScope', '$scope', '$element', '$log'];

function KassaAmountCtrl($rootScope, $scope, $element, $log) {
  var self = this;

  self.isDisabled = false;
  self.model = {};
  self.model.code = $rootScope.OrderCode;
  self.model.clientAmmount = 0;
  self.model.clientSurrender = 0;
  self.model.amount = 0;
  self.cleanOrder = cleanOrder;
  $scope.clientAmmount = null;
  $scope.$watch('clientAmmount', function () {
    $log.debug('$watch -------> model.clientAmmount');
    self.model.clientAmmount = $scope.clientAmmount;
    self.model.clientSurrender = Surrender();
  });

  $rootScope.$watch('orderAmount', function () {
    $log.debug('$watch -------> orderAmount');
    self.model.amount = $rootScope.orderAmount;
    self.model.clientSurrender = Surrender();
  });

  function cleanOrder() {
    $log.debug('------ cleanOrder -----');
    $rootScope.CurrentCategory = 0;
    $rootScope.products.length = 0;
    $rootScope.prod.length = 0;
    $rootScope.orderAmount = 0;
    $scope.clientAmmount = null;
    self.model.clientAmmount = 0;
    self.model.clientSurrender = 0;
  }

  function Surrender() {
    var res = 0;
    if (self.model.clientAmmount > 0) {
      res = self.model.clientAmmount - self.model.amount;
    }
    $log.info('Surrender=' + res);
    return res;
  }
}
/* jshint undef: true, unused: true */
'use strict';
// jscs:disable safeContextKeyword

KassaProductCategoryCtrl.$inject = ['$rootScope', '$scope', '$http', '$element', '$log', 'loadProductsFactory'];

function KassaProductCategoryCtrl($rootScope, $scope, $http, $element, $log, loadProductsFactory) {
  var self = this;

  self.isDisabled = false;
  $scope.newState = newState;
  $scope.goPriv = goPriv;
  $scope.goNext = goNext;
  $scope.lim = '';
  $scope.data = loadProductsFactory.get($rootScope.CurrentCategory, $scope.lim);

  $rootScope.$on('rootScope.emit', function () {
    $log.info('KassaProductCategoryCtrl - rootScope.CurrentCategory');
    $scope.lim = '';
    $scope.data = loadProductsFactory.get($rootScope.CurrentCategory, $scope.lim);
  });

  function goPriv() {
    $scope.lim = $scope.data.nav.priv;
    $scope.data = loadProductsFactory.get($rootScope.CurrentCategory, $scope.lim);
    // $log.info('KassaProductCategoryCtrl priv: ' + $scope.lim);
  }

  function goNext() {
    $scope.lim = $scope.data.nav.next;
    $scope.data = loadProductsFactory.get($rootScope.CurrentCategory, $scope.lim);
    // $log.info('KassaProductCategoryCtrl next: ' + $scope.lim);
  }

  function newState(item) {
    // $scope.lim = id;
    // $rootScope.CurrentCategory = id;
    // $log.info('---------------- KassaProductCategoryCtrl ----------------');
    // $log.info(item);
    $rootScope.products.push(item);
    $rootScope.$emit('rootScope.summa');
  }
}
/* jshint undef: true, unused: true */
/*global angular, app */

/*
 * Line below lets us save `this` as `TC`
 * to make properties look exactly the same as in the template
 */
//jscs:disable safeContextKeyword
'use strict';

CategoryCtrl.$inject = ['$rootScope', '$scope', '$http', '$log', 'loadCategoryFactory'];

function CategoryCtrl($rootScope, $scope, $http, $log, loadCategoryFactory) {
  var self = this;

  self.isDisabled = false;
  self.newState = newState;
  self.current = 0;
  self.items = [];
  // angular.extend(self.items, loadCategoryFactory.GetAll());
  loadCategoryFactory.GetAll().then(function (city) {
    self.items = city;
  });

  function newState(id) {
    // $log.info('Click category ID:' + $rootScope.CurrentCategory + ' curerent:' + self.current + ' id:' + id);
    if (id !== self.current) {
      $rootScope.CurrentCategory = id;
      self.current = id;
    } else {
      $rootScope.CurrentCategory = 0;
      self.current = 0;
      // $log.debug('set category:' + $rootScope.CurrentCategory);
    }
    $rootScope.$emit('rootScope.emit');
  }

  // $http.get($rootScope.Url + '/api/v1/category')
  //   .success(function(data) {
  //     self.items = data.data;
  //   });

  // $log.info(self.items);
}
/* jshint undef: true, unused: true */
/* global angular */

/*
 * Line below lets us save `this` as `self`
 * to make properties look exactly the same as in the template
 */
// jscs:disable safeContextKeyword
'use strict';

KassaProductCtrl.$inject = ['$rootScope', '$scope', '$http', '$element', '$log'];

function KassaProductCtrl($rootScope, $scope, $http, $element, $log) {
  var self = this;

  self.isDisabled = false;
  self.newState = newState;
  self.removeItem = removeItem;
  $scope.items = $rootScope.products;

  function removeItem(id) {
    $rootScope.products.splice(id, 1);
    $rootScope.$emit('rootScope.summa');
    $log.info('remove ID:' + id);
  }

  function newState(id) {
    $rootScope.CurrentCategory = id;
    $rootScope.$emit('rootScope.emit');
    $log.info('Click category ID:' + $rootScope.CurrentCategory);
  }

  // $http.get($rootScope.Url + '/api/v1/category')
  //   .success(function(data) {
  //     $log.info('Load: ' + $rootScope.Url + ' success!');
  //     $scope.items = data.data;
  //     // $log.info(data.data);
  //   });
}
/* jshint undef: true, unused: true */

/*
 * Line below lets us save `this` as `TC`
 * to make properties look exactly the same as in the template
 */
// jscs:disable safeContextKeyword
'use strict';

KassaSearchCtrl.$inject = ['$rootScope', '$log'];

function KassaSearchCtrl($rootScope, $log) {
  var self = this;

  // self.newState = newState;
}
/* jshint undef: true, unused: true */

/*
 * Line below lets us save `this` as `self`
 * to make properties look exactly the same as in the template
 */
// jscs:disable safeContextKeyword
'use strict';