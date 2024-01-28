import { gql } from "@apollo/client";

interface Login {
    success: boolean,
    jwt: string | null
}

interface LoginData {
    login: Login
}

const LOGIN = gql`
    mutation Login($email: String!, $password: String!) {
        login(email: $email, password: $password) {
            success
            jwt
        }
    }
`

export {
    LOGIN, LoginData
};