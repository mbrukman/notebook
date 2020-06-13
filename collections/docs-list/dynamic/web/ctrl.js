// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

angular.module('DocsList', [], function($locationProvider) {
  $locationProvider.html5Mode(true);
});

function DocsListCtrl($scope, $routeParams, $location) {
  $scope.docs = DOCS.files;

  // Searching.
  $scope.location = $location;
  $scope.query = $location.search().q;

  // Display ordering.
  $scope.predicate = 'name';
  $scope.reverse = false;

  // Used in orderBy:identity to sort by string value of tags.
  $scope.identity = function(x) { return x; };

  // If the new value is the same as old value, reverse the sort order.
  // Else, update the predicate to the new value.
  $scope.setPredicate = function(pred) {
    if ($scope.predicate == pred) {
      $scope.reverse = !$scope.reverse;
    } else {
      $scope.predicate = pred;
    }
  };

  var URL_DOCS_PREFIX = 'https://docs.google.com/document/d/';
  var URL_SHEETS_PREFIX = 'https://docs.google.com/spreadsheets/d/';
  var URL_SLIDES_PREFIX = 'https://docs.google.com/presentation/d/';
  var URL_DRAWINGS_PREFIX = 'https://docs.google.com/drawings/d/';
  var URL_SITES_PREFIX = 'https://sites.google.com/';

  var FILETYPE_DOCS = 'docs';
  var FILETYPE_SHEETS = 'sheets';
  var FILETYPE_SLIDES = 'slides';
  var FILETYPE_DRAWINGS = 'drawings';
  var FILETYPE_SITES = 'sites';
  var FILETYPE_PDF = 'pdf';
  var FILETYPE_TEXT = 'text';
  var FILETYPE_UNKNOWN = 'unknown';

  $scope.filetype_ = function(doc) {
    var href = $scope.docUrl(doc);
    if (!href) {
      return null;
    }
    var urlToFiletype = [
      [URL_DOCS_PREFIX, FILETYPE_DOCS],
      [URL_SHEETS_PREFIX, FILETYPE_SHEETS],
      [URL_SLIDES_PREFIX, FILETYPE_SLIDES],
      [URL_DRAWINGS_PREFIX, FILETYPE_DRAWINGS],
      [URL_SITES_PREFIX, FILETYPE_SITES],
    ];
    for (urlFiletype of urlToFiletype) {
      if (href.indexOf(urlFiletype[0]) == 0) {
        return urlFiletype[1];
      }
    }
    return FILETYPE_UNKNOWN;
  };

  $scope.isDoc = function(doc) {
    return $scope.filetype_(doc) == FILETYPE_DOCS;
  };

  $scope.isSheet = function(doc) {
    return $scope.filetype_(doc) == FILETYPE_SHEETS;
  };

  $scope.isSlides = function(doc) {
    return $scope.filetype_(doc) == FILETYPE_SLIDES;
  };

  $scope.isDrawings = function(doc) {
    return $scope.filetype_(doc) == FILETYPE_DRAWINGS;
  };

  $scope.isSites = function(doc) {
    return $scope.filetype_(doc) == FILETYPE_SITES;
  };

  $scope.isPDF = function(doc) {
    return $scope.filetype_(doc) == FILETYPE_PDF;
  };

  $scope.isText = function(doc) {
    return $scope.filetype_(doc) == FILETYPE_TEXT;
  };

  $scope.isUnknown = function(doc) {
    return $scope.filetype_(doc) == FILETYPE_UNKNOWN;
  };

  $scope.docUrl = function(doc) {
    return doc.webViewLink;
  };

  /**
   * Sets the current search query in the text box via the model.
   *
   * @param {string} query
   */
  $scope.search = function(query) {
    $scope.query = query;
  };
}
