export function parse_json(data) {
    return data.text().then(function (text) {
        return text ? JSON.parse(text) : {}
    }).catch((err) => {
        console.log(err);
        return {};
    });
}