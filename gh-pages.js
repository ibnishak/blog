#!/usr/bin/env node

var ghpages = require('gh-pages');

ghpages.publish('public', function(err) {});