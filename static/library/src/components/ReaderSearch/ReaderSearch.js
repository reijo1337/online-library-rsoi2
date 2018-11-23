import React from "react"
import {Button, Form, FormControl, FormGroup} from "react-bootstrap"
import "../../tools"
import {parse_json} from "../../tools";
import ArrearList from "../ArrearsList/ArrearsList"

class ReaderSearch extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            value: '',
            isLoaded: false,
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.url = "http://localhost:5000/getUserArrears";
    }
    render() {
        const arrearsList = this.state.isLoaded && <ArrearList arrears={this.arrears} name={this.state.value}/>;
        return(
            <div>
                <Form onSubmit={this.handleSubmit}>
                    <FormGroup
                        controlId="formBasicText"
                    >
                        <FormControl
                            type="text"
                            placeholder="Имя читателя"
                            value={this.state.value}
                            onChange={this.handleChange}
                        />
                        <FormControl.Feedback />
                        <Button
                            block
                            bsSize="large"
                            type="submit"
                        >
                            Поиск карточки читателя
                        </Button>
                    </FormGroup>
                </Form>
                {arrearsList}
            </div>
        )
    }

    handleChange(event) {
        this.setState({value: event.target.value});
    }

    handleSubmit(event) {
        event.preventDefault();
        const url = this.url + "?name=" + this.state.value;
        const request = async () => {
            const res = await fetch(url);
            if (res.status === 200) {
                this.arrears = await parse_json(res);
                this.setState({ isLoaded: true });
            } else {
                alert("cant load news!");
            }
        };
        request();
    }
}

export default ReaderSearch