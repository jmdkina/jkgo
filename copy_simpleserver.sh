#!/bin/bash

dirs="html/addon"
files="html/css/global.css html/css/jmdkina.css "
files+="html/css/resume.css html/css/shici.css "
files+="html/jmdkina/add.html html/jmdkina/jmdkina.html "
files+="html/shici/add.html html/shici/shici.html "
files+="html/resume/resume.html "
files+="html/404.html "
files+="html/js/jmdkina/add.js html/js/jmdkina/jmdkina.js "
files+="html/js/shici/add.js html/js/shici/shici.js "
files+="html/js/global.js "

source=`pwd`
dst=$source/../bin/snaky-bin

for i in $dirs do
	cp -rfv $source/$i $dst/$i
done

for i in $files do
	cp -rfv $source/$i $dst/$i
done