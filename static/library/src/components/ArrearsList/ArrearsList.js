import React, {Component} from "react"
import {ListGroupItem, ListGroup, Well, Button} from "react-bootstrap"
import Arrear from "../Arrear/Arrear"
import AddArrear from "../AddArrear/AddArrear"
import {parse_json} from "../../tools";

class ArrearsList extends Component {
    constructor(props) {
        super(props);
        let {arrears} = this.props;
        this.state = {
            arrears: arrears,
            show: false
        };

        this.url = "http://localhost:5000/freeBooks";
        this.name = this.props.name
    }

    render() {
        const arrearsList = this.state.arrears.map(ar =>
            <ListGroupItem key={ar.id}>
                <Arrear arrear={ar} handleDel={this.handleDelete}/>
            </ListGroupItem>
        );
        const addArrear = (this.state.show && this.books.length > 0) &&
            <AddArrear
                props={this.books}
                hadleClose={this.handleClose}
                hanldeAddArrear={this.handleAddArrear}
                readerName={this.props.name}
            />;
        return (
            <div>
                <Well bsSize="large">Книги записанные на {this.name}</Well>
                <Button
                    block
                    bsSize="large"
                    onClick={this.handleShow}
                >
                    Добавить новую запись
                </Button>
                {addArrear}
                <ListGroup>
                    {arrearsList}
                </ListGroup>
            </div>
        )
    }

    handleDelete = (id) => {
        let index = -1;
        let arrears = this.state.arrears;
        arrears.forEach(
            function(item, i, arr) {
                if (item.id === id) {
                    index = i;
                }
            });
        arrears.splice(index, 1);
        this.setState({arrears: arrears, show: false});
    };

    handleShow = () => {
        fetch(this.url, {
            method: "get",
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
                if (json.length === 0) {
                    alert("В библиотеке закончились книги =(");
                }
                this.books = json;
                this.setState({ show: true })
            })
            .catch((error) => {
                alert("Cant get free books: " + error.toString());
            });
    };

    handleClose = ()  => {
        this.setState({ show: false });
    };

    handleAddArrear = (arrear) => {
        console.log(arrear);
        let arrears = this.state.arrears;
        arrears.push(arrear);
        this.setState({
            arrears: arrears,
            show: false
        });
    }
}

export default ArrearsList