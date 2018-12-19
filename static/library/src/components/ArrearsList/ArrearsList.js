import React, {Component} from "react"
import {Pager, ListGroupItem, ListGroup, Well, Button} from "react-bootstrap"
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
        this.urlArrears = "http://localhost:5000/getUserArrears";
        this.name = this.props.name;
        this.page = 1;
    }

    render() {
        debugger;
        const arrearsList = this.state.arrears.map(ar =>
            <ListGroupItem key={ar.id}>
                <Arrear arrear={ar} handleDel={this.handleDelete}/>
            </ListGroupItem>
        );

        const addArrear = (this.state.show && this.books.length > 0) &&
            (<AddArrear
                props={this.books}
                hadleClose={this.handleClose}
                hanldeAddArrear={this.handleAddArrear}
                readerName={this.props.name}
            />);

        const disabledNext = this.state.arrears.length !== 5;
        const disabledPrev = this.page === 1;

        const pager = (<Pager>
            <Pager.Item disabled={disabledPrev} previous onClick={this.prevPage}>
                &larr; Previous
            </Pager.Item>
            <Pager.Item disabled={disabledNext} next onClick={this.nextPage}>
                Next &rarr;
            </Pager.Item>
        </Pager>);

        return (
            <div>
                <Well bsSize="small">Книги записанные на {this.name}</Well>
                <Button
                    block
                    bsSize="large"
                    onClick={this.handleShow}
                >
                    Добавить новую запись
                </Button>
                {addArrear}
                {pager}
                <ListGroup>
                    {arrearsList}
                </ListGroup>
            </div>
        )
    }

    handleDelete = (id) => {
        this.setState({show: false}, () => {
            this.loadPage();
        });
    };

    handleShow = () => {
        const token = localStorage.getItem("accessToken");
        fetch(this.url + "?access_token="+token, {
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
                    return res.json();
                }
            })
            .then(json => {
                if (json.error) {
                    throw new Error(json.error);
                }
                if (json.length === 0) {
                    alert("В библиотеке закончились книги =(");
                }
                this.books = json;
                this.setState({ show: true })
            })
            .catch((error) => {
                alert("Cant get free books: " + error.message);
            });
    };

    handleClose = ()  => {
        this.setState({ show: false });
    };

    handleAddArrear = (arrear) => {
        this.setState({show: false}, () => {
            this.loadPage();
        });
    };

    prevPage = () => {
        this.page = this.page > 1 ? this.page - 1 : this.page;
        this.setState({ show: false }, () => {
            this.loadPage();
        });
    };

    nextPage = () => {
        this.page += 1;
        this.setState({ show: false }, () => {
            this.loadPage();
        });
    };

    loadPage = () => {
        const token = localStorage.getItem("accessToken");
        const url = this.urlArrears + "?name=" + this.name + "&page=" + this.page + "&access_token="+token;
        fetch(url)
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
                debugger;
                let arrear = json;
                this.setState({arrears: arrear});

            })
            .catch((error) => {
                alert("Cant get arrears: " + error.message);
            });
    }
}

export default ArrearsList