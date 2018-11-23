export function parse_json(data) {
    return data.text().then(function (text) {
        return text ? JSON.parse(text) : {}
    }).catch((err) => {
        console.log(err);
        return {};
    });
}

export function parse_date(date) {
    const year = date.slice(0, 4);
    const mounth = date.slice(4, 6);
    const day = date.slice(6, 8);
    return day + "." + mounth + "." + year;
}