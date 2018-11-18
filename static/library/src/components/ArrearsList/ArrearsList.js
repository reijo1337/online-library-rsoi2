import React from "react"
import {ListGroupItem, ListGroup} from "react-bootstrap"
import Arrear from "../Arrear/Arrear"

export default function ArrearsList({arrears}) {
    const arrearsList = arrears.map(ar =>
            <ListGroupItem key={ar.id}>
                <Arrear arrear={ar}/>
            </ListGroupItem>
    );
    console.log(arrearsList);
    return (
        <div>
            <ListGroup>
                {arrearsList}
            </ListGroup>
        </div>
    )
}