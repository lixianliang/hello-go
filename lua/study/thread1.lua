
require "socket"

host = "www.w3.org"
file = "/TR/REC-html32.html"
local http = require("socket")
local res = http.request("http://www.baidu.com")
print(res)
