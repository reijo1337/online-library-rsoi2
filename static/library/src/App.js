import React, {Component} from "react"
import MainPage from "./components/MainPage/MainPage";
import Login from "./components/Login/Login";
import {NavItem, Nav} from "react-bootstrap";

const ROUTES = [
    { name: "Личный кабинет" },
    { name: "Управление записями" },
];

class App extends Component{
    constructor() {
        super();
        this.state = {
            route: 0
        };
    }

    render() {
        const route = this.state.route;
        let mainPart;
        if (route === 0) {
            mainPart = <Login/>;
        } else if (route === 1) {
            mainPart = <MainPage/>;
        } else {
            mainPart = <Login/>
        }

        return (
            <div>
                <Nav
                    bsStyle="pills"
                    stacked
                    activeKey={route}
                    onSelect={index => {
                        this.setState({ route: index });
                    }}
                >

                    {ROUTES.map((rout, index) => (
                        <NavItem key={index} eventKey={index}>{rout.name}</NavItem>
                    ))}
                </Nav>
                {mainPart}

        </div >
        )
    }
}

export default App