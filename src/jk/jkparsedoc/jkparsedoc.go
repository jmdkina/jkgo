// Package jkparsedoc test to parse c file to html doc.
//
// See doc.go
package jkparsedoc

import (
	"jk/jklog"
	"os"
	"path/filepath"
	"strings"
)

type JKParseDoc struct {
	files  []string
	prefix string // save position
}

// Parse One file and write to html
func (p *JKParseDoc) parse_one_file(filename string) error {
	ph, err := parse_header_file(filename)
	if err != nil {
		jklog.L().Errorln("parse error : ", err)
		return err
	}
	// p.PrintOut()
	jklog.L().Infoln("Write to file ", filename, " ...")
	ph.SetPrefix(p.prefix)
	err = ph.WriteToHtml()
	if err != nil {
		return err
	}
	return nil
}

// parse dir and set the files to the JKParseDoc.files
func (p *JKParseDoc) parse_dir_files(filename string) error {
	filepath.Walk(filename, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
			jklog.L().Infoln("dir : ", fi.Name(), " of ", path, " with ", filename)
			p.parse_dir_files(filename)
		} else {
			// jklog.L().Infoln("files : ", path, fi.Name())
			if fi.Name()[0] == '.' || !strings.HasSuffix(fi.Name(), ".h") {
				return nil
			}
			p.files = append(p.files, path)
		}
		return nil
	})
	return nil
}

// Call parse_one_file to parse all files
func (p *JKParseDoc) parse_files() error {
	for _, v := range p.files {
		err := p.parse_one_file(v)
		if err != nil {
			jklog.L().Errorln("parse file ", v, " error: ", err)
		}
	}
	return nil
}

// Print out what we find of the dir.
// Debug use.
func (p *JKParseDoc) Output() {
	for _, v := range p.files {
		jklog.L().Infoln("files : ", v)
	}
}

func (p *JKParseDoc) generateIndexFile() error {
	fo, err := os.Create(p.prefix + "/" + "index.html")
	if err != nil {
		return err
	}

	generateHtmlHeader(fo, "index.html", 0)

	// Add some section
	section := "<h5>Doc is generated by jkparsedoc of the author jmdkina@gmail.com</h5>\n"
	generateHtmlString(fo, section)

	// <ul>
	// <li><a href=xxx>xxx</a></li>
	// </ul>
	str := "<ul class='list-unstyled'>\n"
	for _, v := range p.files {
		htmlName := strings.TrimRight(v, ".h") + ".html"
		// TODO: don't add if the dir donesn't exist
		str += "<li><a href='" + htmlName + "'>" + htmlName + "</a></li>\n"
	}
	str += "</ul>\n"
	generateHtmlString(fo, str)

	generateHtmlFooter(fo, 0)

	return nil
}

// Parse start here, give one file or dir
// It will check all file if give dir
// @docsname: create what dir to save docs
func JKParseDocStart(filename string, docsname string) {
	jklog.L().SetLevel(jklog.LEVEL_MUCH)
	jklog.L().Infoln("parse doc start ", filename)

	p := &JKParseDoc{}

	// create docsname if it is not exists.
	err := os.Mkdir(docsname, os.ModeDir)
	if err != nil {
		// docs not exists.
		os.Create(docsname)
	}

	p.prefix = docsname

	dir, err := os.Stat(filename)
	if err != nil {
		return
	}
	if dir.IsDir() {
		err = p.parse_dir_files(filename)
		if err != nil {
			jklog.L().Errorln("pasre dir : ", err)
		}
		// Only dir need generate index file
		p.generateIndexFile()
	} else {
		p.files = append(p.files, filename)
	}
	// p.Output()
	p.parse_files()

	jklog.L().Infoln("All done, exist...")
}