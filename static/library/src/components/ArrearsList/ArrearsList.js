import React, {Component} from "react"
import {ListGroupItem, ListGroup, Well} from "react-bootstrap"
import Arrear from "../Arrear/Arrear"

class ArrearsList extends Component {
    constructor(props) {
        super(props);
        let {arrears} = this.props;
        this.arrears = arrears;
    }

    render() {
        const arrearsList = this.arrears.map(ar =>
            <ListGroupItem key={ar.id}>
                <Arrear arrear={ar} handleDel={this.handleDelete}/>
            </ListGroupItem>
        );
        return (
            <div>
                <Well bsSize="large">Книги записанные на {this.props.name}</Well>
                <ListGroup>
                    {arrearsList}
                </ListGroup>
            </div>
        )
    }

    handleDelete = (id) => {
        console.log(id);
        let index = -1;
        this.arrears.forEach(
            function(item, i, arr) {
                if (item.id === id) {
                    index = i;
                }
            });
        this.arrears.splice(index, 1)
    }
}

export default ArrearsList