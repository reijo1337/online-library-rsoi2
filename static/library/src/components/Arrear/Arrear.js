import React, {Component} from "react"
import "bootstrap/dist/css/bootstrap.css"
import {Button, Panel} from "react-bootstrap"
import "../../tools"
import {parse_date, parse_json} from "../../tools";

class Arrear extends Component{
    constructor(props) {
        super(props);
        const {arrear} = this.props;
        this.arrear = arrear;
        this.url = "http://localhost:5000/arrear"

    }
    render() {
        const text = this.arrear.book_name + " от " + parse_date(this.arrear.start) + " до " + parse_date(this.arrear.end);
        return (
            <Panel>
                {text}
                <Button onClick={this.handleDelete}>
                    Удалить
                </Button>
            </Panel>
        )
    }
    handleDelete = () => {
        console.log("DELETING ARREAR");
        fetch(this.url + "?id=" + String(this.arrear.id), {
            method: "delete",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
        })
            .then( res => {
                if (res.status === 200) {
                    return parse_json(res);
                } else {
                    throw new Error();
                }
            })
            .then(json => {
                this.props.handleDel(this.arrear.id);

            })
            .catch((error) => {
                alert("Cant make arrear: " + error.toString());
            });
    }
}

export default Arrear