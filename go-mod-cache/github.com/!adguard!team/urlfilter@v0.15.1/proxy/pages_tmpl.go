// Code generated by pages DO NOT EDIT.

package proxy

import "text/template"

const blockedPageTemplate = "<!DOCTYPE html>\n" +
"<html>\n" +
"\n" +
"<head>\n" +
"    <meta charset=\"utf-8\" />\n" +
"    <title>{{ .Hostname}}</title>\n" +
"\n" +
"<style type=\"text/css\">\n" +
"body {\n" +
"  background-color: #efefef;\n" +
"  font-family: system-ui, sans-serif;\n" +
"}\n" +
"\n" +
".wrapper {\n" +
"  display: flex;\n" +
"  flex-direction: column;\n" +
"  justify-content: center;\n" +
"  align-items: center;\n" +
"  text-align: center;\n" +
"  min-height: 100vh;\n" +
"  min-width: 320px;\n" +
"  width: 50%;\n" +
"  margin: 0 auto;\n" +
"}\n" +
"\n" +
".ic-error {\n" +
"  height: 100px;\n" +
"  width: 100px;\n" +
"  padding: 0px;\n" +
"  margin: 0px;\n" +
"  background: url('data:image/svg+xml;charset=utf-8,%3C%3Fxml%20version%3D%221.0%22%3F%3E%0A%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22%23000000%22%20viewBox%3D%220%200%20128%20128%22%20width%3D%2264px%22%20height%3D%2264px%22%3E%20%20%20%20%3Cpath%20d%3D%22M%2064%209.4003906%20C%2039.2%209.4003906%2019%2029.600391%2019%2054.400391%20C%2019%2078.200391%2037.6%2097.800781%2061%2099.300781%20L%2061%20115.59961%20C%2061%20117.29961%2062.3%20118.59961%2064%20118.59961%20C%2065.7%20118.59961%2067%20117.29961%2067%20115.59961%20L%2067%2099.300781%20C%2090.4%2097.700781%20109%2078.200391%20109%2054.400391%20C%20109%2029.600391%2088.8%209.4003906%2064%209.4003906%20z%20M%2064%2015.400391%20C%2085.5%2015.400391%20103%2032.900391%20103%2054.400391%20C%20103%2075.900391%2085.5%2093.400391%2064%2093.400391%20C%2042.5%2093.400391%2025%2075.900391%2025%2054.400391%20C%2025%2032.900391%2042.5%2015.400391%2064%2015.400391%20z%20M%2047.300781%2049.400391%20C%2044.500781%2049.400391%2042.300781%2051.600391%2042.300781%2054.400391%20C%2042.300781%2057.200391%2044.500781%2059.400391%2047.300781%2059.400391%20L%2080.699219%2059.400391%20C%2083.499219%2059.400391%2085.699219%2057.200391%2085.699219%2054.400391%20C%2085.699219%2051.600391%2083.499219%2049.400391%2080.699219%2049.400391%20L%2047.300781%2049.400391%20z%22%2F%3E%3C%2Fsvg%3E');\n" +
"  background-size: 100%;\n" +
"  background-repeat: no-repeat;\n" +
"}\n" +
"\n" +
".details {\n" +
"  width: 100%;\n" +
"  margin-top: 25px;\n" +
"}\n" +
"\n" +
".details textarea {\n" +
"  width: 100%;\n" +
"  background-color: #efefef;\n" +
"  font-size: 14px;\n" +
"  font-family: monospace;\n" +
"  text-align: center;\n" +
"\n" +
"  border: none;\n" +
"  overflow: auto;\n" +
"  outline: none;\n" +
"  -webkit-box-shadow: none;\n" +
"  -moz-box-shadow: none;\n" +
"  box-shadow: none;\n" +
"  resize: none;\n" +
"}\n" +
"\n" +
"</style>\n" +
"<script type=\"text/javascript\">\n" +
"/* pages v1.0.0 Tue Nov 12 2019 */\n" +
"(function () {\n" +
"    'use strict';\n" +
"\n" +
"    document.addEventListener('DOMContentLoaded', () => {\n" +
"        const btn = document.getElementById('details-button');\n" +
"        btn.addEventListener('click', () => {\n" +
"            btn.classList.toggle('open');\n" +
"            document.querySelector('.details').classList.toggle('hidden');\n" +
"        });\n" +
"    });\n" +
"\n" +
"}());\n" +
"\n" +
"</script>\n" +
"</head>\n" +
"\n" +
"<body>\n" +
"    <div class=\"wrapper\">\n" +
"        <div class=\"ic-error\"></div>\n" +
"        <h1>Request blocked</h1>\n" +
"        <div class=\"description\">\n" +
"            Request to <strong>{{ .Hostname }}</strong> is blocked by the content blocking rule:\n" +
"        </div>\n" +
"        <div class=\"details\">\n" +
"            <textarea readonly rows=\"4\">{{ .RuleText}}</textarea>\n" +
"        </div>\n" +
"    </div>\n" +
"</body>\n" +
"\n" +
"</html>\n"

var blockedPageTmpl = template.Must(template.New("blockedPage").Parse(blockedPageTemplate))
