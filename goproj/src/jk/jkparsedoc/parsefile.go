package jkparsedoc

import (
	"bufio"
	"errors"
	"jk/jklog"
	"os"
	"strings"
)

// parse // will read the next line, we must check if we read it
// and take the line to next parse
var previousLine = ""

const (
	key_comment_slave = "comment_slave"
)

// Input one file of .h and read the data out put to .html

func parse_get_basename(filename string) string {
	basefilename := filename
	i := strings.LastIndex(filename, "/")
	if i >= 0 {
		basefilename = filename[i+1 : len(filename)]
	}
	i = strings.LastIndex(basefilename, ".")
	if i >= 0 {
		basefilename = basefilename[0:i]
	}
	return basefilename
}

func parse_get_basepath(filename string) string {
	basepath := ""
	i := strings.LastIndex(filename, "/")
	if i > 0 {
		basepath = filename[0:i]
	}
	return basepath
}

type ParseHeaderFile struct {
	all_string   map[string][]string
	baseFileName string
	basePath     string
}

func (p *ParseHeaderFile) exist_comment_slave(key string) (string, bool) {
	lenslave := len(p.all_string[key])
	if lenslave == 0 {
		return "", false
	}
	com_str := p.all_string[key][0]
	if len(com_str) == 0 {
		return "", false // ignore this one for no comment
	}
	return com_str, true
}

