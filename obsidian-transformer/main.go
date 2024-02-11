package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/adrg/frontmatter"
)

// Post represents the blog post fields
type Post struct {
	Title      string
	Date       string
	LastMod    string
	Categories []string
	Body       string
}

func main() {
	var intakefm struct {
		Title string `yaml:"title"`
		ID    string `yaml:"id"`
	}

	file := os.Args[1:][0]
	dat, err := osReadFile(file)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}

	body, err := frontmatter.Parse(bytes.NewReader(dat), &intakefm)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	dateStr := intakefm.ID[:len(intakefm.ID)-3]
	india, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	tm, err := time.ParseInLocation("20060102150405", dateStr, india)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	blogDatelo := "2006-01-02T15:04:05-07:00"
	out := Post{
		Title:      intakefm.Title,
		Date:       tm.Format(blogDatelo),
		LastMod:    time.Now().Format(blogDatelo),
		Categories: []string{},
		Body:       string(body),
	}
	tmpl, err := getTmpl(out)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}

	outFile := generateOutFileName(file)
	f, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	err = tmpl.Execute(f, out)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println(outFile)
}

func osReadFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var size int
	if info, err := f.Stat(); err == nil {
		size64 := info.Size()
		if int64(int(size64)) == size64 {
			size = int(size64)
		}
	}
	size++
	if size < 512 {
		size = 512
	}
	data := make([]byte, 0, size)
	for {
		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
	}
}

func getTmpl(out Post) (*template.Template, error) {
	const postStruct = `---
title: {{ .Title }}
date: {{ .Date }}
lastmod: {{ .LastMod }}
author: Riz

description: 
categories: []
tags: []

draft: false
enableDisqus : false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---

{{ .Body }}`
	tmp := template.New("Template")
	tmpl, err := tmp.Parse(postStruct)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func generateOutFileName(file string) string {
	basename := filepath.Base(file)
	noext := strings.TrimSuffix(basename, filepath.Ext(basename))
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	slug := reg.ReplaceAllString(noext, "-")
	s := strings.ToLower(slug)
	o := filepath.Join("./content/post", s) + ".md"
	return o
}
