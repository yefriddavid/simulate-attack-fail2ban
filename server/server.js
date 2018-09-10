require('dotenv').config()
const express = require('express')
const app = express()
const PORT = process.env.PORT || 81


app.all('*', (req, res) => {
	res.send('Hello World,  I am the server! on ' + PORT + ' port.')
})

app.listen(PORT, () => {
	console.log('Example app listening on port ' + PORT + '!')
})

