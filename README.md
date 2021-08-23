# sample-akita-plugin

Sample plugin for Akita Software CLI. See https://github.com/akitasoftware/akita-cli/blob/main/plugin/README.md for the current documentation on writing plugins.

This plugin rewrites every component of the HTTP path to be "bork".


Used to test https://github.com/akitasoftware/akita-cli/pull/89.

![image](https://user-images.githubusercontent.com/7472514/130297382-d386172d-a063-4884-b864-33671511cd7a.png)

# dynamically loaded

This version demonstrates dynamically loading, to make it work I had to
take the following steps:

  1. Rename the package to 'main' in plugin.go
  2. Redirect akita-cli to the source directory where I compiled the CLI.  (Without this, go will complain about package mismatches.)


