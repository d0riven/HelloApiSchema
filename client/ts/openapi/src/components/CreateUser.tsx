import * as React from "react";
import {CreateUserOutput, CreateUserRequest, DefaultApi} from "../api-client";

export let CreateUser = () => (
    <button onClick={() => {
        const c: DefaultApi = new DefaultApi();
        const input: CreateUserRequest = {
            createUserInput: {
                address: '東京都新宿区西新宿２丁目８−１',
                birthday: new Date(),
                emailAddress: 'taro@example.com',
                lastName: '田中',
                firstName: '太郎',
            },
        };
        c.createUser(input).then((v: CreateUserOutput) => {
            console.log(v);
        });
    }}>CreateUser</button>
);
