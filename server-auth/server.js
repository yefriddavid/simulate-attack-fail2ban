require('dotenv').config()
const PORT = process.env.PORT || 81
const http = require('http')
const auth = require('basic-auth')
const proxy = require('http-proxy').createProxyServer()

// Create server
const server = http.createServer(function (req, res, next) {
  const credentials = auth(req)
  if (!credentials || credentials.name !== 'abcd' || credentials.pass !== '1234') {
    res.statusCode = 401
    res.setHeader('WWW-Authenticate', 'Basic realm="example"')
    res.end('Access denied')
  } else {
    if(typeof process.env.APP_TARGET == "undefined"){
      res.end('Access granted')
    }else{
      //res.end('else')
      next()
      if(process.env.APP_TARGET == ""){
        //res.end('Access granted')
        //next()
        //res.notFound()
        //res.end(proxy.env.APP_TARGET + ' -> Access granted')
      }else{
        /*proxy.web(req, res, {
          target: process.env.APP_TARGET
        })*/
      }
    }
  }
})


server.listen(PORT)

