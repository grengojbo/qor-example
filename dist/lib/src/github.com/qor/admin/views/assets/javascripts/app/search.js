$(function () {

  'use strict';

  var location = window.location;
  var a='a';

  $('.qor-search').each(function () {
    var $this = $(this);
    var $input = $this.find('.qor-search__input');
    var $clear = $this.find('.qor-search__clear');
    var isSearched = !!$input.val();

    var emptySearch = function () {
      var search = location.search.replace(new RegExp($input.attr('name') + '\\=?\\w*'), '');
      if (search == '?'){
        location.href = location.href.split('?')[0]
      } else {
        location.search = location.search.replace(new RegExp($input.attr('name') + '\\=?\\w*'), '');
      }
    }

    $this.closest('.qor-page__header').addClass('has-search');

    $clear.on('click', function () {
      if ($input.val() || isSearched) {
        emptySearch();
      } else {
        $this.removeClass('is-dirty');
      }
    });
  });
});
