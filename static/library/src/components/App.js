import React, {Component} from "react"
import ReaderSearch from "./ReaderSearch/ReaderSearch"
import {PageHeader, } from "react-bootstrap"

class App extends Component{
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