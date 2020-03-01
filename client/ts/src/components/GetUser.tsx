import * as React from "react";
import {DefaultApi, GetUserOutput} from "../api-client";

export let GetUser = () => (
    <button onClick={() => {
        const c: DefaultApi = new DefaultApi();
        c.getUser({id: 1}).then((v: GetUserOutput) => {
            console.log(v);
        });
    }}>GetUser</button>
);