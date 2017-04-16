// Copyright 2017 Catherine Jones. All rights reserved.
// The use of this source code is governed by a 3-clause
// BSD-style license as described in the LICENSE file.

package main

import (
	"github.com/catherinejones/blog"
	"log"
)

// Render the entire site
func RenderSite(context *blog.SiteContext, clean bool) {
	updated, err := blog.RenderTree(context, clean)
	if nil != err {
		log.Print(err.Error())
	}

	if updated && "" != context.Config.BlogPath {
		err = blog.RenderBlogIndex(context)
		if nil != err {
			log.Print(err.Error())
		}

		err = blog.RenderArchives(context)
		if nil != err {
			log.Print(err.Error())
		}

		err = blog.RenderFeed(context)
		if nil != err {
			log.Print(err.Error())
		}
	}
}
