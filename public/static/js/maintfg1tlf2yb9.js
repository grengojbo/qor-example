/* jshint undef: true, unused: true */
/* global angular */
// (function () {
'use strict';

  /**
   * The main Kassa app module
   *
   * @type {angular.Module}
   */
// var app = angular.module('MyApp',['ngMaterial', 'ngMessages', 'ngRoute', 'ngCookies']);
var app = angular.module('MyApp',['ngMaterial', 'ngMessages']);

app.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('{[{').endSymbol('}]}');
});

app.controller('MainCtrl', MainCtrl);
app.controller('MenuCtrl', MenuCtrl);
app.controller('CategoryCtrl', CategoryCtrl);
app.controller('KassaProductCategoryCtrl', KassaProductCategoryCtrl);
app.controller('KassaSearchCtrl', KassaSearchCtrl);
app.controller('KassaProductCtrl', KassaProductCtrl);
app.controller('KassaAmountCtrl', KassaAmountCtrl);

app.directive('clock', Clock);
// app.directive('postRepeatDirective', postRepeatDirective);

app.factory('DataCache', DataCache);
app.factory('UserService', UserService);
app.factory('AuthenticationService', AuthenticationService);
app.factory('loadCategoryFactory', LoadCategory);
app.factory('loadProductsFactory', LoadProducts);
app.factory('orderProductFactory', OrderProduct);

// })();
/* jshint undef: true, unused: true */
/* global angular */
'use strict';

AuthenticationService.$inject = ['$http', '$cookieStore', '$rootScope', '$timeout', 'UserService'];

function AuthenticationService($http, $cookieStore, $rootScope, $timeout, UserService) {
        var service = {};

        service.Login = Login;
        service.SetCredentials = SetCredentials;
        service.ClearCredentials = ClearCredentials;

        return service;

        function Login(username, password, callback) {

            /* Dummy authentication for testing, uses $timeout to simulate api call
             ----------------------------------------------*/
            $timeout(function () {
                var response;
                UserService.GetByUsername(username)
                    .then(function (user) {
                        if (user !== null && user.password === password) {
                            response = { success: true };
                        } else {
                            response = { success: false, message: 'Username or password is incorrect' };
                        }
                        callback(response);
                    });
            }, 1000);

            /* Use this for real authentication
             ----------------------------------------------*/
            //$http.post('/api/authenticate', { username: username, password: password })
            //    .success(function (response) {
            //        callback(response);
            //    });

        }

        function SetCredentials(username, password) {
            var authdata = Base64.encode(username + ':' + password);

            $rootScope.globals = {
                currentUser: {
                    username: username,
                    authdata: authdata
                }
            };

            $http.defaults.headers.common['Authorization'] = 'Basic ' + authdata; // jshint ignore:line
            $cookieStore.put('globals', $rootScope.globals);
        }

        function ClearCredentials() {
            $rootScope.globals = {};
            $cookieStore.remove('globals');
            $http.defaults.headers.common.Authorization = 'Basic';
        }
    }

    // Base64 encoding service used by AuthenticationService
    var Base64 = {

        keyStr: 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=',

        encode: function (input) {
            var output = "";
            var chr1, chr2, chr3 = "";
            var enc1, enc2, enc3, enc4 = "";
            var i = 0;

            do {
                chr1 = input.charCodeAt(i++);
                chr2 = input.charCodeAt(i++);
                chr3 = input.charCodeAt(i++);

                enc1 = chr1 >> 2;
                enc2 = ((chr1 & 3) << 4) | (chr2 >> 4);
                enc3 = ((chr2 & 15) << 2) | (chr3 >> 6);
                enc4 = chr3 & 63;

                if (isNaN(chr2)) {
                    enc3 = enc4 = 64;
                } else if (isNaN(chr3)) {
                    enc4 = 64;
                }

                output = output +
                    this.keyStr.charAt(enc1) +
                    this.keyStr.charAt(enc2) +
                    this.keyStr.charAt(enc3) +
                    this.keyStr.charAt(enc4);
                chr1 = chr2 = chr3 = "";
                enc1 = enc2 = enc3 = enc4 = "";
            } while (i < input.length);

            return output;
        },

        decode: function (input) {
            var output = "";
            var chr1, chr2, chr3 = "";
            var enc1, enc2, enc3, enc4 = "";
            var i = 0;

            // remove all characters that are not A-Z, a-z, 0-9, +, /, or =
            var base64test = /[^A-Za-z0-9\+\/\=]/g;
            if (base64test.exec(input)) {
                window.alert("There were invalid base64 characters in the input text.\n" +
                    "Valid base64 characters are A-Z, a-z, 0-9, '+', '/',and '='\n" +
                    "Expect errors in decoding.");
            }
            input = input.replace(/[^A-Za-z0-9\+\/\=]/g, "");

            do {
                enc1 = this.keyStr.indexOf(input.charAt(i++));
                enc2 = this.keyStr.indexOf(input.charAt(i++));
                enc3 = this.keyStr.indexOf(input.charAt(i++));
                enc4 = this.keyStr.indexOf(input.charAt(i++));

                chr1 = (enc1 << 2) | (enc2 >> 4);
                chr2 = ((enc2 & 15) << 4) | (enc3 >> 2);
                chr3 = ((enc3 & 3) << 6) | enc4;

                output = output + String.fromCharCode(chr1);

                if (enc3 != 64) {
                    output = output + String.fromCharCode(chr2);
                }
                if (enc4 != 64) {
                    output = output + String.fromCharCode(chr3);
                }

                chr1 = chr2 = chr3 = "";
                enc1 = enc2 = enc3 = enc4 = "";

            } while (i < input.length);

            return output;
        }
    };
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

