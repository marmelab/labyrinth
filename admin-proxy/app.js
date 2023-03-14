const POSTGREST_URL=process.env.POSTGREST_URL;

const express = require('express')
const proxy = require('express-http-proxy')
const cookieParser = require('cookie-parser');
const { json } = require('express');

const app = express();
app.use(cookieParser());

app.use(proxy(POSTGREST_URL, {
    proxyReqOptDecorator: function(proxiedRequest, originalRequest) {
        return new Promise(function(resolve, reject) {
            const bearerToken = originalRequest.cookies['_JWT'];
            if (!bearerToken) {
                return reject({status: 401, message: "User is not authenticated."});
            }

            proxiedRequest.headers['Authorization'] = `Bearer ${bearerToken}`,
            resolve(proxiedRequest);
        });
    },
}));

app.use((err, req, res, next) => {
    console.error(err);

    res.status(err.status || 500).json(err);
});

const server = app.listen(3000);
process.on('SIGINT', () => server.close());
