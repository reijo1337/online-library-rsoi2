import React, {Component} from "react"
import "bootstrap/dist/css/bootstrap.css"
import {Button, Panel} from "react-bootstrap"
import "../../tools"
import {parse_date} from "../../tools";

class Arrear extends Component{
    render() {
        const {arrear} = this.props;
        const text = arrear.book_name + " от " + parse_date(arrear.start) + " до " + parse_date(arrear.end);
        return (
            <Panel>
                {text}
                <Button onClick={this.handleDelete(arrear.id)}>
                    Удалить
                </Button>
            </Panel>
        )
    }
    handleDelete = (data) => {
        console.log(data);
        this.props.handleDel(data);
    }
}

export default Arrear