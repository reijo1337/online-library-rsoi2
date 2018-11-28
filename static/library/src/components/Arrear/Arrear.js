import React, {Component} from "react"
import "bootstrap/dist/css/bootstrap.css"
import {Button, Glyphicon, Panel} from "react-bootstrap"
import "../../tools"
import {parse_date, parse_json} from "../../tools"

class Arrear extends Component{
    constructor(props) {
        super(props);
        const {arrear} = this.props;
        this.arrear = arrear;
        this.url = "http://localhost:5000/arrear"

    }
    render() {
        const header = this.arrear.book_name + '. ' + this.arrear.book_author;
        const text = "от " + parse_date(this.arrear.start) + " до " + parse_date(this.arrear.end);
        return (
            <Panel>
                <Panel.Heading>
                    <Panel.Title>{header}</Panel.Title>
                </Panel.Heading>
                    {text}
                    <Button onClick={this.handleDelete}>
                        <Glyphicon glyph="trash" />
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
                    return res.json();
                }
            })
            .then(json => {
                if (json.error) {
                    throw new Error(json.error);
                }
                this.props.handleDel(this.arrear.id);

            })
            .catch((error) => {
                alert("Cant make arrear: " + error.message);
            });
    }
}

export default Arrear