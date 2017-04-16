// Copyright 2017 Catherine Jones. All rights reserved.
// The use of this source code is governed by a 3-clause
// BSD-style license as described in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/catherinejones/blog"
	"os"
)

const SITE_CONF = "site.conf"
const SITE_ROOT = "http"
const SITE_TEMPLATES = "data/templates"

func main() {
	var clean, help bool
	var config_path, site_path, templates_path string

	flag.BoolVar(&clean, "clean", false, "Rebuild the site from scratch.")
	flag.BoolVar(&clean, "c", false, "Short form version of -clean.")
	flag.BoolVar(&help, "help", false, "Display this help text.")
	flag.BoolVar(&help, "h", false, "Display this help text.")
	flag.StringVar(&site_path, "site", SITE_ROOT, "The path to the site's document root.")
	flag.StringVar(&templates_path, "templates", SITE_TEMPLATES, "The path to the site's templates.")

	flag.StringVar(&config_path, "config", SITE_CONF, "The path to the site configuration.")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	config, err := blog.NewContext(site_path, templates_path, config_path)
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	RenderSite(config, clean)
}
