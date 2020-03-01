import * as React from "react";
import {DefaultApi, GetUserOutput, UpdateUserRequest} from "../api-client";

export let UpdateUser = () => (
    <button onClick={() => {
        const c: DefaultApi = new DefaultApi();
        const input: UpdateUserRequest = {
            updateUserInput: {
                // address: '',
                birthday: new Date(),
                // emailAddress: '',
                // firstName: '',
                id: 1,
                // lastName: '',
            },
        };
        c.updateUser(input).then((v: GetUserOutput) => {
            console.log(v);
        });
    }}>UpdateUser</button>
);
