#!/bin/bash

rm dist/resources/app/main*bundle*.js*

set -e 

#webpack --config webpack.config.min.js
webpack --config webpack.config.js

sed -i '' 's|"main|"js/main|' dist/*.html

cp -r htmlpage/* dist/resources/app/
cp node_modules/bootstrap/dist/css/bootstrap.min.css dist/resources/app/css/
