import React, {Component} from "react"
import "bootstrap/dist/css/bootstrap.css"
import {Button} from "react-bootstrap"

class Arrear extends Component{
    constructor(props) {
        super(props)
    }
    render() {
        const {arrear} = this.props;
        console.log(arrear);
        const text = arrear.book_name + " от " + arrear.start + " до " + arrear.end;
        return (
            <div className="card">
                {text}
                <Button>Удалить</Button>
            </div>
        )
    }
}

export default Arrear