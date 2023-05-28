module.exports = function jsonResponse(body, corsOrigin) {
    return {
        statusCode: 200,
        body: JSON.stringify(body),
        headers: {
            'Content-Type': 'application/json',
            'AccessControl-Allow-Origin': corsOrigin
        }
    };
};