// postRepeatDirective.$inject = ['$timeout', '$log',  'TimeTracker'];
/**
   * Отслеживание времени работы директив
   * <tr ng-repeat="item in items" post-repeat-directive>…</tr>
*/
// function postRepeatDirective($timeout, $log, TimeTracker) {
//     return function(scope, element, attrs) {
//       if (scope.$last){
//          $timeout(function(){
//              var timeFinishedLoadingList = TimeTracker.reviewListLoaded();
//              var ref = new Date(timeFinishedLoadingList);
//              var end = new Date();
//              $log.debug("## DOM отобразился за: " + (end - ref) + " ms");
//          });
//        }
//     };
//   }
/* jshint undef: true, unused: true */
/* global angular */

/*
 * Line below lets us save `this` as `self`
 * to make properties look exactly the same as in the template
 */
// jscs:disable safeContextKeyword
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
}

DataCache.$inject = ['$cacheFactory'];

function DataCache($cacheFactory) {
  return $cacheFactory('dataCache', {});
}

/* jshint undef: true, unused: true */
/* global angular */

/*
 * Line below lets us save `this` as `self`
 * to make properties look exactly the same as in the template
 */
// jscs:disable safeContextKeyword
'use strict';

OrderProduct.$inject = ['$rootScope', '$http', '$log'];

