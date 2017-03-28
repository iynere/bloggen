package main

import (
    "github.com/catherinejones/blog"
    "flag"
    "os"
)

const SITE_CONF = "site.conf"

func main() {
    var clean, help bool
    var config_path string

    flag.BoolVar(&clean, "clean", false, "Rebuild the site from scratch.")
    flag.BoolVar(&clean, "c", false, "Short form version of -clean.")
    flag.BoolVar(&help, "help", false, "Display this help text.")
    flag.BoolVar(&help, "h", false, "Display this help text.")

    flag.StringVar(&config_path, "config", SITE_CONF, "The path to the site configuration.")

    flag.Parse()

    if help {
        flag.PrintDefaults()
        os.Exit(0)
    }

    config := blog.LoadConfig(config_path)

    config.Options.Clean = clean

    blog.RenderSite(config)
}
