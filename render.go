// Copyright 2017 Catherine Jones. All rights reserved.
// The use of this source code is governed by a 3-clause
// BSD-style license as described in the LICENSE file.

package main

import (
	"github.com/catherinejones/blog"
	"log"
)

// Render the entire site
func RenderSite(config *blog.Config) {
	updated, err := blog.RenderTree(config)
	if nil != err {
		log.Print(err.Error())
	}

	if updated && "" != config.System.BlogPath {
		err = blog.RenderBlogIndex(config)
		if nil != err {
			log.Print(err.Error())
		}

		err = blog.RenderArchives(config)
		if nil != err {
			log.Print(err.Error())
		}

		err = blog.RenderFeed(config)
		if nil != err {
			log.Print(err.Error())
		}
	}
}
