import React, {Component} from "react"
import {FormControl, FormGroup, ControlLabel, Button, Jumbotron} from "react-bootstrap"
import {parse_json, updater} from "../../tools";
import jwtDecode from "jwt-decode";

class Login extends Component {
    constructor(props) {
        super(props);
        this.url = "http://localhost:5000/auth";
        const token = localStorage.getItem("accessToken");
        if (token === null || token === "") {
            this.state = {
                login: '',
                password: '',
                authorized: false,
            };
            return
        }
        let tokenData = jwtDecode(token);
        let interval = (tokenData.exp - (Date.now().valueOf() / 1000))-10;
        if (interval < 0) {
            localStorage.setItem("accessToken", "");
            localStorage.setItem("refreshToken", "");
            localStorage.setItem("login", "");
            this.state = {
                login: '',
                password: '',
                authorized: false,
            }
        } else {
            const login = localStorage.getItem("login");
            this.state = {
                login: login,
                password: '',
                authorized: true,
            };
        }
    }

    render() {
        let body;
        if (!this.state.authorized) {
            body = <form onSubmit={this.handleSubmit}>
                <FormGroup controlId="login" >
                    <ControlLabel>Логин</ControlLabel>
                    <FormControl
                        autoFocus
                        type="login"
                        value={this.state.login}
                        onChange={this.handleChange}
                    />
                </FormGroup>
                <FormGroup controlId="password" >
                    <ControlLabel>Пароль</ControlLabel>
                    <FormControl
                        value={this.state.password}
                        onChange={this.handleChange}
                        type="password"
                    />
                </FormGroup>
                <Button
                    block
                    bsSize="large"
                    disabled={!this.validateForm()}
                    type="submit"
                >
                    Авторизоваться
                </Button>
            </form>
        } else {
                    body = <Jumbotron>
                        <h1>Вы авторизованы как {this.state.login}</h1>
                        <p>
                            Для управления библиотекой перейдите в раздел "Управление записями"
                        </p>
                    </Jumbotron>
        }
        return (body);
    }

    handleChange = event => {
        this.setState({
            [event.target.id]: event.target.value
        });
    };

    validateForm() {
        return this.state.login.length > 0 && this.state.password.length > 0;
    }

    handleSubmit = event => {
        event.preventDefault();
        const login = this.state.login;
        const password = this.state.password;
        const data = JSON.stringify({
            login: login,
            password: password,
        });
        debugger;
        fetch(this.url, {
            method: "post",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
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
                localStorage.setItem("accessToken", json.accessToken);
                localStorage.setItem("refreshToken", json.refreshToken);
                localStorage.setItem("login", this.state.login);
                clearInterval(this._tokenUpdater);
                const token = json.accessToken;
                let tokenData = jwtDecode(token);
                let interval = (tokenData.exp - (Date.now().valueOf() / 1000))-10;

                this._tokenUpdater = setInterval(updater.bind(this),interval*1000);
                this.setState({authorized: true})
            })
            .catch((error) => {
                alert("Проблемы с доступом в джойказино: " + error.message);
            });
    };
}

export default Login