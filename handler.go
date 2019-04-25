package gindoc

import (
	"bytes"
	"github.com/techoner/gindoc/resources/assets"
	"github.com/techoner/gindoc/resources/source"
	"github.com/techoner/gindoc/resources/views"
	"gopkg.in/russross/blackfriday.v1"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var defaultVersionName = "default"
var docsDir = "storage/docs"
var prefixUri = "docs"
var tmpl *template.Template


type Sidebar struct {
}

func Handler(name string)  []byte {
	version := defaultVersionName
	baseName := "index.html"
	dirName := ""
	versions := make(map[string]string)

	if len(name) > 0 {
		bname := path.Base(name)
		fext := path.Ext(name)
		fname := strings.TrimSuffix(bname, fext)
		dname := path.Dir(name)
		if dname == "." {
			dname = "/"
		}

		fragment := strings.SplitN(dname, "/", 2)

		v := name
		if len(fragment) > 0 {
			v = fragment[0]
		}

		versions = getVersion(v)
		if len(versions) > 0 {
			version = v
		}
		dirName = strings.Replace(dname, version, "", -1)

		if version != fname {
			baseName = bname
		}
	}

	versions = getVersion(version)
	sidebar := getSidebar(version)

	contentFileName := path.Join(dirName, baseName)
	content := getContent(version, contentFileName)

	var buf bytes.Buffer
	t := template.New("")
	tmpl, err := t.Parse(views.Index)
	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(&buf, "gloc", map[string]interface{}{
		"css":                  assets.Index,
		"sidebar":              sidebar,
		"content":              template.HTML(content),
		"versions":             versions,
		"current_version":      version,
		"prefix_uri":           path.Join("/", prefixUri, "/"),
		"basePath":             path.Join("/", prefixUri, version) + "/",
		"contentFileName":      strings.TrimLeft(contentFileName, "/"),
		"default_version_name": defaultVersionName,
	})

	return buf.Bytes()
}

func getVersion(version string) map[string]string {
	versions := make(map[string]string)

	p := getStorageFilePath("versions.yml")

	if !isFile(p) {
		return versions
	}

	versions = yamlParseFile(p)
	_, ok := versions[version]
	if !ok {
		return versions
	}

	_, ok = versions["default"]
	if !ok {
		versions["default"] = "默认版本"
	}

	return versions
}

func getContent(version string, p string) string {
	if version == defaultVersionName {
		version = ""
	}
	p = getStorageFilePath(
		path.Join(version, "_source", strings.TrimSuffix(p, path.Ext(p))),
	)

	p = p + ".md"

	exist := isFile(p)
	var content []byte
	if !exist {
		content = []byte(source.Default)
	}else{
		content, _ = ioutil.ReadFile(p)
	}

	content = blackfriday.MarkdownCommon(content)

	return string(content)
}

func getSidebar(version string) map[string]map[string]string {
	sidebars := make(map[string]map[string]string)
	p := "sidebar.yml"
	if version != defaultVersionName {
		p = version + "/" + p
	}

	p = getStorageFilePath(p)

	if !isFile(p) {
		return sidebars
	}

	data, _ := ioutil.ReadFile(p)

	yaml.Unmarshal(data, &sidebars)

	return sidebars
}

func getStorageFilePath(name string) string {
	return path.Join(docsDir, name)
}

func isFile(p string) bool {

	f, err := os.Stat(p)
	if err != nil {
		return false
	}

	return !f.IsDir()
}

func yamlParseFile(p string) map[string]string {
	versions := make(map[string]string)

	data, _ := ioutil.ReadFile(p)
	yaml.Unmarshal(data, &versions)

	return versions
}
