// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.github.muratbayan.mattermost-file-blocker",
  "name": "File Blocker",
  "description": "This plugin helps block the upload of unauthorized file extensions to the Mattermost server.",
  "version": "0.2.1",
  "min_server_version": "5.12.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "AllowedExtensions",
        "display_name": "Allowed extensions",
        "type": "text",
        "help_text": "Comma separated list of allowed extensions for file attachments",
        "placeholder": "",
        "default": ""
      },
      {
        "key": "ExtensionIsRequired",
        "display_name": "Require extension",
        "type": "bool",
        "help_text": "Set to true if file attachments require having an extension",
        "placeholder": "",
        "default": false
      },
      {
        "key": "CheckMimeType",
        "display_name": "[Experimental] Validate MIME Content",
        "type": "bool",
        "help_text": "Set to true if the plugin should check the MIME content type",
        "placeholder": "",
        "default": false
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
