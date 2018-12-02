import React, {Component} from "react"
import {Panel, Button} from "react-bootstrap"
import  "bootstrap/dist/css/bootstrap.css"
import {parse_json} from "../../tools"

class AddArrear extends Component {
    constructor(props) {
        super(props);

        this.books = props.props;
        this.state = {
            selected_book_id: props.props[0].id,
            selected_book_index: 0,
        };

        this.url = "http://localhost:5000/arrear"
    }

    render() {
        const bookName = this.books[this.state.selected_book_index].name;
        const author = this.books[this.state.selected_book_index].author;
        const header = <div>Запись книги {bookName}. Автор  {author}</div>;

        return (
            <div className="card">
                <Panel>
                    <Panel.Heading>
                        <Panel.Title>{header}</Panel.Title>
                    </Panel.Heading>
                    <Panel.Body>
                        {this.renderDropdownButton()}
                    </Panel.Body>
                    <Panel.Footer>
                        <Button onClick={this.props.hadleClose}>Закрыть</Button>
                        <Button bsStyle="primary" onClick={this.handleAddArrear}>Записать</Button>
                    </Panel.Footer>
                </Panel>
            </div>
        )
    };

    renderDropdownButton = () => {
        const bookssList = this.books.map(bk =>
            <option key={bk.id} value={bk.id}>
                {bk.name}. {bk.author}
            </option>
        );
        return (
            <div>
                <h4> Выберете книгу </h4>
                <select onChange={this.handleBookSelect} id='s1' className="form-control">
                    {bookssList}
                </select>
                {/*<h4> Выберете срок </h4>*/}
                {/*С <input type="date" className="card-group" id='startDate'/>*/}
                {/*По <input type="date" className="card-group" id='endDate'/>*/}
            </div>
        );
    };

    handleBookSelect = () => {
        const a = Number(document.getElementById('s1').value);

        let index = -1;
        this.books.forEach(function(item, i, arr) {
            if (item.id === a) {
                index = i;
            }
        });

        this.setState({
            selected_book_id: a,
            selected_book_index: index
        })
    };

    handleAddArrear = () => {
        // const start = document.getElementById('startDate').value;
        // const end = document.getElementById('endDate').value;
        const data = JSON.stringify({
            reader: this.props.readerName,
            book: this.state.selected_book_id,
        });

        fetch(this.url, {
            method: "post",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },

            //make sure to serialize your JSON body
            body: data
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
                this.props.hanldeAddArrear(json)
            })
            .catch((error) => {
                alert("Cant make arrear: " + error.message);
            });
    }
}

export default AddArrear