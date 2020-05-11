import * as React from "react";
import {GetUser} from "./components/GetUser";
import {UpdateUser} from "./components/UpdateUser";
import {CreateUser} from "./components/CreateUser";

export default class App extends React.Component<any, any> {
    render() {
        return (
            <React.Fragment>
                <GetUser/>
                <UpdateUser/>
                <CreateUser/>
            </React.Fragment>
        );
    }
};
