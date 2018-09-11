require('dotenv').config()
const PORT = process.env.PORT || 81
const http = require('http')
const auth = require('basic-auth')

// Create server
const server = http.createServer(function (req, res) {
  const credentials = auth(req)

  if (!credentials || credentials.name !== 'abcd' || credentials.pass !== '1234') {
    res.statusCode = 401
    res.setHeader('WWW-Authenticate', 'Basic realm="example"')
    res.end('Access denied')
  } else {
    //res.end('Access granted')
    res.notFound()
  }
})


// Listen
server.listen(PORT)