function OrderProduct($rootScope, $http, $log) {
  var self = this,
      apiUrl = $rootScope.Url,
      service = {},
      items = [],
      emptyProduct = {
        id: null,
        code: null,
        price: 0,
        count: 0,
        amount: 0,
        money: 'грн.',
        unit: 'шт.',
        category: null
      },
      product = emptyProduct,
      emptyObj = {
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
        order: 'XXX-XXXXXXX',
        amount: 0
      },
      obj = emptyObj,
      emptySumma = {
        amount: 0,
        money: 'грн.'
      },
      summa = emptySumma;

  service.Add = Add;
  service.Delete = Delete;
  service.GetAll = GetAll;
  service.GetAmount = GetAmount;
  service.Save = Save;
  service.Clean = Clean;

  return service;

  function Add(el) {
    // obj.items.push(el);
    createOrUpdate(el);
    // obj.amount = productCalc();
    productCalc();
  }

  function Delete(i) {
    obj.items.splice(i, 1);
    productCalc();
    // obj.amount = productCalc();
  }

  function GetAll(el) {
    // productCalc();
    return obj.items;
  }

  function GetAmount() {
    $log.debug('GetAmount:' + summa.amount);
    return summa;
  }

  function Save() {
    $log.debug('OrderProduct --> save');
    $log.debug(obj);
    return true;
  }

  function Clean() {
    obj.items.length = 0;
    angular.extend(obj, emptyObj);
    angular.extend(summa, emptySumma);
  }

  // private functions

  function createOrUpdate(el) {
    var isEl = true;
    angular.forEach(obj.items, function(val, key){
      if (val.code === el.code) {
        obj.items[key].amount = val.price * (val.count + 1);
        obj.items[key].count += 1;
        isEl = false;
      }
    });
    if (isEl) {
      el.count = 1;
      el.amount = el.price;
      obj.items.push(el);
    }
  }

  function productCalc() {
      var amount = 0,
      newSumma = emptySumma;
      angular.forEach(obj.items, function(val, key){
        var curAmount = 0;
        curAmount = val.price * val.count;
        amount += curAmount;
        // newSumma.amount += curAmount;
        if (curAmount != val.amount) {
          $log.error('curAmount: ' + curAmount +' != val.amount: ' + val.amount);
        }
      });
      // return amount;
      // amount. = amount;
      // newSumma.amount = amount;
      summa.amount = amount;
      // angular.extend(summa, newSumma);
      $rootScope.$emit('rootScope.summa');
      $log.debug('Set amount:' + summa.amount);
    }

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

LoadProducts.$inject = ['$rootScope', '$http', '$log'];

function LoadProducts($rootScope, $http, $log) {
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
  }
/* jshint undef: true, unused: true */
/* global angular */
'use strict';

UserService.$inject = ['$http'];
function UserService($http) {
  var service = {};

  service.GetAll = GetAll;
        service.GetById = GetById;
        service.GetByUsername = GetByUsername;
        service.Create = Create;
        service.Update = Update;
        service.Delete = Delete;

        return service;

        function GetAll() {
            return $http.get('/api/users').then(handleSuccess, handleError('Error getting all users'));
        }

        function GetById(id) {
            return $http.get('/api/users/' + id).then(handleSuccess, handleError('Error getting user by id'));
        }

        function GetByUsername(username) {
            return $http.get('/api/users/' + username).then(handleSuccess, handleError('Error getting user by username'));
        }

        function Create(user) {
            return $http.post('/api/users', user).then(handleSuccess, handleError('Error creating user'));
        }

        function Update(user) {
            return $http.put('/api/users/' + user.id, user).then(handleSuccess, handleError('Error updating user'));
        }

        function Delete(id) {
            return $http.delete('/api/users/' + id).then(handleSuccess, handleError('Error deleting user'));
        }

        // private functions

        function handleSuccess(res) {
            return res.data;
        }

        function handleError(error) {
            return function () {
                return { success: false, message: error };
            };
        }
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

    // self.productCalc = productCalc;
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

    // $rootScope.$on('rootScope.summa', function() {
    //   $log.debug('MainCtrl - on - rootScope.summa');
    //   self.productCalc();
    // });

    // function productCalc() {
    //   var amount = 0;
    //   angular.forEach($rootScope.products, function(val, key){
    //   var curAmount = 0;
    //   curAmount = val.price * val.count;
    //   amount += curAmount;
    //   if (curAmount != val.amount) {
    //     $log.error('curAmount: ' + curAmount +' != val.amount: ' + val.amount);
    //   }
    //   // if(value.Password == "thomasTheKing") {
    //   //    console.log("username is thomas");
    //   //   }
    //   });
    //   $log.debug('amount:' + amount);
    //   $rootScope.orderAmount = amount;
    // }

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

KassaAmountCtrl.$inject = ['$rootScope', '$scope', '$element', '$log', 'orderProductFactory'];

function KassaAmountCtrl($rootScope, $scope, $element, $log, orderProductFactory) {
  var self = this;

  self.isDisabled = false;
  self.clientSurrender = 0;
  self.saveOrder = saveOrder;
  self.cleanOrder = cleanOrder;

  $scope.clientAmmount = null;
  $scope.clientSumma = orderProductFactory.GetAmount();

  $scope.$watch('clientAmmount', function () {
    $log.debug('$watch -------> clientAmmount');
    self.clientSurrender = Surrender();
  });

  // $rootScope.$watch('clientSumma.amount',function() {
  //   $log.debug('$watch -------> clientSumma');
  //   self.clientSurrender = Surrender();
  // });

  $rootScope.$on('rootScope.summa', function () {
    $log.info('KassaAmountCtrl - rootScope.summa');
    self.clientSurrender = Surrender();
  });

  function saveOrder() {
    // $log.debug('summa:' + $scope.clientSumma);
    var res = orderProductFactory.Save();
    if (res) {
      cleanOrder();
    }
  }

  function cleanOrder() {
    $log.debug('------ cleanOrder -----');
    $rootScope.CurrentCategory = 0;
    $rootScope.products.length = 0;
    $rootScope.prod.length = 0;
    $scope.clientAmmount = null;
    orderProductFactory.Clean();
    angular.extend($scope.clientSumma, { amount: 0, money: 'грн.' });
    self.clientSurrender = 0;
  }

  function Surrender() {
    var res = 0;
    if ($scope.clientAmmount > 0) {
      res = $scope.clientAmmount - $scope.clientSumma.amount;
    }
    $log.info('Surrender=' + res);
    return res;
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
/* jshint undef: true, unused: true */
/* global angular */

/*
 * Line below lets us save `this` as `self`
 * to make properties look exactly the same as in the template
 */
// jscs:disable safeContextKeyword
'use strict';

MenuCtrl.$inject = ['$rootScope', '$scope', '$log'];

function MenuCtrl($rootScope, $scope, $log) {
  var self = this;

  self.isDisabled = false;
}
/* jshint undef: true, unused: true */
/* global angular */

/*
 * Line below lets us save `this` as `self`
 * to make properties look exactly the same as in the template
 */
// jscs:disable safeContextKeyword
'use strict';

KassaProductCtrl.$inject = ['$rootScope', '$scope', '$http', '$element', '$log', 'orderProductFactory'];

function KassaProductCtrl($rootScope, $scope, $http, $element, $log, orderProductFactory) {
  var self = this;

  self.isDisabled = false;
  // self.newState = newState;
  self.removeItem = removeItem;
  // $scope.items = $rootScope.products;
  $scope.items = orderProductFactory.GetAll();
  self.amount = orderProductFactory.GetAmount();

  function removeItem(id) {
    // $rootScope.products.splice(id, 1);
    orderProductFactory.Delete(id);
    // $rootScope.$emit('rootScope.summa');
    // $log.info('remove ID:' + id);
  }

  // function newState(id) {
  //   $rootScope.CurrentCategory = id;
  //   $rootScope.$emit('rootScope.emit');
  //   $log.info('Click category ID:' + $rootScope.CurrentCategory);
  // }

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
/* jshint undef: true, unused: true */
'use strict';
// jscs:disable safeContextKeyword

KassaProductCategoryCtrl.$inject = ['$rootScope', '$scope', '$http', '$element', '$log', 'loadProductsFactory', 'orderProductFactory'];

function KassaProductCategoryCtrl($rootScope, $scope, $http, $element, $log, loadProductsFactory, orderProductFactory) {
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
    // $rootScope.products.push(item);
    orderProductFactory.Add(item);
    // $rootScope.products = obj;
    // $rootScope.$emit('rootScope.summa');
  }
}