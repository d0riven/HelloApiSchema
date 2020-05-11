import * as React from "react";
import {DefaultApi, UpdateUserOutput, UpdateUserRequest} from "../api-client";

export let UpdateUser = () => (
    <button onClick={() => {
        const c: DefaultApi = new DefaultApi();
        const input: UpdateUserRequest = {
            updateUserInput: {
                id: 1,
                address: '静岡県駿東郡小山町桑木',
                birthday: new Date('2020-03-08'),
                emailAddress: 'kin_taro@example.com',
                lastName: '金',
                firstName: '太郎',
            },
        };
        c.updateUser(input).then((v: UpdateUserOutput) => {
            console.log(v);
        });
    }}>UpdateUser</button>
);
