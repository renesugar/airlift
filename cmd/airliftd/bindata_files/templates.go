package bindata_files

import (
	"path/filepath"
	"time"

	"ktkr.us/pkg/vfs/bindata"
)

func init() {
	bindata.RegisterFile(filepath.Join("templates", "content", "config.tmpl"), time.Unix(1528665552, 0), []byte("{{ define \"title\" }} \xe2\x80\xa2 Configure{{ end }}\n\n{{ define \"content\" }}\n  {{ template \"%overview\" . }}\n  {{ template \"%config\" . }}\n  <script src=\"/-/static/common.js\"></script>\n  <script src=\"/-/static/config.js\"></script>\n{{ end }}\n\n{{ define \"%config\" }}\n{{ with $.Data.Data }}\n  <section id=\"section-config\" class=\"floating-section\">\n    <h1>Configuration</h1>\n    <form id=\"config\" autocomplete=\"off\">\n      <div class=\"box\" id=\"host-box\" data-tooltip=\"Returned file links will begin with this domain and path.\" data-tt-pos=\"top\">\n        <label for=\"host\">Base URL</label>\n        <input type=\"text\" id=\"host\" name=\"host\" value=\"{{ .Conf.Host }}\" placeholder=\"i.example.com\">\n      </div>\n      <div class=\"box\" id=\"id-box\">\n        /<span id=\"sample-id\"></span><span id=\"sample-ext\">.ext</span>\n      </div>\n      <div class=\"box\">\n        <label for=\"id-size\">Length of File ID</label>\n        <input type=\"range\" id=\"id-size\" name=\"id-size\" min=\"2\" max=\"12\" value=\"{{ .Conf.HashLen }}\">\n      </div>\n      <div class=\"box checkbox\" data-tooltip=\"Enable to append the original file extension to returned links.\" data-tt-pos=\"left\">\n        <input type=\"checkbox\" id=\"append-ext\" name=\"append-ext\"{{ if .Conf.AppendExt }} checked{{ end }}>\n        <label for=\"append-ext\">Append File Extensions</label>\n      </div>\n      <div class=\"box check-enable\">\n        <input type=\"checkbox\" class=\"hider\" id=\"enable-age-prune\" name=\"enable-age-prune\"{{ if .Conf.MaxAgeEnable }} checked{{ end }}>\n        <label for=\"enable-age-prune\">Limit Upload Age</label>\n        <div class=\"hidee\">\n          <label for=\"max-age\">Maximum Age (Days)</label>\n          <input type=\"number\" id=\"max-age\" name=\"max-age\" value=\"{{ .Conf.Age }}\" min=\"0\"{{ if not .Conf.MaxAgeEnable }} disabled{{ end }}>\n        </div>\n      </div>\n      <div class=\"box check-enable\">\n        <input type=\"checkbox\" class=\"hider\" id=\"enable-size-prune\" name=\"enable-size-prune\"{{ if .Conf.MaxSizeEnable }} checked{{ end }}>\n        <label for=\"enable-size-prune\">Limit Total Uploads Size</label>\n        <div class=\"hidee\">\n          <label for=\"max-size\">Maximum Size (MB)</label>\n          <input type=\"number\" id=\"max-size\" name=\"max-size\" value=\"{{ .Conf.Size }}\" min=\"0\"{{ if not .Conf.MaxSizeEnable }} disabled{{ end }}>\n        </div>\n      </div>\n      <div class=\"box check-enable\" data-tooltip=\"Enable to allow uploads to show Twitter Cards with file previews if applicable.\" data-tt-pos=\"left\">\n        <input type=\"checkbox\" class=\"hider\" id=\"twitter-card\" name=\"twitter-card\"{{ if .Conf.TwitterCardEnable }} checked{{ end }}>\n        <label for=\"twitter-card\">Enable Twitter Cards</label>\n        <div class=\"hidee\">\n          <label for=\"twitter-handle\">Twitter Handle</label>\n          <input type=\"text\" id=\"twitter-handle\" name=\"twitter-handle\" value=\"{{ .Conf.TwitterHandle }}\" required placeholder=\"@handle\"{{ if not .Conf.TwitterCardEnable }} disabled{{ end }}>\n        </div>\n      </div>\n      <div class=\"box check-enable\" data-tooltip=\"Enable to format code text files with syntax highlighting\" data-tt-pos=\"left\">\n        <input type=\"checkbox\" class=\"hider\" id=\"syntax-enable\" name=\"syntax-enable\"{{ if .Conf.SyntaxEnable }} checked{{ end }}>\n        <label for=\"syntax-enable\">Syntax Highlighting</label>\n        <small>\n          <a href=\"https://xyproto.github.io/splash/docs/\" target=\"_blank\">View theme examples</a>\n        </small>\n        <div class=\"hidee\">\n          <label for=\"syntax-theme\">Syntax Theme</label>\n          <select id=\"syntax-theme\" name=\"syntax-theme\">\n            {{ range .SyntaxThemes }}\n              <option value=\"{{ . }}\" {{ if eq . $.Data.Data.Conf.SyntaxTheme }} selected {{ end }} >{{ . }}</option>\n            {{ end }}\n          </select>\n        </div>\n      </div>\n      <div class=\"box\" id=\"directory-box\">\n        <label for=\"directory\">Upload Directory</label>\n        <input type=\"text\" id=\"directory\" name=\"directory\" value=\"{{ .Conf.Directory }}\" placeholder=\"/home/user/uploads\">\n      </div>\n      <div class=\"box\" id=\"newpass-box\" data-tooltip=\"Enter a new password here to change your password.\" data-tt-pos=\"right\">\n        <label for=\"newpass\">New Password</label>\n        <input type=\"password\" id=\"newpass\" name=\"newpass\" placeholder=\"\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\">\n      </div>\n      <div class=\"box\" id=\"newpass-confirm-box\" data-tooltip=\"Confirm new password\" data-tt-pos=\"left\">\n        <label for=\"newpass-confirm\">Confirm New Password</label>\n        <input type=\"password\" id=\"newpass-confirm\" name=\"newpass-confirm\" required placeholder=\"\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\">\n      </div>\n      <button id=\"submit\" type=\"button\">Update configuration</button>\n    </form>\n  </section>\n{{ end }}\n{{ end }}\n\n{{ define \"%overview\" }}\n{{ with $.Data.Data }}\n  <section id=\"section-overview\" class=\"floating-section\">\n    <h1>Overview</h1>\n    <p><strong><a href=\"/-/history/0\">{{ .NumUploads }} upload{{ if ne .NumUploads 1 }}s{{ end }}</a></strong> totalling <strong>{{ .UploadsSize }}</strong>. (<a id=\"purge-all-link\" href=\"javascript:void(0)\">purge</a>)</p>\n    <p>Thumbnail cache is <strong>{{ .ThumbsSize }}</strong>. (<a id=\"purge-thumbs-link\" href=\"javascript:void(0)\">purge</a>)</p>\n  </section>\n{{ end }}\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "content", "default-index.tmpl"), time.Unix(1527653725, 0), []byte("{{ define \"content\" }}\n    <section id=\"front\">\n      <div id=\"big-logo\">\n        <div id=\"big-logo-text\">{{ $.Data.Config.Host }} is powered by <a href=\"https://github.com/moshee/airlift\">Airlift</a>.</div>\n      </div>\n      <div class=\"login-link\"><a href=\"/-/login\">Log in</a></div>\n    </section>\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "content", "errors.tmpl"), time.Unix(1528661634, 0), []byte("{{ define \"400\" }}<!doctype html>\n<html>\n  <head>\n    <title>400</title>\n    <link rel=\"stylesheet\" href=\"/-/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>You're doing it wrong.</h1>\n      {{ if .Err }}<p>{{ .Err }}</p>{{ end }}\n    </div>\n  </body>\n</html>\n{{ end }}\n{{ define \"404\" }}<!doctype html>\n<html>\n  <head>\n    <title>404</title>\n    <link rel=\"stylesheet\" href=\"/-/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>This isn't the page you're looking for.</h1>\n    </div>\n  </body>\n</html>\n{{ end }}\n{{ define \"500\" }}<!doctype html>\n<html>\n  <head>\n    <title>500</title>\n    <link rel=\"stylesheet\" href=\"/-/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>Something went wrong.</h1>\n      {{ if .Err }}<p>{{ .Err }}</p>{{ end }}\n    </div>\n  </body>\n</html>\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "content", "history.tmpl"), time.Unix(1527698648, 0), []byte("{{ define \"title\" }} \xe2\x80\xa2 Uploads{{ end }}\n\n{{ define \"content\" }}\n{{ template \"%history\" . }}\n<script src=\"/-/static/common.js\"></script>\n<script src=\"/-/static/history.js\"></script>\n{{ end }}\n\n{{ define \"%history\" }}\n{{ with $.Data.Data }}\n<section id=\"history\">\n  {{ if len .List | lt 25 }}{{ template \"%pagination\" . }}{{ end }}\n  <ul>\n    {{ range .List }}\n    <li class=\"history-item\" data-id=\"{{ .ID }}\">\n      <a href=\"/{{ .ID }}{{ if $.Data.Data.AppendExt }}{{ .Ext }}{{ end }}\" class=\"upload-link\">{{ if .HasThumb }}<img src=\"/-/thumb/{{ .ID }}.jpg\">{{ else }}<img src=\"/-/static/file.svg\"><div class=\"file-ext-overlay\">{{ .Ext }}</div>{{ end }}</a>\n      <div class=\"history-item-name\" title=\"{{ .Name }}\">{{ .Name }}</div>\n      <div class=\"history-item-data\">{{ .Size }} / <span title=\"{{ .Uploaded.Format \"2006-01-02 15:04:05 MST\" }}\">{{ .Ago }}</span></div>\n      <div class=\"history-item-data\"><a href=\"javascript:\" class=\"delete-upload\">Delete</a></div>\n    </li>\n    {{ end }}\n  </ul>\n  {{ template \"%pagination\" . }}\n</section>\n{{ end }}\n{{ end }}\n\n{{ define \"%pagination\" }}\n<nav class=\"pagination\">\n  <span class=\"prevnext{{ if gt .CurrentPage 1 }} active{{ end }}\"><a href=\"/-/history/{{ .PrevPage }}\">Back</a> \xe2\x80\x94</span>\n  Page {{ .CurrentPage }} of {{ .TotalPages }}\n  <span class=\"prevnext{{ if ne .NextPage 0 }} active{{ end }}\">\xe2\x80\x94 <a href=\"/-/history/{{ .NextPage }}\">Next</a></span>\n</nav>\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "content", "index.tmpl"), time.Unix(1527653732, 0), []byte("{{ define \"content\" }}\n  <section id=\"upload\" class=\"floating-section\">\n    <input type=\"file\" id=\"picker\" name=\"picker[]\" multiple>\n    <div id=\"drop-zone\">\n      <div class=\"progress-bar\"></div>\n      <div id=\"drop-zone-text\">Click/tap/drop/paste</div>\n    </div>\n    <div id=\"uploaded-urls\">\n      <ul></ul>\n    </div>\n  </section>\n  <script src=\"/-/static/common.js\"></script>\n  <script src=\"/-/static/uploader.js\"></script>\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "content", "login.tmpl"), time.Unix(1527653743, 0), []byte("{{ define \"title\" }} \xe2\x80\xa2 Log In{{ end }}\n\n{{ define \"content\" }}\n    <section id=\"section-login\" class=\"floating-section\">\n      <form method=\"post\" action=\"/-/login\" id=\"login\">\n        {{ if $.Data }}<p id=\"message-box\" class=\"bad active\">Incorrect password.</p>{{ end }}\n        <label for=\"password\">Password: </label><input name=\"pass\" id=\"password\" type=\"password\" placeholder=\"password\" autofocus required>\n        <hr>\n        <button type=\"submit\" id=\"submit\">Log in</button>\n      </form>\n    </section>\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "content", "syntax.tmpl"), time.Unix(1528661858, 0), []byte("{{ define \"title\" }}{{ $.Data.Data.Filename }}{{ end }}\n\n{{ define \"content\" }}\n  <main>{{ $.Data.Data.HTML }}</main>\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "content", "twitterbot.tmpl"), time.Unix(1527653775, 0), []byte("{{ define \"content\" }}<!doctype html>\n<html>\n  <head>\n    <title>{{ .ID }}</title>\n    <meta name=\"twitter:card\" content=\"summary_large_image\">\n    <meta name=\"twitter:site\" content=\"{{ .Handle }}\">\n    <meta name=\"twitter:title\" content=\"{{ .ID }}: {{ .Name }}\">\n    <meta name=\"twitter:description\" content=\"{{ .Size }} / uploaded {{ .Uploaded.Format \"2 Jan 2006 15:04\" }}\">\n    <meta name=\"twitter:image\" content=\"http://{{ .Host }}/-/twitterthumb/{{ .ID }}.jpg\">\n  </head>\n  <body>\n    hi twitterbot\n  </body>\n</html>\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "errors", "errors.tmpl"), time.Unix(1443849891, 0), []byte("{{ define \"400\" }}<!doctype html>\n<html>\n  <head>\n    <title>400</title>\n    <link rel=\"stylesheet\" href=\"/-/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>You're doing it wrong.</h1>\n      {{ if .Err }}<p>{{ .Err }}</p>{{ end }}\n    </div>\n  </body>\n</html>\n{{ end }}\n{{ define \"404\" }}<!doctype html>\n<html>\n  <head>\n    <title>404</title>\n    <link rel=\"stylesheet\" href=\"/-/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>This isn't the page you're looking for.</h1>\n    </div>\n  </body>\n</html>\n{{ end }}\n{{ define \"500\" }}<!doctype html>\n<html>\n  <head>\n    <title>500</title>\n    <link rel=\"stylesheet\" href=\"/-/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>Something went wrong.</h1>\n      {{ if .Err }}<p>{{ .Err }}</p>{{ end }}\n    </div>\n  </body>\n</html>\n{{ end }}\n"))
	bindata.RegisterFile(filepath.Join("templates", "layout", "layout.tmpl"), time.Unix(1528663009, 0), []byte("{{ define \"head\" }}\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <link rel=\"shortcut icon\" href=\"/-/static/favicon.png\">\n    <link rel=\"apple-touch-icon\" sizes=\"76x76\" href=\"/-/static/airlift_76x76.png\">\n    <link rel=\"apple-touch-icon\" sizes=\"120x120\" href=\"/-/static/airlift_120x120.png\">\n    <link rel=\"apple-touch-icon\" sizes=\"152x152\" href=\"/-/static/airlift_152x152.png\">\n    <link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/-/static/airlift_180x180.png\">\n    <link rel=\"stylesheet\" href=\"/-/static/style.css\">\n{{ end }}\n\n{{ define \"layout-full\" }}\n<html>\n  <head>\n    <title>Airlift{{ block \"title\" . }}{{ end }}</title>\n    {{ template \"head\" }}\n  </head>\n  <body>\n    <div id=\"message-box\"></div>\n    <nav id=\"nav\">\n      <a href=\"/\">Upload</a> /\n      <a href=\"/-/history/1\">History</a> /\n      <a href=\"/-/config\">Configure</a> /\n      <a href=\"/-/logout\">Log out</a>\n    </nav>\n    {{ block \"content\" $ }}{{ end  }}\n    <div id=\"version\">airliftd {{ $.Data.Version }}</div>\n  </body>\n</html>\n{{ end }}\n\n{{ define \"layout-lite\" }}\n<html>\n  <head>\n    <title>Airlift{{ block \"title\" . }}{{ end }}</title>\n    {{ template \"head\" }}\n  </head>\n  <body>\n    {{ block \"content\" $ }}{{ end  }}\n  </body>\n</html>\n{{ end }}\n\n{{ define \"layout-syntax\" }}\n<html>\n<head>\n  <title>{{ block \"title\" . }}{{ end }}</title>\n  <link rel=\"stylesheet\" href=\"/-/static/syntax.css\">\n  <link rel=\"stylesheet\" href=\"/-/theme/{{ .Data.Data.SyntaxTheme }}.css\">\n</head>\n<body class=\"syntax chroma\">\n  <a href=\"?raw=1\" class=\"raw\">{{ $.Data.Data.Filename }}<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4\"></path><polyline points=\"7 10 12 15 17 10\"></polyline><line x1=\"12\" y1=\"15\" x2=\"12\" y2=\"3\"></line></svg></a>\n  {{ block \"content\" . }}{{ end }}\n</body>\n{{ end }}\n"))
}
