package main

import (
	"embed"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"unsafe"

	"github.com/urfave/cli/v2"
	"github.com/whoisix/ddl2plantuml/constants"
	"github.com/whoisix/ddl2plantuml/driver"
)

//go:embed default.tmpl
var tmpl embed.FS

func main() {
	app := cli.NewApp()
	app.Name = "ddl2plantuml"
	app.Usage = "Convert DDL to PlantUML"
	app.Description = "ddl2plantuml is a tool to generate plantuml ER diagram from database ddl."
	app.Version = constants.Version
	app.Action = action
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "driver",
			Aliases: []string{"d"},
			Usage:   "database driver",
			Value:   "mysql",
		},
		&cli.StringFlag{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "plantuml template file",
		},
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    "ddl sql file, required",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "output directory",
			Value:   ".",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	var d driver.Driver
	switch strings.ToLower(c.String("driver")) {
	case "mysql":
		d = &driver.Mysql{}
	default:
		return cli.Exit("unsupported driver", 1)
	}

	// get template
	tpl, err := getTemplateFile(c.String("template"))
	if err != nil {
		return cli.Exit(err, 1)
	}

	// get ddl from sql file
	ddl, err := ioutil.ReadFile(c.String("file"))
	if err != nil {
		return cli.Exit(err, 1)
	}

	// parse ddl
	tables, err := d.Parse(*(*string)(unsafe.Pointer(&ddl)))
	if err != nil {
		return cli.Exit(err, 1)
	}
	relationship := tables.Relationship()

	// generate plantuml file
	outputFile := path.Join(c.String("output"), "er.puml")
	data := tmplData{
		Tables:       tables,
		Relationship: relationship,
	}

	err = data.generate(tpl, outputFile)
	if err != nil {
		return cli.Exit(err, 1)
	}

	log.Print("file saved: ", outputFile)
	return nil
}

type tmplData struct {
	Tables       []driver.Table
	Relationship []driver.Relation
}

func (t tmplData) generate(tpl *template.Template, outputFile string) error {
	dir := path.Dir(outputFile)
	if !exists(dir) {
		err := mkDir(dir)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(
		outputFile,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	err = tpl.Execute(file, t)
	if err != nil {
		return err
	}
	return nil
}

func getTemplateFile(filepath string) (*template.Template, error) {

	if filepath == "" {
		return template.ParseFS(tmpl, "default.tmpl")
	}

	return template.ParseFiles(filepath)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func mkDir(path string) error {
	dir, _ := os.Getwd()
	return os.MkdirAll(dir+"/"+path, os.ModePerm)
}
