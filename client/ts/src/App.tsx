import * as React from "react";
import {DefaultApi, GetUserOutput, GetUserRequest} from "./api-client";

class GetUserInput implements GetUserRequest {
    constructor(public id:number) {
    }
}

export default class App extends React.Component<any, any> {
    render() {
        return (
            <div>
                <button onClick={() => {
                    let c:DefaultApi = new DefaultApi();
                    c.getUser(new GetUserInput(1)).then((v:GetUserOutput) => {
                        console.log(v);
                    });
                }}>hello
                </button>
                <p>hello</p>
            </div>
        );
    }
};
