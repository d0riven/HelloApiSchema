import * as React from "react";
import {CreateUserRequest, DefaultApi, GetUserOutput, UpdateUserRequest} from "../api-client";

export let CreateUser = () => (
    <button onClick={() => {
        const c: DefaultApi = new DefaultApi();
        const input: CreateUserRequest = {
            createUserInput: {
                address: '東京都新宿区西新宿２丁目８−１',
                birthday: new Date(),
                emailAddress: 'taro@example.com',
                firstName: '田中',
                lastName: '太郎',
            },
        };
        c.createUser(input).then((v: GetUserOutput) => {
            console.log(v);
        });
    }}>CreateUser</button>
);