// filename must be absolute position as we don't known where we are
func parse_header_file(filename string) (*ParseHeaderFile, error) {
	jklog.L().Debugln("parse file : ", filename)
	if !strings.HasSuffix(filename, ".h") {
		return nil, errors.New("Ingore")
	}
	// open the file
	fi, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	buf := bufio.NewReader(fi)

	headerfile_string := ParseHeaderFile{}
	headerfile_string.all_string = make(map[string][]string)
	all_string := headerfile_string.all_string

	headerfile_string.baseFileName = parse_get_basename(filename)
	headerfile_string.basePath = parse_get_basepath(filename)

	jklog.L().Debugln("parse start here...")
	//  Read From file
	for {
		line := previousLine
		if len(line) > 0 {
			// clear previousLine
			previousLine = ""
			// line has be read by other so first parse it
		} else {
			// Need to read
			line, err = buf.ReadString('\n')
			if err != nil {
				break
			}
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == '\n' {
			continue
		}
		if line[len(line)-1] != '\n' {
			line += "\n"
		}

		// jklog.L().Debugln("to parse string : ", line)
		if strings.HasPrefix(line, "#define") {
			jklog.L().Debugln("find define : ", line)
			com_str, bex := headerfile_string.exist_comment_slave(key_comment_slave)
			if !bex {
				continue
			}
			// read the define
			def_string := line
			i := strings.LastIndex(line, ")")
			if i > 0 {
				def_string = line[0 : i+1]
			}
			def_string = com_str + def_string
			if def_string[len(def_string)-1] != '\n' {
				def_string += "\n"
			}
			all_string["define"] = append(all_string["define"], def_string)
			delete(all_string, "comment_slave")
		} else if strings.HasPrefix(line, "enum") {
			jklog.L().Debugln("find enum : ", line)
			// To sure if the enum has comment
			com_str, bex := headerfile_string.exist_comment_slave(key_comment_slave)
			if !bex {
				continue
			}
			// Take enum all out
			enum_str := line
			// If only one line
			if strings.Contains(line, "}") {
				enum_str = com_str + enum_str
				all_string["enum"] = append(all_string["enum"], enum_str)
				delete(all_string, key_comment_slave)
				continue
			}
			for {
				// Read All enum define out
				line, err := buf.ReadString('\n')
				if err != nil {
					break
				}
				if strings.Contains(line, "}") {
					enum_str += line
					break
				}
				enum_str += line
			}
			enum_str = com_str + enum_str
			all_string["enum"] = append(all_string["enum"], enum_str)
			delete(all_string, "comment_slave")
		} else if strings.HasPrefix(line, "typedef") || strings.HasPrefix(line, "struct") {
			jklog.L().Debugln("find typedef : ", line)
			com_str, bex := headerfile_string.exist_comment_slave(key_comment_slave)
			if !bex {
				continue
			}
			// Take all typedef
			typedef_str := line
			// Just a define, no {}
			if strings.Contains(line, ";") {
				typedef_str = com_str + typedef_str
				all_string["typedef"] = append(all_string["typedef"], typedef_str)
				delete(all_string, "comment_slave")
				continue
			}
			// else Read done the typedef
			for {
				line, err := buf.ReadString('\n')
				if err != nil {
					break
				}
				if strings.Contains(line, "}") {
					typedef_str += line
					break
				}
				typedef_str += line
			}
			typedef_str = com_str + typedef_str
			all_string["typedef"] = append(all_string["typedef"], typedef_str)
			delete(all_string, "comment_slave")
		} else if line[0] == '/' && line[1] == '*' {
			jklog.L().Debugln("find comment : ", line)
			if line[2] == '*' || line[2] == '=' {
				// Take all comment
				com_str := line
				if strings.Contains(line, "*/") {
					continue
				}
				for {
					line, err := buf.ReadString('\n')
					if err != nil {
						break
					}
					if strings.Contains(line, "*/") {
						com_str += line
						break
					}
					com_str += line
				}
				delete(all_string, "comment_main")
				all_string["comment_main"] = append(all_string["comment_main"], com_str)
			} else {
				// Take all comment / other comment
				com_str := line
				if strings.Contains(line, "*/") {
					continue
				}
				for {
					line, err := buf.ReadString('\n')
					if err != nil {
						break
					}
					if strings.Contains(line, "*/") {
						com_str += line
						break
					}
					com_str += line
				}
				delete(all_string, key_comment_slave)
				all_string[key_comment_slave] = append(all_string[key_comment_slave], com_str)
			}
		} else if line[0] == '/' && line[1] == '/' {
			jklog.L().Debugln("find // comment: ", line)
			com_str := line
			for {
				line, err := buf.ReadString('\n')
				if err != nil {
					break
				}
				line = strings.TrimSpace(line)
				if len(line) < 3 {
					break
				}
				if line[0] == '\n' {
					break
				}
				// If the next is not // start, it is the last comment
				if !strings.HasPrefix(line, "//") {
					previousLine = line
					break
				}
				if line[0] == '/' && line[1] == '/' {
					lenv := len(line)
					if lenv == 3 && (line[2] == '>' || line[2] == '/') { // The last mark
						break
					}
					com_str += line + "\n"
				}
			}
			delete(all_string, key_comment_slave)
			all_string[key_comment_slave] = append(all_string[key_comment_slave], com_str)
		} else {
			// Go to find function define
			func_str := line
			com_str, bex := headerfile_string.exist_comment_slave(key_comment_slave)
			if !bex {
				continue
			}
			// if strings.HasPrefix(line, "int") || strings.HasPrefix(line, "void") ||
			// strings.HasPrefix(line, "short") || strings.HasPrefix(line, "long") ||
			// strings.HasPrefix(line, "extern") || strings.HasPrefix(line, "char") {
			if strings.Contains(line, "(") {
				if strings.Contains(line, ";") {
					// Only one line
					func_str = com_str + func_str
					all_string["function"] = append(all_string["function"], func_str)
					delete(all_string, key_comment_slave)
					continue
				}
				for {
					line, err := buf.ReadString('\n')
					if err != nil {
						break
					}
					line = strings.TrimSpace(line)
					line += "\n"
					if len(line) < 3 {
						break
					}
					if line[0] == '\n' {
						break
					}
					if strings.Contains(line, ";") {
						func_str += line
						break
					}
				}
				func_str = com_str + func_str
				all_string["function"] = append(all_string["function"], func_str)
				delete(all_string, key_comment_slave)
			}
		}
	}

	jklog.L().Debugln("parse over")
	return &headerfile_string, nil
}

// Output what we parsed out, debug use
func (p *ParseHeaderFile) PrintOut() {
	jklog.L().Debugln("something output here... ")
	for k, v := range p.all_string["define"] {
		jklog.L().Infoln("define: ", k, " -- \n", v)
	}
	jklog.L().Debugln("enum here ...")
	for k, v := range p.all_string["enum"] {
		jklog.L().Infoln("enum: ", k, " -- \n", v)
	}
	jklog.L().Debugln("typedef ... ")
	for k, v := range p.all_string["typedef"] {
		jklog.L().Infoln("typedef : ", k, " -- \n", v)
	}
	jklog.L().Debugln("comment here ... ")
	for k, v := range p.all_string["comment_main"] {
		jklog.L().Infoln("comment main: ", k, " -- \n", v)
	}
	for k, v := range p.all_string["comment_slave"] {
		jklog.L().Infoln("Comment slave : ", k, " -- \n", v)
	}
	jklog.L().Debugln("function here ...")
	for k, v := range p.all_string["function"] {
		jklog.L().Infoln("function: ", k, " -- \n", v)
	}
}

// Write files to html file
// depends on what we parse
func (p *ParseHeaderFile) WriteToHtml() error {
	fo, err := os.Create("bvdoc/" + p.basePath + "/" + p.baseFileName + ".html")
	if err != nil {
		return err
	}
	defer fo.Close()

	jklog.L().Infoln("Write to html file : ", "bvdoc/"+p.basePath+"/"+p.baseFileName+".html")

	// How many levels need to upper
	upper_level := strings.Count(p.basePath, "/")
	upper_level += 1

	generateHtmlHeader(fo, p.baseFileName, upper_level)
	// Write back link
	generateHtmlALink(fo, upper_level, "back")

	// 1. Write comment main
	lencomment := len(p.all_string["comment_main"])
	if lencomment > 0 {
		com_str := p.all_string["comment_main"][0]
		generateHtmlMainComment(fo, com_str)
	}
	fo.WriteString("\n")

	// 2. Write define
	obj := "<h3 class='text-info'><strong>Define</strong> from here</h3>"
	fo.WriteString(obj)
	for _, v := range p.all_string["define"] {
		generateHtmlDefineComment(fo, "", v)
	}
	fo.WriteString("\n")

	// 3. Write enum
	obj = "<h3 class='text-info'><strong>Enum</strong> from here</h3>"
	fo.WriteString(obj)
	for _, v := range p.all_string["enum"] {
		generateHtmlDefineComment(fo, "", v)
	}
	fo.WriteString("\n")

	// 4. Write typedef
	obj = "<h3 class='text-info'><strong>Typedef</strong> from here</h3>"
	fo.WriteString(obj)
	for _, v := range p.all_string["typedef"] {
		generateHtmlDefineComment(fo, "", v)
	}
	fo.WriteString("\n")

	// 5. write function
	obj = "<h3 class='text-info'><strong>Function</strong> from here</h3>"
	fo.WriteString(obj)
	for _, v := range p.all_string["function"] {
		generateHtmlDefineComment(fo, "", v)
	}

	generateHtmlFooter(fo, upper_level)

	return nil
}
