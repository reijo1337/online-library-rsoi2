import React, {Component} from "react"
import ReaderSearch from "./ReaderSearch/ReaderSearch"
import {PageHeader, } from "react-bootstrap"
import  "bootstrap/dist/css/bootstrap.css"

// const Routes = [
//     { name: "main page"},
//     { name: "get arrears" },
//     { name: "delete arrear" },
//     { name: "add arrear" }
// ];


class App extends Component{
    state = {
        route: 0
    };

    render() {
        return (
            <div>
                <PageHeader>
                    Помощник библиотекаря
                </PageHeader>
                <ReaderSearch/>
            </div>
        )
    }
}

export default App