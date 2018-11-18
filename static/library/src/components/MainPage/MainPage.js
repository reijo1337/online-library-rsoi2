import React, {Component} from "react"
import { Button } from "react-bootstrap"

class MainPage extends Component {
    render() {
        const wellStyles = { maxWidth: "100%", margin: '0 auto 10px' };
        return (
            <div className="well" style={wellStyles}>
                <Button bsSize="large" block>
                    Карточки чистателей
                </Button>
            </div>
        )
    }
}

export default MainPage